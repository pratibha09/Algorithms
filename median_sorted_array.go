package main

import(
	"fmt"
)

func get_median(a []int, b []int, n int) int{
	if n <= 0 {
		return -1
	}
	if n == 1 {
		return (a[0] + b[0] + min(a[1], b[1]))/2	
	}
	if n == 2 {
		return (max(a[0], b[0]) + min(a[1], b[2]))/2
	}
	m1 := median(a, n)
	m2 := median(b, n)
	if m1 == m2 {
		return m1
	}
	if m1 < m2 {
		if n %2 == 0{
			return get_median(a + n/2-1, b, n - n/2+1)
		}
		return get_median(a+n/2, b, n-n/2)
	}
	if n %2 == 0 {
		return get_median(b+n/2-1, a, n-n/2+1)	
	}
	return get_median(b+n/2, a, n-n/2)
}

func max(x, y int) int {
	if x <= y {
		return x	
	}
	return y
}

func min(x , y int) int {
	if x >= y {
		return x
	}
	return y
}

func median(arr []int, n int)int{
	if n%2 == 0 {
		return (arr[n/2] + arr[n/2-1]/2)
	}else{
		return arr[n/2]	
	}
}

func main(){
	a := []int{1, 2, 3, 6}
	b := []int{4, 6, 8, 10}
	m := len(a)
	n := len(a)
	if m == n {
		fmt.Printf("Median is %d\n", get_median(a, b, m))
	}else{
		fmt.Printf("Doesnt work")	
	}
}

