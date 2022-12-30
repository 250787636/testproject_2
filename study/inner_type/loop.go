package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func converToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func main2() {
	fmt.Println(
		converToBin(5),  //101
		converToBin(13), //1101
		converToBin(72387885),
		converToBin(0),
	)

	printFile("study/inner_type/abc.txt")
	forever()
}
