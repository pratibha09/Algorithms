package main

import(
	"fmt"
)

func lis(a []int, n int) int{
	lis := make([]int, n)
	var max int = 0
	for i := 0; i < n; i++ {
		lis[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if a[i] > a[j] && lis[i] < lis[j] +1 {
				lis[i] = lis[j]+1
			}
		}
	}
	for i := 0; i < n; i++ {
		if max < lis[i] {
			max = lis[i]
		}
	}
	return max
}

func main(){
	a := []int {10, 9, 2, 5, 3, 7, 101, 18}
	n := len(a)
	fmt.Printf("Length of lis %d\n", lis(a, n))
}

