package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Product struct {
	ID       int
	precio   float64
	cantidad int
}

func main() {

	product1 := Product{
		ID:       1,
		precio:   15000,
		cantidad: 4,
	}
	product2 := Product{
		ID:       2,
		precio:   25000,
		cantidad: 3,
	}
	product3 := Product{
		ID:       3,
		precio:   42000,
		cantidad: 6,
	}
	product4 := Product{
		ID:       4,
		precio:   15000,
		cantidad: 1,
	}

	saveFile(product1, product2, product3, product4)

}

func saveFile(products ...Product) error {
	file, err := os.Create("products.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{"products", "info"})
	if err != nil {
		return err
	}

	data := make(map[string]string)

	for _, p := range products {

		idString := strconv.Itoa(p.ID)
		data["ID"] = idString

		precioString := strconv.FormatFloat(p.precio, 'f', -1, 64)
		data["Precio"] = precioString

		cantidadString := strconv.Itoa(p.cantidad)
		data["Cantidad"] = cantidadString

		fmt.Println(data)

		for clave, valor := range data {
			err := writer.Write([]string{clave, fmt.Sprintf("%s", valor) + "\n"})
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	}

	return nil
}
