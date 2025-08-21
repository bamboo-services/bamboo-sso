package main

import (
	xInit "github.com/bamboo-services/bamboo-base-go/init"
	"github.com/bamboo-services/bamboo-sso/pkg/startup"
)

func main() {
	// 初始化注册器
	register := xInit.Register()
	engine := startup.Register(register)

	err := engine.Run(":8080")
	if err != nil {
		panic("启动服务器失败: " + err.Error())
	}
}
