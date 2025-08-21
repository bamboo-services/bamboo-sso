package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// System 表示系统信息实体，存储系统配置信息的键值对。
type System struct {
	UUID      uuid.UUID `json:"uuid" gorm:"primaryKey;type:uuid;not null;comment:系统信息唯一标识符"`
	Key       string    `json:"key" gorm:"type:varchar(255);not null;uniqueIndex;comment:配置项键名"`
	Value     *string   `json:"value" gorm:"type:varchar(255);comment:配置项值"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
}

// BeforeCreate 在创建 System 记录前自动生成新的 UUID（如果当前 UUID 为空）。
func (s *System) BeforeCreate(_ *gorm.DB) (err error) {
	if s.UUID == uuid.Nil {
		newUUID, err := uuid.NewV7()
		if err != nil {
			return err
		}
		s.UUID = newUUID
	}
	return
}

// BeforeUpdate 在更新 System 记录前自动更新 UpdatedAt 字段。
func (s *System) BeforeUpdate(_ *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return
}
