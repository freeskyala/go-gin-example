package initstart

import (

	logics "github.com/EDDYCJY/go-gin-example/logics"

)
//定时任务
type CronTask struct {
}

func (c *CronTask) InitCrond() {


	//tagService := logic.Crond{}
	//tagService.CrondTest()
	go new(logics.Crond).CrondTest()

	/*
	go new(services.HomeDialogHandler).HomeDialogMQConsumer()
	//翻译Job
	go new(services.GoogleTransService).DoTransTask()
	//同步 redis cache 至 db
	go new(services.SyncCacheTask).DoCacheTask()
*/

}

