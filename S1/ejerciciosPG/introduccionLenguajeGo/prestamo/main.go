package main

import "fmt"

func main() {

	customer1 := map[string]float64{
		"isEmployee": 1,
		"age":        33,
		"experience": 5,
		"salary":     250000,
	}

	/* customer2 := map[string]float64{
		"isEmployee": 1,
		"age":   30,
		"experience":2,
		"salary":50000,
	} */

	/* customer3 := map[string]float64{
		"isEmployee": 1,
		"age":   19,
		"experience": 3,
		"salary":70000,

	} */

	loanValidator(customer1["age"], customer1["isEmployee"], customer1["experience"], customer1["salary"])
}

func loanValidator(age float64, status float64, experience float64, salary float64) {
	loanCondition := age > 22 && status == 1 && experience > 1
	salaryCondition := salary > 100000

	if loanCondition && salaryCondition {
		fmt.Printf(`the loan request was accepted without interest for salary of %f`, salary)
	} else if loanCondition && !salaryCondition {
		fmt.Printf(`the loan request was accepted with interest for salary of %f`, salary)
	} else {
		fmt.Printf(`the loan request was denied with interest for salary of %f`, salary)
	}

}
