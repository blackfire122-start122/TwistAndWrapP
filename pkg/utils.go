package pkg

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"net/http"
)

var InformationPageList []fyne.Window
var CheckoutPageList []fyne.Window
var WorkedPageList []fyne.Window

var ProductList []Product
var Orders []Order

var CheckListProducts []*widget.Check

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
	Status   string
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
