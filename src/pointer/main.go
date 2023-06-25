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
	name  string
	price int
}

type ItemData struct {
	Items []*Item `json:"items"`
}

func NewProduct(productName string, descProduct string, price int) *Product {
	return &Product{
		name:        productName,
		description: descProduct,
		startPrice:  price,
	}
}

func (p *Product) AddItems() *Product {

	return p
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
	return it.name
}

func (it *Item) GetItemPrice() int {
	return it.price
}

func main() {
	items := `{"items":[{"name":"33 Diamond","price":3000},{"name":"73 Diamond","price":20000},{"name":"100 Diamond","price":50000}]}`

	var itemData ItemData
	itemJson := json.Unmarshal([]byte(items), &itemData)
	if itemJson != nil {
		fmt.Println(itemJson)
	}
	fmt.Println()
	fmt.Println(itemData)

	for _, value := range itemData.Items {
		fmt.Println("Name: ", value.name, "Jenis: ", value.price)
	}

	productML := NewProduct("Mobile Legends", "Game Moba - Moontoon", 1000)

	fmt.Println(productML.GetName())
	fmt.Println(productML.GetDesc())
	fmt.Println(productML.GetStartPrice())

	//fmt.Println("LIST ITEM " + productML.GetName())
	//fmt.Println("33 Diamond - " + "3000")
	//fmt.Println("73 Diamond - " + "20000")
	//fmt.Println("100 Diamond - " + "50000")
}
