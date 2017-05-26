package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)
// http请求工具

func GetHttpRequest(url string,getdata string, headersMap map[string]string ) string{
	client := &http.Client{}
	var requestUrl string
	if len(getdata) > 0 {
		requestUrl = url+"?"+getdata
	}else {
		requestUrl = url
	}
	req, err1 := http.NewRequest("GET", requestUrl, nil)
	if err1 != nil {
		fmt.Println(err1.Error())
		return ""
	}
	if headersMap != nil {
		for k, v := range headersMap {
			req.Header.Set(k, v)
		}
	}else {
		req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	}

	resp, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2.Error())
		return ""
	}
	defer resp.Body.Close()
	fmt.Println("response status : ",resp.StatusCode)
	body, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		fmt.Println(err3.Error())
		return ""
	}
	fmt.Println("reponse body : ",string(body))
	return string(body)
	
}

func PostHttpRequest(url string, postdata string, headersMap map[string]string) string {
	client := &http.Client{}
	req, err1 := http.NewRequest("POST", url, strings.NewReader(postdata))
	if err1 != nil {
		fmt.Println(err1.Error())
		return ""
	}
	if headersMap != nil {
		for k, v := range headersMap {
			req.Header.Set(k, v)
		}
	}else {
		req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	}

	resp, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2.Error())
		return ""
	}
	fmt.Println("response status :",resp.Status)
	defer resp.Body.Close()
	body, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		fmt.Println(err3.Error())
		return ""
	}
	var bodyStr string
	bodyStr = string(body)
	fmt.Println("response body : ",bodyStr)
	return bodyStr;
}

func PostHttpRequestText(url string, postdata string) string {
	client := &http.Client{}
	req, err1 := http.NewRequest("POST", url, strings.NewReader(postdata+"\n"))
	if err1 != nil {
		fmt.Println(err1.Error())
		return ""
	}
	resp, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2.Error())
		return ""
	}
	fmt.Println("response status :",resp.Status)
	defer resp.Body.Close()
	body, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		fmt.Println(err3.Error())
		return ""
	}
	var bodyStr string
	bodyStr = string(body)
	fmt.Println("response body : ",bodyStr)
	return bodyStr;
}
