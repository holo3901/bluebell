package redis

//创建redis数据库
import (
	"fmt"
	"xxx/settings"

	"github.com/go-redis/redis"
)

//声明一个全局的rdb变量
var (
	client *redis.Client
	Nil    = redis.Nil
)

//初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			/*	viper.GetString("redis.host"),
					viper.GetInt("redis.port"),
				),
				Password: viper.GetString("redis.password"),
				DB:       viper.GetInt("redis.db"),
				PoolSize: viper.GetInt("redis.pool_size"),*/
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = client.Ping().Result()
	return err
}
func Close() {
	_ = client.Close()
}
