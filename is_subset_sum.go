package main

import(
	"fmt"
)

func is_subset_sum(a []int, n int, sum int) bool{
	if sum == 0 {
		return true
	}
	if n == 0 && sum != 0 {
		return false
	}
	if a[n-1] > sum {
		return is_subset_sum(a, n-1, sum)
	}
	return is_subset_sum(a, n-1, sum) || is_subset_sum(a, n-1, sum-a[n-1])
}

func partition(a []int, n int) bool{
	sum := 0
	for i := 0; i < n; i++ {
		sum += a[i]
	}
	if sum %2 != 0 {
		return false	
	}
	return is_subset_sum(a, n, sum/2)
}

func main(){
	a := [] int{1, 2, 3, 4, 5, 6}
	n := len(a)
	if partition(a, n) == true {	
		fmt.Printf("Can be divided into equal subsets of equal sum\n")
	}else{
		fmt.Printf("Cant be divided!\n")
	}
}

