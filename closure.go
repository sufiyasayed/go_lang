package main

import "fmt"

func closure() func() int {
	return func() int {
		return 1
	}
}
func main() {
	x := closure()
	fmt.Println(x())
}
