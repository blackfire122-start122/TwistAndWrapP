package App

import (
	. "TwistAndWrapP/loginPage"
	. "TwistAndWrapP/settingPage"
	"fyne.io/fyne/v2/app"
)

func RunApp() {
	App := app.New()
	MainWindow := App.NewWindow("Login Page")

	LoginPage(MainWindow)

	MainWindow.ShowAndRun()
}
