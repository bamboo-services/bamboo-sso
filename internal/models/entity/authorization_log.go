package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// AuthorizationLog 表示SSO授权验证日志实体，记录所有授权码验证行为。
//
// 字段说明：
//   - UUID: 日志记录的唯一标识符，由 UUID 表示。
//   - AuthorizationCodeUUID: 关联的授权码UUID，外键。
//   - ApplicationUUID: 关联的应用UUID，外键。
//   - UserUUID: 关联的用户UUID，外键。
//   - RequestIPAddress: 请求IP地址。
//   - RequestUserAgent: 请求User-Agent字符串。
//   - RequestBrowserFingerprint: 请求浏览器指纹哈希值。
//   - IsSuccess: 验证是否成功。
//   - FailureReason: 验证失败原因。
//   - FingerprintMatched: 浏览器指纹是否匹配。
//   - UserAgentMatched: User-Agent是否匹配。
//   - IPMatched: IP地址是否匹配（可配置是否严格检查）。
//   - VerifiedAt: 验证时间。
//   - CreatedAt: 创建记录的时间戳。
type AuthorizationLog struct {
	UUID                      uuid.UUID  `json:"uuid" gorm:"primaryKey;type:uuid;not null;comment:授权日志唯一标识符"`
	AuthorizationCodeUUID     *uuid.UUID `json:"authorization_code_uuid" gorm:"type:uuid;index;comment:关联授权码UUID"`
	ApplicationUUID           uuid.UUID  `json:"application_uuid" gorm:"type:uuid;not null;index;comment:关联应用UUID"`
	UserUUID                  *uuid.UUID `json:"user_uuid" gorm:"type:uuid;index;comment:关联用户UUID"`
	RequestIPAddress          string     `json:"request_ip_address" gorm:"type:varchar(45);not null;comment:请求IP地址"`
	RequestUserAgent          string     `json:"request_user_agent" gorm:"type:text;not null;comment:请求User-Agent"`
	RequestBrowserFingerprint string     `json:"request_browser_fingerprint" gorm:"type:varchar(128);not null;comment:请求浏览器指纹"`
	IsSuccess                 bool       `json:"is_success" gorm:"type:boolean;not null;comment:验证是否成功"`
	FailureReason             *string    `json:"failure_reason" gorm:"type:varchar(255);comment:失败原因"`
	FingerprintMatched        *bool      `json:"fingerprint_matched" gorm:"type:boolean;comment:指纹是否匹配"`
	UserAgentMatched          *bool      `json:"user_agent_matched" gorm:"type:boolean;comment:User-Agent是否匹配"`
	IPMatched                 *bool      `json:"ip_matched" gorm:"type:boolean;comment:IP是否匹配"`
	VerifiedAt                time.Time  `json:"verified_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:验证时间"`
	CreatedAt                 time.Time  `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`

	// 关联关系
	AuthorizationCode *AuthorizationCode `json:"authorization_code,omitempty" gorm:"foreignKey:AuthorizationCodeUUID;references:UUID;comment:关联授权码"`
	Application       *Application       `json:"application,omitempty" gorm:"foreignKey:ApplicationUUID;references:UUID;comment:关联应用"`
	User              *User              `json:"user,omitempty" gorm:"foreignKey:UserUUID;references:UUID;comment:关联用户"`
}

// BeforeCreate 在创建 AuthorizationLog 记录前自动生成新的 UUID（如果当前 UUID 为空）。
func (al *AuthorizationLog) BeforeCreate(_ *gorm.DB) (err error) {
	if al.UUID == uuid.Nil {
		newUUID, err := uuid.NewV7()
		if err != nil {
			return err
		}
		al.UUID = newUUID
	}
	if al.VerifiedAt.IsZero() {
		al.VerifiedAt = time.Now()
	}
	return
}
