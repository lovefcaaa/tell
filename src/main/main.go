package main

import (
	"biz"
	"config"
	"encoding/json"
	"fmt"
	"model"
	"time"
	"asse"
	"util"
)

func main() {
	fmt.Println("app: JiesuanTestEntrance")

	testmodel := config.UAT
	//	testmodel := config.RT

//	env := config.ENVLocal
//		env := config.ENVPreTest
		env := config.Product
	//	env := config.ENVOther

	dealer := config.DealerTestDealer
	//	dealer := config.DealerTestDealer007
	//	dealer := config.DealerTestDealer008
	//	dealer := config.DealerTestDealer009

	broker := config.Broker
	//	broker := config.Broker007

	pay := "0.01"

	fmt.Println("testmodel: ", testmodel)
	fmt.Println("env: ", env)
	fmt.Println("dealer: ", dealer)
	fmt.Println("broker: ", broker)
	fmt.Println("amount: ", pay)
    
	uat(config.ModelRealTime, env, config.Deskey, dealer, broker, pay)
	uat(config.ModelAliPay, env, config.Deskey, dealer, broker, pay)
}

func uat(bizmodel string, env string, key string, dealer string, broker string, pay string) {
	payRespStr, order_id := biz.DealBiz(bizmodel, env, key, dealer, broker, pay)
	var payResp model.PayResponse
	json.Unmarshal([]byte(payRespStr), &payResp)
	if asse.IsSuccess(payResp, pay, order_id) {
		fmt.Println("[INFO]订单接收成功，验证通过")
		util.PushJKdata("PayResp"+bizmodel+"S 1")
	} else {
		fmt.Println("[INFO]订单接收验证失败, ",bizmodel)
		util.PushJKdata("PayResp"+bizmodel+"F 1")
		return
	}
	for i := 0; i < 4; i++ {
		time.Sleep(3000*time.Millisecond)
		_, queryRespStr := biz.QueryOrder(env, key, order_id, config.QueryOrderAPI, bizmodel, dealer)
		var queryOrderResp model.QueryOrderRespDe
		json.Unmarshal([]byte(queryRespStr), &queryOrderResp)
		if asse.IsOrderSuccess(queryOrderResp, pay, order_id) {
			fmt.Println("[INFO]订单支付成功，验证通过")
			util.PushJKdata("QueryResp"+bizmodel+"S 1")
			break
		}else{
			fmt.Println("第",i+1,"次查询验证失败，查询返回的结果是：",queryRespStr)
			if i==3 {
				fmt.Println("4次查询返回的结果是订单都没有打款成功, ",bizmodel)
				util.PushJKdata("QueryResp"+bizmodel+"F 1")
			}
		}
	}
}


