package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	name        string
	description string
	startPrice  int
	items       []*Item
}

type Item struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func NewProduct(productName string, descProduct string, price int) *Product {
	return &Product{
		name:        productName,
		description: descProduct,
		startPrice:  price,
	}
}

func (p *Product) AddItems(jsonData string) error {
	var itemData struct {
		Items []Item `json:"items"`
	}
	err := json.Unmarshal([]byte(jsonData), &itemData)
	if err != nil {
		return err
	}
	p.items = itemData.Items
	return nil
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) GetDesc() string {
	return p.description
}

func (p *Product) GetStartPrice() int {
	return p.startPrice
}

func (p *Product) GetItems() []*Item {
	return p.items
}

func (it *Item) GetItemName() string {
	return it.Name
}

func (it *Item) GetItemPrice() int {
	return it.Price
}

func main() {
	items := `{"items":[{"name":"33 Diamond","price":3000},{"name":"73 Diamond","price":20000},{"name":"100 Diamond","price":50000}]}`

	productML := NewProduct("Mobile Legends", "Game Moba - Moontoon", 1000)
	err := productML.AddItems(items)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("LIST ITEM " + productML.GetName())
	fmt.Println("Description Product " + productML.GetDesc())
	fmt.Println("Product Price ", productML.GetStartPrice())
	fmt.Println()

	for _, item := range productML.GetItems() {
		fmt.Printf("%s - %d\n", item.GetItemName(), item.GetItemPrice())
	}
}
