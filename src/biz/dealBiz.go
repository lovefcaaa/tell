package biz

import (
	"config"
	"fmt"
	"service"
	"strings"
	"time"
	"util"
	"encoding/json"
	"encoding/base64"
	"model"
)

// 测试流执行：组装测试数据、发送请求、判断返回体，具体调用服务层实现

func DealBiz(bizmodel string, env string, deskey string, dealer string, broker string, pay string) (string,string) {
	var dataStr, api, timestampStr, orderId string
	if strings.EqualFold(bizmodel, config.ModelRealTime) {
		dataStr, api, timestampStr, orderId = getRealtimeBizData(dealer, broker, pay)
	}
	if strings.EqualFold(bizmodel, config.ModelWXPay) {
		dataStr, api, timestampStr, orderId = getWXpayData(dealer, broker, pay)
	}
	if strings.EqualFold(bizmodel, config.ModelSixcn) {
		dataStr, api, timestampStr, orderId = getSixcnBizData(dealer, broker, pay)
	}
	if strings.EqualFold(bizmodel, config.ModelAliPay) {
		dataStr, api, timestampStr, orderId = getAlipayData(dealer, broker, pay)
	}
	fmt.Println("request data: ", dataStr)
	var m1 map[string]string
	m1 = make(map[string]string)
	m1["Content-Type"] = "application/x-www-form-urlencoded"
	m1["request-id"] = config.RequestIDInHttpHeader
	m1["dealer-id"] = dealer
	respBodyStr := service.SendPostRquest(env, deskey, dataStr, timestampStr, api, orderId,m1)
	return respBodyStr,orderId
}

func HealthCheck(url string) {
	util.GetHttpRequest(url, "", nil)
}

func QueryOrder(env string, deskey string, orderId string, api string, bizmodel string,dealer string)(string,string) {
	timestamp := time.Now().Unix()
	timestampStr := fmt.Sprintf("%d", timestamp)
	var dataStr0 string
	var dataStr1 string
	if len(bizmodel) > 0 {
		if strings.EqualFold(bizmodel, config.ModelAliPay) {
			dataStr0 = `{"order_id": "` + orderId + `","channel":"支付宝","data_type":"encryption"}`
			dataStr1 = `{"order_id": "` + orderId + `","channel":"支付宝"}`
		} else if strings.EqualFold(bizmodel, config.ModelRealTime) {
			dataStr0 = `{"order_id": "` + orderId + `","channel":"银行卡","data_type":"encryption"}`
			dataStr1 = `{"order_id": "` + orderId + `","channel":"银行卡"}`
		} else if strings.EqualFold(bizmodel, config.ModelWXPay) {
			dataStr0 = `{"order_id": "` + orderId + `","channel":"微信","data_type":"encryption"}`
			dataStr1 = `{"order_id": "` + orderId + `","channel":"微信"}`
		} else {
			dataStr0 = `{"order_id": "` + orderId + `","data_type":"encryption"}`
			dataStr1 = `{"order_id": "` + orderId + `"}`
		}
	} else {
		dataStr0 = `{"order_id": "` + orderId + `","data_type":"encryption"}`
		dataStr1 = `{"order_id": "` + orderId + `"}`
	}
	var m1 map[string]string
	m1 = make(map[string]string)
	m1["Content-Type"] = "application/x-www-form-urlencoded"
	m1["request-id"] = config.RequestIDInHttpHeader
	m1["dealer-id"] = dealer
	fmt.Println("query data : "+dataStr0)
	respStr0 := service.SendQueryOrderRquest(env, deskey, dataStr0, timestampStr, api,m1)
	if strings.Contains(respStr0, "0000") {
		var respData model.QueryOrderResp
		json.Unmarshal([]byte(respStr0), &respData)
		decrypt(respData.Data,deskey)
	}
	
	fmt.Println("query data : "+dataStr1)
	respStr1 := service.SendQueryOrderRquest(env, deskey, dataStr1, timestampStr, api,m1)
	return respStr0,respStr1
}

//实时打款接口测试数据
func getRealtimeBizData(dealer string, broker string, pay string) (string, string, string, string) {
	timestamp := time.Now().Unix()
	timestampStr := fmt.Sprintf("%d", timestamp)
	orderId := "WZYTEST" + timestampStr
	dataStr := `{"pay": "` + pay + `","order_id": "` + orderId + `","reserve_id": "49892543 201703-1","dealer_id": "` + dealer + `","broker_id": "` + broker + `","anchor_id": "49892545","real_name": "魏增艺", "card_no": "6214850101469267",  "phone_no": "18701433307","id_card": "350403198801131017", "is_company":0,"is_abroad":0,"is_bankbook":0}`
	api := config.RealTimeAPI
	return dataStr, api, timestampStr, orderId
}

//六间房模式打款测试数据
func getSixcnBizData(dealer string, broker string, pay string) (string, string, string, string) {
	timestamp := time.Now().Unix()
	timestampStr := fmt.Sprintf("%d", timestamp)
	//	orderId := "WZYTEST" + timestampStr
	dataStr := `{"pay": "` + pay + `","order_id": "` + "201704170066277439135" + `","reserve_id": "49892543 201703-1","dealer_id": "` + dealer + `","broker_id": "` + broker + `","anchor_id": "66277439","real_name": "姜宇", "card_no": "6217000990006452169","phone_no": "15245847317","id_card": "230224198808243912","verify_mode": "sixcn","detail_id":"20170410066277439001","trans_batch_id": "N20170417152915","notes":"66277439 201704-1","is_company":0,"is_abroad":0,"is_bankbook":0,"issue_id":"20170419"}`
	api := config.SixcnAPI
	return dataStr, api, timestampStr, "201704170066277439135"
}

//支付宝打款测试数据
func getAlipayData(dealer string, broker string, pay string) (string, string, string, string) {
	timestamp := time.Now().Unix()
	timestampStr := fmt.Sprintf("%d", timestamp)
	orderId := "WZYTEST" + timestampStr
	dataStr := `{"pay": "` + pay + `","order_id": "` + orderId + `", "dealer_id": "` + dealer + `","broker_id": "` + broker + `","real_name": "魏增艺","card_no": "weizengyipp@gmail.com","id_card": "350403198801131017"}`
	api := config.AlipayAPI
	return dataStr, api, timestampStr, orderId
}

//微信红包打款测试数据
func getWXpayData(dealer string, broker string, pay string) (string, string, string, string) {
	timestamp := time.Now().Unix()
	timestampStr := fmt.Sprintf("%d", timestamp)
	orderId := "WZYTEST" + timestampStr
	dataStr := `{"pay": "` + pay + `","order_id": "` + orderId + `", "dealer_id": "` + dealer + `","broker_id": "` + broker + `","real_name": "魏增艺","openid": "`+config.WXOPIDWZY+`","comment":"增艺测试微信","id_card":"350403198801131017"}`
	api := config.WXPayAPI
	return dataStr, api, timestampStr, orderId
}

func decrypt(datastr string,deskey string){
	base64De,err := base64.StdEncoding.DecodeString(datastr)
     if err != nil{
     	fmt.Println(err.Error())
     }
     keyByte := []byte(deskey)
//     fmt.Println(base64De)
//     fmt.Println(keyByte)
     org,_ := util.TripleDesDecrypt(base64De, keyByte)
     fmt.Println("decrypted data : ",byteString(org))
}

func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}
