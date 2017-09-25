package main

import(
	"fmt"
)

func max(a int, b int)int{
	if a > b{
		return a
	}
	return b
}

func palindromic_sequence(a string, i int, j int)int{
	if i == j {
		return 1
	}
	if a[i] == a[j] && i+1 == j {
		return 2
	}
	if a[i] == a[j] {
		return palindromic_sequence(a, i+1, j-1)+2
	}
	return max(palindromic_sequence(a, i, j-1), palindromic_sequence(a, i+1, j))
}

func main(){
	a := "GEEKSFORGEEKS"
	n := len(a)
	fmt.Printf("The length of the lps is %d\n", palindromic_sequence(a, 0, n-1))
}

