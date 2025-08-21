package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// UserProfile 表示用户详细资料实体，存储用户的扩展信息。
//
// 字段说明：
//   - UUID: 用户资料的唯一标识符，由 UUID 表示。
//   - UserUUID: 关联的用户UUID，外键。
//   - Nickname: 用户昵称，可选字段。
//   - Avatar: 用户头像URL，可选字段。
//   - Gender: 用户性别，0-未知，1-男性，2-女性。
//   - Birthday: 用户生日，可选字段。
//   - Country: 国家，可选字段。
//   - Province: 省份/州，可选字段。
//   - City: 城市，可选字段。
//   - Bio: 个人简介，可选字段。
//   - CreatedAt: 创建记录的时间戳。
//   - UpdatedAt: 最后更新时间戳。
type UserProfile struct {
	UUID      uuid.UUID  `json:"uuid" gorm:"primaryKey;type:uuid;not null;comment:用户资料唯一标识符"`
	UserUUID  uuid.UUID  `json:"user_uuid" gorm:"type:uuid;not null;uniqueIndex;comment:关联用户UUID"`
	Nickname  *string    `json:"nickname" gorm:"type:varchar(50);comment:用户昵称"`
	Avatar    *string    `json:"avatar" gorm:"type:varchar(500);comment:用户头像URL"`
	Gender    int        `json:"gender" gorm:"type:smallint;default:0;comment:性别(0-未知,1-男,2-女)"`
	Birthday  *time.Time `json:"birthday" gorm:"type:date;comment:生日"`
	Country   *string    `json:"country" gorm:"type:varchar(50);comment:国家"`
	Province  *string    `json:"province" gorm:"type:varchar(50);comment:省份/州"`
	City      *string    `json:"city" gorm:"type:varchar(50);comment:城市"`
	Bio       *string    `json:"bio" gorm:"type:text;comment:个人简介"`
	CreatedAt time.Time  `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`

	// 关联关系
	User *User `json:"user,omitempty" gorm:"foreignKey:UserUUID;references:UUID;constraint:OnDelete:CASCADE;comment:关联用户"`
}

// BeforeCreate 在创建 UserProfile 记录前自动生成新的 UUID（如果当前 UUID 为空）。
func (up *UserProfile) BeforeCreate(tx *gorm.DB) (err error) {
	if up.UUID == uuid.Nil {
		newUUID, err := uuid.NewV7()
		if err != nil {
			return err
		}
		up.UUID = newUUID
	}
	return
}

// BeforeUpdate 在更新 UserProfile 记录前自动更新 UpdatedAt 字段。
func (up *UserProfile) BeforeUpdate(tx *gorm.DB) (err error) {
	up.UpdatedAt = time.Now()
	return
}
