package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// LoginLog 表示用户登录日志实体，记录所有登录行为。
//
// 字段说明：
//   - UUID: 日志记录的唯一标识符，由 UUID 表示。
//   - UserUUID: 关联的用户UUID，外键（可为空，记录登录失败的情况）。
//   - LoginType: 登录类型（password-密码登录，third_party-第三方登录）。
//   - ProviderUUID: 第三方提供商UUID，外键（第三方登录时使用）。
//   - IPAddress: 登录IP地址。
//   - UserAgent: 用户浏览器User-Agent字符串。
//   - BrowserFingerprint: 浏览器指纹哈希值。
//   - IsSuccess: 登录是否成���。
//   - FailureReason: 登录失败原因。
//   - LoginAt: 登录时间。
//   - CreatedAt: 创建记录的时间戳。
type LoginLog struct {
	UUID               uuid.UUID  `json:"uuid" gorm:"primaryKey;type:uuid;not null;comment:登录日志唯一标识符"`
	UserUUID           *uuid.UUID `json:"user_uuid" gorm:"type:uuid;index;comment:关联用户UUID(可为空)"`
	LoginType          string     `json:"login_type" gorm:"type:varchar(20);not null;comment:登录类型(password/third_party)"`
	ProviderUUID       *uuid.UUID `json:"provider_uuid" gorm:"type:uuid;comment:第三方提供商UUID(第三方登录时)"`
	IPAddress          string     `json:"ip_address" gorm:"type:varchar(45);not null;comment:登录IP地址"`
	UserAgent          string     `json:"user_agent" gorm:"type:text;not null;comment:用户浏览器User-Agent"`
	BrowserFingerprint string     `json:"browser_fingerprint" gorm:"type:varchar(128);comment:浏览器指纹哈希"`
	IsSuccess          bool       `json:"is_success" gorm:"type:boolean;not null;comment:是否成功"`
	FailureReason      string     `json:"failure_reason" gorm:"type:varchar(255);comment:失败原因"`
	LoginAt            time.Time  `json:"login_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:登录时间"`
	CreatedAt          time.Time  `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`

	// 关联关系
	User     *User               `json:"user,omitempty" gorm:"foreignKey:UserUUID;references:UUID;comment:关联用户"`
	Provider *ThirdPartyProvider `json:"provider,omitempty" gorm:"foreignKey:ProviderUUID;references:UUID;comment:关联第三方提供商"`
}

// BeforeCreate 在创建 LoginLog 记录前自动生成新的 UUID（如果当前 UUID 为空）。
func (ll *LoginLog) BeforeCreate(tx *gorm.DB) (err error) {
	if ll.UUID == uuid.Nil {
		newUUID, err := uuid.NewV7()
		if err != nil {
			return err
		}
		ll.UUID = newUUID
	}
	if ll.LoginAt.IsZero() {
		ll.LoginAt = time.Now()
	}
	return
}
