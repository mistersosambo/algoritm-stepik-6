package main

import "fmt"

func main() {
	var a int
	count := 0
	fmt.Scan(&a)
	for j := a; j != 0; j /= 10 {
		count += j % 10
	}
	fmt.Print(count)
}
