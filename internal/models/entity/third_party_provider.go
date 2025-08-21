package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// ThirdPartyProvider 表示第三方登录提供商实体，存储第三方平台的配置信息。
//
// 字段说明：
//   - UUID: 提供商的唯一标识符，由 UUID 表示。
//   - Name: 提供商名称，如 "QQ"、"微信"、"Github"。
//   - Code: 提供商代码，用于程序识别，如 "qq"、"wechat"、"github"。
//   - ClientID: 第三方平台分配的客户端ID。
//   - ClientSecret: 第三方平台分配的客户端密钥。
//   - AuthURL: 授权地址。
//   - TokenURL: 获取Token的地址。
//   - UserInfoURL: 获取用户信息的地址。
//   - Scope: 请求的权限范围。
//   - RedirectURL: 回调地址。
//   - IsEnabled: 是否启用该提供商，默认为 true。
//   - SortOrder: 排序顺序，用于前端显示。
//   - CreatedAt: 创建记录的时间戳。
//   - UpdatedAt: 最后更新时间戳。
type ThirdPartyProvider struct {
	UUID         uuid.UUID `json:"uuid" gorm:"primaryKey;type:uuid;not null;comment:第三方提供商唯一标识符"`
	Name         string    `json:"name" gorm:"type:varchar(50);not null;comment:提供商名称"`
	Code         string    `json:"code" gorm:"type:varchar(30);not null;uniqueIndex;comment:提供商代码"`
	ClientID     string    `json:"client_id" gorm:"type:varchar(255);not null;comment:客户端ID"`
	ClientSecret string    `json:"-" gorm:"type:varchar(255);not null;comment:客户端密钥"`
	AuthURL      string    `json:"auth_url" gorm:"type:varchar(500);not null;comment:授权地址"`
	TokenURL     string    `json:"token_url" gorm:"type:varchar(500);not null;comment:Token地址"`
	UserInfoURL  string    `json:"user_info_url" gorm:"type:varchar(500);not null;comment:用户信息地址"`
	Scope        string    `json:"scope" gorm:"type:varchar(200);comment:权限范围"`
	RedirectURL  string    `json:"redirect_url" gorm:"type:varchar(500);not null;comment:回调地址"`
	IsEnabled    bool      `json:"is_enabled" gorm:"type:boolean;not null;default:true;comment:是否启用"`
	SortOrder    int       `json:"sort_order" gorm:"type:integer;default:0;comment:排序顺序"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`

	// 关联关系
	WechatAccounts []*UserThirdPartyWechat `json:"wechat_accounts,omitempty" gorm:"foreignKey:ProviderUUID;references:UUID;constraint:OnDelete:CASCADE;comment:微信用户账号"`
	GithubAccounts []*UserThirdPartyGithub `json:"github_accounts,omitempty" gorm:"foreignKey:ProviderUUID;references:UUID;constraint:OnDelete:CASCADE;comment:Github用户账号"`
	QQAccounts     []*UserThirdPartyQQ     `json:"qq_accounts,omitempty" gorm:"foreignKey:ProviderUUID;references:UUID;constraint:OnDelete:CASCADE;comment:QQ用户账号"`
}

// BeforeCreate 在创建 ThirdPartyProvider 记录前自动生成新的 UUID（如果当前 UUID 为空）。
func (tpp *ThirdPartyProvider) BeforeCreate(_ *gorm.DB) (err error) {
	if tpp.UUID == uuid.Nil {
		newUUID, err := uuid.NewV7()
		if err != nil {
			return err
		}
		tpp.UUID = newUUID
	}
	return
}

// BeforeUpdate 在更新 ThirdPartyProvider 记录前自动更新 UpdatedAt 字段。
func (tpp *ThirdPartyProvider) BeforeUpdate(_ *gorm.DB) (err error) {
	tpp.UpdatedAt = time.Now()
	return
}
