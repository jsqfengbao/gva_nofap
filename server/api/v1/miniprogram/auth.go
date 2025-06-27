package miniprogram

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	miniprogramReq "github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram/request"
	miniprogramRes "github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram/response"
	miniprogramService "github.com/flipped-aurora/gin-vue-admin/server/service/miniprogram"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthApi struct{}

var authService = new(miniprogramService.AuthService)
var userService = new(miniprogramService.UserService)

// WxLogin 微信小程序登录
// @Tags      MiniProgram
// @Summary   微信小程序登录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      miniprogramReq.WxLoginRequest                                true  "微信登录请求"
// @Success   200   {object}  response.Response{data=miniprogramRes.WxLoginResponse,msg=string}  "登录成功"
// @Router    /miniprogram/auth/wx-login [post]
func (a *AuthApi) WxLogin(c *gin.Context) {
	var loginReq miniprogramReq.WxLoginRequest
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		global.GVA_LOG.Error("微信登录请求参数解析失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 参数验证
	if loginReq.Code == "" {
		response.FailWithMessage("登录凭证不能为空", c)
		return
	}

	// 调用微信API获取session_key和openid
	wxResult, err := authService.GetWxSessionInfo(loginReq.Code)
	if err != nil {
		global.GVA_LOG.Error("微信登录失败", zap.Error(err))
		response.FailWithMessage("微信登录失败", c)
		return
	}

	// 解密用户信息
	var userInfo *miniprogramReq.WxUserInfo
	if loginReq.EncryptedData != "" && loginReq.Iv != "" {
		userInfo, err = authService.DecryptWxUserInfo(loginReq.EncryptedData, loginReq.Iv, wxResult.SessionKey)
		if err != nil {
			global.GVA_LOG.Error("用户信息解密失败", zap.Error(err))
			response.FailWithMessage("用户信息解析失败", c)
			return
		}
	}

	// 查找或创建用户
	user, err := userService.FindOrCreateWxUser(wxResult.Openid, wxResult.Unionid, userInfo)
	if err != nil {
		global.GVA_LOG.Error("用户创建失败", zap.Error(err))
		response.FailWithMessage("用户创建失败", c)
		return
	}

	// 生成JWT token
	token, err := authService.GenerateToken(user.ID)
	if err != nil {
		global.GVA_LOG.Error("token生成失败", zap.Error(err))
		response.FailWithMessage("登录失败", c)
		return
	}

	// 返回登录结果
	loginRes := miniprogramRes.WxLoginResponse{
		Token: token,
		User:  user,
	}

	response.OkWithData(loginRes, c)
}

// Login 普通用户登录
// @Tags      MiniProgram
// @Summary   普通用户登录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      miniprogramReq.LoginRequest                                true  "登录请求"
// @Success   200   {object}  response.Response{data=miniprogramRes.LoginResponse,msg=string}  "登录成功"
// @Router    /miniprogram/auth/login [post]
func (a *AuthApi) Login(c *gin.Context) {
	var loginReq miniprogramReq.LoginRequest
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 参数验证
	if loginReq.Phone == "" || loginReq.Password == "" {
		response.FailWithMessage("手机号和密码不能为空", c)
		return
	}

	// 验证用户登录
	user, err := userService.ValidateUserLogin(loginReq.Phone, loginReq.Password)
	if err != nil {
		global.GVA_LOG.Error("用户登录验证失败", zap.Error(err))
		response.FailWithMessage("手机号或密码错误", c)
		return
	}

	// 生成JWT token
	token, err := authService.GenerateToken(user.ID)
	if err != nil {
		global.GVA_LOG.Error("token生成失败", zap.Error(err))
		response.FailWithMessage("登录失败", c)
		return
	}

	// 返回登录结果
	loginRes := miniprogramRes.LoginResponse{
		Token: token,
		User:  user,
	}

	response.OkWithData(loginRes, c)
}

// Register 用户注册
// @Tags      MiniProgram
// @Summary   用户注册
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      miniprogramReq.RegisterRequest                                true  "注册请求"
// @Success   200   {object}  response.Response{data=miniprogramRes.RegisterResponse,msg=string}  "注册成功"
// @Router    /miniprogram/auth/register [post]
func (a *AuthApi) Register(c *gin.Context) {
	var registerReq miniprogramReq.RegisterRequest
	err := c.ShouldBindJSON(&registerReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 参数验证
	if registerReq.Phone == "" || registerReq.Password == "" || registerReq.Nickname == "" {
		response.FailWithMessage("手机号、密码和昵称不能为空", c)
		return
	}

	// 创建用户
	user, err := userService.CreateUser(&registerReq)
	if err != nil {
		global.GVA_LOG.Error("用户注册失败", zap.Error(err))
		response.FailWithMessage("注册失败: "+err.Error(), c)
		return
	}

	// 返回注册结果
	registerRes := miniprogramRes.RegisterResponse{
		User: user,
	}

	response.OkWithData(registerRes, c)
}

// RefreshToken 刷新token
// @Tags      MiniProgram
// @Summary   刷新用户token
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      miniprogramReq.RefreshTokenRequest                                true  "刷新token请求"
// @Success   200   {object}  response.Response{data=miniprogramRes.RefreshTokenResponse,msg=string}  "刷新成功"
// @Router    /miniprogram/auth/refresh-token [post]
func (a *AuthApi) RefreshToken(c *gin.Context) {
	var refreshReq miniprogramReq.RefreshTokenRequest
	err := c.ShouldBindJSON(&refreshReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 验证并刷新token
	newToken, err := authService.RefreshToken(refreshReq.Token)
	if err != nil {
		response.FailWithMessage("token刷新失败", c)
		return
	}

	refreshRes := miniprogramRes.RefreshTokenResponse{
		Token: newToken,
	}

	response.OkWithData(refreshRes, c)
}

// ClearAuthFailures 清除认证失败记录（开发测试用）
func (a *AuthApi) ClearAuthFailures(c *gin.Context) {
	clientIP := c.ClientIP()

	if global.GVA_REDIS != nil {
		key := "auth_fail:" + clientIP
		err := global.GVA_REDIS.Del(c, key).Err()
		if err != nil {
			global.GVA_LOG.Error("清除认证失败记录失败", zap.Error(err))
			response.FailWithMessage("清除失败", c)
			return
		}
		global.GVA_LOG.Info("清除认证失败记录成功", zap.String("ip", clientIP))
		response.OkWithMessage("认证失败记录已清除", c)
	} else {
		response.FailWithMessage("Redis未配置", c)
	}
}
