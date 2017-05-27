package util

import (
	"fmt"
	"config"
)

//向普罗米修斯推数据
func PushJKdata(pushdata string) {
	fmt.Println("push data: ", pushdata)
	_,status :=PostHttpRequestText(config.PrometheusURL+"/metrics/job/monitor_jiesuan", pushdata)
	if status - 202 != 0 {
		fmt.Println("[INFO]数据推送Prometheus失败")
	}else{
		fmt.Println("[INFO]数据推送Prometheus成功")
	}
}
