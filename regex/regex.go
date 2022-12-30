package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is ccmouse1@gmail.com@aac.com
email is abc@def.org
email2 is kk@qq.com
email3 is ddd@abc.com.cn
`

// .+ 任意一个或多个
// .* 任意零个或多个
// [a-zA-Z0-9] 小写字母 大写字母 数字

func main() {
	re := regexp.MustCompile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)(\.[a-zA-Z0-9.]+)`)
	mathch := re.FindAllStringSubmatch(text, -1)
	for _, m := range mathch {
		fmt.Println(m)
	}
}
