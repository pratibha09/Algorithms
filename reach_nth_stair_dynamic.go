package main

import(
	"fmt"
)

func number_of_ways_to_reach_nth_stair(n int, m int) int {
	res := make([]int, n)
	res[0] = 1
	res[1] = 1
	for i := 2; i < n; i++ {
		res[i] = 0
		for j := 1; j <= m && j <= i; j++ {
			res[i] += res[i-j]
		}
	}
	return res[n-1]
}

func count_ways(s int, m int) int {
	return number_of_ways_to_reach_nth_stair(s+1, m)
}

func main() {
	var inputs int
	fmt.Scanln(&inputs)
	for i := 0; i < inputs; i++ {
	var m, n int
	fmt.Scan(&m)
	fmt.Scan(&n)
	
	fmt.Printf("%d\n", count_ways(m, n))
	}
}

