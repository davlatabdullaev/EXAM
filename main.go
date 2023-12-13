package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	ID       string `json:"id"`
	Category string `json:"category"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

type Basket struct {
	ID       string    `json:"id"`
	Products []Product `json:"products"`
	Total    int       `json:"total"`
}

type Customer struct {
	ID        string  `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Cash      int     `json:"cash"`
	Basket    Basket  `json:"basket"`
}

func main() {
	// JSON faylni o'qish
	fileContent, err := os.ReadFile("customers.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}
  

	// JSON dan structga o'tkazish
	var readCustomers []Customer
	err = json.Unmarshal(fileContent, &readCustomers)
	if err != nil {
		fmt.Println("Error unmarshall :", err)
		return
	}

		//task_1(readCustomers)
		//task_2(readCustomers)
		//task_3(readCustomers)
		//task_4(readCustomers)
		//task_5(readCustomers)
		//task_6(readCustomers)
		//task_8(readCustomers)
}
