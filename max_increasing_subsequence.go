package main

import(
	"fmt"
)

/*const(
	//a = [] int{1, 101, 2, 3, 100, 4, 5}
	n = len(a)
)*/

func max_increasing_subsequence(a []int, n int)int{
	if n == 0{
		return 0
	}
	max := 0
	var m[n] int
	for i := 0; i<n; i++{
		m[i]= a[i]
	}

	for i:= 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if a[i] > a[j] && m[i] < m[j] + a[i]{
				m[i] = m[j]+a[i]
			}
		}
	}
	for i := 0; i < n; i++{
		if max < m[i]{
			max = m[i]		
		}
	}
	return max
}

func main(){
	a := [] int{1, 101, 2, 3, 100, 4, 5}
	n := len(a)
	fmt.Printf("Sum of max sum increasing subsequence is %d", max_increasing_subsequence(a, n))
}

