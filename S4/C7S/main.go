package main

import (
	"fmt"
	"time"
)
func main() {

	devolucion := make(chan string)
	prestamo := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			devolucion <- "devolución procesada"
		}

	}()

	go func() {
		for {
			time.Sleep(time.Second / 2)
			prestamo <- "prestamo procesado"
		}
	}()


	go func() {
		for devProcesada := range devolucion {
			fmt.Println(devProcesada);
		}
	
	}()

	go func() {
		for presProcesado := range prestamo {
			fmt.Println(presProcesado);
		}
	
	}()


	// otra forma
	select {
	case msg := <- devolucion:
		fmt.Println(msg)
	}

	fmt.Scanln()

}
