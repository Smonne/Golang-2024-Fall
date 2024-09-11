package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func productToJSON(p Product) (string, error) {
	jsonData, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func jsonToProduct(jsonString string) (Product, error) {
	var product Product
	err := json.Unmarshal([]byte(jsonString), &product)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func reformat() {
	product := Product{Name: "Laptop", Price: 999.99, Quantity: 10}

	jsonString, err := productToJSON(product)
	if err != nil {
		fmt.Println("Error converting product to JSON:", err)
	} else {
		fmt.Println("Product in JSON format:", jsonString)
	}

	jsonData := `{"name":"Smartphone","price":599.99,"quantity":20}`
	decodedProduct, err := jsonToProduct(jsonData)
	if err != nil {
		fmt.Println("Error decoding JSON to product:", err)
	} else {
		fmt.Printf("Decoded Product: %+v\n", decodedProduct)
	}
}
