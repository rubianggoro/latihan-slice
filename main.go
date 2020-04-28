package main

import "fmt"

func main() {
	s := make([]int, 1)
	fmt.Printf("s = %v\n", s)

	for num := 1; num < 20; num++ {
		i := 0
		for z := 1; z < 20; z++ {
			if num%z == 0 {
				i++
			}
		}
		if i == 2 && num > 1 {
			s = append(s, num)
		}
	}
	fmt.Printf("s = %v\n", s)

}
