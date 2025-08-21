package startup

import (
	xInit "github.com/bamboo-services/bamboo-base-go/init"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type reg struct {
	serv *xInit.Reg    // 服务实例，提供必要的依赖和配置
	db   *gorm.DB      // 数据库连接实例，用于与数据库进行交互
	rdb  *redis.Client // Redis 客户端实例，用于与 Redis 数据库进行交互
}

// New 创建一个新的 reg 实例并初始化其必要的依赖项。输入参数 serv 必须是有效的 *xInit.Reg 实例。
func New(serv *xInit.Reg) *reg {
	return &reg{
		serv: serv,
	}
}

// Register 初始化并注册一个新的 reg 实例，将其绑定到提供的 *xInit.Reg 服务实例中。
func Register(serv *xInit.Reg) *gin.Engine {
	reg := New(serv)

	// 初始化内容注册
	reg.DatabaseStartup()
	reg.RedisStartup()

	// 注册上下文

	return reg.serv.Serve
}
