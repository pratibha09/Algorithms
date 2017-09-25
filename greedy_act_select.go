package main

import(
	"fmt"
)

func activity_selection(start []int, finish []int, n int){
	fmt.Printf("Activities chosen are: \n")
	i := 0
	for j := 1; j < n; j++ {
		if start[j] >= finish[i] {
			fmt.Printf("%d\n", j)
			i = j
		}
	}
}

func main(){
	start := [] int{1, 3, 0, 5, 8, 5}
	finish := [] int{2, 4, 6, 7, 9, 9}
	n := len(start)
	activity_selection(start, finish, n)
}

