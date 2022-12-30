package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestZ(t *testing.T) {
	//var Envelope struct {
	//	XMLName xml.Name `xml:"Envelope"`
	//	Soap    string   `xml:"soap,attr"`
	//	Body    struct {
	//		QueryEntUserByPhoneResponse struct {
	//			Ns2    string `xml:"ns2,attr"`
	//			Return struct {
	//				Code string `xml:"code"`
	//				Data struct {
	//					EntUser struct {
	//						DepartMentId   string `xml:"departMentId"`
	//						DepartMentName string `xml:"departMentName"`
	//						Email          string `xml:"email"`
	//						IsValid        string `xml:"isValid"`
	//						MobilePhone    string `xml:"mobilePhone"`
	//						OrgFullName    string `xml:"orgFullName"`
	//						OrgId          string `xml:"orgId"`
	//						OrgName        string `xml:"orgName"`
	//						PortalUserId   string `xml:"portalUserId"`
	//						PortalUserName string `xml:"portalUserName"`
	//						Position       string `xml:"position"`
	//						Sex            string `xml:"sex"`
	//						UserCategory   int    `xml:"userCategory"`
	//						UserId         string `xml:"userId"`
	//					} `xml:"entUser"`
	//				} `xml:"data"`
	//				Msg string `xml:"msg"`
	//			} `xml:"return"`
	//		} `xml:"queryEntUserByPhoneResponse"`
	//	} `xml:"Body"`
	//}

	var Html struct {
		XMLName  xml.Name `xml:"html"`
		Xmlns    string   `xml:"xmlns,attr"`
		Envelope struct {
			Soap string `xml:"soap,attr"`
			Body struct {
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
								UserCategory   string `xml:"userCategory"`
								UserId         string `xml:"userId"`
							} `xml:"entUser"`
						} `xml:"data"`
						Msg string `xml:"msg"`
					} `xml:"return"`
				} `xml:"queryEntUserByPhoneResponse"`
			} `xml:"Body"`
		} `xml:"Envelope"`
	}
	file, _ := ioutil.ReadFile("a.xml")
	err := xml.Unmarshal(file, &Html)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", Html)
}
