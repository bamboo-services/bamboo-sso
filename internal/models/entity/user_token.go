package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// UserToken 表示用户在 SSO 系统中的访问令牌和刷新令牌，用于维持用户登录状态。
//
// 字段说明：
//   - UUID: 令牌记录的唯一标识符。
//   - UserUUID: 关联的用户唯一标识符。
//   - AccessToken: 访问令牌，用于短期身份验证。
//   - RefreshToken: 刷新令牌，用于获取新的访问令牌。
//   - AccessTokenExpiresAt: 访问令牌过期时间。
//   - RefreshTokenExpiresAt: 刷新令牌过期时间。
//   - DeviceInfo: 设备信息（可选），用于记录登录设备。
//   - IPAddress: 登录时的 IP 地址。
//   - UserAgent: 登录时的用户代理信息。
//   - IsRevoked: 令牌是否已被撤销。
//   - LastUsedAt: 最后使用时间。
//   - CreatedAt: 创建时间。
//   - UpdatedAt: 更新时间。
type UserToken struct {
	UUID                  uuid.UUID  `json:"uuid" gorm:"primaryKey;type:uuid;not null;comment:令牌记录唯一标识符"`
	UserUUID              uuid.UUID  `json:"user_uuid" gorm:"type:uuid;not null;index;comment:用户唯一标识符"`
	AccessToken           string     `json:"access_token" gorm:"type:varchar(255);not null;uniqueIndex;comment:访问令牌"`
	RefreshToken          string     `json:"refresh_token" gorm:"type:varchar(255);not null;uniqueIndex;comment:刷新令牌"`
	AccessTokenExpiresAt  time.Time  `json:"access_token_expires_at" gorm:"type:timestamp;not null;comment:访问令牌过期时间"`
	RefreshTokenExpiresAt time.Time  `json:"refresh_token_expires_at" gorm:"type:timestamp;not null;comment:刷新令牌过期时间"`
	DeviceInfo            *string    `json:"device_info" gorm:"type:varchar(255);comment:设备信息"`
	IPAddress             *string    `json:"ip_address" gorm:"type:varchar(45);comment:登录IP地址"`
	UserAgent             *string    `json:"user_agent" gorm:"type:text;comment:用户代理信息"`
	IsRevoked             bool       `json:"is_revoked" gorm:"type:boolean;not null;default:false;comment:是否已撤销"`
	LastUsedAt            *time.Time `json:"last_used_at" gorm:"type:timestamp;comment:最后使用时间"`
	CreatedAt             time.Time  `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt             time.Time  `json:"updated_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`

	// 关联关系
	User *User `json:"user,omitempty" gorm:"foreignKey:UserUUID;references:UUID;constraint:OnDelete:CASCADE;comment:关联用户"`
}

// BeforeCreate 在创建 UserToken 记录前自动生成新的 UUID（如果当前 UUID 为空）。
func (ut *UserToken) BeforeCreate(_ *gorm.DB) (err error) {
	if ut.UUID == uuid.Nil {
		newUUID, err := uuid.NewV7()
		if err != nil {
			return err
		}
		ut.UUID = newUUID
	}
	return
}

// BeforeUpdate 在更新 UserToken 记录前自动更新 UpdatedAt 字段。
func (ut *UserToken) BeforeUpdate(_ *gorm.DB) (err error) {
	ut.UpdatedAt = time.Now()
	return
}

// IsAccessTokenExpired 检查访问令牌是否已过期。
func (ut *UserToken) IsAccessTokenExpired() bool {
	return time.Now().After(ut.AccessTokenExpiresAt) || ut.IsRevoked
}

// IsRefreshTokenExpired 检查刷新令牌是否已过期。
func (ut *UserToken) IsRefreshTokenExpired() bool {
	return time.Now().After(ut.RefreshTokenExpiresAt) || ut.IsRevoked
}

// IsValid 检查令牌记录是否有效（未撤销且刷新令牌未过期）。
func (ut *UserToken) IsValid() bool {
	return !ut.IsRevoked && !ut.IsRefreshTokenExpired()
}
