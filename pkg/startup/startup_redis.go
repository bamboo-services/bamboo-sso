package startup

import (
	xConsts "github.com/bamboo-services/bamboo-base-go/constants"
	"github.com/redis/go-redis/v9"
	"strconv"
)

// RedisStartup 初始化 Redis 客户端实例并将其配置为服务的 Redis 数据库连接。
func (r *reg) RedisStartup() {
	r.serv.Logger.Named(xConsts.LogINIT).Info("初始化 Redis 客户端")
	getConfig := r.serv.Config

	rdb := redis.NewClient(&redis.Options{
		Addr:     getConfig.Nosql.Host + ":" + strconv.Itoa(getConfig.Nosql.Port),
		Password: getConfig.Nosql.Pass,
		DB:       getConfig.Nosql.Database,
	})
	r.rdb = rdb
}
