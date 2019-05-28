package main

import "fmt"

type abhi int
var x abhi
var y int
func main()  {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	y = int(x)
	fmt.Printf("%T\n", y)

	fmt.Println(x, y)



}
