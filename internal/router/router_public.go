package router

// RouterPublic 注册公共访问的路由入口点。
//
// 路径 "/public/ping" 可用于提供基本的公共服务，如可达性测试。
func (r *router) RouterPublic() {
	group := r.group.Group("/public")

	{
		group.GET("/ping")
	}
}
