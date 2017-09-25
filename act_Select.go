package main

import(
	"fmt"
)

func act_select(a []int, b []int, n int){
	i := 0
	fmt.Printf("%d\n", i)
	for j := 1; j < n; j++ {
		if a[j] >= b[i] {
			fmt.Printf("%d\n", j)
			i = j
		}
	}
}

func main(){
	a := []int{1, 3, 0, 5, 8, 5}
	b := []int{2, 4, 6, 7, 9, 9}
	n := len(a)
	act_select(a, b, n)
	
}

