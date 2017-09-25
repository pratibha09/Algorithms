package main

import(
	"fmt"
)

func lcs_dynamic(a, string, b string, m int, n int) string{
	
	if m == 0 || n == 0 {
		return 
	}else if a[m-1] == b[n-1]{
		return 1 + a[m-1]
	}else{
		return max()
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
	fmt.Printf("Length of LCS is %d\n", lcs_dynamic(a,b,m,n))

}
