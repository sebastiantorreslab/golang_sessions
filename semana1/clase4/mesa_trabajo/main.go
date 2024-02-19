package main

import (
	"fmt"
)



type Product struct {
	ID int
	Name string
	Price float64
	Description string
	Category string
}

var Products = []Product{
    {1,"aceite",5000,"uso industrial","supermercado"},
    {2,"arroz",5000,"uso industrial","supermercado"},
    {3,"cebolla",2000,"uso industrial","supermercado"},
	{4,"cebolla",6000,"uso industrial","supermercado"},
}

func (p *Product) Save() {
	Products = append(Products, *p)

}

func (p *Product) GetAll() {
	for _, product := range Products {
		fmt.Printf("\nProducto %d, Nombre: %s,Precio: %2f, Categoria: %s\n",product.ID,product.Name,product.Price,product.Category);
	}
	
}

func getById(id int) *Product{
	for _, product := range Products {
        if product.ID == id {
            return &product
        }
    }
    return nil

}


func main(){

	// mostrar todos los productos
	fmt.Printf("\nMostrar todos los productos\n")
	p := Product{}
    p.GetAll()


	fmt.Printf("\nTraer producto por ID\n")

	p1 := getById(2)
	fmt.Printf("\nProducto %d, Nombre: %s,Precio: %2f, Categoria: %s\n",p1.ID,p1.Name,p1.Price,p1.Category);

	fmt.Printf("\nGuardar producto\n")
	p2 := Product{5,"mantequilla",5000,"uso industrial","supermercado"}
	p2.Save()


	fmt.Printf("\nMostrar todos los productos con nuevo agregado\n")
	p.GetAll();
	
}