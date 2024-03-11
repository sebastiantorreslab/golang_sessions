package main

import (
	"errors"
	"fmt"
)

/*
Impuestos de salario
En la función main, definir una variable llamada salary y asignarle un valor de tipo “int”.
Crear un error con errors.New() con el mensaje “salario muy bajo" y lanzarlo en caso de que salary
sea menor o igual a 10.000.
La validación debe ser hecha con la función Is() dentro del main.
*/

const lowSalary = 10000

var ErrSalary = errors.New("Error:")

func main() {
	salary := 1000
	err := validarSalario(1000)
	if errors.Is(err, ErrSalary) {
		fmt.Printf("Error: el mínimo imponible es de 150.000 y el salario ingresado es %v", salary)
	}
}

func validarSalario(salary int) error {
	if salary <= lowSalary {
		return fmt.Errorf("%w ", ErrSalary)
	}
	return nil
}
