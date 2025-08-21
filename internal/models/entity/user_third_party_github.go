package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// UserThirdPartyGithub 表示用户Github账号绑定实体，存储用户与Github平台账号的绑定关系。
//
// Github登录特有字段：
//   - GithubID: Github用户ID（数字）
//   - Login: Github用户名
//   - NodeID: Github Node ID
//   - Type: 用户类型（User/Organization）
//   - SiteAdmin: 是否为站点管理员
//   - Company: 所属公司
//   - Blog: 博客地址
//   - Location: 地理位置
//   - Bio: 个人简介
//   - PublicRepos: 公开仓库数量
//   - PublicGists: 公开Gist数量
//   - Followers: 关注者数量
//   - Following: 关注的人数量
//   - TwitterUsername: Twitter用户名
//   - HirableAvailable: 是否可雇佣
//
// 字段说明：
//   - UUID: 绑定记录的唯一标识符，由 UUID 表示。
//   - UserUUID: 关联的用户UUID，外键。
//   - ProviderUUID: 关联的Github提供商UUID，外键。
//   - GithubID: Github用户ID。
//   - Login: Github用户名。
//   - NodeID: Github Node ID。
//   - Avatar: Github用户头像。
//   - GravatarID: Gravatar ID。
//   - Type: 用户类型。
//   - SiteAdmin: 是否为站点管理员。
//   - Name: 用户真实姓名。
//   - Company: 所属公司。
//   - Blog: 博客地址。
//   - Location: 地理位置。
//   - Email: 邮箱地址。
//   - Bio: 个人简介。
//   - TwitterUsername: Twitter用户名。
//   - PublicRepos: 公开仓库数量。
//   - PublicGists: 公开Gist数量。
//   - Followers: 关注者数量。
//   - Following: 关注的人数量。
//   - HirableAvailable: 是否可雇佣。
//   - AccessToken: Github访问令牌（加密存储）。
//   - TokenType: 令牌类型。
//   - Scope: 权限范围。
//   - IsActive: 绑定是否激活，默认为 true。
//   - FirstBindAt: 首次绑定时间。
//   - LastLoginAt: 最后一次通过Github登录的时间。
//   - CreatedAt: 创建记录的时间戳。
//   - UpdatedAt: 最后更新时间戳。
type UserThirdPartyGithub struct {
	UUID             uuid.UUID  `json:"uuid" gorm:"primaryKey;type:uuid;not null;comment:Github绑定记录唯一标识符"`
	UserUUID         uuid.UUID  `json:"user_uuid" gorm:"type:uuid;not null;index;comment:关联用户UUID"`
	ProviderUUID     uuid.UUID  `json:"provider_uuid" gorm:"type:uuid;not null;index;comment:关联Github提供商UUID"`
	GithubID         int64      `json:"github_id" gorm:"type:bigint;not null;uniqueIndex;comment:Github用户ID"`
	Login            string     `json:"login" gorm:"type:varchar(100);not null;index;comment:Github用户名"`
	NodeID           *string    `json:"node_id" gorm:"type:varchar(50);comment:Github Node ID"`
	Avatar           *string    `json:"avatar" gorm:"type:varchar(500);comment:Github用户头像"`
	GravatarID       *string    `json:"gravatar_id" gorm:"type:varchar(50);comment:Gravatar ID"`
	Type             *string    `json:"type" gorm:"type:varchar(20);default:'User';comment:用户类型(User/Organization)"`
	SiteAdmin        bool       `json:"site_admin" gorm:"type:boolean;default:false;comment:是否为站点管理员"`
	Name             *string    `json:"name" gorm:"type:varchar(100);comment:用户真实姓名"`
	Company          *string    `json:"company" gorm:"type:varchar(100);comment:所属公司"`
	Blog             *string    `json:"blog" gorm:"type:varchar(200);comment:博客地址"`
	Location         *string    `json:"location" gorm:"type:varchar(100);comment:地理位置"`
	Email            *string    `json:"email" gorm:"type:varchar(100);comment:邮箱地址"`
	Bio              *string    `json:"bio" gorm:"type:text;comment:个人简介"`
	TwitterUsername  *string    `json:"twitter_username" gorm:"type:varchar(50);comment:Twitter用户名"`
	PublicRepos      int        `json:"public_repos" gorm:"type:integer;default:0;comment:公开仓库数量"`
	PublicGists      int        `json:"public_gists" gorm:"type:integer;default:0;comment:公开Gist数量"`
	Followers        int        `json:"followers" gorm:"type:integer;default:0;comment:关注者数量"`
	Following        int        `json:"following" gorm:"type:integer;default:0;comment:关注的人数量"`
	HirableAvailable bool       `json:"hirable_available" gorm:"type:boolean;default:false;comment:是否可雇佣"`
	AccessToken      *string    `json:"-" gorm:"type:text;comment:Github访问令牌(加密)"`
	TokenType        *string    `json:"token_type" gorm:"type:varchar(20);default:'bearer';comment:令牌类型"`
	Scope            *string    `json:"scope" gorm:"type:varchar(200);comment:权限范围"`
	IsActive         bool       `json:"is_active" gorm:"type:boolean;not null;default:true;comment:是否激活"`
	FirstBindAt      time.Time  `json:"first_bind_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:首次绑定时间"`
	LastLoginAt      *time.Time `json:"last_login_at" gorm:"type:timestamp;comment:最后登录时间"`
	CreatedAt        time.Time  `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt        time.Time  `json:"updated_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`

	// 关联关系
	User     *User               `json:"user,omitempty" gorm:"foreignKey:UserUUID;references:UUID;constraint:OnDelete:CASCADE;comment:关联用户"`
	Provider *ThirdPartyProvider `json:"provider,omitempty" gorm:"foreignKey:ProviderUUID;references:UUID;constraint:OnDelete:CASCADE;comment:关联Github提供商"`
}

// BeforeCreate 在创建 UserThirdPartyGithub 记录前自动生成新的 UUID（如果当前 UUID 为空）。
func (utpg *UserThirdPartyGithub) BeforeCreate(tx *gorm.DB) (err error) {
	if utpg.UUID == uuid.Nil {
		newUUID, err := uuid.NewV7()
		if err != nil {
			return err
		}
		utpg.UUID = newUUID
	}
	if utpg.FirstBindAt.IsZero() {
		utpg.FirstBindAt = time.Now()
	}
	return
}

// BeforeUpdate 在更新 UserThirdPartyGithub 记录前自动更新 UpdatedAt 字段。
func (utpg *UserThirdPartyGithub) BeforeUpdate(tx *gorm.DB) (err error) {
	utpg.UpdatedAt = time.Now()
	return
}
