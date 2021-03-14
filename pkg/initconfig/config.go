package initconfig

import (
	lib "github.com/EDDYCJY/go-gin-example/pkg/librarys"
	"github.com/go-ini/ini"
	"log"
	"time"
)
type DataTest struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}
var DataTestConfig = &DataTest{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerConfig = &Server{}


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

var cfg *ini.File
// Setup initialize the configuration instance
func InitConfig() {
	cfg, _ = ini.Load(".env")
	mapTo("database_test", DataTestConfig)
	mapTo("server", ServerConfig)
	mapTo("app", AppConfig)
	lib.P(DataTestConfig)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
