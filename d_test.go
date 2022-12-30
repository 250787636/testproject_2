package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
	"testing"
)

func Test_d(t *testing.T) {
	//apiKey := "741b1789-e132-4f6a-86ec-6a2f6f45b5b6"
	//apiKey := "3c7c6a00-6581-4d16-9d21-ae0245365917"
	//apiKey := "741b1789-e132-4f6a-86ec-6a2f6f45b5b6"
	//apiKey := "27fade0f-c166-497d-aa32-1fe951e8811f"
	apiKey := "829bd229-01d5-4842-9cb6-d61483da1f56"
	//apiSecret := "a0033c08-43cf-4203-9b87-1df0b8c5fe25"
	//apiSecret := "7f998490-b0c5-45e1-82d5-5091a5a42dab"
	//apiSecret := "a0033c08-43cf-4203-9b87-1df0b8c5fe25"
	//apiSecret := "64793974-e521-455e-81f1-bb1d07ed5e3d"
	apiSecret := "60c853f5-b8ed-49ea-b9bb-0fa84c0cddfe"
	//apiKey := "e9e52b98-890a-457f-83c6-7376cb32412e"
	//apiSecret := "d3f83e2a-a1c2-4739-825f-48ab802d53a1"
	//username := "testl"
	//policyConfig := "{\"addition\":{\"antixml\":false,\"compresspng\":false,\"nolog\":false},\"antiTamp\":{\"enable\":false,\"polling\":true,\"ptrace\":false},\"cpu\":[\"armeabi\",\"armeabi-v7a\",\"arm64-v8a\"],\"dataEnc\":{\"enable\":false,\"dataencxor\":false,\"dataencbind\":false,\"rule\":[]},\"dex\":{\"v1\":{\"enable\":false,\"dexfast\":true,\"fun\":[]},\"v2\":{\"enable\":false,\"dexdcrp\":true,\"dexdmrp\":false,\"fakeclass\":false,\"fun\":[]},\"v4\":{\"enable\":false,\"dexbind\":false,\"allvmp\":false,\"so\":\"\",\"fun\":[]},\"obfstr\":{\"enable\":false,\"obffilter\":[]},\"antiRe\":{\"dexhunter\":false,\"antidex2jar\":true,\"antijadx\":true,\"shelldexhelper\":true}},\"integrity\":{\"enable\":false,\"rule\":[\"*\",\"!AndroidManifest.xml\"]},\"resEnc\":{\"enable\":false,\"zipres\":false,\"rule\":[]},\"runtime\":{\"pappmo\":false,\"proot\":false,\"psimulator\":false,\"antibooster\":false,\"pmem\":false,\"hooktools\":false,\"sig\":false,\"antijnject\":true,\"hijack\":{\"enable\":false,\"activity\":[]},\"safescreen\":{\"enable\":false,\"activity\":[]}},\"so\":{\"enable\":false,\"ver\":\"\",\"bind\":false,\"clear\":false,\"enc\":false,\"file\":[]},\"u3d\":{\"enable\":false,\"version\":\"\",\"unityver\":\"\",\"dll\":[]},\"version\":\"ver7.1.2_TAND_cc_211104.1.docker\"}"
	//m := map[string]interface{}{
	////"policy_name":"中间平台通用策略2",
	////"policy_status":"已启用",
	//"username":username,
	//"policy_id":2028,
	//"upload_type":2,
	//}
	//添加用户
	//m := map[string]interface{}{
	//"username":"普通用户_3",
	//"user_type":"0",
	//"email":"1",
	//"phone":"1",
	//"is_active":true,
	//"expire_time":"2006-01-02 15:04:05",
	//"realname":"测试用户_3",
	//"company":"棒棒",
	//"company_id":"1",
	//}
	//修改用户
	//m := map[string]interface{}{
	//	"username":   "测试的",
	//	"user_type":   "4",
	//	"email":       "1",
	//	"phone":       "1",
	//	"is_active":   true,
	//	"expire_time": "2006-01-02 15:04:05",
	//	"real_name":   "测试用户",-
	//	"company":     "棒棒",
	//	"company_id":  "1",
	//}
	// 删除用户
	//m := map[string]interface{}{
	//	"username":   "测试的",
	//}

	// 天工上传h5加固任务
	//m := map[string]interface{}{
	//	//"username":"cesirenyuan-测试人员",
	//	"username":"cesirenyuan-测试人员",
	//	"app_code":"systemH5-农业银行h5加固模块",
	//	"upload_type":0,
	//	"policy_id":1,
	//	"upload_url":"https://www.python.org/ftp/python/pc/32python.zip",
	//}
	//m := map[string]interface{}{
	//	"username":    "gutianxu-谷天旭",
	//	"app_code":    "gtx_谷",
	//	"app_model":    "111",
	//	"upload_type": "2",
	//}
	m := map[string]interface{}{
		"username": "testabc",
	}
	//m := map[string]interface{}{
	//	"username":  "gutianxu-谷天旭",
	//	"h5info_id": 170,
	//}

	//天工下载h5加固包
	//m := map[string]interface{}{
	//	"h5info_id":170,
	//	"download_type":1,
	//}
	fmt.Println("sign:", hmacSha1(apiSecret, concatParam(m, apiKey)))
}

func hmacSha1(secret, text string) string {
	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write([]byte(text))
	return hex.EncodeToString(mac.Sum(nil))
}

func concatParam(m map[string]interface{}, apiKey string) string {
	result := apiKey
	keyList := make([]string, 0)
	for k, _ := range m {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		result = result + fmt.Sprintf("%v", m[k])
	}
	return strings.Trim(result, "&")
}
