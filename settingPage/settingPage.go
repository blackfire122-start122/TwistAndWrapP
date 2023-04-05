package settingPage

import (
	. "TwistAndWrapP/checkoutPage"
	. "TwistAndWrapP/entrys"
	. "TwistAndWrapP/informationPage"
	. "TwistAndWrapP/pkg"
	. "TwistAndWrapP/workedPage"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func SettingPage(MainWindow fyne.Window) {
	numberInfoEntry := NewNumericalEntry()
	numberInfoEntry.SetPlaceHolder("Number of information pages")

	numberCheckoutEntry := NewNumericalEntry()
	numberCheckoutEntry.SetPlaceHolder("Number of checkout pages")

	numberWorkEntry := NewNumericalEntry()
	numberWorkEntry.SetPlaceHolder("Number of worked pages")

	for _, product := range ProductList {
		CheckListProducts = append(CheckListProducts, widget.NewCheck(product.Name, func(bool) {}))
	}

	var items []fyne.CanvasObject
	for _, c := range CheckListProducts {
		items = append(items, c)
	}

	ProductCheckList := container.NewVBox(items...)

	newContent := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Setting Profile"),

		numberInfoEntry,
		numberCheckoutEntry,
		numberWorkEntry,
		ProductCheckList,

		container.New(layout.NewHBoxLayout(),
			widget.NewButton("Logout", func() {
				// send request logout
				LoginPage(MainWindow)
			}),
			widget.NewButton("Apply", func() {
				CreateWindows(numberInfoEntry, "Information Page", &InformationPageList, InformationPage)
				CreateWindows(numberCheckoutEntry, "Checkout Page", &CheckoutPageList, CheckoutPage)
				CreateWindows(numberWorkEntry, "Worked Page", &WorkedPageList, WorkedPage)
			}),
		),
	)
	MainWindow.SetTitle("Setting Page")
	MainWindow.SetContent(newContent)
}

func CreateWindows(entry *NumericalEntry, title string, pageList *[]fyne.Window, funcPage func(Window fyne.Window)) {
	var entryValue int64
	entryValue, _ = strconv.ParseInt(entry.Text, 10, 64)

	for _, item := range *pageList {
		item.Close()
	}

	*pageList = []fyne.Window{}

	for i := int64(0); i < entryValue; i++ {
		*pageList = append(*pageList, App.NewWindow(title))

		funcPage((*pageList)[i])
		(*pageList)[i].Show()
	}
}
