package startup

import (
	"fmt"
	xConsts "github.com/bamboo-services/bamboo-base-go/constants"
	"github.com/bamboo-services/bamboo-sso/internal/models/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var tableEntity = []interface{}{
	&entity.RoleEntity{},
	&entity.UserEntity{},
	&entity.UserProfileEntity{},
	&entity.UserRoleEntity{},
	&entity.ThirdPartyProviderEntity{},
	&entity.UserThirdPartyWechatEntity{},
	&entity.UserThirdPartyGithubEntity{},
	&entity.UserThirdPartyQQEntity{},
	&entity.ApplicationEntity{},
	&entity.AuthorizationCodeEntity{},
	&entity.LoginLogEntity{},
	&entity.AuthorizationLogEntity{},
}

// DatabaseStartup 初始化数据库连接并配置为服务的主数据库实例。
// 如果数据库连接失败，函数将会因 panic 终止程序。
func (r *reg) DatabaseStartup() {
	r.serv.Logger.Named(xConsts.LogINIT).Info("初始化 PgSQL 数据库连接")
	getConfig := r.serv.Config

	// 数据库连接
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s  TimeZone=Asia/Shanghai",
		getConfig.Database.Host,
		getConfig.Database.Port,
		getConfig.Database.User,
		getConfig.Database.Pass,
		getConfig.Database.Name,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("[DB] 数据库连接失败: " + err.Error())
	}

	// 自动迁移数据库表
	err = db.AutoMigrate(tableEntity...)
	if err != nil {
		panic("[DB] 数据库自动迁移失败: " + err.Error())
	} else {
		r.serv.Logger.Named(xConsts.LogINIT).Debug("数据库自动迁移成功")
	}

	r.db = db
}
