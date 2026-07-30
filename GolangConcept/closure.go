package main

import "fmt"

func counter() func() int {
	var count int = 0
	return func() int {
		count += 1
		return count
	}
}

func changenum(num *int) {
	*num = 5
	fmt.Println(*num)
	fmt.Println("Memory address", &num)
}

func main() {
	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())

	num := 1

	changenum(&num)
	fmt.Println("Memory address", &num)
	fmt.Println("After changeNum in main", num)

}
