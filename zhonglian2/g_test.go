package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestC(t *testing.T) {
	payload := strings.NewReader(
		`<soap:Envelope` +
			` xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">` +
			//`<soap:Header>` +
			//`<msghead>` +
			//`<appId>` + appId + `</appId>` +
			//`<appSecret>` + appSecret + `</appSecret>` +
			//`<token>` + finalToken + `</token>` +
			//`<reqTimestamp>` + reqTimestamp + `</reqTimestamp>` +
			//`<functionId>` + functionId + `</functionId>` +
			//`<signature>` + signature + `</signature>` +
			//`<version>` + version + `</version>` +
			//`</msghead>` +
			//`</soap:Header>` +
			//`<soap:Body>` +
			//`<ns2:` + functionId +
			//` xmlns:ns2="http://webservice.cfms.richinfo.cn/">` +
			//`<phone>` + phoneNum + `</phone>` +
			//`</ns2:` + functionId + `>` +
			`</soap:Body>` +
			`</soap:Envelope>`)
	fmt.Printf("%+v", payload)

	//str := GetString()
	//url := "http://127.0.0.1:8888/aaa"
	//method := "POST"
	//client := &http.Client{}
	//req, err := http.NewRequest(method, url, str)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//req.Header.Add("Content-Type", "application/xml")
	//
	//res, err := client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer res.Body.Close()
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))
}

func GetString() *strings.Reader {
	payload := strings.NewReader(
		`<soap:Envelope` +
			` xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">` +
			`<soap:Header>` +
			`<msghead>` +
			`<appId>54208878</appId>` +
			`<appSecret>e85489a5c0884495a6f87ff55589bd8</appSecret>` +
			`<token>CE0000001657613335317bc6262267f584d41b3fe72061dbcab3</token>` +
			`<reqTimestamp>1658307562</reqTimestamp>` +
			`<functionId>queryEntUserByPhone</functionId>` +
			`<signature>fsere3</signature>` +
			`<version>1.3.7</version>` +
			`</msghead>` +
			`</soap:Header>` +
			`<soap:Body>` +
			`<ns2:queryEntUserByPhone` +
			` xmlns:ns2="http://webservice.cfms.richinfo.cn/">` +
			`<phone>18860230359</phone>` +
			`</ns2:queryEntUserByPhone>` +
			`</soap:Body>` +
			`</soap:Envelope>`)
	return payload
}
