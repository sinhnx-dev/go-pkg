package main

import (
	"fmt"

	"github.com/sinhnx-dev/go-pkg/numbers"
)

func main() {
	num := 13
	if numbers.IsPrime(num) {
		fmt.Printf("%d is prime\n", num)
	} else {
		fmt.Printf("%d is not prime\n", num)
	}
}
