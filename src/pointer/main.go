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

func NewProduct(productName string, descProduct string, price int) *Product {
	return &Product{
		name:        productName,
		description: descProduct,
		startPrice:  price,
	}
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

type ItemData struct {
	Items []*Item `json:"items"`
}

type Item struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (it *Item) GetItemName() string {
	return it.Name
}

func (it *Item) GetItemPrice() int {
	return it.Price
}

func (p *Product) AddItems(itemsv []*Item) *Product {
	for _, ItemValue := range itemsv {
		p.items = append(p.items, ItemValue)
	}
	return p
}

func (p *Product) GetItems() []*Item {
	return p.items
}

func main() {
	productML := NewProduct("Mobile Legends", "Game Moba - Moontoon", 1000)

	items := "{\"items\":[{\"name\":\"33 Diamond\",\"price\":3000},{\"name\":\"73 Diamond\",\"price\":20000}]}"

	itemN := new(ItemData)

	itemJson := json.Unmarshal([]byte(items), itemN)
	if itemJson != nil {
		fmt.Println(itemJson)
	}

	productML.AddItems(itemN.Items)

	for _, muncul := range productML.items {
		fmt.Printf("Name Game : %s - Descript Product: %s - Price Product: %d \n Product Items Name : %s \n Product Items Price : %d \n", productML.GetName(), productML.GetDesc(), productML.GetStartPrice(), muncul.GetItemName(), muncul.GetItemPrice())
	}

}
