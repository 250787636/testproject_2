package main

import (
	"bytes"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"
	"time"
)

func TestC(t *testing.T) {
	//first  := goutils.UUID()
	//second := goutils.UUID()
	//
	//err := CP("cc.json","media/"+first)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//zipName := fmt.Sprintf("./media/%s/%s.zip", first, second)
	//err = ZipFile("./media",zipName)
	//if err !=nil {
	//	fmt.Println(err)
	//}

	//a := []string{"1", "2", "1", "3", "2", "3"}
	//c := RemoveDuplicatesAndEmpty(a)
	//fmt.Println(c)

	//dirName := "logsFile"
	//err := os.Mkdir(dirName, os.ModePerm)
	//if err != nil {
	//	fmt.Println(err)
	//}
	// 一天前
	//now := time.Now()
	//currentTime, _ := time.ParseDuration("-24h")
	//oneDayBefore := now.Add(currentTime)
	//timStr := oneDayBefore.Format("2006-01-02")
	//fieName :=dirName+"/"+timStr + ".log"
	////fieName := "2022-08-17" + ".log"
	//file, err := os.OpenFile(fieName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//file.WriteString("张三李四王五")
	//file.WriteString("\n")
	//
	//err = file.Close()
	//if err != nil {
	//	fmt.Println(err)
	//}

	layout := "2006-01-02"
	layout2 := "2006-01-02 15:04:05"

	now := time.Now()
	currentTime, _ := time.ParseDuration("-24h")
	//获取当前日期的前一天
	oneDayBefore := now.Add(currentTime)
	// 当前时间
	//targetDate := now.Format(layout2)
	targetDate := "2022-08-24 00:05:00"
	// 获取前一天的 午夜 点时刻
	startDate := oneDayBefore.Format(layout) + " " + "23:59:59"
	// 获取当前日期的 午夜 10分
	endDate := now.Format(layout) + " " + "00:10:00"

	fmt.Println(startDate)
	fmt.Println(endDate)

	targetTime, err := time.Parse(layout2, targetDate)
	if err != nil {
		fmt.Println(err)
	}
	startDateTime, err := time.Parse(layout2, startDate)
	if err != nil {
		fmt.Println(err)
	}
	endDateTime, err := time.Parse(layout2, endDate)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(startDateTime.Before(targetTime) && targetTime.Before(endDateTime))
	fieName := "logsFile/2022-08-23.log"
	if startDateTime.Before(targetTime) && targetTime.Before(endDateTime) {
		err = uploadFiletoFtp(fieName)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func uploadFiletoFtp(file string) error {
	c, err := ftp.Dial("127.0.0.1:21")
	if err != nil {
		return err
	}
	defer c.Quit()
	err = c.Login("test", "123456")
	if err != nil {
		return err
	}
	err = c.MakeDir("ipName")
	if err != nil {
		log.Println("文件夹已经存在,", err)
	}

	err = c.ChangeDir("ipName")
	if err != nil {
		log.Println("进入文件夹失败", err)
	}
	log.Println("进入文件夹成功")

	// 上传文件
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println("读取文件失败,", err.Error())
	}
	err = c.Stor(filepath.Base(file), bytes.NewBuffer(b))
	if err != nil {
		log.Println(err.Error())
	}
	// 读取文件
	//r, err := c.Retr(file)
	//defer r.Close()
	//buf, err := ioutil.ReadAll(r)
	//log.Println(string(buf))

	// 退出
	return nil
}

func RemoveDuplicatesAndEmpty(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

//func CP(file,target string) error {
//	err := os.MkdirAll(filepath.Dir(target), os.ModePerm)
//	if err != nil {
//		return err
//	}
//	_, err = exec.Command("/bin/sh", []string{
//		"-c",fmt.Sprintf(fmt.Sprintf("cp -i %s %s",file,target)),
//	}...).Output()
//	if err != nil {
//		return err
//	}
//	return err
//}
//
//func ZipFile(dir, zipName string) error {
//	order := fmt.Sprintf("cd %s && zip -r ./%s ./* ", dir, zipName)
//	fmt.Println(order)
//	err := exec.Command("/bin/sh", "-c", order).Run()
//	if err != nil {
//		order := fmt.Sprintf("cd %s && touch %s", dir, zipName)
//		err = exec.Command("/bin/sh", "-c", order).Run()
//	}
//	return err
//}

//固件给源文件打压缩包那个地方的方法有问题，内存泄漏，你可以用上面两个方法，
//一个把源文件复制到某个目录下，另一个方法可以给这个目录打压缩包
