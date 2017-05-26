package config

import (

)

//通用配置数据常量

const(
	
//	商户签名密钥
	Signkey = "uuu9b4fad7894fbce1df0b30eee48766"
	
//	商户业务数据加密密钥
	Deskey  = "877432188765432112345678"
	
//	dealer
	DealerTestDealer = "testdealer"
	DealerTestDealer007 = "testdealer007"
	DealerTestDealer008 = "testdealer008"
	DealerTestDealer009 = "testdealer009"
	
//	http header
	DealerInHttpHeader = "testdealer"
	RequestIDInHttpHeader = "mengmengda"
	
//	broker
	Broker="testbroker"
	Broker007="testbroker007"
	
//	打款模式
	ModelRealTime = "RealTime"
	ModelWXPay = "Wxpay"
	ModelAliPay = "Alipay"
	ModelSixcn = "Sixcn"    
	
//	测试环境
	ENVLocal = "local"
	ENVPreTest = "pretest"
	ENVOther = "other"
	
	
//	测试环境URL前缀
	LocalURLPre = "http://api.jiesuan.local:8083"
	PreTestURLPre = "https://api-jiesuan-stagexx.yunzhanghu.com"
	Product = "https://api-jiesuan.yunzhanghu.com"
	Other = "http://172.16.1.215:8082"
	
	
//	接口地址
	RealTimeAPI = "/api/payment/v1/order-realtime"
	SixcnAPI = "/api/payment/v1/order-realtime"
	AlipayAPI = "/api/payment/v1/order-alipay"
	WXPayAPI = "/api/payment/v1/order-wxpay"
	QueryOrderAPI = "/api/payment/v1/query-realtime-order"
	HealthCheck = "/api/local/health"
	
	UAT = "uat"
	RT ="rt"
	
	WXOPIDXF = "og3_ZwB1PRLMRiaA4JNme1U0K1UY"
	WXOPIDZYQ = "og3_ZwPeLen47Ehy-6h8mod3VyQQ"
	WXOPIDWZY = "og3_ZwMJdprTD0k2P3gF-86Bn68o"
	
//	查询接口场景
	QueryAlipay = "QueryAlipay"
	QueryRealtime = "QueryRealtime"
	QueryWXpay = "QueryWXpay"
	RealtimeWithoutChannel="RealtimeWithoutChannel"
	QueryAlipayWithoutChannel = "QueryAlipayWithoutChannel"
	QueryWXpayWithoutChannel = "QueryWXpayWithoutChannel"
)

