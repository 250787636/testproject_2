package main

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestFF(t *testing.T) {
	//c := []int{1, 2, 3}
	//a := num(c, 6)
	//fmt.Println(a)
	//res := reverse(9000)
	//fmt.Println(res)
	//isTrue := isPalindrome(-121)
	//isTrue := isPalindrome2(12345678987653321)
	//num := romanToInt2("MCMXCIV")
	//arr := []string{"flower", "flow", "flo"}
	//num := longestCommonPrefix(arr)
	//fmt.Println(num)
	//final := GreedyAlgorithm("5478624", 2)
	final := isValid("(){}[]")
	fmt.Println(final)
}

// 两数之和
func num(x []int, target int) []int {
	valMap := make(map[int]int, 0)
	for index, value := range x {
		if mapIndex, ok := valMap[target-value]; ok {
			return []int{mapIndex, index}
		}
		valMap[value] = index
	}
	return []int{-1, -1}
}

const (
	MaxInt32 = 1<<31 - 1
	MinInt32 = -1 << 31
)

// 计数反转
func reverse(x int) int {
	var res = 0
	for x != 0 {
		var t = x % 10

		if res > MaxInt32/10 || res < MinInt32/10 {
			return 0
		}
		res = res*10 + t
		x /= 10
	}
	return res
}

// 回文整数
func isPalindrome(x int) (isRight bool) {
	isRight = true
	arr1 := make([]int, 0)
	var num1 int
	var index int
	if x < 0 {
		isRight = false
		goto END
	}
	for x != 0 {
		num1 = x % 10
		if x > math.MaxInt32 || x < math.MinInt32 {
			isRight = false
			goto END
		}
		arr1 = append(arr1, num1)
		x /= 10
	}

	if len(arr1)%2 != 0 {
		index = (len(arr1) - 1) / 2
	} else {
		index = len(arr1) / 2
	}

	for i := 0; i < index; i++ {
		j := len(arr1) - 1 - i
		if arr1[i] != arr1[j] {
			isRight = false
			goto END
		}
	}
END:
	return isRight
}
func isPalindrome2(x int) bool {
	num := x
	if num < 0 {
		return false
	}

	num_2 := 0
	for num != 0 {
		num_2 = num_2*10 + num%10
		num /= 10
	}
	if num_2 == x {
		return true
	}
	return false
}

// 罗马转int
func romanToInt(s string) int {
	num := 0
	romaSp := make(map[string]int, 0)
	romaSp["IV"] = 4
	romaSp["IX"] = 9
	romaSp["XL"] = 40
	romaSp["XC"] = 90
	romaSp["CD"] = 400
	romaSp["CM"] = 900

	romaNo := make(map[string]int, 0)
	romaNo["I"] = 1
	romaNo["V"] = 5
	romaNo["X"] = 10
	romaNo["L"] = 50
	romaNo["C"] = 100
	romaNo["D"] = 500
	romaNo["M"] = 1000

	for k, v := range romaSp {
		if strings.Contains(s, k) {
			num += strings.Count(s, k) * v
			s = strings.ReplaceAll(s, k, "")
		}
	}
	for k, v := range romaNo {
		if strings.Contains(s, k) {
			num += strings.Count(s, k) * v
			s = strings.ReplaceAll(s, k, "")
		}
	}
	return num
}

// 最长公约数前缀
func longestCommonPrefix(strs []string) string {
	res, short := "", ""
	if len(strs) > 0 {
		res, short = strs[0], strs[0]
	}
	for _, str := range strs {
		if len(str) < len(short) {
			short, str = str, short
		}
		temp := short
		for v := range short {
			if str[v] != short[v] {
				temp = short[:v]
				break
			}
		}
		if len(temp) < len(res) {
			res = temp
		}
	}
	return res
}

// 贪心算法
func GreedyAlgorithm(nums string, k int) string {
	for i := 0; i < k; i++ {
		final := ""
		arr := strings.Split(nums, "")
		num := arr[0]
		index := 0

		for i := 1; i < len(arr); i++ {
			if num > arr[i] {
				index = i - 1
				break
			} else {
				num = arr[i]
			}
		}
		if index == len(arr)-1 {
			arr[index] = ""
		}
		for i := index; i < len(arr)-1; i++ {
			arr[i] = arr[i+1]
		}
		arr[len(arr)-1] = ""
		for _, v := range arr {
			final += v
		}
		nums = final
	}
	return nums
}

// 有效中括号
func isValid(s string) bool {
	isTrue := true

	arr := strings.Split(s, "")
	if len(arr)%2 != 0 {
		isTrue = false
	}

	fmt.Println(arr)
	return isTrue
}
