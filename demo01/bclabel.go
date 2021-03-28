package main

import "fmt"

func main() {
	fmt.Println("break:")
BEXIT:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				break BEXIT
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println("continue:")
CEXIT:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue CEXIT
			}
			fmt.Println(i, j)
		}
	}
}
