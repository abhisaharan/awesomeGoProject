package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 6
	fmt.Println("Hello world")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Printf("%d\t\t%b", x, x)
}
