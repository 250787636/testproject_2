package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"regexp"
	"strings"
	"time"
)

//基础类
type Base struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-" gorm:"index"`
}

//客户账号关联
type ClientUser struct {
	Base
	ClientKey    string `gorm:"type:varchar(255);index:client_key" json:"client_key"`     // key
	ClientName   string `gorm:"type:varchar(255)" json:"user"`                            //客户名称
	UserAccount  string `gorm:"type:varchar(255);index:user_account" json:"user_account"` //账号
	UserRealName string `gorm:"type:varchar(255)" json:"user_real_name"`                  //用户名
	Status       int    `gorm:"type:int(11)" json:"status"`                               //状态：是否激活
	RoleID       uint   `gorm:"type:int(11)" json:"role_id"`                              // 权限ID
	RoleName     string `gorm:"type:varchar(255)" json:"role_name"`                       // 权限名
	Email        string `gorm:"type:varchar(255)" json:"email"`                           //邮箱
	IsAdmin      bool   `gorm:"type:int(11)" json:"is_admin"`                             // 是否是管理员
	IsLock       bool   `gorm:"type:int(11)" json:"is_lock"`                              // 是否锁定
	Products     string `gorm:"type:longtext" json:"products"`                            // 产品限制 TODO 加密
	ApiSignature string `gorm:"type:varchar(255)" json:"-"`
	SubmitType   int    `gorm:"type:tinyint(4)" json:"submit_type"`
}

func main() {
	fileName := "aimrsk-2022-09-20T15-11-13.999.csv"

	startIndex := strings.Index(fileName, "T")
	lastIndex := strings.LastIndex(fileName, ".")
	finalName := fileName[:startIndex] + fileName[lastIndex:]
	fmt.Println(finalName)

}

func GetPhoneNum() {
	arr := []string{"18860230359", "1034054160@qq.com", "yingjian.yan@bangcle.com", "18008098277"}
	regRuler := "^1[345789]{1}\\d{9}$"
	reg := regexp.MustCompile(regRuler)
	var arr2 []string
	for _, value := range arr {
		if reg.MatchString(value) {
			arr2 = append(arr2, value)
		}
	}
	fmt.Println(arr2)
}

func TheClock() {
	theClock := cron.New()
	theClock2 := cron.New()
	fmt.Println(time.Now())
	if _, err := theClock.AddFunc("0/1 * * * ?", func() {
		fmt.Println(time.Now())
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
	}); err != nil {
		fmt.Println(err)
	}
	if _, err := theClock2.AddFunc(fmt.Sprintf("@every %ds", 6), func() {
		fmt.Println("每6秒执行一次")
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
	}); err != nil {
		fmt.Println(err)
	}
	theClock.Start()
	//theClock2.Start()
}

func DataBase() {
	//db, err := gorm.Open(mysql.Open("root:www@admin@2020@tcp(172.16.102.58:33060)/ssp?charset=utf8mb4&parseTime=True&loc=Local"),&gorm.Config{
	//	NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
	//})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//var userAccountArr []string
	//// 手机正则
	////regRuler := "^1[345789]{1}\\d{9}$"
	////reg := regexp.MustCompile(regRuler)
	//if err := db.Model(ClientUser{}).Select("user_account").Find(&userAccountArr).Error; err != nil {
	//	fmt.Println(err)
	//}
	//
	//for _, value := range userAccountArr {
	//	//if reg.MatchString(value.UserAccount) {
	//
	//	//}
	//	fmt.Println(value)
	//}
}
