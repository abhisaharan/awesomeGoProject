package main

import "fmt"

func main()  {
	fav := "blonde"
	switch fav {
	case "brunette":
		println("should not print")
	case "blonde":
		println("should print")
	}

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}


