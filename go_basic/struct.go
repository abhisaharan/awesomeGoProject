package main

import "fmt"

type person struct {
	first string
	last string
}

type profession struct {
	person
	job bool
}

func main()  {
	p1:= person{
		first: "abhi",
		last: "saharan",
	}

	p2:= person{
		first: "tik",
		last: "tok",
	}

	p3:= profession{
		person: person{
			first: "Mary",
			last: "Jane",
		},
		job: true,
	}

	// anonymous struct
	a1 := struct {
		first string
		last string
	}{
		first: "Kit",
		last: "Kat",
	}

	fmt.Println(p1.first, p1.last)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(a1)
}