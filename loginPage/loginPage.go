package LoginPage

import (
	. "TwistAndWrapP/pkg"
	. "TwistAndWrapP/settingPage"
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"net/http"
	//"net/http/cookiejar"
)

func LoginPageFirstCall(MainWindow fyne.Window) {
	LoginPage = LoginPageFirstCall

	Client = &http.Client{}

	//jar, err := cookiejar.New(nil)
	//
	//if err != nil {
	//	// error handling
	//}

	idEntry := widget.NewEntry()
	passEntry := widget.NewPasswordEntry()

	submitButton := widget.NewButton("Submit", func() {
		// send request login bar
		SettingPage(MainWindow)
	})

	loginPage := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Please enter your ID and password:"),
		idEntry,
		passEntry,
		submitButton,
	)

	GetAndSetAllData()

	MainWindow.SetTitle("Login Page")
	MainWindow.SetContent(loginPage)
}

func GetAndSetAllData() {
	getJson("http://localhost:8080/getAllProducts", &ProductList)
}

func getJson(url string, target any) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := Client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		fmt.Printf("Error while parsing response body: %s\n", err.Error())
		return
	}
}
