package main

import "fmt"

func main ()  {
	//its an slice built on top of arrays
	x1 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y1 := []int{52, 53, 54, 55}
	z1 := x1[:3]
	z1 = append(z1, x1[6:]...)
	x1 = append(x1, y1...)
	//for i, v := range z1 {
	//	fmt.Printf("this is the index nd value %v %v \n", i, v)
	//}

	x2 := make([]string, 50, 65)
	x2 = append(x2, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`, )
	//for i, v := range x2 {
	//	fmt.Printf("this is the index %v and state %v \n", i, v)
	//}

	//x3 := []string{"James", "Bond", "Shaken, not stirred"}
	//x31 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	//xx3 := [][]string{x3, x31}
	//for i, xs := range xx3 {
	//	fmt.Println("record: ", i)
	//	for j, val := range xs {
	//		fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
	//	}
	//}

	x4 := map[string][]string{
		"bond_james": {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": { `James Bond`, `Literature`, `Computer Science`},
	}
	x4["abhi"] = []string{"python", "cisco", "edyta" }
	delete(x4, "bond_james")
	//fmt.Println(x4)
	for k, v := range x4 {
		fmt.Println("this is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}


