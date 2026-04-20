package service

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	miniprogramReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/model/request"
	jwt "github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type AuthService struct{}

// WxSessionResult 微信会话信息
type WxSessionResult struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// JWTClaims JWT声明
type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

const (
	// JWT密钥，生产环境应该从配置文件读取
	jwtSecret = "nofap_miniprogram_jwt_secret_key"
	// Token过期时间（7天）
	tokenExpireDuration = time.Hour * 24 * 7
	// 微信API URL
	wxApiURL = "https://api.weixin.qq.com/sns/jscode2session"
)

// GetWxSessionInfo 获取微信会话信息
func (s *AuthService) GetWxSessionInfo(code string) (*WxSessionResult, error) {
	// 这里应该从配置文件读取AppID和AppSecret
	appID := global.GVA_CONFIG.MiniProgram.AppID
	appSecret := global.GVA_CONFIG.MiniProgram.AppSecret

	// 临时使用测试配置，生产环境需要在配置文件中设置
	if appID == "" {
		appID = "test_app_id" // 需要在配置文件中设置真实的AppID
	}
	if appSecret == "" {
		appSecret = "test_app_secret" // 需要在配置文件中设置真实的AppSecret
	}

	if appID == "test_app_id" || appSecret == "test_app_secret" {
		global.GVA_LOG.Warn("使用测试微信小程序配置，请在配置文件中设置真实的AppID和AppSecret")
	}

	url := fmt.Sprintf("%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		wxApiURL, appID, appSecret, code)

	resp, err := http.Get(url)
	if err != nil {
		global.GVA_LOG.Error("微信API请求失败", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("读取微信API响应失败", zap.Error(err))
		return nil, err
	}

	var result WxSessionResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		global.GVA_LOG.Error("解析微信API响应失败", zap.Error(err))
		return nil, err
	}

	if result.ErrCode != 0 {
		global.GVA_LOG.Error("微信API返回错误",
			zap.Int("错误码", result.ErrCode),
			zap.String("错误信息", result.ErrMsg),
		)
		return nil, fmt.Errorf("微信API错误: %s", result.ErrMsg)
	}

	return &result, nil
}

// DecryptWxUserInfo 解密微信用户信息
func (s *AuthService) DecryptWxUserInfo(encryptedData, iv, sessionKey string) (*miniprogramReq.WxUserInfo, error) {
	// Base64解码
	sessionKeyBytes, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}

	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}

	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}

	// AES解密
	block, err := aes.NewCipher(sessionKeyBytes)
	if err != nil {
		return nil, err
	}

	if len(encryptedBytes) < aes.BlockSize {
		return nil, errors.New("密文长度不足")
	}

	mode := cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(encryptedBytes, encryptedBytes)

	// 去除PKCS7填充
	encryptedBytes = removePKCS7Padding(encryptedBytes)

	// 解析JSON
	var userInfo miniprogramReq.WxUserInfo
	err = json.Unmarshal(encryptedBytes, &userInfo)
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}

// removePKCS7Padding 去除PKCS7填充
func removePKCS7Padding(data []byte) []byte {
	length := len(data)
	if length == 0 {
		return data
	}
	unpadding := int(data[length-1])
	if unpadding > length {
		return data
	}
	return data[:(length - unpadding)]
}

// GenerateToken 生成JWT token
func (s *AuthService) GenerateToken(userID uint) (string, error) {
	claims := &JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "nofap-miniprogram",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

// ValidateToken 验证JWT token
func (s *AuthService) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的token")
}

// RefreshToken 刷新token
func (s *AuthService) RefreshToken(tokenString string) (string, error) {
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	// 检查token是否即将过期（剩余时间少于1天）
	if claims.ExpiresAt.Time.Sub(time.Now()) > time.Hour*24 {
		return tokenString, nil // token还很新，无需刷新
	}

	// 生成新token
	return s.GenerateToken(claims.UserID)
}

// GetUserIDFromToken 从token中获取用户ID
func (s *AuthService) GetUserIDFromToken(tokenString string) (uint, error) {
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
}
