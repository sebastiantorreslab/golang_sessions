package main

import (
	"fmt"
)


func main(){

	//Ejercicio 1 Listado de nombres
	
	students := []string {"Benajamin","Nahuel","Brenda","Marcos","Pedro","Axel","Alez","Dolores","Federico","Hernan","Leandro","Eduardo","Duvrashka"};

 	students = append(students, "Gabriela");

 	fmt.Println(students);


	//Ejercicio 2 - que edad tiene

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	
	fmt.Printf(`La edad de Benjamin es %d años`,employees["Benjamin"]);

	var counter int;
	for key, value := range employees {
		if(value > 21 && key != ""){
			counter++
		}
		fmt.Printf(`Los trabajadores mayores a  21 años son %d`,counter);
	}

	employees["Federico"] = 25;
	fmt.Println(employees);
	delete(employees,"Pedro");

	fmt.Println("Contenido del mapa:")
	for key, value := range employees {
        fmt.Printf("Clave: %s, Valor: %d\n", key, value)
    }

	

	
}

















