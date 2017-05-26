package util

import (
	"fmt"
	"config"
)

//向普罗米修斯推数据
func PushJKdata(pushdata string) {
	fmt.Println("push data: ", pushdata)
	PostHttpRequestText(config.PrometheusURL+"/metrics/job/monitor_jiesuan", pushdata)
}
