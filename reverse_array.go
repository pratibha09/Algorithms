package main

import(
	"fmt"
)

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
		fmt.Printf("%d\n", a[i])
	}
}

func main(){
	a := []int{2, 3, 4, 5, 6, 7}
	printarray(a)
	fmt.Printf("Reversed array\n")
	reverse_array(a, 0, 5)
	printarray(a)
}

