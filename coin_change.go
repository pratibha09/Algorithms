package main

import(
	"fmt"
)

func coin_change(a []int, m int, n int) int{
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	if m <= 0 && n >= 1 {
		return 0
	}
	return coin_change(a, m-1, n) + coin_change(a, m, n-a[m-1])
}

func main(){
	a := []int{3, 2, 5}
	m := len(a)
	fmt.Printf("%d\n", coin_change(a,m, 11))
	
}

