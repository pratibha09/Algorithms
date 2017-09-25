package main

import(
	"fmt"
)

func equillibrium(a []int) int{
	var left_sum, right_sum int
	for i := 0; i < len(a); i++ {
		left_sum = 0
		right_sum = 0 
		for j := 0; j < i; j++ {
			left_sum += a[j]
		}
		for j := i+1; j < len(a); j++ {
			right_sum += a[j]
		}
		if left_sum == right_sum {
			return i
		} 
	}
	return -1
}

func main(){
	a := []int{-7, 1, 5, 2, -4, 3, 0}
	fmt.Printf("%d\n", equillibrium(a))
}

