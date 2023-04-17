package pkg

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"net/http"
)

var ProductList []Product

var CheckListProducts []*widget.Check

var OrderListId []uint64

var App fyne.App

var Client *http.Client

type Product struct {
	Id    string `json:"Id"`
	Image string `json:"Image"`
	Name  string `json:"Name"`
}

type ProductOrder struct {
	Product Product
	Count   uint8
}

type Order struct {
	Id       uint64
	Products []ProductOrder
}

func GetProducts() []Product {
	var products []Product
	for i, obj := range CheckListProducts {
		if obj.Checked {
			products = append(products, ProductList[i])
		}
	}
	return products
}

func GenerateIdOrderList() uint64 {
	var id uint64 = 0
	for contains(OrderListId, id) {
		id++
	}
	OrderListId = append(OrderListId, id)
	return id
}

func contains(slice []uint64, item uint64) bool {
	for _, val := range slice {
		if val == item {
			return true
		}
	}
	return false
}
