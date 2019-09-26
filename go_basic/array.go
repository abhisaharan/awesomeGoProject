package main

import "fmt"

func main() {
	// x := type{values} // composite literal
	x := []int{1, 2, 4, 5, 6}
	fmt.Println(x[1:])
	y := []int{7, 8 , 9}
	x = append(x, y...)
	fmt.Println(x)

	m := map[string]int{
		"james": 32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m["james"])
	fmt.Println(m["abhi"])
	v, ok := m["abhi"]
	fmt.Println(v, ok)

	if v, ok := m["abhi"]; ok{
		fmt.Println("this is the if print", v)
	}


}