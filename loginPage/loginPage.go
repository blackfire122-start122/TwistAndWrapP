package LoginPage

import (
	. "TwistAndWrapP/pkg"
	. "TwistAndWrapP/settingPage"
	"bytes"
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"net/http"
	"net/http/cookiejar"
)

func FirstCallLoginPage(MainWindow fyne.Window) {
	LoginPage = FirstCallLoginPage

	Client = &http.Client{}

	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{
		Jar: jar,
	}

	idEntry := widget.NewEntry()
	passEntry := widget.NewPasswordEntry()
	errorLabel := widget.NewLabel("Error data")

	submitButton := widget.NewButton("Submit", func() {
		type LoginResponse struct {
			Login string
		}

		type LoginRequest struct {
			IdBar    string
			Password string
		}

		url := "http://localhost:8080/loginBar"
		loginData := LoginRequest{
			IdBar:    idEntry.Text,
			Password: passEntry.Text,
		}

		jsonData, err := json.Marshal(loginData)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error creating HTTP request:", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending HTTP request:", err)
			return
		}

		defer resp.Body.Close()

		var data LoginResponse

		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			fmt.Printf("Error while parsing response body: %s\n", err.Error())
		}

		if data.Login == "OK" {
			SettingPage(MainWindow)
		} else {
			errorLabel.Show()
		}
	})

	errorLabel.Hide()

	loginPage := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Please enter your ID and password:"),
		idEntry,
		passEntry,
		submitButton,
		errorLabel,
	)

	GetAndSetAllData()

	MainWindow.SetTitle("Login Page")
	MainWindow.SetContent(loginPage)
}

func GetAndSetAllData() {
	GetJson("http://localhost:8080/getAllProducts", &ProductList)
}

func GetJson(url string, target any) {
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
