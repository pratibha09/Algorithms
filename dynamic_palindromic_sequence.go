package main

import(
	"fmt"
)

const(
	a = "GEEKSFORGEEKS"
	n = len(a)
)

func max(a int, b int)int{
	if a > b{
		return a
	}
	return b
}

func palindromic_sequence(a string)int{
	var L[n][n] int
	for i := 0; i < n; i++ {
		L[i][i] = 1
	}
	for cl := 2; cl <= n; cl++ {
		for i := 0; i < n-cl+1; i++ {
			j := i+cl-1
			if a[i] == a[j] && cl ==2 {
				L[i][j]=2			
			}else if a[i] == a[j]{
				L[i][j] = L[i+1][j-1] +2
			}else{
				L[i][j] = max(L[i][j-1], L[i+1][j])
			}
		}
	}
	return L[0][n-1]
}

func main(){
	fmt.Printf("The length of the lps is %d\n", palindromic_sequence(a))
}

