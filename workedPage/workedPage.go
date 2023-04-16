package workedPage

import (
	. "TwistAndWrapP/pkg"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

var ListWorkedPage []*WorkedPage

func NewWorkedPage(w fyne.Window, callbackComplete func(Id uint64)) {
	page := &WorkedPage{WindowPage: w, callbackComplete: callbackComplete}
	page.SetWindowContent()
	ListWorkedPage = append(ListWorkedPage, page)
	w.Show()
}

type WorkedPage struct {
	ListWork         *fyne.Container
	WindowPage       fyne.Window
	callbackComplete func(Id uint64)
}

func (w *WorkedPage) Window() fyne.Window {
	return w.WindowPage
}

func (w *WorkedPage) SetWindowContent() {
	w.ListWork = container.NewVBox()
	w.WindowPage.SetContent(w.ListWork)
}

func (w *WorkedPage) createDropdownCheckList(items []string) fyne.CanvasObject {
	checkboxes := make([]*widget.Check, len(items))
	for i, item := range items {
		checkboxes[i] = widget.NewCheck(item, nil)
	}

	var itemsCheckboxes []fyne.CanvasObject
	for _, c := range checkboxes {
		itemsCheckboxes = append(itemsCheckboxes, c)
	}

	dropdownContent := container.NewVBox(itemsCheckboxes...)
	dropdownContent.Hide()

	dropdownButton := widget.NewButton("Select Items", func() {
		if dropdownContent.Hidden {
			dropdownContent.Show()
		} else {
			dropdownContent.Hide()
		}
	})

	return container.NewVBox(
		dropdownButton,
		dropdownContent,
		layout.NewSpacer(),
	)
}

func (w *WorkedPage) CreateOrderItem(o Order) *fyne.Container {
	numberLabel := widget.NewLabel(strconv.FormatUint(o.Id, 10))

	var productNameList []string
	var item *fyne.Container

	for _, product := range o.Products {
		productNameList = append(productNameList, product.Name)
	}

	dropdownCheckList := w.createDropdownCheckList(productNameList)

	statuses := []string{"need work", "worked", "end"}

	selectList := widget.NewSelect(statuses, func(s string) {
		if s == "end" {
			w.callbackComplete(o.Id)
			w.ListWork.Remove(item)
		}
		for _, page := range ListWorkedPage {
			for _, work := range page.ListWork.Objects {
				if cont, ok := work.(*fyne.Container); ok {
					if numberLabel, ok := cont.Objects[0].(*widget.Label); ok {
						if number, err := strconv.ParseUint(numberLabel.Text, 10, 64); err == nil && number == o.Id {
							if selectList, ok := cont.Objects[2].(*widget.Select); ok {
								selectList.Selected = s
								if s == "end" {
									page.ListWork.Remove(work)
								}
								page.Window().Content().Refresh()
								break
							}
						}
					}
				}
			}
		}
	})

	selectList.Selected = statuses[0]

	item = container.New(
		layout.NewHBoxLayout(),
		numberLabel,
		dropdownCheckList,
		selectList,
	)

	return item
}

func (w *WorkedPage) AddOrder(o Order) {
	w.ListWork.Add(w.CreateOrderItem(o))
}
