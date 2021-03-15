package logging

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/librarys"
	"github.com/EDDYCJY/go-gin-example/pkg/initconfig"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Level int

var (
	F *os.File
	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// getLogFilePath get the log file save path
func GetLogFilePath() string {
	return fmt.Sprintf("%s%s", initconfig.AppConfig.RuntimeRootPath, initconfig.AppConfig.LogSavePath)
}

// getLogFileName get the save name of the log file
func GetLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		initconfig.AppConfig.LogSaveName,
		time.Now().Format(initconfig.AppConfig.TimeFormat),
		initconfig.AppConfig.LogFileExt,
	)
}
// Setup initialize the log instance
func InitLog() {
	var err error
	filePath := GetLogFilePath()
	fileName := GetLogFileName()
	F, err = librarys.FileMustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// Debug output logs at debug level
func LogDebug(v ...interface{}) {
	LogsetPrefix(DEBUG)
	logger.Println(v)
}

// Info output logs at info level
func LogInfo(v ...interface{}) {
	LogsetPrefix(INFO)
	logger.Println(v)
}

// Warn output logs at warn level
func LogWarn(v ...interface{}) {
	LogsetPrefix(WARNING)
	logger.Println(v)
}

// Error output logs at error level
func LogError(v ...interface{}) {
	LogsetPrefix(ERROR)
	logger.Println(v)
}

// Fatal output logs at fatal level
func LogFatal(v ...interface{}) {
	LogsetPrefix(FATAL)
	logger.Fatalln(v)
}

// setPrefix set the prefix of the log output
func LogsetPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
