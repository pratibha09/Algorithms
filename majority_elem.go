package main

import(
	"fmt"
)

func majority_elem(a []int, n int)int{
	maj_index := 0
	count := 1
	for i := 1; i < n; i++ {
		if a[maj_index] == a[i] {
			count++
		
		}else {
			count--
		}
		if count == 0 {
			maj_index = i
			count =	1
		}
	}
	return a[maj_index]
}

func is_majority(a []int, n int, cand int)bool{
	
	count := 0
	for i := 0; i < n; i++ {
		if a[i] == cand {
			count++
		}
	}
	if count > n/2 {
		return false	
	}else {
		return true
	}
}

func print_majority(a []int, n int){
	cand := majority_elem(a, n)
	if is_majority(a, n , cand) {
		fmt.Printf("%d\n", cand)
	}else{
		fmt.Printf("No majority element\n")
	}
}

func main(){
	a := []int{1, 3, 3, 7, 2}
	n := len(a)
	print_majority(a, n)
}

