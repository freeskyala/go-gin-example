package logics

import (
	"encoding/json"
	"github.com/EDDYCJY/go-gin-example/pkg/help"
	"github.com/EDDYCJY/go-gin-example/pkg/redisClient"
	"time"
)

type Crond struct {


}

func (t *Crond) CrondTest()  {
	var homeDialogMQKey string = "quarantineSignTipsMQ"
	redisClient := new(redisClient.RedisClientHandler)
	/*
	producer, _ := json.Marshal(map[string]interface{}{
		"type":     "newVersion",
 		"time":     time.Now().Unix(),
	})
	redisClient.RedisSelect(0).RedisRPush(homeDialogMQKey, string(producer))
	 */
	for {
		msgJson, err := redisClient.RedisSelect(0).RedisLPop(homeDialogMQKey)
		if err != nil {
			help.P("队列无消息:")
			time.Sleep(1 * time.Second)
			continue
		}
		go t.SendMsgToYunxin(msgJson)
	}
}
func (t *Crond) SendMsgToYunxin(msgJson string) {
	var data = make(map[string]interface{}, 0)
	json.Unmarshal([]byte(msgJson), &data)
	help.P("ssss",data)
}

