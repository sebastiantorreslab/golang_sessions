package main

import (
	"fmt"
)


func main(){

	//Ejercicio 1 impuestos de salario
	var employees =  map[string]float64{"a":50000,"b":26000,"c":45000,"d":250000};
	taxSalary(employees);

	//Ejercicio 2 calcular promedio

	fmt.Printf(`elpromedio de las notas es %d`,avgSalary(1,5,4,4,3,5,3,2,2,4,5))

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

// ejercicio 2


func avgSalary(califications ...int) int{
	sum := 0
	q := len(califications)
	avg := 0
	for  i := 0 ; i < q; i++ {
		sum += califications[i]
	}
	 avg = sum/q
		return avg
}