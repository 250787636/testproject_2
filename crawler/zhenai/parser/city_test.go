package parser

import (
	"fmt"
	"regexp"
	"testing"
)

var a = `http://album.zhenai.com/u/[0-9]+`

func TestCity(t *testing.T) {
	var str = `<a href="http://album.zhenai.com/u/1677398857" target="_blank">愿得一人心的美好</a>`
	re := regexp.MustCompile(a)
	matches := re.FindAllSubmatch([]byte(str), -1)
	for _, m := range matches {
		fmt.Println(m)
	}
}
