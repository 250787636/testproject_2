package main_test

import (
	"fmt"
	uuid "github.com/google/uuid"
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	//arr := []int{3, 1, 2, 5, 4, 9, 7, 2}
	//dic := make(map[int]int)
	//for _, v := range arr {
	//	if dic[v]==0 {
	//		dic[v] = 1
	//	}else {
	//		dic[v] = dic[v] +1
	//	}
	//}
	//fmt.Println(dic)
	//arrc := push(arr, 2, 10)
	//arrc := del(arr, 2)
	//fmt.Println(arrc)

	a := uuid.New()
	fmt.Println(a)
}

func push(arr []int, index int, num int) (res []int) {
	arr2 := make([]int, len(arr)+1)
	for i, _ := range arr {
		if i < index {
			arr2[i] = arr[i]
		} else if i == index {
			arr2[index] = num
			arr2[i+1] = arr[i]
		} else {
			arr2[i+1] = arr[i]
		}
	}

	return arr2
}

func del(arr []int, index int) (res []int) {
	arr2 := make([]int, len(arr)-1)
	for i, _ := range arr {
		if i > index {
			arr2[i-1] = arr[i]
		} else if i == index {
			arr2[index] = arr[i+1]
		} else {
			arr2[i] = arr[i]
		}
	}

	return arr2
}

type Tree struct {
	data          int
	LeftChildren  *Tree
	RightChildren *Tree
}

func Test2(t *testing.T) {
	li := []int{3, 2, 8, 0, 0, 10, 0, 0, 8, 0, 4}
	treeNode := createNode(li)
	fmt.Println("前序遍历:")
	preTraveral(*treeNode)
	//fmt.Println("中序遍历:")
	//inTraveral(*treeNode)
	//fmt.Println("后序遍历:")
	//postTraveral(*treeNode)
}

//创建节点
func createNode(nodes []int) (t *Tree) {
	if len(nodes) == 0 {
		return
	}
	data := getData(nodes)
	nodes = getArray(nodes)
	node := new(Tree)
	if data != 0 {
		node.data = data
		node.LeftChildren = createNode(nodes)
		node.RightChildren = createNode(nodes)
	}
	return node
}

// 前序遍历
func preTraveral(node Tree) {
	if reflect.DeepEqual(node, Tree{}) {
		return
	}
	fmt.Println(node.data)
	if reflect.DeepEqual(node, Tree{data: node.data, LeftChildren: nil, RightChildren: nil}) {
		return
	}
	preTraveral(*node.LeftChildren)
	preTraveral(*node.RightChildren)
}

// 中中序列遍历
func inTraveral(node Tree) {
	if reflect.DeepEqual(node, Tree{}) {
		return
	}
	inTraveral(*node.LeftChildren)
	fmt.Println(node.data)
	inTraveral(*node.RightChildren)
}

// 后序遍历
func postTraveral(node Tree) {
	if reflect.DeepEqual(node, Tree{}) {
		return
	}
	postTraveral(*node.LeftChildren)
	postTraveral(*node.RightChildren)
	fmt.Println(node.data)
}

func getData(array []int) (data int) {
	for i, _ := range array {
		if array[i] != 0 {
			data = array[i]
			break
		}
	}
	return data
}

func getArray(array []int) (arr []int) {
	for i := 0; i < len(array); i++ {
		if array[i] != 0 {
			arr = make([]int, len(array)-1-i)
			for j := 0; j < len(arr); j++ {
				arr[j] = array[len(array)-len(arr)+j]
			}
			break
		}
	}
	return arr
}

func Test3(t *testing.T) {
	//arr := []int{5, 8, 6, 3, 9, 2, 1, 7}
	arr := []int{2, 3, 4, 5, 6, 7, 8, 1, 11, 19, 15, 35, 12, 0, 99}
	res := bubble(arr)
	fmt.Println(res)
}

func bubble(arr []int) []int {
	lastExchangeIndex := 0
	sortLength := len(arr) - 1
	for i := 0; i < len(arr)-1; i++ {
		isSort := true
		for j := 0; j < sortLength; j++ {
			if arr[j] > arr[j+1] {
				c := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = c
				isSort = false
				lastExchangeIndex = j
			}
		}
		sortLength = lastExchangeIndex
		if isSort {
			break
		}
	}
	return arr
}

func chickWine(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		//isSoretd := true
		// 奇数轮
		for j := i; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				//isSoretd = false
			}
		}
		//if isSoretd {
		//	break
		//}

		//isSoretd = true
		// 偶数轮
		for j := len(arr) - i - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
				//isSoretd = false
			}
		}
		//if isSoretd {
		//	break
		//}
	}
	return arr
}

func Test4(t *testing.T) {
	//arr := []int{3,3,6,5,3,2,8,1}
	arr := []int{4, 4, 6, 5, 3, 2, 8, 1}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func partition(arr []int, startIndex int, endIndex int) int {
	privot := arr[startIndex]
	left := startIndex
	right := endIndex
	for left != right {
		for left < right && arr[right] > privot {
			right--
		}
		for left < right && arr[left] <= privot {
			left++
		}
		if left < right {
			p := arr[left]
			arr[left] = arr[right]
			arr[right] = p
		}
	}
	// pivot和指针重合点交换
	arr[startIndex] = arr[left]
	arr[left] = privot
	return left
}

func partition2(arr []int, startIndex int, endIndex int) int {
	pivot := arr[startIndex]
	mark := startIndex

	for i := startIndex + 1; i <= endIndex; i++ {
		if arr[i] < pivot {
			mark++
			p := arr[mark]
			arr[mark] = arr[i]
			arr[i] = p
		}
	}
	arr[startIndex] = arr[mark]
	arr[mark] = pivot
	return mark
}

func quickSort(arr []int, startIndex int, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	pivotIndex := partition(arr, startIndex, endIndex)
	quickSort(arr, startIndex, pivotIndex-1)
	quickSort(arr, startIndex+1, pivotIndex)
}

// 计数排序
func Test7(t *testing.T) {
	arr := []int{4, 4, 6, 5, 3, 2, 8, 1, 7, 5, 6, 0, 10}
	sortedArr := countSort2(arr)
	fmt.Println(sortedArr)
}

func countSort(arr []int) []int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	countArr := make([]int, max+1)
	for i := 0; i < len(arr); i++ {
		countArr[arr[i]]++
	}
	index := 0
	sortedArr := make([]int, len(arr))
	for i := 0; i < len(countArr); i++ {
		for j := 0; j < countArr[i]; j++ {
			sortedArr[index] = i
			index++
		}
	}
	return sortedArr
}

func countSort2(arr []int) []int {
	max := arr[0]
	min := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	d := max - min
	arr2 := make([]int, d+1)
	for i := 0; i < len(arr); i++ {
		arr2[arr[i]-min]++
	}
	for i := 1; i < len(arr2); i++ {
		arr2[i] += arr2[i-1]
	}
	sorteddArr := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		sorteddArr[arr2[arr[i]-min]-1] = arr[i]
		arr2[arr[i]-min]--
	}
	return sorteddArr
}

func TestG(t *testing.T) {
	//a:=method(25,5)
	//b:=method(100,80)
	c := method(27, 14)
	//fmt.Println(a)
	//fmt.Println(b)
	fmt.Println(c)
}
func method(a int, b int) int {
	var small int
	var big int
	if a > b {
		small = b
		big = a
	} else {
		small = a
		big = b
	}
	if big%small == 0 {
		return small
	}
	return method(b%small, small)
}
