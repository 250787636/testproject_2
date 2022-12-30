package main

import "fmt"

func mains() {
	s1 := []int{1, 2, 3}
	s2 := make([]int, 3)
	fmt.Println(s1)
	fmt.Println(s2)
	copy(s2, s1)
	fmt.Println(s1)
	fmt.Println(s2)

	s2 = append(s2[:2], s2[3:]...)
	fmt.Println(s2)
}
