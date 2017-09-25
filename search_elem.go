package main

import(
	"fmt"
)

func search_elem(a []int, x int){
	for i := 0; i < len(a); i++ {
		if a[i] == x {
			fmt.Printf("The element %d is present at %d\n:", x, a[i])
		}else {
			fmt.Printf("The element %d isn't present %d\n:", x, a[i])
		}
	}
}

func reverse_array(a []int, start int, end int){
	var temp int
	for start < end {
		temp = a[start]
		a[start] = a[end]
		a[end] = temp
		start++
		end--
	}
}

func printarray(a []int){
	for i := 0; i < len(a); i++ {
		fmt.Printf("\n%d\n", a[i])
	}
}

func main(){
	a := []int{2, 3, 4, 5, 6, 7}
	search_elem(a, 5)
	reverse_array(a, 0, 5)
	printarray(a)
}
