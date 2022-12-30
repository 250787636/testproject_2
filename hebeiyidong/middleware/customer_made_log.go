package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"os"
	"strings"
	"time"
)

type LogValue struct {
	ID          string `json:"id" comment:"唯一标识"`
	TimeStamp   string `json:"time_stamp" comment:"日期"`
	IsSuccess   string `json:"is_success" comment:"响应是否成功"`
	Account     string `json:"account" comment:"登录账号"`
	ServerIpv4  string `json:"server_ipv4" comment:"服务端Ipv4"`
	ServerIpv6  string `json:"server_ipv6" comment:"服务端Ipv6"`
	AccountIpv4 string `json:"account_ipv4" comment:"用户端Ipv4"`
	SystemName  string `json:"system_name" comment:"系统名称"`
	LogsType    int    `json:"logs_type" comment:"日志类型 1:登录日志，2:操作日志"`
	LoginType   int    `json:"login_type" comment:"登录类型 1:登入，2:登出"`
	UseModule   string `json:"use_module" comment:"操作模块"`
	UseType     int    `json:"use_type" comment:"操作类型"`
	UseContent  string `json:"use_content" comment:"操作内容"`
	UseResult   string `json:"use_result" comment:"操作结果"`
}

var serviceIps = map[string]interface{}{
	"ServerIpv4": "172.27.223.202",
	"ServerIpv6": "fe70::e0ac:28be:asdf:2c54%8",
	"SystemName": "测评平台",
}

var useTypesMap = map[string]interface{}{
	"增加":   1,
	"删除":   2,
	"修改":   3,
	"查询":   4,
	"其他":   5,
	"复制":   6,
	"上传":   7,
	"下载":   8,
	"导入":   9,
	"导出":   10,
	"会话跟踪": 11,
	"登录":   12,
	"登出":   13,
}

// 操作类型数组
var addArray = []string{"CREATE", "ADD"}
var delArray = []string{"DELETE"}
var updateArray = []string{"UPDATE", "SAVE", "CHANGE", "EDIT", "SET"}
var listArray = []string{"LIST", "QUEUE", "GET", "HOMEPAGE", "DETAIL", "QUERY"}
var uploadArray = []string{"UPLOAD"}
var downloadArray = []string{"DETECT", "DOWNLOAD"}
var importArray = []string{"IMPORT"}
var exportArray = []string{"EXPORT", "SEARCH"}

// 模块统计数组
var useModuleMap = map[string]interface{}{
	"/TASK":     "检测任务",
	"/TEMPLATE": "任务模板",
	"/REPORT":   "任务报告",
	"/LIC":      "系统授权",
	"/SSO":      "定制",
}

func GenerationLogs(c *gin.Context) {
	c.Next()
	// 获取状态码
	statusCode := c.Writer.Status()
	// 获取请求接口
	fullPath := c.FullPath()
	var this LogValue
	this.ID = uuid.New().String()
	this.TimeStamp = time.Now().Format("20060102150405")
	this.Account = "张三"
	this.ServerIpv4 = serviceIps["ServerIpv4"].(string)
	this.ServerIpv6 = serviceIps["ServerIpv6"].(string)
	this.AccountIpv4 = c.ClientIP()
	this.SystemName = serviceIps["SystemName"].(string)
	this.LogsType, this.LoginType = logTypeFilter(fullPath)
	this.UseType, this.UseModule, this.UseContent, this.UseResult, this.IsSuccess = getUseData(this.LogsType, statusCode, fullPath)

	context := makeContext(this)
	creSuccess, timeStr := createFileContext(context)
	if creSuccess {
		fmt.Println("已记录日志" + timeStr)
	}
}

// 根据请求接口号 判断 日志类型 || 登录类型
func logTypeFilter(str string) (logsType, loginType int) {
	isLog := strings.Contains(strings.ToUpper(str), "LOG")
	isLogOut := strings.Contains(strings.ToUpper(str), "LOGOUT")
	if isLog {
		// "登录日志"
		logsType = 1
		// 登入
		loginType = 1
		if isLogOut {
			//登出
			loginType = 2
		}
	} else {
		// 操作日志
		logsType = 2
		loginType = 0
	}
	return logsType, loginType
}

// 操作类型 操作内容 操作结果 响应信息
func getUseData(logsType, statusCode int, fullpath string) (useType int, useModule, useContent, useResult, isSuccess string) {
	// 状态码控制响应信息
	if statusCode == 200 {
		isSuccess = "成功"
		useResult = "成功"
	} else {
		isSuccess = "失败"
		useResult = "失败"
	}

	// 登录日志 无需操作结果信息所以置空
	if logsType == 1 {
		useModule = ""
		useType = 0
		useContent = ""
		useResult = ""
	} else {
		// 操作日志 响应信息非必填所以置空
		useModule = getUseModule(fullpath)
		useType = useTypesMap[getUseType(fullpath)].(int)
		useContent = useModule + getUseType(fullpath)
		isSuccess = ""
	}

	return useType, useModule, useContent, useResult, isSuccess
}

// 获取操作模块
func getUseModule(fullPath string) string {
	for k, v := range useModuleMap {
		if strings.Contains(strings.ToUpper(fullPath), k) {
			return v.(string)
		}
	}
	return "其他模块"
}

// 获取操作类型
func getUseType(fullPath string) string {
	// 不同数组中进行查询
	for _, v := range addArray {
		if strings.Contains(strings.ToUpper(fullPath), v) {
			return "增加"
		}
	}
	for _, v := range delArray {
		if strings.Contains(strings.ToUpper(fullPath), v) {
			return "删除"
		}
	}
	for _, v := range updateArray {
		if strings.Contains(strings.ToUpper(fullPath), v) {
			return "修改"
		}
	}

	for _, v := range listArray {
		if strings.Contains(strings.ToUpper(fullPath), v) {
			return "查询"
		}
	}
	for _, v := range uploadArray {
		if strings.Contains(strings.ToUpper(fullPath), v) {
			return "上传"
		}
	}

	for _, v := range downloadArray {
		if strings.Contains(strings.ToUpper(fullPath), v) {
			return "下载"
		}
	}

	for _, v := range importArray {
		if strings.Contains(strings.ToUpper(fullPath), v) {
			return "导入"
		}
	}
	for _, v := range exportArray {
		if strings.Contains(strings.ToUpper(fullPath), v) {
			return "导出"
		}
	}
	return "其他"
}

// 生成获取的定制日志数据
func makeContext(this LogValue) (context string) {
	// 登录日志
	if this.LogsType == 1 {
		context = fmt.Sprintf("%s|%s|%s|%s|%s|%s|%d|%s|%d|%s",
			this.ID, this.TimeStamp, this.Account, this.ServerIpv4, this.ServerIpv6, this.AccountIpv4,
			this.LogsType, this.SystemName, this.LoginType, this.IsSuccess)
	} else {
		//操作日志
		context = fmt.Sprintf("%s|%s|%s|%s|%s|%s|%d|%s|%s|%d|%s|%s",
			this.ID, this.TimeStamp, this.Account, this.ServerIpv4, this.ServerIpv6, this.AccountIpv4,
			this.LogsType, this.SystemName, this.UseModule, this.UseType, this.UseContent, this.UseResult)
	}
	return context
}

// 日志文件创建
func createFileContext(context string) (bool, string) {
	timStr := time.Now().Format("2006-01-02")
	fieName := "logsFile/" + timStr + ".log"
	file, err := os.OpenFile(fieName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	file.WriteString(context)
	file.WriteString("\n")

	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	return true, timStr
}
