package main

import (
	"fmt"
)

type Cliente struct {
	nombre string
	telefono string
	next *Cliente
}

var primerCliente *Cliente


func main() {

	AddCliente("Pepe", "555")
	AddCliente("Juan", "556")
	AddCliente("Rodrigo", "557")
	AddCliente("David", "558")

	ListaClientes()

}


func AddCliente(nombre string, telefono string){

	var cliente = Cliente{
		nombre: nombre,
		telefono: telefono,
		next:     nil,
	}

	if primerCliente == nil {
		primerCliente = &cliente
		return
	}

	var nextCliente *Cliente = primerCliente

	for nextCliente.next != nil {
		nextCliente = nextCliente.next
	}
	nextCliente.next = &cliente
}

func ListaClientes() {
	
	if primerCliente == nil{
		return

	}

	var nextCliente *Cliente = primerCliente
	for nextCliente.next != nil{
		fmt.Printf("Nombre :%s Teléfono %s \n",nextCliente.nombre,nextCliente.telefono)
		nextCliente = nextCliente.next
	}

	fmt.Printf("Nombre :%s Teléfono %s \n",nextCliente.nombre,nextCliente.telefono)

}