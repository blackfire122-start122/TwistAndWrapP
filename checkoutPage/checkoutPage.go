package checkoutPage

import (
	. "TwistAndWrapP/pkg"
	. "TwistAndWrapP/workedPage"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var ListCheckoutPage []*CheckoutPage

func NewCheckoutPage(w fyne.Window) Page {
	page := CheckoutPage{WindowPage: w}
	page.SetWindowContent()
	ListCheckoutPage = append(ListCheckoutPage, &page)
	w.Show()
	return page
}

type CheckoutPage struct {
	WindowPage fyne.Window
}

func (c CheckoutPage) Window() fyne.Window {
	return c.WindowPage
}

func (c CheckoutPage) SetWindowContent() {
	products := GetProducts()
	productCheckboxes := make([]*widget.Check, len(products))
	for i, p := range products {
		productCheckboxes[i] = widget.NewCheck(p.Name, nil)
	}

	var items []fyne.CanvasObject
	for _, c := range productCheckboxes {
		items = append(items, c)
	}

	btnAdd := widget.NewButton("Add order", func() {
		var productsOrder []Product
		for i, checkbox := range productCheckboxes {
			if checkbox.Checked {
				productsOrder = append(productsOrder, products[i])
			}
		}

		for _, page := range ListWorkedPage {
			page.Window().RequestFocus()
			page.AddOrder(Order{Id: 0, Products: productsOrder})
		}
	})

	productList := container.NewVBox(items...)
	content := container.NewVBox(
		productList,
		btnAdd,
	)

	c.WindowPage.SetContent(content)
}
