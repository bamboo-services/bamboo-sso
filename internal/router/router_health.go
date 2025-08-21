package router

// RouterHealth 注册用于监控服务健康状态的路由。
//
// 路径 "/health/ping" 提供基本健康检查功能，通常用于负载均衡器探测服务状态。
func (r *router) RouterHealth() {
	group := r.group.Group("/health")

	{
		group.GET("/ping")
	}
}
