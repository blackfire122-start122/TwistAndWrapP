package checkoutPage

import (
	. "TwistAndWrapP/informationPage"
	. "TwistAndWrapP/pkg"
	. "TwistAndWrapP/workedPage"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
)

var ListCheckoutPage []*CheckoutPage

func NewCheckoutPage(w fyne.Window) Page {
	page := &CheckoutPage{WindowPage: w}
	page.SetWindowContent()
	ListCheckoutPage = append(ListCheckoutPage, page)
	w.Show()
	return page
}

type CheckoutPage struct {
	OrdersList *fyne.Container
	WindowPage fyne.Window
}

func (c *CheckoutPage) Window() fyne.Window {
	return c.WindowPage
}

func (c *CheckoutPage) SetWindowContent() {
	go func() {
		for {
			for _, page := range ListInformationPage {
				for _, order := range page.ListOrders.Objects {
					if cont, ok := order.(*fyne.Container); ok {
						if statusLabel, ok := cont.Objects[1].(*widget.Label); ok {
							if statusLabel.Text == "Complete" {
								if numberLabel, ok := cont.Objects[0].(*widget.Label); ok {
									if number, err := strconv.ParseUint(numberLabel.Text, 10, 64); err == nil {
										c.SetStatusCompleteOrder(number)
										c.Window().Content().Refresh()
									}
								}
							}
						}
					}
				}
			}
		}
	}()

	products := GetProducts()
	productCheckboxes := make([]*widget.Check, len(products))
	for i, p := range products {
		productCheckboxes[i] = widget.NewCheck(p.Name, nil)
	}

	var items []fyne.CanvasObject
	for _, c := range productCheckboxes {
		items = append(items, c)
	}

	c.OrdersList = container.NewHBox()

	btnAdd := widget.NewButton("Add order", func() {
		var productsOrder []Product
		for i, checkbox := range productCheckboxes {
			if checkbox.Checked {
				productsOrder = append(productsOrder, products[i])
			}
		}

		order := Order{Id: GenerateIdOrderList(), Products: productsOrder}

		for _, page := range ListWorkedPage {
			page.AddOrder(order)
			page.Window().Content().Refresh()
		}

		for _, page := range ListInformationPage {
			page.AddOrder(order)
			page.Window().Content().Refresh()
		}

		var item *fyne.Container

		numberLabel := widget.NewLabel(strconv.FormatUint(order.Id, 10))
		statusLabel := widget.NewLabel("worked")
		btnGive := widget.NewButton("Give", func() {
			for _, page := range ListInformationPage {
				page.DeleteOrder(order.Id)
				page.Window().Content().Refresh()
			}
			c.OrdersList.Remove(item)
		})

		item = container.New(
			layout.NewHBoxLayout(),
			numberLabel,
			statusLabel,
			btnGive,
		)

		c.OrdersList.Add(item)

	})

	findEntry := widget.NewEntry()
	findEntry.OnChanged = func(s string) {
		for _, checkbox := range productCheckboxes {
			if strings.Contains(strings.ToLower(checkbox.Text), strings.ToLower(s)) {
				checkbox.Show()
			} else {
				checkbox.Hide()
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

func (c *CheckoutPage) SetStatusCompleteOrder(id uint64) {
	for _, order := range c.OrdersList.Objects {
		if cont, ok := order.(*fyne.Container); ok {
			if numberLabel, ok := cont.Objects[0].(*widget.Label); ok {
				if number, err := strconv.ParseUint(numberLabel.Text, 10, 64); err == nil && number == id {
					if statusLabel, ok := cont.Objects[1].(*widget.Label); ok {
						statusLabel.SetText("Complete")
						break
					}
				}
			}
		}
	}
}

// GenerateIdOrderList ToDo: Need optimize
func GenerateIdOrderList() uint64 {
	var id uint64 = 0
	checkId(&id)

	OrderListId = append(OrderListId, id)

	return id
}

func checkId(id *uint64) {
	for _, val := range OrderListId {
		if val == *id {
			*id += 1
			checkId(id)
		}
	}
}
