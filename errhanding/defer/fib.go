package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for i := 0; i < 15 && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
}
