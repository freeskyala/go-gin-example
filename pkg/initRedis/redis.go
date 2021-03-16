package initRedis

import (
	global2 "github.com/EDDYCJY/go-gin-example/pkg/global"
	"github.com/EDDYCJY/go-gin-example/pkg/initconfig"
	"github.com/gomodule/redigo/redis"
	"time"
)

var BaseRedisPool *redis.Pool

type RedisPool struct {

}

func (r *RedisPool) InitRedisDb() {
	r.Select(0)
}

//获取切库连接池中的连接
func (r *RedisPool) Select (db int) redis.Conn  {
	if global2.RedisPoolConn == nil {
		global2.RedisPoolConn = make(map[int]*redis.Pool)
	}
	if _, ok := global2.RedisPoolConn[db]; !ok{
		global2.RedisPoolConn[db] = r.CreateRedisConn(db)
	}
	return global2.RedisPoolConn[db].Get()
}

func (r RedisPool) CreateRedisConn(db int) *redis.Pool {
	//var c *redis.Pool
	maxIdle := 50
	maxActive := 1024
	MaxIdleTimeout := 30
	ConnectTimeout := 1
	ReadTimeout := 2
	// 建立连接池
	BaseRedisPool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(MaxIdleTimeout) * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp",initconfig.RedisConfig.Host,
				redis.DialPassword(initconfig.RedisConfig.Password),
				redis.DialDatabase(0),
				redis.DialConnectTimeout(time.Duration(ConnectTimeout)*time.Second),
				redis.DialReadTimeout(time.Duration(ReadTimeout)*time.Second),
				redis.DialWriteTimeout(time.Duration(ReadTimeout)*time.Second))
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("SELECT", db); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
	return BaseRedisPool
}


