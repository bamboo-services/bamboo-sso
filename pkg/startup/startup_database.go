package startup

import (
	"fmt"
	xConsts "github.com/bamboo-services/bamboo-base-go/constants"
	xUtil "github.com/bamboo-services/bamboo-base-go/utility"
	"github.com/bamboo-services/bamboo-sso/internal/models/entity"
	"github.com/bamboo-services/bamboo-sso/pkg/config"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
)

var tableEntity = []interface{}{
	&entity.Role{},
	&entity.User{},
	&entity.UserProfile{},
	&entity.UserRole{},
	&entity.ThirdPartyProvider{},
	&entity.UserThirdPartyWechat{},
	&entity.UserThirdPartyGithub{},
	&entity.UserThirdPartyQQ{},
	&entity.Application{},
	&entity.AuthorizationCode{},
	&entity.LoginLog{},
	&entity.AuthorizationLog{},
}

// prepare 表示服务的初始化准备阶段，包括注册器和基础数据的初始化。
//
// 字段说明：
// - reg: 注册器实例，提供服务、数据库和缓存相关的功能。
// - init: 基础数据初始化实例，用于配置初始化数据库所需的数据。
type prepare struct {
	init *config.InitializeData // init 是初始化数据实例，用于准备基础数据
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   xUtil.DefaultIfBlank(getConfig.Database.Prefix, "xlf_"), // 表前缀
			SingularTable: true,                                                    // 使用单数表名
		},
	})
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

	// 检查是否启用 Debug 模式
	if getConfig.Xlf.Debug {
		r.serv.Logger.Named(xConsts.LogINIT).Debug("数据库连接开启 Debug 模式")
		db = db.Debug()
	}

	// 初始化基础数据
	getPrepare := &prepare{init: config.New(db, r.serv.Logger)}

	// 使用 WaitGroup 并发执数据的初始化
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() { defer wg.Done(); getPrepare.PrepareRole() }()
	go func() { defer wg.Done(); getPrepare.PrepareApplication() }()
	wg.Wait()
	r.serv.Logger.Named(xConsts.LogINIT).Info("基础数据初始化完成")

	r.db = db
}

// PrepareRole 初始化系统的默认角色数据。
//
// 调用此方法时，将在角色表中检查是否存在预定义的角色。若角色不存在，则创建以下默认角色：
// - "SUPER_ADMIN": 超级管理员，拥有所有权限。
// - "ADMIN": 管理员，仅次于超级管理员，拥有管理权限。
// - "USER": 普通用户，具有基本访问权限。
// 此方法用于系统初始化阶段以确保基础角色数据完整性。
func (p *prepare) PrepareRole() {
	superAdminDesc := "拥有所有权限的超级管理员角色"
	adminDesc := "拥有管理权限的管理员角色"
	userDesc := "普通用户角色，拥有基本的访问权限"

	p.init.RoleInit(
		&entity.Role{Name: "SUPER_ADMIN", DisplayName: "超级管理员", Description: &superAdminDesc},
		&entity.Role{Name: "ADMIN", DisplayName: "管理员", Description: &adminDesc},
		&entity.Role{Name: "USER", DisplayName: "普通用户", Description: &userDesc},
	)
}

// PrepareApplication 初始化系统的默认应用数据。
//
// 调用此方法时，将在应用表中检查是否存在预定义的应用数据。若应用不存在，则创建以下默认应用：
// - "Bamboo SSO": 单点登录服务应用，提供基础 SSO 功能支持。
// 此方法用于系统初始化阶段以确保基础应用数据的完整性。
func (p *prepare) PrepareApplication() {
	// 默认应用相关
	defaultApplicationRedirectJson, err := jsoniter.MarshalToString([]string{
		"http://localhost:2233",
		"http://localhost:4173",
		"http://localhost:3000",
		"http://localhost:4000",
	})
	defaultApplicationAllowedOriginsJson, err := jsoniter.MarshalToString([]string{
		"http://localhost:2233",
		"http://localhost:4173",
		"http://localhost:3000",
		"http://localhost:4000",
	})

	// 测试应用相关
	demoApplicationRedirectJson, err := jsoniter.MarshalToString([]string{
		"*",
	})
	demoApplicationAllowedOriginsJson, err := jsoniter.MarshalToString([]string{
		"*",
	})

	if err != nil {
		panic(err)
	}

	// 初始化默认应用
	defaultDesc := "Bamboo SSO 应用"
	demoDesc := "用于测试的应用"

	p.init.ApplicationInit(
		&entity.Application{
			Name:              "默认应用",
			Description:       &defaultDesc,
			ApplicationID:     "10000",
			ApplicationSecret: xUtil.GenerateSecurityKey(),
			RedirectURIs:      &defaultApplicationRedirectJson,
			AllowedOrigins:    &defaultApplicationAllowedOriginsJson,
		},
		&entity.Application{
			Name:              "测试应用",
			Description:       &demoDesc,
			ApplicationID:     "10001",
			ApplicationSecret: xUtil.GenerateSecurityKey(),
			RedirectURIs:      &demoApplicationRedirectJson,
			AllowedOrigins:    &demoApplicationAllowedOriginsJson,
		},
	)
}
