package main

import(
	"fmt"
)

func lcs(a string, b string, m int, n int) int{
	if m == 0 || n == 0 {
		return 0
	}
	if a[m-1] == b[n-1] {
		return 1 + lcs(a, b, m-1, n-1)
	}else {
		return max(lcs(a, b, m, n-1), lcs(a, b, m-1, n))
	}
}

func max(a, b int)int{
	if a > b{
		return a
	}
	return b
}

func main(){
	a := "ABCDBAASDASD"
	b := "asdb"
	m := len(a)
	n := len(b)
	fmt.Printf("Length of LCS is %d\n", lcs(a, b, m, n))

}

