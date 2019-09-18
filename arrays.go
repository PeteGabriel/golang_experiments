package main

import (
	"fmt"
)

func main() {
	//reverse using an array
	a := []int{0, 1, 2, 3, 4, 5}
	reverse(a)
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//reverse using a pointer to an array
	aP := &[]int{0, 1, 2, 3, 4, 5}
	reverseP(aP)
	fmt.Println(*aP) // "[5 4 3 2 1 0]"
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseP(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
