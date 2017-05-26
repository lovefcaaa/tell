package util

import (
	"fmt"
)

//向普罗米修斯推数据
func PushJKdata(pushdata string) {
	fmt.Println("push data: ", pushdata)
	PostHttpRequestText("http://10.10.1.70:30001/metrics/job/monitor_jiesuan", pushdata)
}
