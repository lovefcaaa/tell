package model

import (
	"strings"
)

//查询订单，明文返回体
type QueryOrderRespDe struct {
	Code       string
	Data       QueryData
	Message    string
	Request_id string
}

func (q *QueryOrderRespDe) Equal(q1 QueryOrderRespDe) bool {
	if strings.EqualFold(q.Code, q1.Code) &&
		strings.EqualFold(q.Message, q1.Message) &&
		strings.EqualFold(q.Request_id, q1.Request_id) &&
		(q.Data.Equal(q1.Data)) {
		return true
	} else {
		return false
	}
}

//data数据定义
type QueryData struct {
	Amount   string
	Card_no  string
	Order_id string
	Pay      string
	Ref      string
	Status   string
}

func (d *QueryData) Equal(d1 QueryData) bool {
	if strings.EqualFold(d.Amount, d1.Amount) &&
		strings.EqualFold(d.Card_no, d1.Card_no) &&
		strings.EqualFold(d.Order_id, d1.Order_id) &&
		strings.EqualFold(d.Pay, d.Pay) &&
		strings.EqualFold(d.Ref, d.Ref) &&
		strings.EqualFold(d.Status, d.Status) {
		return true
	} else {
		return false
	}
}
