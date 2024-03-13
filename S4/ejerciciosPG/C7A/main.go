package main

import "fmt"

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

type Servicio struct {
	nombre            string
	precio            float64
	minutosTrabajados int
}

type Mantenimiento struct {
	nombre string
	precio float64
}

func main() {
	// punteros y canales en go
	//productos servicios y mantenimientos

	productos := make(chan float64)
	servicios := make(chan float64)
	mantenimientos := make(chan float64)

	producto1 := Producto{
		nombre:   "Pintura",
		precio:   45000,
		cantidad: 4,
	}
	producto2 := Producto{
		nombre:   "Cemento",
		precio:   55000,
		cantidad: 2,
	}
	producto3 := Producto{
		nombre:   "Arena",
		precio:   15000,
		cantidad: 5,
	}
	producto4 := Producto{
		nombre:   "Sellante",
		precio:   23000,
		cantidad: 1,
	}
	servicio1 := Servicio{
		nombre:            "Limpieza",
		precio:            23000,
		minutosTrabajados: 50,
	}
	servicio2 := Servicio{
		nombre:            "Arreglo",
		precio:            3000,
		minutosTrabajados: 35,
	}
	servicio3 := Servicio{
		nombre:            "Revisión",
		precio:            3000,
		minutosTrabajados: 20,
	}
	mantenimiento1 := Mantenimiento{
		nombre: "Revisión",
		precio: 2000,
	}
	mantenimiento2 := Mantenimiento{
		nombre: "Revisión",
		precio: 5000,
	}
	mantenimiento3 := Mantenimiento{
		nombre: "Revisión",
		precio: 7000,
	}

	go func(p ...Producto) {
		defer close(productos)

		number := sumarProductos(producto1, producto2, producto3, producto4)
		productos <- number

	}(producto1, producto2, producto3, producto4)

	p := <-productos
	fmt.Printf("El valor total de los productos es $%.3f \n", p)

	go func(s ...Servicio) {

		defer close(servicios)

		number := sumarServicios(servicio1, servicio2, servicio3)
		servicios <- number

	}()

	s := <-servicios
	fmt.Printf("El valor total de los servicios es $%.3f \n", s)

	go func(m ...Mantenimiento) {

		defer close(mantenimientos)

		number := sumarMantenimiento(mantenimiento1, mantenimiento2, mantenimiento3)
		mantenimientos <- number

	}()

	m := <-mantenimientos
	fmt.Printf("El valor total de los mantenimientos es $%.3f \n", m)

}

func sumarProductos(p ...Producto) float64 {
	var total float64

	for _, product := range p {
		total += product.precio * float64(product.cantidad)
	}

	return total
}

func sumarServicios(s ...Servicio) float64 {
	var total float64
	var minutos int

	for _, service := range s {
		minutos = service.minutosTrabajados
		if minutos > 30 {
			total += service.precio * float64(minutos)
		} else {
			total += service.precio
		}

	}
	return total
}

func sumarMantenimiento(m ...Mantenimiento) float64 {

	var total float64

	for _, service := range m {
		total += service.precio
	}

	return total

}
