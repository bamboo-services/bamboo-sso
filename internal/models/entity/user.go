package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// User 表示系统中的用户实体，存储用户基础信息。
//
// 字段说明：
//   - UUID: 用户的唯一标识符，由 UUID 表示。
//   - Username: 用户名，必须唯一。
//   - Email: 邮箱地址，必须唯一，可用于登录。
//   - Phone: 手机号，可选字段。
//   - PasswordHash: 加密后的密码哈希值。
//   - IsActive: 用户是否激活，默认为 true。
//   - LastLoginAt: 最后登录时间。
//   - CreatedAt: 创建记录的时间戳。
//   - UpdatedAt: 最后更新时间戳。
type User struct {
	UUID         uuid.UUID  `json:"uuid" gorm:"primaryKey;type:uuid;not null;comment:用户唯一标识符"`
	Username     string     `json:"username" gorm:"type:varchar(50);not null;uniqueIndex;comment:用户名"`
	Email        string     `json:"email" gorm:"type:varchar(100);not null;uniqueIndex;comment:邮箱地址"`
	Phone        *string    `json:"phone" gorm:"type:varchar(20);comment:手机号"`
	PasswordHash string     `json:"-" gorm:"type:varchar(255);not null;comment:密码哈希值"`
	IsActive     bool       `json:"is_active" gorm:"type:boolean;not null;default:true;comment:是否激活"`
	LastLoginAt  *time.Time `json:"last_login_at" gorm:"type:timestamp;comment:最后登录时间"`
	CreatedAt    time.Time  `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`

	// 关联关系
	Profile            *UserProfile            `json:"profile,omitempty" gorm:"foreignKey:UserUUID;references:UUID;constraint:OnDelete:CASCADE;comment:用户详细资料"`
	Roles              []*Role                 `json:"roles,omitempty" gorm:"many2many:user_roles;comment:用户角色"`
	WechatAccounts     []*UserThirdPartyWechat `json:"wechat_accounts,omitempty" gorm:"foreignKey:UserUUID;references:UUID;constraint:OnDelete:CASCADE;comment:微信账号"`
	GithubAccounts     []*UserThirdPartyGithub `json:"github_accounts,omitempty" gorm:"foreignKey:UserUUID;references:UUID;constraint:OnDelete:CASCADE;comment:Github账号"`
	QQAccounts         []*UserThirdPartyQQ     `json:"qq_accounts,omitempty" gorm:"foreignKey:UserUUID;references:UUID;constraint:OnDelete:CASCADE;comment:QQ账号"`
	AuthorizationCodes []*AuthorizationCode    `json:"authorization_codes,omitempty" gorm:"foreignKey:UserUUID;references:UUID;constraint:OnDelete:CASCADE;comment:授权码"`
}

// BeforeCreate 在创建 User 记录前自动生成新的 UUID（如果当前 UUID 为空）。
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	if u.UUID == uuid.Nil {
		newUUID, err := uuid.NewV7()
		if err != nil {
			return err
		}
		u.UUID = newUUID
	}
	return
}

// BeforeUpdate 在更新 User 记录前自动更新 UpdatedAt 字段。
func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
