package settingPage

import (
	. "TwistAndWrapP/checkoutPage"
	. "TwistAndWrapP/entrys"
	. "TwistAndWrapP/informationPage"
	. "TwistAndWrapP/pkg"
	. "TwistAndWrapP/workedPage"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func SettingPage(MainWindow fyne.Window, LoginPage func(MainWindow fyne.Window)) {
	numberInfoEntry := NewNumericalEntry()
	numberInfoEntry.SetPlaceHolder("Number of information pages")

	numberCheckoutEntry := NewNumericalEntry()
	numberCheckoutEntry.SetPlaceHolder("Number of checkout pages")

	numberWorkEntry := NewNumericalEntry()
	numberWorkEntry.SetPlaceHolder("Number of worked pages")

	numberInfoEntry.Text = "1"     // ToDo: Default value
	numberCheckoutEntry.Text = "1" // ToDo: Default value
	numberWorkEntry.Text = "1"     // ToDo: Default value

	for _, product := range ProductList {
		CheckListProducts = append(CheckListProducts, widget.NewCheck(product.Name, func(bool) {}))
	}

	var items []fyne.CanvasObject
	for _, c := range CheckListProducts {
		c.Checked = true // ToDo: Default value
		items = append(items, c)
	}
	productLabel := widget.NewLabel("Select products what yours have: ")
	ProductCheckList := container.NewHScroll(container.NewHBox(items...))

	newContent := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Setting Profile"),

		numberInfoEntry,
		numberCheckoutEntry,
		numberWorkEntry,
		productLabel,
		ProductCheckList,

		container.New(layout.NewHBoxLayout(),
			widget.NewButton("Logout", func() {
				err := Conn.Close()
				if err != nil {
					fmt.Println(err)
				}
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

				CreateWindows(numberInfoEntry, "Information Page", func(w fyne.Window) {
					NewInformationPage(w)
				})
				CreateWindows(numberCheckoutEntry, "Checkout Page", func(w fyne.Window) {
					NewCheckoutPage(w, func(order Order) {
						for _, page := range ListWorkedPage {
							page.AddOrder(order)
							page.Window().Content().Refresh()
						}

						for _, page := range ListInformationPage {
							page.AddOrder(order)
							page.Window().Content().Refresh()
						}
					}, func(Id uint64) {
						for _, page := range ListInformationPage {
							page.DeleteOrder(Id)
							page.Window().Content().Refresh()
						}
					})
				})
				CreateWindows(numberWorkEntry, "Worked Page", func(w fyne.Window) {
					NewWorkedPage(w, func(Id uint64) {
						for _, page := range ListInformationPage {
							page.SetStatusCompleteOrder(Id)
							page.Window().Content().Refresh()
						}
						for _, page := range ListCheckoutPage {
							page.SetStatusCompleteOrder(Id)
							page.Window().Content().Refresh()
						}
					})
				})
			}),
		),
	)
	MainWindow.SetTitle("Setting Page")
	MainWindow.SetContent(newContent)
}

func CreateWindows(entry *NumericalEntry, title string, funcCreatePage func(w fyne.Window)) {
	var entryValue int64
	entryValue, _ = strconv.ParseInt(entry.Text, 10, 64)

	for i := int64(0); i < entryValue; i++ {
		funcCreatePage(fyne.CurrentApp().NewWindow(title))
	}
}
