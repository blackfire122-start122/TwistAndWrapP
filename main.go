package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	MainWindow := myApp.NewWindow("Login Page")

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
	newContent := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Welcome!"),
		widget.NewButton("Logout", func() {
			// send request logout
			LoginPage(MainWindow)
		}),
	)
	MainWindow.SetTitle("Setting Page")
	MainWindow.SetContent(newContent)
}
