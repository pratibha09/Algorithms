package main

import(
	"fmt"
)

func mincost(a [3][3]int, m int, n int)int {
	if n < 0 || m < 0 {
		return 1
	}else if m == 0 && n == 0{
		return a[m][n]
	}else {
		return a[m][n] + min(mincost(a, m-1, n-1), mincost(a, m-1, n), mincost(a, m, n-1))
	}
}

func min(x, y, z int)int {
	if x < y {
		if x < z {
			return x
		}
		return z
	}else if y < z {
		return y
	}
	return z
}

func main(){
	var a = [3][3]int{{1, 2, 3}, {4, 8, 2}, {1, 5, 3}}
	fmt.Printf("%d\n", mincost(a, 2, 2))
}

