package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestFile(T *testing.T) {
	os.MkdirAll("data/scan_upload", os.ModePerm)
	os.MkdirAll("datastorage/ak/upload/scan_upload", os.ModePerm)
	go TickerDel(1 * time.Second)
	for {

	}
}
func TickerDel(clock time.Duration) {
	ticker := time.NewTicker(clock)
	for range ticker.C {
		fmt.Println("开工了！！！！！")
		err := getFile("datastorage/ak/upload/scan_upload",
			"data/scan_upload")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getFile(movpath, deldir string) error {
	now := time.Now()
	// 超过30天放入待删除路径
	err := filepath.Walk(movpath, func(path string, info fs.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		// 文件真实路径
		relpath, _ := filepath.Rel(movpath, path)
		//30*time.Hour*24
		if isMov := now.Sub(info.ModTime()) > time.Second; isMov {
			afterPath := deldir + "/" + relpath
			err := os.MkdirAll(filepath.Dir(afterPath), os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
			err = os.Rename(path, afterPath)
			if err != nil {
				fmt.Println(err)
			}
			// 删除初始文件夹
			err = os.Remove(filepath.Dir(path))
			if err != nil {
				fmt.Println(err)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 超过365天进行清除
	err = filepath.Walk(deldir, func(path string, info fs.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		//365*time.Hour*24
		if isMov := now.Sub(info.ModTime()) > 1*time.Second; isMov {
			err := os.Remove(path)
			if err != nil {
				fmt.Println(err)
			}
		}
		err = os.Remove(filepath.Dir(path))
		if err != nil {
			fmt.Println(err)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
