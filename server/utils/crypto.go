package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// CryptoConfig 加密配置
type CryptoConfig struct {
	SecretKey string // 32字节的密钥
	Salt      string // 盐值
}

// DefaultCryptoConfig 默认加密配置
var DefaultCryptoConfig = CryptoConfig{
	SecretKey: "nofap-miniprogram-secret-key-32", // 32字节
	Salt:      "nofap-salt-2025",
}

// AESEncrypt AES加密
func AESEncrypt(plaintext string, config ...CryptoConfig) (string, error) {
	cfg := DefaultCryptoConfig
	if len(config) > 0 {
		cfg = config[0]
	}

	// 确保密钥长度为32字节
	key := make([]byte, 32)
	copy(key, []byte(cfg.SecretKey))

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 使用GCM模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 生成随机nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// 加密
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	// 返回base64编码的结果
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AESDecrypt AES解密
func AESDecrypt(ciphertext string, config ...CryptoConfig) (string, error) {
	cfg := DefaultCryptoConfig
	if len(config) > 0 {
		cfg = config[0]
	}

	// 解码base64
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// 确保密钥长度为32字节
	key := make([]byte, 32)
	copy(key, []byte(cfg.SecretKey))

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("密文太短")
	}

	nonce, ciphertext_bytes := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext_bytes, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// HashPassword 密码哈希（使用bcrypt）
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword 验证密码
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// MaskSensitiveData 敏感数据脱敏
func MaskSensitiveData(data string, dataType string) string {
	switch dataType {
	case "phone":
		return maskPhone(data)
	case "email":
		return maskEmail(data)
	case "idcard":
		return maskIDCard(data)
	case "name":
		return maskName(data)
	case "address":
		return maskAddress(data)
	default:
		return maskDefault(data)
	}
}

// maskPhone 手机号脱敏
func maskPhone(phone string) string {
	if len(phone) != 11 {
		return phone
	}
	return phone[:3] + "****" + phone[7:]
}

// maskEmail 邮箱脱敏
func maskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return email
	}

	username := parts[0]
	domain := parts[1]

	if len(username) <= 3 {
		return username + "@" + domain
	}

	return username[:2] + "***" + username[len(username)-1:] + "@" + domain
}

// maskIDCard 身份证脱敏
func maskIDCard(idcard string) string {
	if len(idcard) != 18 {
		return idcard
	}
	return idcard[:6] + "********" + idcard[14:]
}

// maskName 姓名脱敏
func maskName(name string) string {
	runes := []rune(name)
	if len(runes) <= 1 {
		return name
	}
	if len(runes) == 2 {
		return string(runes[0]) + "*"
	}
	return string(runes[0]) + strings.Repeat("*", len(runes)-2) + string(runes[len(runes)-1])
}

// maskAddress 地址脱敏
func maskAddress(address string) string {
	if len(address) <= 10 {
		return address[:3] + "***"
	}
	return address[:6] + "***" + address[len(address)-4:]
}

// maskDefault 默认脱敏
func maskDefault(data string) string {
	if len(data) <= 4 {
		return strings.Repeat("*", len(data))
	}
	return data[:2] + strings.Repeat("*", len(data)-4) + data[len(data)-2:]
}

// SHA256Hash SHA256哈希
func SHA256Hash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// SanitizeInput 输入清理
func SanitizeInput(input string) string {
	// 移除危险字符
	input = strings.ReplaceAll(input, "<", "&lt;")
	input = strings.ReplaceAll(input, ">", "&gt;")
	input = strings.ReplaceAll(input, "\"", "&quot;")
	input = strings.ReplaceAll(input, "'", "&#x27;")
	input = strings.ReplaceAll(input, "&", "&amp;")

	// 移除控制字符
	re := regexp.MustCompile(`[\x00-\x1f\x7f-\x9f]`)
	input = re.ReplaceAllString(input, "")

	return strings.TrimSpace(input)
}

// IsValidInput 验证输入是否安全
func IsValidInput(input string, maxLength int) bool {
	// 长度检查
	if len(input) > maxLength {
		return false
	}

	// 检查是否包含危险模式
	dangerousPatterns := []string{
		`<script`,
		`javascript:`,
		`vbscript:`,
		`onload=`,
		`onerror=`,
		`onclick=`,
		`<iframe`,
		`<object`,
		`<embed`,
		`union.*select`,
		`drop.*table`,
		`delete.*from`,
		`insert.*into`,
	}

	lowerInput := strings.ToLower(input)
	for _, pattern := range dangerousPatterns {
		if matched, _ := regexp.MatchString(pattern, lowerInput); matched {
			return false
		}
	}

	return true
}
