package main

import(
	"fmt"
)

const(
	n = 9
)

func f(n int)int{
	var f[n] int
	f[0] = 0
	f[1] = 1
	for i := 2; i < n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func main(){
	fmt.Printf("%d\n", f(n))

}

