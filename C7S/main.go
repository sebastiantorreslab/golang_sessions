package main

import "fmt"

func main() {

	c_par := make(chan int)
	c_impar := make(chan int)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go par(c_par)
	go impar(c_impar)

	for _, n := range numbers {
		if n%2 == 0 {
			c_par <- n
		} else if n%2 != 0 {
			c_impar <- n
		}
	}

}

func par(c chan int) {
	for {
		num := <-c
		fmt.Printf("número par : %v ", num)
	}

}

func impar(c chan int) {
	for {
		num := <-c
		fmt.Printf("número impar : %v ", num)
	}

}
