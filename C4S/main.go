package main

import "fmt"

func main() {

	// Administrar productos

	producto1 := factory("p", 1000)
	fmt.Println("Precio producto pequeño:", producto1.Precio())

	producto2 := factory("m", 2000)
	fmt.Println("Precio producto mediano:", producto2.Precio())

	producto3 := factory("g", 3000)
	fmt.Println("Precio producto grande:", producto3.Precio())

}

type Producto interface {
	Precio() float64
}

type Pequeno struct {
	precio float64
}

type Mediano struct {
	precio float64
}

type Grande struct {
	precio float64
}

func (p Pequeno) Precio() float64 {
	return p.precio
}

func (m Mediano) Precio() float64 {
	return m.precio * (1 + 0.03)
}

func (g Grande) Precio() float64 {
	return (g.precio * (1 + 0.06)) + 2500
}

func factory(tipo string, precio float64) Producto {
	switch tipo {
	case "p":
		return Pequeno{precio: precio}
	case "m":
		return Mediano{precio: precio}
	case "g":
		return Grande{precio: precio}
	}
	return nil
}
