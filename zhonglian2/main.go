package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/aaa", sso)
	r.Run(":8888")
}

//type Envelope struct {
//	XMLName xml.Name `xml:"Envelope"`
//	Soap    string   `xml:"soap,attr"`
//	Header  struct {
//		Msghead struct {
//			AppId        string `xml:"appId"`
//			AppSecret    string `xml:"appSecret"`
//			Token        string `xml:"token"`
//			ReqTimestamp string `xml:"reqTimestamp"`
//			FunctionId   string `xml:"functionId"`
//			Signature    string `xml:"signature"`
//			Version      string `xml:"version"`
//		} `xml:"msghead"`
//	} `xml:"Header"`
//	Body struct {
//		QueryEntUserByPhone struct {
//			Ns2   string `xml:"ns2,attr"`
//			Phone string `xml:"phone"`
//		} `xml:"queryEntUserByPhone"`
//	} `xml:"Body"`
//}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Soap    string   `xml:"soap,attr"`
	Body    struct {
		QueryEntUserByPhoneResponse struct {
			Ns2    string `xml:"ns2,attr"`
			Return struct {
				Code string `xml:"code"`
				Data struct {
					EntUser struct {
						DepartMentId   string `xml:"departMentId"`
						DepartMentName string `xml:"departMentName"`
						Email          string `xml:"email"`
						IsValid        string `xml:"isValid"`
						MobilePhone    string `xml:"mobilePhone"`
						OrgFullName    string `xml:"orgFullName"`
						OrgId          string `xml:"orgId"`
						OrgName        string `xml:"orgName"`
						PortalUserId   string `xml:"portalUserId"`
						PortalUserName string `xml:"portalUserName"`
						Position       string `xml:"position"`
						Sex            string `xml:"sex"`
						UserCategory   int    `xml:"userCategory"`
						UserId         string `xml:"userId"`
					} `xml:"entUser"`
				} `xml:"data"`
				Msg string `xml:"msg"`
			} `xml:"return"`
		} `xml:"queryEntUserByPhoneResponse"`
	} `xml:"Body"`
}

func sso(r *gin.Context) {
	//var c Envelope
	//all, _ := ioutil.ReadAll(r.Request.Body)
	//err := xml.Unmarshal(all, &c)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(c)
	//r.JSON(200, gin.H{
	//	"msg": "请求成功",
	//})
	var en Envelope
	data := GetData()
	err := xml.Unmarshal(data, &en)
	if err != nil {
		fmt.Println(err)
	}
	r.XML(200, en)
}

func GetData() []byte {
	payload := []byte(
		`<soap:Envelope` +
			` xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">` +
			`<soap:Body>` +
			`<ns2:queryEntUserByPhoneResponse` +
			` xmlns:ns2="http://webservice.cfms.richinfo.cn/">` +
			`<return>` +
			` <code>0</code>` +
			`<data>` +
			`<entUser>` +
			`<departMentId>775629752</departMentId>` +
			` <departMentName>彩讯科技股份有限公司</departMentName>` +
			`<email>yejunpei@richinfo.cn</email>` +
			`<isValid>1</isValid>` +
			`<mobilePhone>18860230359</mobilePhone>` +
			`<orgFullName>彩讯科技股份有限公司</orgFullName>` +
			`<orgId>775629752</orgId>` +
			`<orgName>彩讯科技股份有限公司</orgName>` +
			`<portalUserId>dwyejunpei</portalUserId>` +
			`<portalUserName>叶君培</portalUserName>` +
			`<position>1</position>` +
			`<sex>1</sex>` +
			`<userCategory>2</userCategory>` +
			`<userId>10100</userId>` +
			`</entUser>` +
			`</data>` +
			`<msg>根据手机号查询实体（自有或合作伙伴）人员信息成功</msg>` +
			`</return>` +
			`</ns2:queryEntUserByPhoneResponse>` +
			`</soap:Body>` +
			`</soap:Envelope>`)
	return payload
}
