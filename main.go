package main

import (
	. "TwistAndWrapP/entrys"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

var InformationPageList []fyne.Window
var App fyne.App

func main() {
	App = app.New()
	MainWindow := App.NewWindow("Login Page")

	LoginPage(MainWindow)

	MainWindow.ShowAndRun()
}

func LoginPage(MainWindow fyne.Window) {
	idEntry := widget.NewEntry()
	passEntry := widget.NewPasswordEntry()

	submitButton := widget.NewButton("Submit", func() {
		// send request login bar
		settingPage(MainWindow)
	})

	loginPage := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Please enter your ID and password:"),
		idEntry,
		passEntry,
		submitButton,
	)

	MainWindow.SetTitle("Login Page")
	MainWindow.SetContent(loginPage)
}

func settingPage(MainWindow fyne.Window) {
	numberInfoEntry := NewNumericalEntry()
	numberInfoEntry.SetPlaceHolder("Number of information pages")

	newContent := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Setting Profile"),
		numberInfoEntry,

		container.New(layout.NewHBoxLayout(),
			widget.NewButton("Logout", func() {
				// send request logout
				LoginPage(MainWindow)
			}),
			widget.NewButton("Apply", func() {
				var infoEntryValue int64
				infoEntryValue, _ = strconv.ParseInt(numberInfoEntry.Text, 10, 64)

				for i := int64(0); i < infoEntryValue; i++ {
					InformationPageList = append(InformationPageList, App.NewWindow("Information Page"))
					InformationPage(InformationPageList[i])
					InformationPageList[i].Show()
				}
			}),
		),
	)
	MainWindow.SetTitle("Setting Page")
	MainWindow.SetContent(newContent)
}

func InformationPage(Window fyne.Window) {
	numberList := []string{"1", "2", "3", "4", "5"}
	statusList := []string{"Connected", "Disconnected", "Connected", "Disconnected", "Connected"}

	var listItems []fyne.CanvasObject

	for i := 0; i < len(numberList); i++ {
		numberLabel := widget.NewLabel(numberList[i])
		statusLabel := widget.NewLabel(statusList[i])

		listItems = append(listItems, container.New(
			layout.NewHBoxLayout(),
			numberLabel,
			statusLabel,
		))
	}

	scrollContainer := container.NewVScroll(container.NewVBox(listItems...))

	Window.SetContent(scrollContainer)
}
