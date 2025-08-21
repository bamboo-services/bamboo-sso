package router

import "github.com/gin-gonic/gin"

type router struct {
	group *gin.RouterGroup
}

func RegisterRoute(engine *gin.Engine) {
	group := engine.Group("api/v1")

	r := &router{group: group}

	// 路由注册
	r.RouterHealth()
	r.RouterPublic()
}
