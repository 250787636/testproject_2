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

//?????????
type Base struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

// ?????????
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

////??????????????????
//type ClientUser struct {
//	Base
//	ClientKey    string     `json:"client_key"`                                     // key
//	ClientName   string     `json:"user" gorm:"size:50;index:user_idx"`             //????????????
//	UserAccount  string     `json:"user_account" gorm:"size:50;index:account_idx" ` //??????
//	UserRealName string     `json:"user_real_name"`                                 //?????????
//	Status       t.EnumType `json:"status" gorm:"size:5"`                           //?????????????????????
//	RoleID       uint       `json:"role_id"`                                        // ??????ID
//	RoleName     string     `json:"role_name"`                                      // ?????????
//	Email        string     `json:"email" gorm:"column:email"`                      //??????
//	IsAdmin      bool       `json:"is_admin"`                                       // ??????????????????
//	IsLock       bool       `json:"is_lock"`                                        // ????????????
//	Products     string     `json:"przoducts" gorm:"type:text"`                     // ???????????? TODO ??????
//	Signature    string     `json:"signature"`                                      // ??????
//}

type ApsUser struct {
	ID uint `json:"-"`
	//UID       string `json:"uid" gorm:"size:50;index:uid_idx"`          //??????UID. ID??????????????????????????????
	Name      string `json:"name" gorm:"unique;index:name_idx"` //?????????????????????
	RealName  string `json:"real_name"`                         //??????????????????
	Password  string `json:"-" gorm:"size:64"`                  //????????????
	LastLogin string `json:"last_login" gorm:"size:19"`         //??????????????????
	Locked    int    `json:"-"`                                 //?????????????????????
	Fails     int    `json:"-"`                                 //??????????????????
}

type MetaRole struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Permissions string `json:"permissions" gorm:"type:text"`
	ClientKey   string `json:"client_key" gorm:"size:64"`
}

// ???????????????
type ModPermission struct {
	Resource   string   `json:"resources"`
	Operations []string `json:"operations"`
}

// ????????????
type RoleUser struct {
	OldUser       []CoreUser
	RoleList      []string
	AdminRole     MetaRole
	OtherRole     MetaRole
	AdminRoleID   uint
	OtherRoleID   uint
	HasEditResult bool
}

// ?????????????????? ????????????
func MakeRole() (permissions string) {
	pmList := []ModPermission{
		// ??????
		{
			Resource:   "index",
			Operations: []string{},
		},
		// ??????
		{
			Resource:   "index-home",
			Operations: []string{},
		},
		// ????????????
		{
			Resource:   "SecService",
			Operations: []string{},
		},
		// ????????????
		{
			Resource:   "Prod_FwAudit",
			Operations: []string{},
		},
		// ????????????
		{
			Resource:   "Prod_FwAudit_Create",
			Operations: []string{},
		},
		// ??????
		{
			Resource:   "Prod_FwAudit_preview",
			Operations: []string{},
		},
		// ????????????
		{
			Resource:   "Prod_FwAudit_compare",
			Operations: []string{},
		},
		// ??????????????????
		{
			Resource:   "Prod_SoftwareComponentSecurity",
			Operations: []string{},
		},
		// ????????????
		{
			Resource:   "SoftwareComponentSecurity_list",
			Operations: []string{},
		},
		// ????????????????????????????????????
		//{
		//	Resource:   "Prod_NetworkAudit",
		//	Operations: []string{},
		//},
		// ????????????
		//{
		//	Resource:   "NetworkAudit_Add",
		//	Operations: []string{},
		//},
		// ?????????????????? ??????
		//{
		//	Resource:   "NetworkAudit_View",
		//	Operations: []string{},
		//},
		//  ???????????? ??????
		//{
		//	Resource:   "NetworkAudit_Bug_View",
		//	Operations: []string{},
		//},

		// ????????????
		{
			Resource:   "TaskTemplate",
			Operations: []string{},
		},
		// ??????????????????
		{
			Resource:   "Prod_FwAudit_Template",
			Operations: []string{},
		},
		// ????????????????????????
		//{
		//	Resource:   "Prod_NetworkAudit_Template",
		//	Operations: []string{},
		//},

		//  ????????????
		{
			Resource:   "system",
			Operations: []string{},
		},
		// ????????????
		{
			Resource:   "AccountMng",
			Operations: []string{},
		},
		// ????????????
		{
			Resource:   "PermMng",
			Operations: []string{},
		},
		// ????????????
		{
			Resource:   "about",
			Operations: []string{},
		},
		// ????????????
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

	// ?????????
	var coreUserArray []CoreUser
	var singleCoreUser PartCoreUser
	var emptyPartCoreUser []PartCoreUser
	// ??????????????????
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
			singleCoreUser.IsValid = "??????"
		} else {
			singleCoreUser.IsValid = "??????"
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
	cell.Value = "??????"
	cell = row.AddCell()
	cell.Value = "?????????"
	cell = row.AddCell()
	cell.Value = "????????????"
	cell = row.AddCell()
	cell.Value = "????????????"
	cell = row.AddCell()
	cell.Value = "??????"
	cell = row.AddCell()
	cell.Value = "????????????"
	cell = row.AddCell()
	cell.Value = "????????????"
	cell = row.AddCell()
	cell.Value = "????????????"
	cell = row.AddCell()
	cell.Value = "????????????"
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
		file.Save("????????????2.0.xlsx")
	}

	// ??????????????????
	//for k, v := range coreUserArray {
	//	if v.RoleType == "superadmin" { // ???????????????
	//		superAdmin = append(superAdmin, coreUserArray[k])
	//	} else if v.RoleType == "normaluser" { // ????????????
	//		normaluser = append(normaluser, coreUserArray[k])
	//	} else { // ???????????????
	//		//systemadmin = coreUserArray[k]
	//	}
	//}

	// ??????????????????ApsUser ??????
	//for _, v := range coreUserArray {
	//	// ????????????realName ?????????????????????????????????username?????????
	//	realName := v.Realname
	//	if v.Realname =="" {
	//		realName = v.Username
	//	}
	//	// apsUser???
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

	// ???????????????json??????
	//file, _ := ioutil.ReadFile("./aa.json")
	// ?????????????????????
	//for _, admin := range superAdmin {
	//	// clientUser???
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

	// ????????????
	//permission := MakeRole()
	//// ???????????????
	//for _, user := range normaluser {
	//	// metaRole???
	//	if err = db.Save(&MetaRole{
	//		Name:        user.Username,
	//		Permissions: permission,
	//		ClientKey:   "",
	//	}).Error; err != nil {
	//		fmt.Println(err)
	//	}
	//}
	// ??????metaRole????????????id??????RoleID ?????????????????????
	//for _, user := range normaluser {
	//	var metaRole MetaRole
	//	db.First(&metaRole).Where("name = ?", user.Username)
	//	// clientUser???
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
//		// ????????????
//		for key, file := range files {
//			fil, _ := file.Open()
//			context, _ := io.ReadAll(fil)
//			udids[key] = context
//		}
//		results := make(map[int]string)
//		for i := 0; i < len(udids); i++ {
//			// ??????????????????
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
	//????????????
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
	licFormatError := errors.New("SPP license????????????")
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

//???????????????????????????????????????????????????????????????????????????????????????????????????????????????
//??????????????????????????????
// ?????????????????????
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
			"failure_msg": "??????token",
		})
	}
}
