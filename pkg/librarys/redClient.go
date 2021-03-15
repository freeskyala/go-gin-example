package librarys

import (
	"errors"
	"fmt"
	global "github.com/EDDYCJY/go-gin-example/pkg"
	"github.com/EDDYCJY/go-gin-example/pkg/initstart"
	"github.com/gomodule/redigo/redis"
)

type RedisClientHandler struct {
	DB redis.Conn
}

//关闭连接池
func (r *RedisClientHandler) RedisClose()  {
	if len(global.RedisPoolConn)  != 0{
		for db, con := range global.RedisPoolConn {
			con.Close()
			fmt.Println(db, "redis pool closed")
		}
	}
}

//获取切库连接池中的连接
func (r *RedisClientHandler) RedisSelect (db int) *RedisClientHandler  {
	if global.RedisPoolConn == nil {
		global.RedisPoolConn = make(map[int]*redis.Pool)
	}
	if _, ok := global.RedisPoolConn[db]; !ok{
		global.RedisPoolConn[db] = new(initstart.RedisPool).CreateRedisConn(db)
	}
	r.DB = global.RedisPoolConn[db].Get()
	return r
}

func (r *RedisClientHandler) RedisSet(key string, value interface{})  (bool, error) {
	defer r.DB.Close()
	//value, _ := json.Marshal(value)
	_, err := r.DB.Do("set", key, value)
	if err != nil{
		return  false, errors.New("set failed")
	}
	return  true, err
}


func (r *RedisClientHandler) RedisRPush(key string, value string)  (bool, error) {
	defer r.DB.Close()
	_, err := r.DB.Do("RPUSH", key, value)
	if err != nil{
		return  false, errors.New("RPUSH failed")
	}
	return  true, err
}


func (r *RedisClientHandler) RedisLPop(key string)  (string, error) {
	defer r.DB.Close()
	str, err := redis.String(r.DB.Do("LPOP", key))
	return str, err
}

//查询结果为string的
func (r *RedisClientHandler) RedisGet(key string) (string) {
	defer r.DB.Close()
	str, err := redis.String(r.DB.Do("get", key))
	if err != nil{
		errors.New("set failed")
	}
	return str
}


func (r *RedisClientHandler) RedisIncr(key string) (int64, error) {
	defer r.DB.Close()
	res, err := r.DB.Do("incr", key)
	if err != nil{
		return  0, errors.New("set failed")
	}
	return  res.(int64), err
}

func (r *RedisClientHandler) RedisTtl(key string) (int64, error) {
	defer r.DB.Close()
	res, err := r.DB.Do("ttl", key)
	if err != nil{
		return  0, errors.New("get ttl failed")
	}
	return  res.(int64), err
}

func (r *RedisClientHandler) RedisExpire(key string, time int) {
	defer r.DB.Close()
	_, err := r.DB.Do("expire", key, time)
	if err != nil{
		return
	}
	return
}


func (r *RedisClientHandler) RedisSmembers(key string)(reply interface{}, err error) {
	defer r.DB.Close()
	reply , err = redis.Strings(r.DB.Do("smembers", key))
	if err != nil{
		return
	}
	return
}

func (r *RedisClientHandler) RedisScard(key string)(reply int, err error) {
	defer r.DB.Close()
	reply , err = redis.Int(r.DB.Do("scard", key))
	if err != nil{
		return
	}
	return
}

func (r *RedisClientHandler) RedisSismember(key string, member interface{})(reply bool, err error) {
	defer r.DB.Close()
	reply , err = redis.Bool(r.DB.Do("sismember", key, member))
	if err != nil{
		return
	}
	return
}

func (r *RedisClientHandler) RedisDel(key string) {
	defer r.DB.Close()
	_ , err := redis.Bool(r.DB.Do("del", key))
	if err != nil{
		return
	}
	return
}

func (r *RedisClientHandler) RedisSREM(key, val string) {
	defer r.DB.Close()
	_ , err := redis.Int(r.DB.Do("SREM", key, val))
	if err != nil{
		return
	}
	return
}

func (r *RedisClientHandler) RedisSadd(key string, args ...interface{}) {
	defer r.DB.Close()
	r.DB.Send("MULTI")
	for _,i := range args {
		r.DB.Send("sadd", key, i)
	}
	_ , err := r.DB.Do("EXEC")
	if err != nil{
		return
	}
	return
}

func (r *RedisClientHandler) RedisHGetAll(key string)(reply map[string]string, err error) {
	defer r.DB.Close()
	reply , err = redis.StringMap(r.DB.Do("hgetall", key))
	if err != nil{
		return
	}
	return
}

func (r *RedisClientHandler) RedisHMSet(key string, values map[string]interface{}){
	defer r.DB.Close()
	_ , err := r.DB.Do("hmset", redis.Args{}.Add(key).AddFlat(values)...)
	if err != nil{
		return
	}
	return
}

func (r *RedisClientHandler) RedisHExists(key string, member interface{})(reply bool, err error) {
	defer r.DB.Close()
	reply , err = redis.Bool(r.DB.Do("HEXISTS", key, member))
	if err != nil{
		return
	}
	return
}