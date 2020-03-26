package main

import (
	"fmt"
)

var j int = 0

func main() {
	fmt.Println("Please input your decimal")
	fmt.Scan(&j)
	fmt.Printf("j = %d, in binary format =%b\n", j, j)
}
