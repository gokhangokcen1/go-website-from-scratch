package main

import "fmt"

func EnBuyuk(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(EnBuyuk(3, 5))
}
