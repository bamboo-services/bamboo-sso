package startup

import (
	xConsts "github.com/bamboo-services/bamboo-base-go/constants"
	"github.com/gin-gonic/gin"
)

type handler struct {
	reg *reg
}

// ContextRegister 注册数据库与缓存的上下文到服务中。
//
// 此方法通过中间件将数据库和 Redis 客户端实例绑定到请求生命周期的上下文中。
// 它确保系统在处理每个请求时都可以访问预配置的数据库和缓存实例。
func (r *reg) ContextRegister() {
	r.serv.Logger.Named("Context").Info("注册数据库与缓存的上下文")

	// 初始化注册器
	handler := &handler{reg: r}

	// 注册系统上下文处理函数
	r.serv.Serve.Use(handler.handlerContext)
}

// handlerContext 将数据库和 Redis 客户端实例绑定到请求上下文中以便后续处理使用。
func (h *handler) handlerContext(c *gin.Context) {
	c.Set(xConsts.ContextDatabase, h.reg.db)
	c.Set(xConsts.ContextRedisClient, h.reg.rdb)
	c.Next()
}
