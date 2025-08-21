package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// AuthorizationCode 表示SSO授权码实体，这是你的核心设计实现。
//
// 设计特点：
//   - 几小时有效期的授权码（非一次性）
//   - 绑定用户信息、应用信息、浏览器指纹等安全信息
//   - 支持多重验证：User-Agent、浏览器指纹、IP地址
//   - 防止授权码被恶意使用
//
// 字段说明：
//   - UUID: 授权码记录的唯一标识符，由 UUID 表示。
//   - Code: 授权码值，分发给客户端的实际码值。
//   - UserUUID: 关联的用户UUID，外键。
//   - ApplicationUUID: 关联的应用UUID，外键。
//   - UserAgent: 用户浏览器User-Agent字符串。
//   - BrowserFingerprint: 浏览器指纹哈希值。
//   - IPAddress: 用户IP地址。
//   - ExpiresAt: 授权码过期时间（几小时后）。
//   - IsActive: 授权码是否有效。
//   - UsageCount: 使用次数统计。
//   - LastUsedAt: 最后使用时间。
//   - CreatedAt: 创建记录的时间戳。
//   - UpdatedAt: 最后更新时间戳。
type AuthorizationCode struct {
	UUID               uuid.UUID  `json:"uuid" gorm:"primaryKey;type:uuid;not null;comment:授权码记录唯一标识符"`
	Code               string     `json:"code" gorm:"type:varchar(128);not null;uniqueIndex;comment:授权码值"`
	UserUUID           uuid.UUID  `json:"user_uuid" gorm:"type:uuid;not null;index;comment:关联用户UUID"`
	ApplicationUUID    uuid.UUID  `json:"application_uuid" gorm:"type:uuid;not null;index;comment:关联应用UUID"`
	UserAgent          string     `json:"user_agent" gorm:"type:text;not null;comment:用户浏览器User-Agent"`
	BrowserFingerprint string     `json:"browser_fingerprint" gorm:"type:varchar(128);not null;comment:浏览器指纹哈希"`
	IPAddress          string     `json:"ip_address" gorm:"type:varchar(45);not null;comment:用户IP地址"`
	ExpiresAt          time.Time  `json:"expires_at" gorm:"type:timestamp;not null;comment:过期时间"`
	IsActive           bool       `json:"is_active" gorm:"type:boolean;not null;default:true;comment:是否有效"`
	UsageCount         int        `json:"usage_count" gorm:"type:integer;not null;default:0;comment:使用次数"`
	LastUsedAt         *time.Time `json:"last_used_at" gorm:"type:timestamp;comment:最后使用时间"`
	CreatedAt          time.Time  `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt          time.Time  `json:"updated_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`

	// 关联关系
	User        *User        `json:"user,omitempty" gorm:"foreignKey:UserUUID;references:UUID;constraint:OnDelete:CASCADE;comment:关联用户"`
	Application *Application `json:"application,omitempty" gorm:"foreignKey:ApplicationUUID;references:UUID;constraint:OnDelete:CASCADE;comment:关联应用"`
}

// BeforeCreate 在创建 AuthorizationCode 记录前自动生成新的 UUID（如果当前 UUID 为空）。
func (ac *AuthorizationCode) BeforeCreate(tx *gorm.DB) (err error) {
	if ac.UUID == uuid.Nil {
		newUUID, err := uuid.NewV7()
		if err != nil {
			return err
		}
		ac.UUID = newUUID
	}
	return
}

// BeforeUpdate 在更新 AuthorizationCode 记录前自动更新 UpdatedAt 字段。
func (ac *AuthorizationCode) BeforeUpdate(tx *gorm.DB) (err error) {
	ac.UpdatedAt = time.Now()
	return
}

// IsExpired 检查授权码是否已过期
func (ac *AuthorizationCode) IsExpired() bool {
	return time.Now().After(ac.ExpiresAt)
}

// IsValid 检查授权码是否有效（未过期且激活）
func (ac *AuthorizationCode) IsValid() bool {
	return ac.IsActive && !ac.IsExpired()
}

// IncrementUsage 增加使用次数并更新最后使用时间
func (ac *AuthorizationCode) IncrementUsage() {
	ac.UsageCount++
	now := time.Now()
	ac.LastUsedAt = &now
}
