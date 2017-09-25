package main

import(
	"fmt"
)

const(
	a = [3][3]int{{1, 2, 3}, {4, 8, 2}, {1, 5, 3}}
)

func min_path(a [3][3]int, m int, n int) int{
	var tc[m+1][n+1] int
	tc[0][0] = a[0][0]

	//initialising the column
	for i := 1; i <= m; i++ {   
		tc[i][0] = tc[i-1][0] + a[i][0]
	}
	//initialising the row
	for j := 0; j <= n; j++{
		tc[0][j] = tc[0][j-1] + a[0][j]
	}

	for i := 1; i <= m; i++ {
		for j := 0; j <= n; j++ {
			tc[i][j] = min(tc[i-1][j], tc[i][j-1], tc[i-1][j-1]) +a[i][j]
		}
	}
	return tc[m][n]
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
	//var a = [3][3]int{{1, 2, 3}, {4, 8, 2}, {1, 5, 3}}
	fmt.Printf("%d\n", min_path(a, 2, 2))
}

