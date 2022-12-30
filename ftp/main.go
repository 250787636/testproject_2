package main

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	//c, err := ftp.Dial("127.0.0.1:21")
	//fmt.Println(err)
	//defer c.Quit()

	// 可以重复连接，创建不同的实例
	//c1, err := ftp.Dial("127.0.0.1:21")
	//fmt.
	//defer c1.Quit()

	//err = c.Login("test", "123456")
	//fmt.Println(err)

	// 新建文件
	//reader := bytes.NewReader([]byte("12356546"))
	//err = c.Stor("1.txt", reader)
	//fmt.

	// 创建文件夹，可以嵌套 dir1/dir11
	//err = c.MakeDir("dir1")
	//fmt.Println(err)
	//fmt.Println("创建文件夹成功")
	//
	//// 进入文件夹，可以嵌套 dir1/dir11
	//err = c.ChangeDir("dir1")
	//fmt.Println(err)
	//fmt.Println("进入文件夹成功")

	// 新建文件
	//reader := bytes.NewReader([]byte("12356546"))
	//err = c.Stor("1.txt", reader)
	//fmt.Println(err)
	//fmt.Println("创建文件成功")
	//
	//dir, err := c.CurrentDir()
	//fmt.Println(err)
	//fmt.Println("current dir: ", dir)

	// 返回上一级
	//err = c.ChangeDirToParent()
	//fmt.Println(err)
	//fmt.Println(c.CurrentDir())

	//entries, err := c.List("/")
	//fmt.fmt.Println(err)
	layout := "2006-01-02"
	now := time.Now()
	currentTime, _ := time.ParseDuration("-24h")
	oneDayBefore := now.Add(currentTime)
	rightDate := oneDayBefore.Format(layout)
	dirFiles, _ := TPFuncReadDirFiles("logsFile")
	for _, file := range dirFiles {
		if strings.Contains(file, rightDate) {
			fmt.Println(true)
		}
	}
	fmt.Printf("%s", dirFiles)
}
func TPFuncReadDirFiles(dir string) ([]string, error) {
	var files []string
	//方法一
	var walkFunc = func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	}
	err := filepath.Walk(dir, walkFunc)
	return files, err
}

func EnsureFtpDirExist(c *ftp.ServerConn, dir string) error {
	// 这里不能直接 MakeDir，有权限问题
	_, err := c.List(dir)
	if err != nil {
		if er := c.MakeDir(dir); er != nil {
			if er.Error() == "550 Directory with same name already exists." {
				return nil
			}
			return er
		}
	}
	return nil
}
