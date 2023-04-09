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
				for _, item := range ListWorkedPage {
					item.Window().Close()
				}

				for _, item := range ListInformationPage {
					item.Window().Close()
				}

				for _, item := range ListCheckoutPage {
					item.Window().Close()
				}

				ListWorkedPage = []*WorkedPage{}
				ListInformationPage = []*InformationPage{}
				ListCheckoutPage = []*CheckoutPage{}

				CreateWindows(numberInfoEntry, "Information Page", NewInformationPage)
				CreateWindows(numberCheckoutEntry, "Checkout Page", NewCheckoutPage)
				CreateWindows(numberWorkEntry, "Worked Page", NewWorkedPage)
			}),
		),
	)
	MainWindow.SetTitle("Setting Page")
	MainWindow.SetContent(newContent)
}

func CreateWindows(entry *NumericalEntry, title string, funcCreatePage func(w fyne.Window) Page) {
	var entryValue int64
	entryValue, _ = strconv.ParseInt(entry.Text, 10, 64)

	for i := int64(0); i < entryValue; i++ {
		funcCreatePage(App.NewWindow(title))
	}
}
