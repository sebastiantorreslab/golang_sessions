package main

import (
	"fmt"
	"os"
)

type SalaryError struct {
	message string
}

func (e *SalaryError) Error() string {
	return e.message
}

func main() {

	// manejo de errores en go
	// Ejercicio 1 - Impuestos de salario
	/*
			En la función main , definir una variable salary y asignarle un valor de tipo “int”. Luego, crear
		un error personalizado con un struct que implemente Error() con el mensaje “Error: el salario
		ingresado no alcanza el mínimo imponible” y lanzarlo en caso de que salary sea menor a
		150.000. De lo contrario, imprimir por consola el mensaje “Debe pagar impuesto”.

	*/

	salary := 200000

	if salary < 150000 {
		err := &SalaryError{"Error: el salario ingresado no alcanza el mínimo imponible"}
		fmt.Println(err.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}

	// Ejercicio 2 - Datos de clientes

	/*
			Un estudio contable necesita acceder a los datos de sus empleados para poder realizar
		distintas liquidaciones. Para ello, cuentan con todo el detalle necesario en un archivo TXT .
		1. Desarrollar el código necesario para leer los datos de un archivo llamado
		“customers.txt”. Sin embargo, debemos tener en cuenta que la empresa no nos ha
		pasado el archivo a leer por el programa. Dado que no contamos con el archivo
		necesario, se obtendrá un error. En tal caso, el programa deberá arrojar un panic al
		intentar leer un archivo que no existe, mostrando el mensaje “El archivo indicado no
		fue encontrado o está dañado”.

		2. Más allá de eso, deberá siempre imprimirse por consola “Ejecución finalizada”.
	*/

	filename := "customers.txt"
	readCustomersData(filename)

}

func readCustomersData(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}
	defer func() {
		file.Close()
		fmt.Println("Ejecución finalizada")
	}()
}
