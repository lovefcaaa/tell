package service

import (
	"fmt"
	"net/http"
	"strings"
	"util"
	"config"
)

//服务层：加密数据、组装请求URL，发送交易请求、判断返回体

func SendPostRquest(env string, deskey string,dataStr string, timestampStr string,api string,orderId string,httpHeader map[string]string) string{
	var requestURLPre string
	requestURLPre = getTestURLPreL(env)
	requestURL := requestURLPre + api
	bodystr := encrypted(deskey, dataStr, timestampStr)
	var respBodyStr string
	respBodyStr = sendPost(requestURL, bodystr,httpHeader)
	return respBodyStr
}

func SendQueryOrderRquest(env string, deskey string,dataStr string, timestampStr string,api string, httpHeader map[string]string) string{
	var requestURLPre string
	requestURLPre = getTestURLPreL(env)
	requestURL := requestURLPre + api
	bodystr := encrypted(deskey, dataStr, timestampStr)
	var respBodyStr string
	respBodyStr = sendGet(requestURL, bodystr,httpHeader)
	return respBodyStr
}



func encrypted(deskey string,dataStr string, timestampStr string) string {
	origData := []byte(dataStr)
	var base64Ecode, sign string
	base64Ecode, sign = util.BaseSignYHKP(deskey,origData, timestampStr)
	fmt.Println("base64Ecode : " + base64Ecode)
	fmt.Println("sign : " + sign)
	fmt.Println("timestamp : " + timestampStr)
	var r http.Request
	r.ParseForm()
	r.Form.Add("data", base64Ecode)
	r.Form.Add("sign", sign)
	r.Form.Add("timestamp", timestampStr)
	r.Form.Add("mess", deskey)
	r.Form.Add("sign_type", "sha256")
	bodystr := strings.TrimSpace(r.Form.Encode())
	fmt.Println("bosystr : ",bodystr)
	return bodystr
}

func sendPost(requestURL string, bodystr string,m1 map[string]string) string{
	fmt.Println("request URL : " + requestURL)
	var respBodyStr string
	respBodyStr = util.PostHttpRequest(requestURL, bodystr, m1)
	return respBodyStr
}

func sendGet(requestURL string, bodystr string,m1 map[string]string) string{
	fmt.Println("request URL : " + requestURL)
	var respBodyStr string
	respBodyStr = util.GetHttpRequest(requestURL, bodystr, m1)
	return respBodyStr
}

func getTestURLPreL(env string) string {
	var requestURLPre string
	if strings.EqualFold(config.ENVLocal, env) {
		requestURLPre = config.LocalURLPre
	} else if strings.EqualFold(config.ENVPreTest, env) {
		requestURLPre = config.PreTestURLPre
	}else if strings.EqualFold(config.Product, env) {
		requestURLPre = config.Product
	}else {
		requestURLPre = config.Other
	}
	return requestURLPre
}

func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}
