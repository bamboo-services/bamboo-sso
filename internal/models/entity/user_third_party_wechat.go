package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// UserThirdPartyWechatEntity 表示用户微信账号绑定实体，存储用户与微信平台账号的绑定关系。
//
// 微信登录特有字段：
//   - UnionID: 微信开放平台唯一标识，用于统一用户身份
//   - OpenID: 微信公众平台唯一标识，每个公众号下唯一
//   - SessionKey: 小程序会话密钥（如果是小程序登录）
//   - City/Province/Country: 微信用户地理位置信息
//   - Language: 用户语言
//   - Subscribe: 是否关注公众号
//   - SubscribeTime: 关注公众号时间
//
// 字段说明：
//   - UUID: 绑定记录的唯一标识符，由 UUID 表示。
//   - UserUUID: 关联的用户UUID，外键。
//   - ProviderUUID: 关联的微信提供商UUID，外键。
//   - UnionID: 微信开放平台唯一标识。
//   - OpenID: 微信公众平台唯一标识。
//   - SessionKey: 小程序会话密钥（加密存储）。
//   - Nickname: 微信用户昵称。
//   - Avatar: 微信用户头像。
//   - Gender: 性别（1-男性，2-女性，0-未知）。
//   - City: 城市。
//   - Province: 省份。
//   - Country: 国家。
//   - Language: 用户语言。
//   - Subscribe: 是否关注公众号。
//   - SubscribeTime: 关注公众号时间。
//   - AccessToken: 微信访问令牌（加密存储）。
//   - RefreshToken: 微信刷新令牌（加密存储）。
//   - TokenExpiresAt: 访问令牌过期时间。
//   - IsActive: 绑定是否激活，默认为 true。
//   - FirstBindAt: 首次绑定时间。
//   - LastLoginAt: 最后一次通过微信登录的时间。
//   - CreatedAt: 创建记录的时间戳。
//   - UpdatedAt: 最后更新时间戳。
type UserThirdPartyWechatEntity struct {
	UUID           uuid.UUID  `json:"uuid" gorm:"primaryKey;type:uuid;not null;comment:微信绑定记录唯一标识符"`
	UserUUID       uuid.UUID  `json:"user_uuid" gorm:"type:uuid;not null;index;comment:关联用户UUID"`
	ProviderUUID   uuid.UUID  `json:"provider_uuid" gorm:"type:uuid;not null;index;comment:关联微信提供商UUID"`
	UnionID        string     `json:"union_id" gorm:"type:varchar(100);uniqueIndex;comment:微信开放平台唯一标识"`
	OpenID         string     `json:"open_id" gorm:"type:varchar(100);not null;index;comment:微信公众平台唯一标识"`
	SessionKey     string     `json:"-" gorm:"type:varchar(255);comment:小程序会话密钥(加密)"`
	Nickname       string     `json:"nickname" gorm:"type:varchar(100);comment:微信用户昵称"`
	Avatar         string     `json:"avatar" gorm:"type:varchar(500);comment:微信用户头像"`
	Gender         int        `json:"gender" gorm:"type:smallint;default:0;comment:性别(0-未知,1-男,2-女)"`
	City           string     `json:"city" gorm:"type:varchar(50);comment:城市"`
	Province       string     `json:"province" gorm:"type:varchar(50);comment:省份"`
	Country        string     `json:"country" gorm:"type:varchar(50);comment:国家"`
	Language       string     `json:"language" gorm:"type:varchar(10);default:'zh_CN';comment:用户语言"`
	Subscribe      bool       `json:"subscribe" gorm:"type:boolean;default:false;comment:是否关注公众号"`
	SubscribeTime  *time.Time `json:"subscribe_time" gorm:"type:timestamp;comment:关注公众号时间"`
	AccessToken    string     `json:"-" gorm:"type:text;comment:微信访问令牌(加密)"`
	RefreshToken   string     `json:"-" gorm:"type:text;comment:微信刷新令牌(加密)"`
	TokenExpiresAt *time.Time `json:"token_expires_at" gorm:"type:timestamp;comment:访问令牌过期时间"`
	IsActive       bool       `json:"is_active" gorm:"type:boolean;not null;default:true;comment:是否激活"`
	FirstBindAt    time.Time  `json:"first_bind_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:首次绑定时间"`
	LastLoginAt    *time.Time `json:"last_login_at" gorm:"type:timestamp;comment:最后登录时间"`
	CreatedAt      time.Time  `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt      time.Time  `json:"updated_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`

	// 关联关系
	User     *UserEntity               `json:"user,omitempty" gorm:"foreignKey:UserUUID;references:UUID;constraint:OnDelete:CASCADE;comment:关联用户"`
	Provider *ThirdPartyProviderEntity `json:"provider,omitempty" gorm:"foreignKey:ProviderUUID;references:UUID;constraint:OnDelete:CASCADE;comment:关联微信提供商"`
}

// BeforeCreate 在创建 UserThirdPartyWechatEntity 记录前自动生成新的 UUID（如果当前 UUID 为空）。
func (utpw *UserThirdPartyWechatEntity) BeforeCreate(tx *gorm.DB) (err error) {
	if utpw.UUID == uuid.Nil {
		newUUID, err := uuid.NewV7()
		if err != nil {
			return err
		}
		utpw.UUID = newUUID
	}
	if utpw.FirstBindAt.IsZero() {
		utpw.FirstBindAt = time.Now()
	}
	return
}

// BeforeUpdate 在更新 UserThirdPartyWechatEntity 记录前自动更新 UpdatedAt 字段。
func (utpw *UserThirdPartyWechatEntity) BeforeUpdate(tx *gorm.DB) (err error) {
	utpw.UpdatedAt = time.Now()
	return
}

// TableName 指定表名
func (utpw *UserThirdPartyWechatEntity) TableName() string {
	return "user_third_party_wechats"
}
