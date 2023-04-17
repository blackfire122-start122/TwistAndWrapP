package checkoutPage

import (
	. "TwistAndWrapP/entrys"
	. "TwistAndWrapP/pkg"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
)

var ListCheckoutPage []*CheckoutPage

func NewCheckoutPage(w fyne.Window, callbackAddOrder func(order Order), callbackGive func(Id uint64)) {
	page := &CheckoutPage{WindowPage: w, callbackAddOrder: callbackAddOrder, callbackGive: callbackGive}
	page.SetWindowContent()
	ListCheckoutPage = append(ListCheckoutPage, page)
	w.Show()
}

type CheckoutPage struct {
	OrdersList       *fyne.Container
	WindowPage       fyne.Window
	callbackAddOrder func(order Order)
	callbackGive     func(Id uint64)
	products         []Product
}

func (c *CheckoutPage) Window() fyne.Window {
	return c.WindowPage
}

func (c *CheckoutPage) SetWindowContent() {
	c.products = GetProducts()
	productCheckboxes := make([]*widget.Check, len(c.products))
	productEntries := make([]*NumericalEntry, len(c.products))
	for i, p := range c.products {
		productEntries[i] = NewNumericalEntry()
		productEntries[i].PlaceHolder = "Count"
		productEntries[i].Disable()

		checkBoxFunc := func(i int) func(bool) {
			return func(b bool) {
				if b {
					productEntries[i].Enable()
				} else {
					productEntries[i].Disable()
				}
			}
		}(i)

		productCheckboxes[i] = widget.NewCheck(p.Name, checkBoxFunc)
	}

	var items []fyne.CanvasObject
	for i := range productCheckboxes {
		items = append(items, container.NewVBox(productCheckboxes[i], productEntries[i]))
	}

	c.OrdersList = container.NewHBox()

	btnAdd := widget.NewButton("Add order", func() {
		var productsOrder []ProductOrder
		for i, checkbox := range productCheckboxes {
			if checkbox.Checked {
				count, err := strconv.ParseUint(productEntries[i].Text, 10, 8)
				if err != nil {
					fmt.Println(err)
				}
				if count <= 0 {
					count = 1
				}
				productsOrder = append(productsOrder, ProductOrder{Product: c.products[i], Count: uint8(count)})
				checkbox.SetChecked(false)
				productEntries[i].SetText("")
				productEntries[i].Disable()
			}
		}
		c.AddOrder(productsOrder, GenerateIdOrderList())
	})

	findEntry := widget.NewEntry()
	findEntry.OnChanged = func(s string) {
		for i, _ := range productCheckboxes {
			if strings.Contains(strings.ToLower(productCheckboxes[i].Text), strings.ToLower(s)) {
				productCheckboxes[i].Show()
				productEntries[i].Show()
			} else {
				productCheckboxes[i].Hide()
				productEntries[i].Hide()
			}
		}
	}

	productList := container.NewHBox(items...)
	content := container.NewVBox(
		findEntry,
		productList,
		btnAdd,
		container.NewHScroll(c.OrdersList),
	)

	c.WindowPage.SetContent(content)
}

func (c *CheckoutPage) AddOrder(productsOrder []ProductOrder, Id uint64) {
	order := Order{Id: Id, Products: productsOrder}

	c.callbackAddOrder(order)

	var item *fyne.Container

	numberLabel := widget.NewLabel(strconv.FormatUint(order.Id, 10))
	statusLabel := widget.NewLabel("worked")
	btnGive := widget.NewButton("Give", func() {
		c.callbackGive(order.Id)
		c.OrdersList.Remove(item)
	})

	btnGive.Disable()

	item = container.New(
		layout.NewHBoxLayout(),
		numberLabel,
		statusLabel,
		btnGive,
	)

	c.OrdersList.Add(item)
}

func (c *CheckoutPage) SetStatusCompleteOrder(id uint64) {
	for _, order := range c.OrdersList.Objects {
		if cont, ok := order.(*fyne.Container); ok {
			if numberLabel, ok := cont.Objects[0].(*widget.Label); ok {
				if number, err := strconv.ParseUint(numberLabel.Text, 10, 64); err == nil && number == id {
					if statusLabel, ok := cont.Objects[1].(*widget.Label); ok {
						if btnGive, ok := cont.Objects[2].(*widget.Button); ok {
							statusLabel.SetText("Complete")
							btnGive.Enable()
							break
						}
					}
				}
			}
		}
	}
}
