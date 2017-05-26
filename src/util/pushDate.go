package util

import (
	"fmt"
)

func PushJKdata(pushdata string) {
	fmt.Println("push data: ", pushdata)
	PostHttpRequestText("http://10.10.1.70:30001/metrics/job/monitor_jiesuan", pushdata)
}
