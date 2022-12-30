package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	jsoniter "github.com/json-iterator/go"
	"github.com/pikanezi/mapslice"
	"github.com/tealeg/xlsx"
	"github.com/tidwall/gjson"
	goutils "github.com/typa01/go-utils"
	"github.com/wumansgy/goEncrypt"
	"io"
	"io/ioutil"
	"kmesh2/utils/rsa"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

//func main1() {
//	//s,err :=os.Stat("./111")
//	//fmt.Println(s.ModTime())
//	//fmt.Println(time.Now())
//	//
//
//	//fmt.Println(time.Now().Sub(s.ModTime()).Hours())
//	filepath.Walk("./aa", func(path string, info fs.FileInfo, err error) error {
//		if !info.IsDir() {
//			fmt.Println(path)
//			fmt.Println(info.Name())
//			if time.Now().Sub(info.ModTime()).Hours() > 12 {
//				os.Remove(path)
//			}
//		}
//		return nil
//	})
//
//	filepath.WalkDir("./aa", func(path string, d fs.DirEntry, err error) error {
//		//fmt.Println(path)
//		fmt.Println(d.Name())
//		return nil
//	})
//
//	//fmt.Println(err)
//}
func main2() {
	store := captcha.NewMemoryStore(100, time.Minute)
	var wait sync.WaitGroup
	wait.Add(90000)
	for i := 0; i < 90000; i++ {
		go func() {
			defer wait.Done()
			m := make(map[string]interface{})
			f, err := os.Create("captcha.png")
			if err != nil {
				fmt.Println(err)
			}
			defer f.Close()

			var w io.WriterTo
			id := captcha.New()

			d := captcha.RandomDigits(4)
			m["captchaId"] = id
			store.Set(id, d)

			w = captcha.NewImage(id, d, 240, 80)
			_, err = w.WriteTo(f)
			if err != nil {
				fmt.Println(err)
				return
			}

			ff, _ := ioutil.ReadFile("captcha.png")
			digits := base64.StdEncoding.EncodeToString(ff)
			m["digits"] = digits
			//fmt.Println(digits,"222222")
		}()
	}
	wait.Wait()
}

//go func() {
//	defer wait.Done()
//	m := make(map[string]interface{})
//	//f, err := os.Create("captcha.png")
//	//if err != nil {
//	//	fmt.Println(err)
//	//}
//	//defer f.Close()
//
//	var w io.WriterTo
//	id := captcha.New()
//
//	d := captcha.RandomDigits(4)
//	m["captchaId"] = id
//	store.Set(id, d)
//
//	w = captcha.NewImage(id, d, 240,80)
//	var a bytes.Buffer
//
//	_ ,err := w.WriteTo(&a)
//	//_, err = w.WriteTo(f)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	//ff, _ := ioutil.ReadFile("captcha.png")
//	digits := base64.StdEncoding.EncodeToString(a.Bytes())
//	m["digits"] = digits
//	//fmt.Println(digits,"222222")
//}()

//基础类
type Base struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

// 总数据
type CoreUser struct {
	ID                  int64     `json:"id" `
	CreateTime          time.Time `json:"create_time"`
	ExtendAttribute1    string    `json:"extend_attribute_1" `
	ExtendAttribute2    string    `json:"extend_attribute_2" `
	ExtendAttribute3    string    `json:"extend_attribute_3" `
	ExtendAttribute4    string    `json:"extend_attribute_4" `
	ExtendAttribute5    string    `json:"extend_attribute_5" `
	ExtendAttributes    string    `json:"extend_attribute_s" gorm:"type:longtext"`
	UpdateTime          time.Time `json:"update_time"`
	Creator             string    `json:"creator" `
	CreatorID           int64     `json:"creator_id" `
	CustomerId          int64     `json:"customer_id" `
	DataPermissionType  uint      `json:"data_permission_type" `
	DepartmentId        int64     `json:"department_id" `
	DepartmentId1       int64     `json:"department_id1" `
	Email               string    `json:"email" `
	IsPasswordChanged   []byte    `json:"is_password_changed" `
	IsActive            []byte    `json:"is_active" `
	LastLoginIp         string    `json:"last_login_ip" `
	LastLoginTime       time.Time `json:"last_login_time"`
	Password            string    `json:"password" `
	Phone               string    `json:"phone" `
	Realname            string    `json:"realname"`
	RoleId              int64     `json:"role_id" `
	Username            string    `json:"username" `
	RoleType            string    `json:"role_type" `
	UDataPermissionType int       `json:"u_data_permission_type" `
	IsValid             []byte    `json:"is_valid" `
	AndroidStudioId     int       `json:"android_studio_id" `
	ApiSignature        string    `json:"api_signature" `
	AuthType            string    `json:"auth_type" `
	CepingApps          int       `json:"ceping_apps" `
	CepingTimes         int       `json:"ceping_times" `
	Company             string    `json:"company" `
	ContactEmail        string    `json:"contact_email" `
	Enable              []byte    `json:"enable" `
	IosEndtime          time.Time `json:"ios_endtime" `
	IosStudioId         int       `json:"ios_studio_id"`
	IotEndtime          time.Time `json:"iot_endtime"`
	IpaApps             int       `json:"ipa_apps"`
	IpaTimes            int       `json:"ipa_times"`
	IsFirstLogin        []byte    `json:"is_first_login"`
	LastUseday          int       `json:"last_useday"`
	NetEndtime          int       `json:"net_endtime"`
	NetTime             int       `json:"net_time"`
	PlatForm            string    `json:"plat_form"`
	PlatForms           string    `json:"plat_forms"`
	ReportFormats       string    `json:"report_formats"`
	ReportIosLanguage   int       `json:"report_ios_language"`
	ReportLanguage      int       `json:"report_language"`
	ReportLanguages     string    `json:"report_languages"`
	ReportTemplate      string    `json:"report_template"`
	SoureceEndtime      time.Time `json:"sourece_endtime"`
	SourceTimes         int       `json:"source_times"`
	Special             string    `json:"special"`
	TEndtime            time.Time `json:"t_endtime"`
	TRegtime            time.Time `json:"t_regtime"`
	TUpdatetime         time.Time `json:"t_updatetime"`
}

type PartCoreUser struct {
	ID         string
	CreateTime string
	UpdateTime string
	Creator    string
	Email      string
	Username   string
	RoleType   string
	IsValid    string
	EndTime    string
}

////客户账号关联
//type ClientUser struct {
//	Base
//	ClientKey    string     `json:"client_key"`                                     // key
//	ClientName   string     `json:"user" gorm:"size:50;index:user_idx"`             //客户名称
//	UserAccount  string     `json:"user_account" gorm:"size:50;index:account_idx" ` //账号
//	UserRealName string     `json:"user_real_name"`                                 //用户名
//	Status       t.EnumType `json:"status" gorm:"size:5"`                           //状态：是否激活
//	RoleID       uint       `json:"role_id"`                                        // 权限ID
//	RoleName     string     `json:"role_name"`                                      // 权限名
//	Email        string     `json:"email" gorm:"column:email"`                      //邮箱
//	IsAdmin      bool       `json:"is_admin"`                                       // 是否是管理员
//	IsLock       bool       `json:"is_lock"`                                        // 是否锁定
//	Products     string     `json:"przoducts" gorm:"type:text"`                     // 产品限制 TODO 加密
//	Signature    string     `json:"signature"`                                      // 签名
//}

type ApsUser struct {
	ID uint `json:"-"`
	//UID       string `json:"uid" gorm:"size:50;index:uid_idx"`          //用户UID. ID不会在业务系统中使用
	Name      string `json:"name" gorm:"unique;index:name_idx"` //用户名，即账号
	RealName  string `json:"real_name"`                         //用户真实姓名
	Password  string `json:"-" gorm:"size:64"`                  //用户密码
	LastLogin string `json:"last_login" gorm:"size:19"`         //最后一次登录
	Locked    int    `json:"-"`                                 //账号是否被锁定
	Fails     int    `json:"-"`                                 //登录失败次数
}

type MetaRole struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Permissions string `json:"permissions" gorm:"type:text"`
	ClientKey   string `json:"client_key" gorm:"size:64"`
}

// 对应的路由
type ModPermission struct {
	Resource   string   `json:"resources"`
	Operations []string `json:"operations"`
}

// 权限操作
type RoleUser struct {
	OldUser       []CoreUser
	RoleList      []string
	AdminRole     MetaRole
	OtherRole     MetaRole
	AdminRoleID   uint
	OtherRoleID   uint
	HasEditResult bool
}

// 除了漏洞检查 其余全开
func MakeRole() (permissions string) {
	pmList := []ModPermission{
		// 主页
		{
			Resource:   "index",
			Operations: []string{},
		},
		// 总览
		{
			Resource:   "index-home",
			Operations: []string{},
		},
		// 安全服务
		{
			Resource:   "SecService",
			Operations: []string{},
		},
		// 固件检测
		{
			Resource:   "Prod_FwAudit",
			Operations: []string{},
		},
		// 提交检测
		{
			Resource:   "Prod_FwAudit_Create",
			Operations: []string{},
		},
		// 预览
		{
			Resource:   "Prod_FwAudit_preview",
			Operations: []string{},
		},
		// 固件对比
		{
			Resource:   "Prod_FwAudit_compare",
			Operations: []string{},
		},
		// 软件组件安全
		{
			Resource:   "Prod_SoftwareComponentSecurity",
			Operations: []string{},
		},
		// 固件列表
		{
			Resource:   "SoftwareComponentSecurity_list",
			Operations: []string{},
		},
		// 网络漏洞检测部分暂时去掉
		//{
		//	Resource:   "Prod_NetworkAudit",
		//	Operations: []string{},
		//},
		// 提交检测
		//{
		//	Resource:   "NetworkAudit_Add",
		//	Operations: []string{},
		//},
		// 网络漏洞检测 预览
		//{
		//	Resource:   "NetworkAudit_View",
		//	Operations: []string{},
		//},
		//  漏洞详情 预览
		//{
		//	Resource:   "NetworkAudit_Bug_View",
		//	Operations: []string{},
		//},

		// 任务模板
		{
			Resource:   "TaskTemplate",
			Operations: []string{},
		},
		// 固件检测模板
		{
			Resource:   "Prod_FwAudit_Template",
			Operations: []string{},
		},
		// 网络漏洞检测模板
		//{
		//	Resource:   "Prod_NetworkAudit_Template",
		//	Operations: []string{},
		//},

		//  用户管理
		{
			Resource:   "system",
			Operations: []string{},
		},
		// 账号管理
		{
			Resource:   "AccountMng",
			Operations: []string{},
		},
		// 角色管理
		{
			Resource:   "PermMng",
			Operations: []string{},
		},
		// 关于平台
		{
			Resource:   "about",
			Operations: []string{},
		},
		// 平台介绍
		{
			Resource:   "introduce",
			Operations: []string{},
		},
	}

	permissions, _ = jsoniter.MarshalToString(pmList)
	return permissions
}

func main666() {
	db, err := gorm.Open("mysql", "root:www@admin@2020@(192.168.138.224:33060)/ssp?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	db.SingularTable(true)
	defer db.Close()

	// 总数据
	var coreUserArray []CoreUser
	var singleCoreUser PartCoreUser
	var emptyPartCoreUser []PartCoreUser
	// 进行账户分类
	//var superAdmin []CoreUser
	//var normaluser []CoreUser
	//var systemadmin model.CoreUser

	if err := db.Find(&coreUserArray).Error; err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(coreUserArray); i++ {
		if coreUserArray[i].RoleType == "systemadmin" {
			continue
		}
		singleCoreUser.ID = strconv.FormatInt(coreUserArray[i].ID, 10)
		singleCoreUser.Username = coreUserArray[i].Username
		singleCoreUser.RoleType = coreUserArray[i].RoleType
		singleCoreUser.Email = coreUserArray[i].Email
		singleCoreUser.Creator = coreUserArray[i].Creator
		singleCoreUser.CreateTime = coreUserArray[i].CreateTime.Format("2006-01-02 15:04:05")
		singleCoreUser.UpdateTime = coreUserArray[i].UpdateTime.Format("2006-01-02 15:04:05")
		singleCoreUser.EndTime = coreUserArray[i].ExtendAttribute1
		if coreUserArray[i].IsValid[0] == 0 {
			singleCoreUser.IsValid = "禁用"
		} else {
			singleCoreUser.IsValid = "启用"
		}
		emptyPartCoreUser = append(emptyPartCoreUser, singleCoreUser)
	}
	id, _ := mapslice.ToStrings(emptyPartCoreUser, "ID")
	createTime, _ := mapslice.ToStrings(emptyPartCoreUser, "CreateTime")
	updateTime, _ := mapslice.ToStrings(emptyPartCoreUser, "UpdateTime")
	creator, _ := mapslice.ToStrings(emptyPartCoreUser, "Creator")
	email, _ := mapslice.ToStrings(emptyPartCoreUser, "Email")
	username, _ := mapslice.ToStrings(emptyPartCoreUser, "Username")
	roleType, _ := mapslice.ToStrings(emptyPartCoreUser, "RoleType")
	isValid, _ := mapslice.ToStrings(emptyPartCoreUser, "IsValid")
	endTime, _ := mapslice.ToStrings(emptyPartCoreUser, "EndTime")

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("Sheet1")
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "编号"
	cell = row.AddCell()
	cell.Value = "账号名"
	cell = row.AddCell()
	cell.Value = "账号类型"
	cell = row.AddCell()
	cell.Value = "账号状态"
	cell = row.AddCell()
	cell.Value = "邮箱"
	cell = row.AddCell()
	cell.Value = "创建者名"
	cell = row.AddCell()
	cell.Value = "创建时间"
	cell = row.AddCell()
	cell.Value = "更新时间"
	cell = row.AddCell()
	cell.Value = "到期时间"
	for i := 0; i < len(id); i++ {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = id[i]

		cell = row.AddCell()
		cell.Value = username[i]

		cell = row.AddCell()
		cell.Value = roleType[i]

		cell = row.AddCell()
		cell.Value = isValid[i]

		cell = row.AddCell()
		cell.Value = email[i]

		cell = row.AddCell()
		cell.Value = creator[i]

		cell = row.AddCell()
		cell.Value = createTime[i]

		cell = row.AddCell()
		cell.Value = updateTime[i]

		cell = row.AddCell()
		cell.Value = endTime[i]
		file.Save("迁移用户2.0.xlsx")
	}

	// 进行账号分类
	//for k, v := range coreUserArray {
	//	if v.RoleType == "superadmin" { // 超级管理员
	//		superAdmin = append(superAdmin, coreUserArray[k])
	//	} else if v.RoleType == "normaluser" { // 普通用户
	//		normaluser = append(normaluser, coreUserArray[k])
	//	} else { // 系统管理员
	//		//systemadmin = coreUserArray[k]
	//	}
	//}

	// 用户全部存入ApsUser 表中
	//for _, v := range coreUserArray {
	//	// 老数据无realName 当出现此种情况时需要用username来填充
	//	realName := v.Realname
	//	if v.Realname =="" {
	//		realName = v.Username
	//	}
	//	// apsUser表
	//	if err = db.Save(&ApsUser{
	//		Name:      v.Username,
	//		RealName:  realName,
	//		Password:  v.Password,
	//		LastLogin: v.LastLoginTime.Format("2006-01-02 15:04:05"),
	//		Locked:    0,
	//		Fails:     0,
	//	}).Error;err!=nil{
	//		fmt.Println(err)
	//	}
	//}

	// 获取默认的json字段
	//file, _ := ioutil.ReadFile("./aa.json")
	// 客户管理员账号
	//for _, admin := range superAdmin {
	//	// clientUser表
	//	if err = db.Save(&ClientUser{
	//		Base: Base{
	//			CreatedAt: admin.CreateTime,
	//			UpdatedAt: admin.UpdateTime,
	//			DeletedAt: nil,
	//		},
	//		ClientKey:    "",
	//		ClientName:   "root",
	//		UserAccount:  admin.Username,
	//		UserRealName: admin.Username,
	//		Status:       t.EnumType(admin.IsValid[0]),
	//		RoleID:       0,
	//		RoleName:     admin.Username,
	//		Email:        admin.Email,
	//		IsAdmin:      true,
	//		IsLock:       false,
	//		Products:     string(file),
	//		Signature:    "",
	//	}).Error; err != nil {
	//		fmt.Println(err)
	//	}
	//}

	// 获取权限
	//permission := MakeRole()
	//// 客户子账号
	//for _, user := range normaluser {
	//	// metaRole表
	//	if err = db.Save(&MetaRole{
	//		Name:        user.Username,
	//		Permissions: permission,
	//		ClientKey:   "",
	//	}).Error; err != nil {
	//		fmt.Println(err)
	//	}
	//}
	// 需要metaRole的自生成id作为RoleID 所以分两次遍历
	//for _, user := range normaluser {
	//	var metaRole MetaRole
	//	db.First(&metaRole).Where("name = ?", user.Username)
	//	// clientUser表
	//	if err = db.Save(&ClientUser{
	//		Base: Base{
	//			CreatedAt: user.CreateTime,
	//			UpdatedAt: user.UpdateTime,
	//			DeletedAt: nil,
	//		},
	//		ClientKey:    "",
	//		ClientName:   user.Creator,
	//		UserAccount:  user.Username,
	//		UserRealName: user.Username,
	//		Status:       t.EnumType(user.IsValid[0]),
	//		RoleID:       metaRole.ID,
	//		RoleName:     user.Username,
	//		Email:        user.Email,
	//		IsAdmin:      false,
	//		IsLock:       false,
	//		Products:     string(file),
	//		Signature:    "",
	//	}).Error; err != nil {
	//		fmt.Println(err)
	//	}
	//}
}

//func main10() {
//	r := gin.Default()
//	r.MaxMultipartMemory = 8 << 20
//	r.GET("/file", func(c *gin.Context) {
//		form, _ := c.MultipartForm()
//		files := form.File["file"]
//		udids := make(map[int]interface{})
//		// 文件获取
//		for key, file := range files {
//			fil, _ := file.Open()
//			context, _ := io.ReadAll(fil)
//			udids[key] = context
//		}
//		results := make(map[int]string)
//		for i := 0; i < len(udids); i++ {
//			// 进行结果解码
//			result := DecryptLic(udids[i].([]byte))
//			results[i] = result
//		}
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"state":     200,
//			"clientkey": results,
//		})
//	})
//	r.Run(":8080")
//}

const (
	//系统私钥
	SYSTEM_PRV_KEY_NAME = "system.pem"
	HEAD                = "-----BEGIN SSP LICENSE KEY-----\n"
	BOTTOM              = "\n-----END SSP LICENSE KEY-----"
	ADD                 = "\n+\n"
	DEFAULT_URL         = "/data/bin/program"
)

var systemPrivateKey string

func DecryptLic(data []byte) (res string) {
	licStr := data
	if len(licStr) == 0 {
		fmt.Println("err")
		return
	}
	b, err := ioutil.ReadFile(SYSTEM_PRV_KEY_NAME)

	systemPrivateKey = string(b)

	result, err := decryptSSPLicense(licStr, []byte(systemPrivateKey))
	if err != nil {
		fmt.Println(err)
	}
	resudid := gjson.Get(result, "udid")

	return resudid.Str
}

func decryptSSPLicense(licData, priKey []byte) (string, error) {
	licStr := string(licData)
	result := ""
	licFormatError := errors.New("SPP license格式错误")
	if !strings.Contains(licStr, HEAD) || !strings.Contains(licStr, BOTTOM) {
		return result, licFormatError
	}
	licStr = strings.Replace(licStr, HEAD, "", -1)
	licStr = strings.Replace(licStr, BOTTOM, "", -1)
	partArray := strings.SplitN(licStr, ADD, -1)
	if len(partArray) != 2 {
		return result, licFormatError
	}

	encryptedData := partArray[0]
	encryptedKey := partArray[1]
	de64EncryptedKey, err := base64.StdEncoding.DecodeString(encryptedKey)
	if err != nil {
		return result, err
	}
	de64EncryptedData, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return result, err
	}
	aesKeyBytes, err := rsa.RsaDecrypt(de64EncryptedKey, priKey)
	if err != nil {
		return result, err
	}
	resultBytes, err := goEncrypt.AesCtrDecrypt(de64EncryptedData, aesKeyBytes)
	if err != nil {
		return result, err
	}
	result = string(resultBytes)
	return result, nil
}

func main4() {
	second := goutils.UUID()
	//a := "media\\33a2f05cd5af9e49bc6126c7800ca6e1.zip"
	//b := strings.ReplaceAll(a,"\\","/")
	//fmt.Println(b)
	err := CP("media/33a2f05cd5af9e49bc6126c7800ca6e1.zip", "media/cc/")
	if err != nil {
		fmt.Println(err)
	}

	zipName := fmt.Sprintf("%s.zip", second)
	err = ZipFile("media/cc/", zipName)
	if err != nil {
		fmt.Println(err)
	}
}

func CP(file, target string) error {
	err := os.MkdirAll(filepath.Dir(target), os.ModePerm)
	if err != nil {
		return err
	}
	_, err = exec.Command("/bin/sh", []string{
		"-c", fmt.Sprintf(fmt.Sprintf("cp -i %s %s", file, target)),
	}...).Output()
	if err != nil {
		return err
	}
	return err
}

func ZipFile(dir, zipName string) error {
	order := fmt.Sprintf("cd %s && zip -r ./%s ./* ", dir, zipName)
	//fmt.Println(order)
	err := exec.Command("/bin/sh", "-c", order).Run()
	if err != nil {
		order := fmt.Sprintf("cd %s && touch %s", dir, zipName)
		err = exec.Command("/bin/sh", "-c", order).Run()
	}
	return err
}

//固件给源文件打压缩包那个地方的方法有问题，内存泄漏，你可以用上面两个方法，
//一个把源文件复制到某
// 石化的单点登录
func main() {
	r := gin.Default()
	r.POST("/dmcwebapi/api/dmc/ReqToken", sso)
	r.Run(":8888")
}

type reciveData struct {
	Type    string `json:"type"`
	Imtoken string `json:"imtoken"`
}

func sso(r *gin.Context) {
	var data reciveData
	err := r.ShouldBindJSON(&data)
	if err != nil {
		fmt.Println(err)
	}
	if data.Imtoken == "31f0692eb40840aab95591cd1faf8be9" {
		r.JSON(http.StatusOK, gin.H{
			"status": "1",
			"uid":    "wangfl123",
		})
	} else if data.Imtoken == "11f0692eb40840aab95591cd1faf8be8" {
		r.JSON(http.StatusOK, gin.H{
			"status": "1",
			"uid":    "zhangm.sdsy",
		})
	} else if data.Imtoken == "21f0692eb40840aab95591cd1faf8be7" {
		r.JSON(http.StatusOK, gin.H{
			"status": "1",
			"uid":    "xuy.sdsy",
		})
	} else if data.Imtoken == "41f0692eb40840aab95591cd1faf8be6" {
		r.JSON(http.StatusOK, gin.H{
			"status": "1",
			"uid":    "malf69",
		})
	} else {
		r.JSON(http.StatusOK, gin.H{
			"status":      "0",
			"failure_msg": "非法token",
		})
	}
}
