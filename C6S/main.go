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
		fmt.Println("Error creando archivo:", err)
		return
	}
	fmt.Println("Archivo creado:", category)

	/* category1, err := os.Open("./categories.csv")
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
			break
		}
		fmt.Println(linea)
	} */

	var productos []Product = []Product{
		{1, "Producto 1", 10.99, 5, "Muebles"},
		{2, "Producto 2", 20.99, 3, "Automotor"},
		{3, "Producto 3", 30.99, 2, "Herramientas"},
		{4, "Producto 4", 50.99, 3, "Herramientas"},
		{5, "Producto 5", 12.99, 6, "Herramientas"},
	}

	c := getCategoryByName("Herramientas")
	c.assingProducts(productos)

	c1 := getCategoryByName("Automotor")
	c1.assingProducts(productos)

	c2 := getCategoryByName("Muebles")
	c2.assingProducts(productos)

	fmt.Println("--------ver categories-------")
	fmt.Println(c)
	fmt.Println()
	fmt.Println(c1)
	fmt.Println()
	fmt.Println(c2)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println("--------ver inventories actualizados-------")
	fmt.Println()
	fmt.Println()

	updateInventories(c, c1, c2)

	fmt.Println()
	fmt.Println()

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

var categoryList []Category = []Category{
	{0, "Electrodomésticos", []Product{}},
	{0, "Muebles", []Product{}},
	{0, "Herramientas", []Product{}},
	{0, "Pinturas", []Product{}},
	{0, "Aberturas", []Product{}},
	{0, "Construcción", []Product{}},
	{0, "Automotor", []Product{}},
}

var inventories []Category

func (c *Category) assingProducts(products []Product) {

	for i := 0; i < len(products); i++ {
		pCategory := products[i].category
		if c.name == pCategory {
			c.products = append(c.products, products[i])
		}
	}

}

func getCategoryByName(category string) Category {
	if category != "" {
		for _, c := range inventories {
			if strings.ToLower(c.name) == strings.ToLower(category) {
				return c
			}
		}
	}

	return Category{}

}

func createFile() (*os.File, error) {

	file, err := os.Create("categories.csv")
	if err != nil {
		fmt.Println("Error creating the file ")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var cat Category

	var id = 0
	for i := 0; i < len(categoryList); i++ {

		id++
		cat.id = id
		cat.name = categoryList[i].name

		inventories = append(inventories, cat)
		writer.Write(strings.Split(strconv.Itoa(cat.id)+","+cat.name+",", "\n"))

	}
	return file, nil
}

func updateInventories(categories ...Category) {

	inventories = categories
	fmt.Println(inventories)

}

func createEstimaciones() {

}
