package main

import (
	"fmt"
)





func main(){

	categoria,salario := calculateSalary("a",1560);
	fmt.Printf(`El salario para el empleado de categoria %s es de %f`,categoria,salario);

}

func calculateSalary(category string, minutos int64) (string,float64){

	var horasMes = (minutos/60); 

	salary := map[string]float64{
		"a": 3000,
		"b": 1500,
		"c": 1000,
	}

	var response float64;

	if(category == "a"){
		response = salary["a"] * float64(horasMes) * (1+0.5)
	}else if(category == "b"){
		response = salary["b"] * float64(horasMes) * (1+0.2)
	}else if (category == "c"){
		response = salary["c"] * float64(horasMes)
	}

	return category,response;

}