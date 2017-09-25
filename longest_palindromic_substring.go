package main

import(
	"fmt"
)

func longest_pal_substring(str []rune) int{
	maxlength := 1
	start := 0
	length := len(str)
	var low, high int
	
	for i := 1; i < length; i++ {
		low = i - 1
		high = i
		for low >= 0 && high < length && str[low] == str[high] {
			if high -low +1 > maxlength {
				start = low
				maxlength = high -low+1
			}
			low = low -1
			high = high+1
		}
		low = i -1 
		high = i+1
		for low >= 0 && high < length && str[low] == str[high] {
			if high -low+1 > maxlength {
				start = low
				maxlength = high-low+1
			}
			low = low-1
			high = high+1
		}
	}
	fmt.Printf("Longest palindrome substring is \n")
	printsubstr(str, start, start+maxlength-1)
	return maxlength
}

func printsubstr(str []rune, low int, high int) {
	for i := low; i <= high; i++ {
		fmt.Printf("%s", string(str[i]))
	}
}


func main(){
	str := []rune("aaa")
	ans := longest_pal_substring(str)
	fmt.Printf("\nLength is %d\n", int(ans))

}

