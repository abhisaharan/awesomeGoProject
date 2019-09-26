package main

import (
	"fmt"
)

const (
	_ = iota
	year1 = iota * 2
	year2 = iota * 2
	year3 = iota * 3

)
func main()  {
	x := 255
	fmt.Printf("%d\t%b\t%#X\n",x,x,x)
	b := x << 1
	fmt.Printf("%d\t%b\t%#X\n", b, b, b)
	fmt.Println(year1, year2, year3)



}