package config

import (
	"database/sql"
	"errors"
	xConsts "github.com/bamboo-services/bamboo-base-go/constants"
	"github.com/bamboo-services/bamboo-sso/internal/models/entity"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// InitializeData 表示初始化数据的结构体。
//
// InitializeData 包含一个 GORM 数据库连接实例，
// 用于执行数据库相关的操作，如表的初始化、数据迁移以及数据填充。
//
// 字段说明：
//   - db: GORM 数据库连接实例，用于与数据库交互。
//   - log: 日志记录器实例，用于记录初始化过程中的日志信息。
type InitializeData struct {
	db  *gorm.DB
	log *zap.Logger
}

// New 创建并返回一个新的 InitializeData 实例。
func New(db *gorm.DB, logger *zap.Logger) *InitializeData {
	return &InitializeData{
		db:  db,
		log: logger,
	}
}

// RoleInit 检查并初始化系统中缺失的角色数据。
//
// 参数 getEntity 是一组指针，指向需要检测或创建的角色实体。
// 如果传入的角色在数据库中不存在，则会创建默认的角色记录。
// 当角色已存在时，不会重复创建，避免数据库冗余。
//
// 方法使用逻辑：
//   - 首先检查每个角色的名称是否已存在于数据库。
//   - 若角色不存在，则记录在批量插入列表中以优化数据库操作。
//   - 最后，统一插入所有需要创建的角色记录以减少数据库压力。
//
// 注意：创建操作会忽略已存在的记录，并直接略过处理。
func (i *InitializeData) RoleInit(getEntity ...*entity.Role) {
	db := i.db
	log := i.log

	var noneRoleList []*entity.Role

	// 检查并创建默认角色
	for _, roleEntity := range getEntity {
		var role entity.Role
		if err := db.Where("name = @name", sql.Named("name", roleEntity.Name)).First(&role).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Named(xConsts.LogINIT).Sugar().Debugf("角色 %s 不存在，创建默认角色", roleEntity.Name)
				noneRoleList = append(noneRoleList, roleEntity)
			} else {
				log.Named(xConsts.LogINIT).Sugar().Debugf("角色 %s 已存在，跳过创建", roleEntity.Name)
			}
		}
	}

	// 批量创建角色「统一插入减少数据库操作压力」
	if len(noneRoleList) > 0 {
		db.Create(noneRoleList)
	}
}

func (i *InitializeData) ApplicationInit(getEntity ...*entity.Application) {

}
