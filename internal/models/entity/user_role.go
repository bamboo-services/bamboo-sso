package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// UserRole 表示用户角色关联实体，用于多对多关系映射。
//
// 字段说明：
//   - UUID: 关联记录的唯一标识符，由 UUID 表示。
//   - UserUUID: 关联的用户UUID，外键。
//   - RoleUUID: 关联的角色UUID，外键。
//   - AssignedBy: 分配者UUID，记录是谁分配的此角色。
//   - AssignedAt: 分配时间。
//   - ExpiresAt: 角色过期时间，可选字段。
//   - IsActive: 角色是否激活，默认为 true。
//   - CreatedAt: 创建记录的时间戳。
//   - UpdatedAt: 最后更新时间戳。
type UserRole struct {
	UUID       uuid.UUID  `json:"uuid" gorm:"primaryKey;type:uuid;not null;comment:用户角色关联唯一标识符"`
	UserUUID   uuid.UUID  `json:"user_uuid" gorm:"type:uuid;not null;index;comment:关联用户UUID"`
	RoleUUID   uuid.UUID  `json:"role_uuid" gorm:"type:uuid;not null;index;comment:关联角色UUID"`
	AssignedBy *uuid.UUID `json:"assigned_by" gorm:"type:uuid;comment:分配者UUID"`
	AssignedAt time.Time  `json:"assigned_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:分配时间"`
	ExpiresAt  *time.Time `json:"expires_at" gorm:"type:timestamp;comment:过期时间"`
	IsActive   bool       `json:"is_active" gorm:"type:boolean;not null;default:true;comment:是否激活"`
	CreatedAt  time.Time  `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`

	// 关联关系
	User           *User `json:"user,omitempty" gorm:"foreignKey:UserUUID;references:UUID;constraint:OnDelete:CASCADE;comment:关联用户"`
	Role           *Role `json:"role,omitempty" gorm:"foreignKey:RoleUUID;references:UUID;constraint:OnDelete:CASCADE;comment:关联角色"`
	AssignedByUser *User `json:"assigned_by_user,omitempty" gorm:"foreignKey:AssignedBy;references:UUID;comment:分配者"`
}

// BeforeCreate 在创建 UserRole 记录前自动生成新的 UUID（如果当前 UUID 为空）。
func (ur *UserRole) BeforeCreate(_ *gorm.DB) (err error) {
	if ur.UUID == uuid.Nil {
		newUUID, err := uuid.NewV7()
		if err != nil {
			return err
		}
		ur.UUID = newUUID
	}
	if ur.AssignedAt.IsZero() {
		ur.AssignedAt = time.Now()
	}
	return
}

// BeforeUpdate 在更新 UserRole 记录前自动更新 UpdatedAt 字段。
func (ur *UserRole) BeforeUpdate(_ *gorm.DB) (err error) {
	ur.UpdatedAt = time.Now()
	return
}
