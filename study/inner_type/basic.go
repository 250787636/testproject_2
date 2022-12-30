package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/cmplx"
)

func euler() {
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
	//cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

const filename = "abc.txt"

func consts() {
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}
func enums() {
	const (
		cpp = iota
		//java
		_
		python
		golang
		javascript
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	//fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
func file() {
	filename := "study/inner_type/abc.txt"
	if res, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%s\n", res)
	}
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main1() {
	//euler()
	//triangle()
	//consts()
	//enums()
	//file()
	fmt.Println(grade(101))
}
