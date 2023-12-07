package main

import (
	"fmt"

	"github.com/me885/linq-for-go"
)

func main() {
	//Where on slice of ints
	x := []int{1, 2, 3, 4, 5}

	y := linq.SliceWhere(x, func(i int) bool { return i%2 == 0 })

	fmt.Println(y)

	//Select on slice of ints
	x2 := []int{1, 2, 3, 4, 5}

	y2 := linq.SliceSelect(x2, func(i int) int { return i * 2 })

	fmt.Println(y2)

	//Select on slice of structs
	type Person struct {
		Name string
		Age  int
	}

	x3 := []Person{
		{Name: "Bob", Age: 30},
		{Name: "Bill", Age: 36},
	}

	y3 := linq.SliceSelect(x3, func(i Person) string { return i.Name })

	fmt.Println(y3)

	//First on slice of ints
	x4 := []int{1, 2, 3, 4, 5}

	y4, err := linq.SliceFirst(x4)

	fmt.Println(y4, err)

	//Single on slice of length 1
	x5 := []string{"HELLO"}

	y5, err := linq.SliceSingle(x5)

	fmt.Println(y5, err)

	//Single on longer slice
	x6 := []string{"HELLO", "GOOD", "MORNING"}

	y6, err := linq.SliceSingle(x6)

	fmt.Println(y6, err)
}
