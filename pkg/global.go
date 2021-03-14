package global

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

var(
	VAFFLE_DB  *gorm.DB
	TEST_DB  *gorm.DB
	REDIS redis.Conn
)
