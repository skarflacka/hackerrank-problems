package main

import (
	"fmt"

)

func isFibo(n int) string{
	fib := make([]int, 0)
	fib = append(fib, 0, 1) 

	for i := 0; i < 80; i++ {
		fib = append(fib, fib[i] + fib[i+1])
	}

	for _, v := range fib {
		if n == v {
			return "IsFibo"
		}
	}
	return "IsNotFibo"
}


func main() {
	var t int
	fmt.Scanf("%d", &t)

	for foo := 0; foo <= t; foo++ {
		var n int
		fmt.Scanf("%d", &n)

		r := isFibo(n)
		fmt.Println(r)
	}
}