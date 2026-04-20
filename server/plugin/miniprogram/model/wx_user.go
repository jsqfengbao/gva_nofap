package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// WxUser 微信用户表
type WxUser struct {
	global.GVA_MODEL
	Openid       string    `json:"openid" gorm:"uniqueIndex;size:100;comment:微信openid"`
	Unionid      string    `json:"unionid" gorm:"size:100;comment:微信unionid"`
	Phone        string    `json:"phone" gorm:"uniqueIndex;size:11;comment:手机号"`
	Password     string    `json:"-" gorm:"size:255;comment:密码"`
	Nickname     string    `json:"nickname" gorm:"size:50;comment:用户昵称"`
	AvatarUrl    string    `json:"avatarUrl" gorm:"type:text;comment:头像URL"`
	Gender       int       `json:"gender" gorm:"default:0;comment:性别:0未知,1男,2女"`
	City         string    `json:"city" gorm:"size:50;comment:城市"`
	Province     string    `json:"province" gorm:"size:50;comment:省份"`
	Country      string    `json:"country" gorm:"size:50;comment:国家"`
	PrivacyLevel int       `json:"privacyLevel" gorm:"default:1;comment:隐私级别:1低,2中,3高"`
	Status       int       `json:"status" gorm:"default:1;comment:状态:0禁用,1正常"`
	LastLoginAt  time.Time `json:"lastLoginAt" gorm:"comment:最后登录时间"`
}

// TableName 设置表名
func (WxUser) TableName() string {
	return "nofap_wx_users"
}
