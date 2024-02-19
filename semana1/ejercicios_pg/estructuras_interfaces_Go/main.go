package main

import (
	"fmt"
)

type Alumno struct{
	nombre string
	apellido string
	dni int
	fecha string
}



func (a Alumno) Detalle(){
	
	fmt.Printf("\nDetalle del alumno:\n")
	fmt.Printf("\nNombre: %s \nApellido %s \nDNI: %v \nFecha: %s", a.nombre,a.apellido,a.dni,a.fecha)
}


func main(){

	alumno := Alumno {
		nombre: "Sebastian",
		apellido: "Torres",
		dni: 1245,
		fecha: "25/10/2023",
	}

	alumno1 := Alumno {
		nombre: "Carlos",
		apellido: "Crespo",
		dni: 8875654,
		fecha: "29/01/2015",
	}

	Alumno.Detalle(alumno);
	Alumno.Detalle(alumno1);

}