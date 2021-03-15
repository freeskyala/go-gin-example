package initconfig

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)


type App struct {
	PageSize  int
	PrefixUrl string
	RuntimeRootPath string
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppConfig = &App{}


type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerConfig = &Server{}


type DataTest struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}
var DataTestConfig = &DataTest{}


type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisConfig = &Redis{}



type Elasti struct {
	Host        string
	MaxIdleConns     int
	IdleConnTimeout   int
	MaxConnsPerHost   int
}

var ElasticConfig = &Elasti{}

var cfg *ini.File
// Setup initialize the configuration instance
func InitConfig() {
	cfg, _ = ini.Load(".env")
	mapTo("database_test", DataTestConfig)
	mapTo("server", ServerConfig)
	mapTo("app", AppConfig)
	mapTo("redis", RedisConfig)
	mapTo("elastic", ElasticConfig)
	ServerConfig.ReadTimeout = ServerConfig.ReadTimeout * time.Second
	ServerConfig.WriteTimeout = ServerConfig.WriteTimeout * time.Second
	//lib.P(DataTestConfig)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
