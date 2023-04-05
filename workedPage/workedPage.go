package workedPage

import (
	. "TwistAndWrapP/pkg"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func WorkedPage(Window fyne.Window) {
	// test data
	Orders = []Order{{Id: 0, Status: "need work", Products: ProductList}, {Id: 6, Status: "need work", Products: ProductList}, {Id: 89, Status: "need work", Products: ProductList}}
	//

	var listItems []fyne.CanvasObject

	for _, o := range Orders {
		listItems = append(listItems, CreateOrderItem(o))
	}

	scrollContainer := container.NewVScroll(container.NewVBox(listItems...))

	Window.SetContent(scrollContainer)
}

func createDropdownCheckList(items []string) fyne.CanvasObject {
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

func CreateOrderItem(o Order) *fyne.Container {
	numberLabel := widget.NewLabel(strconv.FormatUint(o.Id, 10))

	var productNameList []string

	for _, product := range o.Products {
		productNameList = append(productNameList, product.Name)
	}

	dropdownCheckList := createDropdownCheckList(productNameList)

	statuses := []string{"need work", "worked", "end"}

	selectList := widget.NewSelect(statuses, nil)
	selectList.Selected = statuses[0]

	return container.New(
		layout.NewHBoxLayout(),
		numberLabel,
		dropdownCheckList,
		selectList,
	)
}
