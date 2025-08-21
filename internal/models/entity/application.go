package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Application 表示接入SSO的客户端应用实体，存储第三方应用的配置信息。
//
// 字段说明：
//   - UUID: 应用的唯一标识符，由 UUID 表示。
//   - Name: 应用名称。
//   - Description: 应用描述信息。
//   - ApplicationID: 应用标识符，分发给客户端用于身份识别。
//   - ApplicationSecret: 应用密钥，用于验证客户端身份。
//   - RedirectURIs: 允许的回调地址列表，JSON数组格式。
//   - AllowedOrigins: 允许的来源域名列表，JSON数组格式。
//   - LogoURL: 应用Logo地址。
//   - HomepageURL: 应用主页地址。
//   - PrivacyPolicyURL: 隐私政策地址。
//   - TermsOfServiceURL: 服务条款地址。
//   - IsActive: 应用是否激活，默认为 true。
//   - CreatedBy: 创建者UUID。
//   - CreatedAt: 创建记录的时间戳。
//   - UpdatedAt: 最后更新时间戳。
type Application struct {
	UUID              uuid.UUID `json:"uuid" gorm:"primaryKey;type:uuid;not null;comment:应用唯一标识符"`
	Name              string    `json:"name" gorm:"type:varchar(100);not null;comment:应用名称"`
	Description       string    `json:"description" gorm:"type:text;comment:应用描述"`
	ApplicationID     string    `json:"application_id" gorm:"type:varchar(50);not null;uniqueIndex;comment:应用标识符"`
	ApplicationSecret string    `json:"-" gorm:"type:varchar(255);not null;comment:应用密钥"`
	RedirectURIs      string    `json:"redirect_uris" gorm:"type:jsonb;comment:允许的回调地址(JSON数组)"`
	AllowedOrigins    string    `json:"allowed_origins" gorm:"type:jsonb;comment:允许的来源域名(JSON数组)"`
	LogoURL           string    `json:"logo_url" gorm:"type:varchar(500);comment:应用Logo地址"`
	HomepageURL       string    `json:"homepage_url" gorm:"type:varchar(500);comment:应用主页地址"`
	PrivacyPolicyURL  string    `json:"privacy_policy_url" gorm:"type:varchar(500);comment:隐私政策地址"`
	TermsOfServiceURL string    `json:"terms_of_service_url" gorm:"type:varchar(500);comment:服务条款地址"`
	IsActive          bool      `json:"is_active" gorm:"type:boolean;not null;default:true;comment:是否激活"`
	CreatedBy         uuid.UUID `json:"created_by" gorm:"type:uuid;comment:创建者UUID"`
	CreatedAt         time.Time `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`

	// 关联关系
	AuthorizationCodes []*AuthorizationCode `json:"authorization_codes,omitempty" gorm:"foreignKey:ApplicationUUID;references:UUID;constraint:OnDelete:CASCADE;comment:授权码"`
	Creator            *User                `json:"creator,omitempty" gorm:"foreignKey:CreatedBy;references:UUID;comment:创建者"`
}

// BeforeCreate 在创建 Application 记录前自动生成新的 UUID（如果当前 UUID 为空）。
func (a *Application) BeforeCreate(tx *gorm.DB) (err error) {
	if a.UUID == uuid.Nil {
		newUUID, err := uuid.NewV7()
		if err != nil {
			return err
		}
		a.UUID = newUUID
	}
	return
}

// BeforeUpdate 在更新 Application 记录前自动更新 UpdatedAt 字段。
func (a *Application) BeforeUpdate(tx *gorm.DB) (err error) {
	a.UpdatedAt = time.Now()
	return
}
