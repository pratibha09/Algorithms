package main

import(
	"fmt"
)

const (
	a = "sunday"
	b = "funday"
	m = len(a)
	n = len(b)
)

func edit_distance(a string, b string) int{
	var dp[m+1][n+1]int
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 {
				dp[i][j] = 0
			}else if j == 0 {
				dp[i][j] = i
			}else if (a[i-1] == b[j-1]) {
				dp[i][j] = dp[i-1][j-1]
			}else{
				dp[i][j] = 1+min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1])
			}
		}
	}
	return dp[m][n]
}

func min(a, b, c int)int{
	if a < b{
		if a < c{
			return a
		}
		return c	
	}else if b < c {
		return b
	}
	return c
}

func main(){
	fmt.Println(edit_distance(a, b))
}

