package model

import (
	"strings"
)

type PayResponse struct {
	Code       string
	Data       PayData
	Message    string
	Request_id string
}

func (pr1 *PayResponse) Equal(pr2 PayResponse) bool{
	if strings.EqualFold(pr1.Code, pr2.Code) &&
		strings.EqualFold(pr1.Message, pr2.Message) &&
		strings.EqualFold(pr1.Request_id, pr2.Request_id) &&
		pr1.Data.Equal(pr1.Data) {
		return true
	} else {
		return false
	}

}

type PayData struct {
	Order_id string
	Pay      string
	Ref      string
}

func (pd1 *PayData) Equal(pd2 PayData) bool {
	if strings.EqualFold(pd1.Order_id, pd2.Order_id) &&
		strings.EqualFold(pd1.Pay, pd2.Pay) &&
		strings.EqualFold(pd1.Order_id, pd2.Order_id) &&
		strings.EqualFold(pd1.Ref, pd2.Ref) {
		return true
	} else {
		return false
	}
}
