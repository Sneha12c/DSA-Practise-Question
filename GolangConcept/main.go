package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(1 + 1)
	// string
	fmt.Println("hello golang")
	// bool
	fmt.Println(true)
	fmt.Println(false)
	// floats
	fmt.Println(10.5)
	fmt.Println(7.0 / 3.0)

	// var name string = "golang"
	// infer
	// var name = "golang"
	// var isAdult bool = true

	// var age int = 30

	// shorthand syntax
	name := "golang"
	fmt.Println(name)

	// var name string
	// name = "golang"

	// var price float32 = 50.5
	// var price = 50.5
	// price := 50.5

	// fmt.Println(price)

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is a weekend")
	default:
		fmt.Println("it is weekday")

	}

	var nums [4]bool

	nums[2] = true
	fmt.Println(nums)

	var nums2 []int
	fmt.Println(nums2)

	var nums3 []int
	nums3 = append(nums3, 1)
	// fmt.Println(nums3[2]) // error

	mp := make(map[string]string)

	mp["hland"] = "bb"
	mp["hlandbb"] = "bbjbchje"
	fmt.Println(mp)

	v, ok := mp["hland"]
	fmt.Println(v, ok)

	for k, v := range mp {
		fmt.Println(k, v)
	}

	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
