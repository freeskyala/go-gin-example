package initstart

import (
	global "github.com/EDDYCJY/go-gin-example/pkg"
	"github.com/EDDYCJY/go-gin-example/pkg/initconfig"
	//"github.com/EDDYCJY/go-gin-example/pkg/initconfig"
	elastic "github.com/olivere/elastic/v6"
	"net"
	"net/http"
	"sync"
	"time"
)

var TEST_ES_STATUS = false

type ElasticSearch struct {

}

func (this * ElasticSearch) InitDefaultEs() {
	defer func() {
		if err := recover(); err != nil {
			//TODO panic happened, need log
		}
	}()
	//加锁，防止多个对象调用初始化方法
	lock := sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()
	if TEST_ES_STATUS {
		return
	}
	maxIdle := initconfig.ElasticConfig.MaxIdleConns
	idleTimeout := initconfig.ElasticConfig.MaxIdleConns
	maxConnPerHost := initconfig.ElasticConfig.MaxConnsPerHost
	//修改底层net/http配置，避免反复创建连接，导致TIME_WAIT问题
	httpClient := &http.Client{}
	httpClient.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 60 * time.Second,
		}).DialContext,
		MaxIdleConns:          maxIdle,
		MaxIdleConnsPerHost:   maxIdle,
		MaxConnsPerHost:       maxConnPerHost,
		IdleConnTimeout:       time.Duration(idleTimeout) * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	//创建连接对象
	client, _ := elastic.NewClient(elastic.SetHttpClient(httpClient),
		elastic.SetSniff(false),
		elastic.SetURL(initconfig.ElasticConfig.Host),
		elastic.SetHealthcheck(false),
		//设置错误日志输出
		//elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		////设置info日志输出
		//elastic.SetInfoLog(log.New(os.Stdout, "Goods ES Info log: ", log.LstdFlags)),
		////设置trace日志输出
		//elastic.SetTraceLog(log.New(os.Stdout, "Goods ES Trace log: ", log.LstdFlags)),
	)
	global.ES = client

	TEST_ES_STATUS = true
}

