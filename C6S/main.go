package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Ejercicio - mayorista

	/*
		Se necesita un software desarrollado en Go que imprima las estimaciones de lo que
		generarían, los productos de cada categoría en un mayorista, en ventas de productos para
		el hogar. Para eso se detallarán los pasos para conseguirlo:

	*/

	category, err := createFile()
	if err != nil {
		return
	}
	fmt.Println("Archivo creado:", category)

	category1, err := os.Open("./categories.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo CSV:", err)
		return
	}
	defer category1.Close()

	// Crear un lector CSV
	lector := csv.NewReader(category1)

	// Leer y mostrar cada línea del archivo
	for {
		linea, err := lector.Read()
		if err != nil {
			break // Se ha alcanzado el final del archivo
		}
		fmt.Println(linea)
	}

}

type Product struct {
	code            int
	name            string
	currentPrice    float64
	currentQunatity int
	category        string
}

type Category struct {
	id       int
	name     string
	products []Product
}

var categoryList []string = []string{
	"Electrodomésticos",
	"Muebles",
	"Herramientas",
	"Pinturas",
	"Aberturas",
	"Construcción",
	"Automotor",
}

var productos []Product = []Product{
	{1, "Producto 1", 10.99, 5, "Muebles"},
	{2, "Producto 2", 20.99, 3, "Automoviles"},
	{3, "Producto 3", 30.99, 2, "Herramientas"},
	{4, "Producto 3", 50.99, 3, "Herramientas"},
}

func createFile() (*os.File, error) {

	file, err := os.Create("categories.csv")
	if err != nil {
		fmt.Println("Error creating the file ")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	var cat1 = Category{}

	var id = 0
	for i := 0; i < len(categoryList); i++ {
		id++
		cat1.id = id
		cat1.name = categoryList[i]
		writer.Write(strings.Split(strconv.Itoa(cat1.id)+","+cat1.name+",", ";"))

	}

	return file, nil
}
