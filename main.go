package main

import (
	"fmt"
	xInit "github.com/bamboo-services/bamboo-base-go/init"
	"github.com/bamboo-services/bamboo-sso/internal/router"
	"github.com/bamboo-services/bamboo-sso/pkg/startup"
)

// main 是程序的入口点，负责初始化服务注册器、路由和启动 HTTP 服务器。
func main() {
	// 初始化注册器
	register := xInit.Register()
	engine := startup.Register(register)

	// 注册路由
	router.RegisterRoute(engine)

	// 启动服务器
	startPort := 2233
	if register.Config.Xlf.Port != nil {
		startPort = *register.Config.Xlf.Port
	}
	err := engine.Run(fmt.Sprintf("%s:%d", register.Config.Xlf.Host, startPort))
	if err != nil {
		panic("启动服务器失败: " + err.Error())
	}
}
