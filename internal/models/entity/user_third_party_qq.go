package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// UserThirdPartyQQEntity 表示用户QQ账号绑定实体，存储用户与QQ平台账号的绑定关系。
//
// QQ登录特有字段：
//   - OpenID: QQ用户唯一标识
//   - UnionID: QQ开放平台统一用户标识（如果应用申请了UnionID）
//   - QQNumber: QQ号码（需要特殊权限才能获取）
//   - Gender: 性别
//   - Province: 省份
//   - City: 城市
//   - Year: 出生年份
//   - Constellation: 星座
//   - IsLost: 判断QQ用户是否为认证用户
//   - Figureurl: 用户头像链接（30x30像素）
//   - Figureurl1: 用户头像链接（50x50像素）
//   - Figureurl2: 用户头像链接（100x100像素）
//   - FigureurlQQ1: QQ头像链接（40x40像素）
//   - FigureurlQQ2: QQ头像链接（100x100像素）
//   - IsVip: 是否为QQ会员
//   - VipLevel: QQ会员等级
//   - IsYellowVip: 是否为黄钻会员
//   - YellowVipLevel: 黄钻等级
//
// 字段说明：
//   - UUID: 绑定记录的唯一标识符，由 UUID 表示。
//   - UserUUID: 关联的用户UUID，外键。
//   - ProviderUUID: 关联的QQ提供商UUID，外键。
//   - OpenID: QQ用户唯一标识。
//   - UnionID: QQ开放平台统一用户标识。
//   - QQNumber: QQ号码。
//   - Nickname: QQ用户昵称。
//   - Gender: 性别。
//   - Province: 省份。
//   - City: 城市。
//   - Year: 出生年份。
//   - Constellation: 星座。
//   - IsLost: 是否为认证用户。
//   - Figureurl: 30x30像素头像。
//   - Figureurl1: 50x50像素头像。
//   - Figureurl2: 100x100像素头像。
//   - FigureurlQQ1: 40x40像素QQ头像。
//   - FigureurlQQ2: 100x100像素QQ头像。
//   - IsVip: 是否为QQ会员。
//   - VipLevel: QQ会员等级。
//   - IsYellowVip: 是否为黄钻会员。
//   - YellowVipLevel: 黄钻等级。
//   - AccessToken: QQ访问令牌（加密存储）。
//   - RefreshToken: QQ刷新令牌（加密存储）。
//   - TokenExpiresAt: 访问令牌过期时间。
//   - IsActive: 绑定是否激活，默认为 true。
//   - FirstBindAt: 首次绑定时间。
//   - LastLoginAt: 最后一次通过QQ登录的时间。
//   - CreatedAt: 创建记录的时间戳。
//   - UpdatedAt: 最后更新时间戳。
type UserThirdPartyQQEntity struct {
	UUID           uuid.UUID  `json:"uuid" gorm:"primaryKey;type:uuid;not null;comment:QQ绑定记录唯一标识符"`
	UserUUID       uuid.UUID  `json:"user_uuid" gorm:"type:uuid;not null;index;comment:关联用户UUID"`
	ProviderUUID   uuid.UUID  `json:"provider_uuid" gorm:"type:uuid;not null;index;comment:关联QQ提供商UUID"`
	OpenID         string     `json:"open_id" gorm:"type:varchar(100);not null;uniqueIndex;comment:QQ用户唯一标识"`
	UnionID        string     `json:"union_id" gorm:"type:varchar(100);comment:QQ开放平台统一用户标识"`
	QQNumber       string     `json:"qq_number" gorm:"type:varchar(20);comment:QQ号码"`
	Nickname       string     `json:"nickname" gorm:"type:varchar(100);comment:QQ用户昵称"`
	Gender         string     `json:"gender" gorm:"type:varchar(10);comment:性别(男/女)"`
	Province       string     `json:"province" gorm:"type:varchar(50);comment:省份"`
	City           string     `json:"city" gorm:"type:varchar(50);comment:城市"`
	Year           string     `json:"year" gorm:"type:varchar(4);comment:出生年份"`
	Constellation  string     `json:"constellation" gorm:"type:varchar(20);comment:星座"`
	IsLost         bool       `json:"is_lost" gorm:"type:boolean;default:false;comment:是否为认证用户"`
	Figureurl      string     `json:"figureurl" gorm:"type:varchar(500);comment:用户头像30x30"`
	Figureurl1     string     `json:"figureurl_1" gorm:"type:varchar(500);comment:用户头像50x50"`
	Figureurl2     string     `json:"figureurl_2" gorm:"type:varchar(500);comment:用户头像100x100"`
	FigureurlQQ1   string     `json:"figureurl_qq_1" gorm:"type:varchar(500);comment:QQ头像40x40"`
	FigureurlQQ2   string     `json:"figureurl_qq_2" gorm:"type:varchar(500);comment:QQ头像100x100"`
	IsVip          bool       `json:"is_vip" gorm:"type:boolean;default:false;comment:是否为QQ会员"`
	VipLevel       int        `json:"vip_level" gorm:"type:smallint;default:0;comment:QQ会员等级"`
	IsYellowVip    bool       `json:"is_yellow_vip" gorm:"type:boolean;default:false;comment:是否为黄钻会员"`
	YellowVipLevel int        `json:"yellow_vip_level" gorm:"type:smallint;default:0;comment:黄钻等级"`
	AccessToken    string     `json:"-" gorm:"type:text;comment:QQ访问令牌(加密)"`
	RefreshToken   string     `json:"-" gorm:"type:text;comment:QQ刷新令牌(加密)"`
	TokenExpiresAt *time.Time `json:"token_expires_at" gorm:"type:timestamp;comment:访问令牌过期时间"`
	IsActive       bool       `json:"is_active" gorm:"type:boolean;not null;default:true;comment:是否激活"`
	FirstBindAt    time.Time  `json:"first_bind_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:首次绑定时间"`
	LastLoginAt    *time.Time `json:"last_login_at" gorm:"type:timestamp;comment:最后登录时间"`
	CreatedAt      time.Time  `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt      time.Time  `json:"updated_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`

	// 关联关系
	User     *UserEntity               `json:"user,omitempty" gorm:"foreignKey:UserUUID;references:UUID;constraint:OnDelete:CASCADE;comment:关联用户"`
	Provider *ThirdPartyProviderEntity `json:"provider,omitempty" gorm:"foreignKey:ProviderUUID;references:UUID;constraint:OnDelete:CASCADE;comment:关联QQ提供商"`
}

// BeforeCreate 在创建 UserThirdPartyQQEntity 记录前自动生成新的 UUID（如果当前 UUID 为空）。
func (utpq *UserThirdPartyQQEntity) BeforeCreate(tx *gorm.DB) (err error) {
	if utpq.UUID == uuid.Nil {
		newUUID, err := uuid.NewV7()
		if err != nil {
			return err
		}
		utpq.UUID = newUUID
	}
	if utpq.FirstBindAt.IsZero() {
		utpq.FirstBindAt = time.Now()
	}
	return
}

// BeforeUpdate 在更新 UserThirdPartyQQEntity 记录前自动更新 UpdatedAt 字段。
func (utpq *UserThirdPartyQQEntity) BeforeUpdate(tx *gorm.DB) (err error) {
	utpq.UpdatedAt = time.Now()
	return
}

// TableName 指定表名
func (utpq *UserThirdPartyQQEntity) TableName() string {
	return "user_third_party_qqs"
}
