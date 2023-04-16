package App

import (
	. "TwistAndWrapP/loginPage"
	. "TwistAndWrapP/pkg"
	"fyne.io/fyne/v2/app"
)

func RunApp() {
	App = app.New()
	MainWindow := App.NewWindow("Login Page")

	LoginPage(MainWindow)

	MainWindow.ShowAndRun()
}
