package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Role 表示系统中的角色实体，用于定义权限和业务角色。
//
// 字段说明：
//   - UUID: 角色的唯一标识符，由 UUID 表示。
//   - Name: 角色名称，必须唯一。
//   - Description: 角色的描述信息，可选字段。
//   - CreatedAt: 创建记录的时间戳。
//   - UpdatedAt: 最后更新时间戳。
type Role struct {
	UUID        uuid.UUID `json:"uuid" gorm:"primaryKey;type:uuid;not null;comment:角色唯一标识符"`
	Name        string    `json:"name" gorm:"type:varchar(50);not null;uniqueIndex;comment:角色名称"`
	DisplayName string    `json:"display_name" gorm:"type:varchar(100);not null;comment:角色显示名称"`
	Description *string   `json:"description" gorm:"type:varchar(255);comment:角色描述信息"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
}

// BeforeCreate 在创建 Role 记录前自动生成新的 UUID（如果当前 UUID 为空）。
func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	if r.UUID == uuid.Nil {
		newUUID, err := uuid.NewV7()
		if err != nil {
			return err
		}
		r.UUID = newUUID
	}
	return
}

// BeforeUpdate 在更新 Role 记录前自动更新 UpdatedAt 字段。
func (r *Role) BeforeUpdate(tx *gorm.DB) (err error) {
	r.UpdatedAt = time.Now()
	return
}
