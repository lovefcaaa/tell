package asse

import (
	"model"
	"strings"
)

func IsOrderSuccess(queryOrderResp model.QueryOrderRespDe, pay string, order_id string) bool {
	if strings.EqualFold(queryOrderResp.Code, "0000") &&
		strings.EqualFold(queryOrderResp.Message, "操作成功") &&
		strings.EqualFold(queryOrderResp.Data.Amount, pay) &&
		strings.EqualFold(queryOrderResp.Data.Order_id, order_id) &&
		strings.EqualFold(queryOrderResp.Data.Status, "已打款") {
		return true
	} else {
		return false
	}
}

func IsSuccess(payResp model.PayResponse, pay string, order_id string) bool {
	if strings.EqualFold(payResp.Code, "0000") &&
		strings.EqualFold(payResp.Data.Order_id, order_id) &&
		strings.EqualFold(payResp.Data.Pay, pay) &&
		strings.EqualFold(payResp.Message, "操作成功") {
		return true
	} else {
		return false
	}
}