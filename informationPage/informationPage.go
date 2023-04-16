package informationPage

import (
	. "TwistAndWrapP/pkg"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

var ListInformationPage []*InformationPage

func NewInformationPage(w fyne.Window) {
	page := &InformationPage{WindowPage: w}
	page.SetWindowContent()
	ListInformationPage = append(ListInformationPage, page)
	w.Show()
}

type InformationPage struct {
	ListOrders *fyne.Container
	WindowPage fyne.Window
}

func (i *InformationPage) Window() fyne.Window {
	return i.WindowPage
}

func (i *InformationPage) SetWindowContent() {
	i.ListOrders = container.NewVBox()
	i.WindowPage.SetContent(i.ListOrders)
}

func (i *InformationPage) CreateOrderItem(o Order) *fyne.Container {
	var item *fyne.Container

	numberLabel := widget.NewLabel(strconv.FormatUint(o.Id, 10))
	statusLabel := widget.NewLabel("worked")

	item = container.New(
		layout.NewHBoxLayout(),
		numberLabel,
		statusLabel,
	)

	return item
}

func (i *InformationPage) AddOrder(o Order) {
	i.ListOrders.Add(i.CreateOrderItem(o))
}

func (i *InformationPage) DeleteOrder(id uint64) {
	for _, order := range i.ListOrders.Objects {
		if cont, ok := order.(*fyne.Container); ok {
			if numberLabel, ok := cont.Objects[0].(*widget.Label); ok {
				if number, err := strconv.ParseUint(numberLabel.Text, 10, 64); err == nil && number == id {
					i.ListOrders.Remove(order)
					for i, v := range OrderListId {
						if v == id {
							OrderListId = append(OrderListId[:i], OrderListId[i+1:]...)
							break
						}
					}
					break
				}
			}
		}
	}
}

func (i *InformationPage) SetStatusCompleteOrder(id uint64) {
	for _, order := range i.ListOrders.Objects {
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
