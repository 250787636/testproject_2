package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"os"
	"testing"
)

func TestGGGGGG(t *testing.T) {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	text, err := os.ReadFile("示例.et")
	if err != nil {
		fmt.Println(err)
	}
	var files = []struct {
		Name, Body string
	}{
		{"示例.et", string(text)},
	}

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			fmt.Println(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			fmt.Println(err)
		}
	}
	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}
	f, err := os.OpenFile("示例.zip", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	buf.WriteTo(f)

	f2, err := excelize.OpenFile("./示例.zip")
	if err != nil {
		log.Fatal(err)
	}

	//读取某个单元格的值
	value := f2.GetCellValue("示例", "B2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)

	//读取某个表单的所有数据
	//rows := f2.GetRows("示例")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, row := range rows {
	//	for _, value := range row {
	//		fmt.Printf("\t%s", value)
	//	}
	//	fmt.Println()
	//}
}
