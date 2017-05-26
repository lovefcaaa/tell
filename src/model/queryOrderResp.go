package model

import (
	"fmt"
	"strings"
)

//查询订单接口，密文返回体
type QueryOrderResp struct {
	Code       string
	Data       string
	Message    string
	Request_id string
}

func (q *QueryOrderResp) Equal(q1 QueryOrderResp) bool {
	if strings.EqualFold(q.Code, q1.Code) &&
		strings.EqualFold(q.Data, q1.Data) &&
		strings.EqualFold(q.Message, q1.Message) &&
		strings.EqualFold(q.Request_id, q1.Request_id) {
		return true
	} else {
		return false
	}
}
func (q *QueryOrderResp) printData() {
	fmt.Println(q.Code)
	fmt.Println(q.Data)
	fmt.Println(q.Message)
	fmt.Println(q.Request_id)
}
