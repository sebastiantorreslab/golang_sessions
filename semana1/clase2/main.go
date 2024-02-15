package main

import (
	"fmt"
)


func main(){

	//Ejercicio 1 salary tax
	var employees =  map[string]float64{"a":50000,"b":26000,"c":45000,"d":250000};
	taxSalary(employees);

	//Ejercicio 2 avg



}


func taxSalary(employees map[string]float64){
	
	baseTax:= 0.17
	extraTax:= 0.27

	for key, value := range employees{
		if(value > 150000){
			salary := value * (1-extraTax)
			fmt.Printf(`el salario para %s es %f`,key,salary)
		}else if(value > 50000){
			salary := value * (1-baseTax)
			fmt.Printf(`el salario para %s es %f`,key,salary)
		} else {
			fmt.Printf(`el salario para %s es %f`,key,value)
		}

	}


}



func avgSalary(califications ...int) int{

	for int i  = 0 ; i < califications; i++ {

	}



return 0.0
}