package main

import (
	"fmt"

	"github.com/hotaka-makiuchi/golang/package/sub1"
	"github.com/hotaka-makiuchi/golang/package/sub2"
)

func main() {
	fmt.Printf("hello, world\n")

	sub1.Module1()
	sub2.Module2()
}
