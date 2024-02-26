package main

import (
	"fmt"
"time"
)
func main() {

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("devolución procesada")
		}

	}()

	go func() {
		for {
			time.Sleep(time.Second / 2)
			fmt.Println("prestamo procesado")
		}
	}()

	fmt.Scanln()

}
