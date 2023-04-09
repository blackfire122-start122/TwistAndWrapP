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

var LoginPage func(Window fyne.Window)

var Client *http.Client

type Product struct {
	Id    string `json:"Id"`
	Image string `json:"Image"`
	Name  string `json:"Name"`
}

type Order struct {
	Id       uint64
	Products []Product
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

type Page interface {
	Window() fyne.Window
	SetWindowContent()
}