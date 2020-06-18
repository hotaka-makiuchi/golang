package main

import (
	"fmt"

	"github.com/hotaka-makiuchi/golang/test/hello"
)

func main() {
	s := hello.GetHello("テスト")
	fmt.Println(s)
}
