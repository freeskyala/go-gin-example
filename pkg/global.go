package global

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	elastic "github.com/olivere/elastic/v6"

)

var(
	VAFFLE_DB  *gorm.DB
	TEST_DB  *gorm.DB
	REDIS redis.Conn
	RedisPoolConn map[int]*redis.Pool
	ES *elastic.Client
)
