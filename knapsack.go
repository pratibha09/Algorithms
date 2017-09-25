package main

import(
	"fmt"
)

func max(a, b int)int {
	if a >b {
		return a
	}
	return b
}

func knapsack(a []int, b []int, W int, n int)int{
	if W == 0 || n == 0{
		return 0
	}
	if b[n-1] > W {
		return knapsack(a, b, W, n-1)
	}
	return max(a[n-1]+ knapsack(a, b, W-b[n-1], n-1), knapsack(a, b, W, n-1))
	
}

func main(){
	a := []int{60, 100, 120}
	b := []int{10, 20, 30}
	W := 50
	n := len(a)
	
	fmt.Printf("%d\n",knapsack(a, b, W, n))
}


