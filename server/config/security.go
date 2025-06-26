package config

import "time"

// Security 安全配置
type Security struct {
	// 加密配置
	Encryption EncryptionConfig `mapstructure:"encryption" json:"encryption" yaml:"encryption"`

	// 速率限制配置
	RateLimit RateLimitConfig `mapstructure:"rate-limit" json:"rate-limit" yaml:"rate-limit"`

	// CORS配置
	CORS CORSConfig `mapstructure:"cors" json:"cors" yaml:"cors"`

	// JWT配置
	JWT JWTSecurityConfig `mapstructure:"jwt" json:"jwt" yaml:"jwt"`

	// 输入验证配置
	InputValidation InputValidationConfig `mapstructure:"input-validation" json:"input-validation" yaml:"input-validation"`

	// 审计日志配置
	AuditLog AuditLogConfig `mapstructure:"audit-log" json:"audit-log" yaml:"audit-log"`

	// 敏感数据配置
	SensitiveData SensitiveDataConfig `mapstructure:"sensitive-data" json:"sensitive-data" yaml:"sensitive-data"`
}

// EncryptionConfig 加密配置
type EncryptionConfig struct {
	SecretKey string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`
	Salt      string `mapstructure:"salt" json:"salt" yaml:"salt"`
	Algorithm string `mapstructure:"algorithm" json:"algorithm" yaml:"algorithm"` // AES-256-GCM
}

// RateLimitConfig 速率限制配置
type RateLimitConfig struct {
	Enabled     bool          `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	WindowSize  time.Duration `mapstructure:"window-size" json:"window-size" yaml:"window-size"`
	MaxRequests int           `mapstructure:"max-requests" json:"max-requests" yaml:"max-requests"`

	// API特定限制
	AuthAPI struct {
		WindowSize  time.Duration `mapstructure:"window-size" json:"window-size" yaml:"window-size"`
		MaxRequests int           `mapstructure:"max-requests" json:"max-requests" yaml:"max-requests"`
	} `mapstructure:"auth-api" json:"auth-api" yaml:"auth-api"`

	// 小程序API限制
	MiniprogramAPI struct {
		WindowSize  time.Duration `mapstructure:"window-size" json:"window-size" yaml:"window-size"`
		MaxRequests int           `mapstructure:"max-requests" json:"max-requests" yaml:"max-requests"`
	} `mapstructure:"miniprogram-api" json:"miniprogram-api" yaml:"miniprogram-api"`
}

// CORSConfig CORS配置
type CORSConfig struct {
	AllowedOrigins   []string `mapstructure:"allowed-origins" json:"allowed-origins" yaml:"allowed-origins"`
	AllowedMethods   []string `mapstructure:"allowed-methods" json:"allowed-methods" yaml:"allowed-methods"`
	AllowedHeaders   []string `mapstructure:"allowed-headers" json:"allowed-headers" yaml:"allowed-headers"`
	ExposedHeaders   []string `mapstructure:"exposed-headers" json:"exposed-headers" yaml:"exposed-headers"`
	AllowCredentials bool     `mapstructure:"allow-credentials" json:"allow-credentials" yaml:"allow-credentials"`
	MaxAge           int      `mapstructure:"max-age" json:"max-age" yaml:"max-age"`
}

// JWTSecurityConfig JWT安全配置
type JWTSecurityConfig struct {
	BlacklistEnabled bool          `mapstructure:"blacklist-enabled" json:"blacklist-enabled" yaml:"blacklist-enabled"`
	RefreshThreshold time.Duration `mapstructure:"refresh-threshold" json:"refresh-threshold" yaml:"refresh-threshold"`
	MaxFailAttempts  int           `mapstructure:"max-fail-attempts" json:"max-fail-attempts" yaml:"max-fail-attempts"`
	LockoutDuration  time.Duration `mapstructure:"lockout-duration" json:"lockout-duration" yaml:"lockout-duration"`
}

// InputValidationConfig 输入验证配置
type InputValidationConfig struct {
	MaxRequestSize   int64    `mapstructure:"max-request-size" json:"max-request-size" yaml:"max-request-size"`
	AllowedFileTypes []string `mapstructure:"allowed-file-types" json:"allowed-file-types" yaml:"allowed-file-types"`
	MaxFileSize      int64    `mapstructure:"max-file-size" json:"max-file-size" yaml:"max-file-size"`

	// 字段长度限制
	FieldLimits map[string]int `mapstructure:"field-limits" json:"field-limits" yaml:"field-limits"`

	// XSS防护
	XSSProtection struct {
		Enabled           bool     `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
		StrictMode        bool     `mapstructure:"strict-mode" json:"strict-mode" yaml:"strict-mode"`
		AllowedTags       []string `mapstructure:"allowed-tags" json:"allowed-tags" yaml:"allowed-tags"`
		AllowedAttributes []string `mapstructure:"allowed-attributes" json:"allowed-attributes" yaml:"allowed-attributes"`
	} `mapstructure:"xss-protection" json:"xss-protection" yaml:"xss-protection"`

	// SQL注入防护
	SQLInjectionProtection struct {
		Enabled    bool     `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
		StrictMode bool     `mapstructure:"strict-mode" json:"strict-mode" yaml:"strict-mode"`
		Patterns   []string `mapstructure:"patterns" json:"patterns" yaml:"patterns"`
	} `mapstructure:"sql-injection-protection" json:"sql-injection-protection" yaml:"sql-injection-protection"`
}

// AuditLogConfig 审计日志配置
type AuditLogConfig struct {
	Enabled    bool   `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	LogLevel   string `mapstructure:"log-level" json:"log-level" yaml:"log-level"`
	LogPath    string `mapstructure:"log-path" json:"log-path" yaml:"log-path"`
	MaxSize    int    `mapstructure:"max-size" json:"max-size" yaml:"max-size"` // MB
	MaxBackups int    `mapstructure:"max-backups" json:"max-backups" yaml:"max-backups"`
	MaxAge     int    `mapstructure:"max-age" json:"max-age" yaml:"max-age"` // days
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`

	// 记录的事件类型
	Events struct {
		Login         bool `mapstructure:"login" json:"login" yaml:"login"`
		Logout        bool `mapstructure:"logout" json:"logout" yaml:"logout"`
		DataAccess    bool `mapstructure:"data-access" json:"data-access" yaml:"data-access"`
		DataModify    bool `mapstructure:"data-modify" json:"data-modify" yaml:"data-modify"`
		AdminAction   bool `mapstructure:"admin-action" json:"admin-action" yaml:"admin-action"`
		SecurityEvent bool `mapstructure:"security-event" json:"security-event" yaml:"security-event"`
	} `mapstructure:"events" json:"events" yaml:"events"`
}

// SensitiveDataConfig 敏感数据配置
type SensitiveDataConfig struct {
	// 需要加密的字段
	EncryptedFields []string `mapstructure:"encrypted-fields" json:"encrypted-fields" yaml:"encrypted-fields"`

	// 需要脱敏的字段
	MaskedFields map[string]string `mapstructure:"masked-fields" json:"masked-fields" yaml:"masked-fields"`

	// 数据保留策略
	DataRetention struct {
		UserData      time.Duration `mapstructure:"user-data" json:"user-data" yaml:"user-data"`
		LogData       time.Duration `mapstructure:"log-data" json:"log-data" yaml:"log-data"`
		TemporaryData time.Duration `mapstructure:"temporary-data" json:"temporary-data" yaml:"temporary-data"`
	} `mapstructure:"data-retention" json:"data-retention" yaml:"data-retention"`

	// 数据导出限制
	ExportLimits struct {
		MaxRecords    int           `mapstructure:"max-records" json:"max-records" yaml:"max-records"`
		MaxSize       int64         `mapstructure:"max-size" json:"max-size" yaml:"max-size"` // bytes
		RateLimit     time.Duration `mapstructure:"rate-limit" json:"rate-limit" yaml:"rate-limit"`
		RequiredRoles []string      `mapstructure:"required-roles" json:"required-roles" yaml:"required-roles"`
	} `mapstructure:"export-limits" json:"export-limits" yaml:"export-limits"`
}

// GetDefaultSecurityConfig 获取默认安全配置
func GetDefaultSecurityConfig() Security {
	return Security{
		Encryption: EncryptionConfig{
			SecretKey: "nofap-miniprogram-secret-key-32",
			Salt:      "nofap-salt-2025",
			Algorithm: "AES-256-GCM",
		},
		RateLimit: RateLimitConfig{
			Enabled:     true,
			WindowSize:  time.Minute,
			MaxRequests: 100,
			AuthAPI: struct {
				WindowSize  time.Duration `mapstructure:"window-size" json:"window-size" yaml:"window-size"`
				MaxRequests int           `mapstructure:"max-requests" json:"max-requests" yaml:"max-requests"`
			}{
				WindowSize:  time.Minute,
				MaxRequests: 10,
			},
			MiniprogramAPI: struct {
				WindowSize  time.Duration `mapstructure:"window-size" json:"window-size" yaml:"window-size"`
				MaxRequests int           `mapstructure:"max-requests" json:"max-requests" yaml:"max-requests"`
			}{
				WindowSize:  time.Minute,
				MaxRequests: 60,
			},
		},
		CORS: CORSConfig{
			AllowedOrigins:   []string{"https://servicewechat.com"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Token", "X-User-Id"},
			ExposedHeaders:   []string{"Content-Length", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           86400,
		},
		JWT: JWTSecurityConfig{
			BlacklistEnabled: true,
			RefreshThreshold: 24 * time.Hour,
			MaxFailAttempts:  10,
			LockoutDuration:  time.Hour,
		},
		InputValidation: InputValidationConfig{
			MaxRequestSize:   10 * 1024 * 1024, // 10MB
			AllowedFileTypes: []string{".jpg", ".jpeg", ".png", ".gif"},
			MaxFileSize:      5 * 1024 * 1024, // 5MB
			FieldLimits: map[string]int{
				"username":    50,
				"nickname":    100,
				"content":     2000,
				"description": 500,
				"title":       200,
			},
			XSSProtection: struct {
				Enabled           bool     `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
				StrictMode        bool     `mapstructure:"strict-mode" json:"strict-mode" yaml:"strict-mode"`
				AllowedTags       []string `mapstructure:"allowed-tags" json:"allowed-tags" yaml:"allowed-tags"`
				AllowedAttributes []string `mapstructure:"allowed-attributes" json:"allowed-attributes" yaml:"allowed-attributes"`
			}{
				Enabled:           true,
				StrictMode:        true,
				AllowedTags:       []string{"p", "br", "strong", "em"},
				AllowedAttributes: []string{"class"},
			},
			SQLInjectionProtection: struct {
				Enabled    bool     `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
				StrictMode bool     `mapstructure:"strict-mode" json:"strict-mode" yaml:"strict-mode"`
				Patterns   []string `mapstructure:"patterns" json:"patterns" yaml:"patterns"`
			}{
				Enabled:    true,
				StrictMode: true,
				Patterns: []string{
					"union.*select",
					"drop.*table",
					"delete.*from",
					"insert.*into",
					"update.*set",
				},
			},
		},
		AuditLog: AuditLogConfig{
			Enabled:    true,
			LogLevel:   "info",
			LogPath:    "./log/audit.log",
			MaxSize:    100,
			MaxBackups: 30,
			MaxAge:     90,
			Compress:   true,
			Events: struct {
				Login         bool `mapstructure:"login" json:"login" yaml:"login"`
				Logout        bool `mapstructure:"logout" json:"logout" yaml:"logout"`
				DataAccess    bool `mapstructure:"data-access" json:"data-access" yaml:"data-access"`
				DataModify    bool `mapstructure:"data-modify" json:"data-modify" yaml:"data-modify"`
				AdminAction   bool `mapstructure:"admin-action" json:"admin-action" yaml:"admin-action"`
				SecurityEvent bool `mapstructure:"security-event" json:"security-event" yaml:"security-event"`
			}{
				Login:         true,
				Logout:        true,
				DataAccess:    false,
				DataModify:    true,
				AdminAction:   true,
				SecurityEvent: true,
			},
		},
		SensitiveData: SensitiveDataConfig{
			EncryptedFields: []string{"phone", "email", "real_name", "id_card", "address"},
			MaskedFields: map[string]string{
				"phone":     "phone",
				"email":     "email",
				"real_name": "name",
				"id_card":   "idcard",
				"address":   "address",
			},
			DataRetention: struct {
				UserData      time.Duration `mapstructure:"user-data" json:"user-data" yaml:"user-data"`
				LogData       time.Duration `mapstructure:"log-data" json:"log-data" yaml:"log-data"`
				TemporaryData time.Duration `mapstructure:"temporary-data" json:"temporary-data" yaml:"temporary-data"`
			}{
				UserData:      365 * 24 * time.Hour, // 1年
				LogData:       90 * 24 * time.Hour,  // 90天
				TemporaryData: 24 * time.Hour,       // 1天
			},
			ExportLimits: struct {
				MaxRecords    int           `mapstructure:"max-records" json:"max-records" yaml:"max-records"`
				MaxSize       int64         `mapstructure:"max-size" json:"max-size" yaml:"max-size"`
				RateLimit     time.Duration `mapstructure:"rate-limit" json:"rate-limit" yaml:"rate-limit"`
				RequiredRoles []string      `mapstructure:"required-roles" json:"required-roles" yaml:"required-roles"`
			}{
				MaxRecords:    10000,
				MaxSize:       100 * 1024 * 1024, // 100MB
				RateLimit:     time.Hour,
				RequiredRoles: []string{"admin", "data_export"},
			},
		},
	}
}
