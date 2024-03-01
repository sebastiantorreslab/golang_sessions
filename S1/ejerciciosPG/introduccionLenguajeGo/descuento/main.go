package main

import "fmt"

func main() {

	clothes := []string{"t-shirt", "pants", "socks", "jacket", "wallet"}
	fmt.Println("Total price with discounts: ");
	fmt.Println((setDiscount(clothes)));

}

// precio y descuento

func setDiscount(clothes []string) float64 {

	var price float64;
	price = 2000;

	discount := map[string]float64{
		"t-shirt": 0.1,
		"pants":   0.2,
		"other":   0.3,
	}

	 var output float64;

	for i := 0; i < len(clothes); i++ {
		if (clothes[i] == "t-shirt"){
			output += price * (1-discount[clothes[i]])
		}else if(clothes[i] == "pants") {
			output += price * (1-discount[clothes[i]])
		}else{
			output+= price * (1-discount["other"])
		}
	}

	return output;

}
