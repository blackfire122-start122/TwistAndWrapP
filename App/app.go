package App

import (
	. "TwistAndWrapP/loginPage"
	"fyne.io/fyne/v2/app"
)

func RunApp() {
	App := app.New()
	MainWindow := App.NewWindow("Login Page")

	LoginPage(MainWindow)

	MainWindow.ShowAndRun()
}
