package LoginPage

import (
	. "TwistAndWrapP/checkoutPage"
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
	"strconv"
	"time"
)

type MsgCreateOrder struct {
	FoodIdCount map[uint64]uint8
	Time        string
}

func LoginPage(MainWindow fyne.Window) {
	Client = &http.Client{}

	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	Client = &http.Client{
		Jar: jar,
	}

	idEntry := widget.NewEntry()
	passEntry := widget.NewPasswordEntry()
	errorLabel := widget.NewLabel("")

	submitButton := widget.NewButton("Submit", func() {
		type LoginResponse struct {
			Login string
		}

		type LoginRequest struct {
			IdBar    string
			Password string
		}

		url := "http://localhost/api/bar/loginBar"
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

		resp, err := Client.Do(req)
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
			IdBar = idEntry.Text
			for _, cookie := range resp.Cookies() {
				err := ConnectToWebsocketServer(cookie, func(message Message) (uint64, error) {
					var productsOrder []ProductOrder
					var msgProductsId MsgCreateOrder

					if err := json.Unmarshal([]byte(message.Msg), &msgProductsId); err != nil {
						fmt.Println(err)
						return 0, err
					}

					for _, product := range ProductList {
						pId, err := strconv.ParseUint(product.Id, 10, 64)
						if err != nil {
							fmt.Println(err)
							return 0, err
						}

						for id, count := range msgProductsId.FoodIdCount {
							if pId == id {
								productsOrder = append(productsOrder, ProductOrder{Product: product, Count: count})
							}
						}
					}
					t, err := time.Parse("15:04", msgProductsId.Time)

					if err != nil {
						fmt.Println("Error parsing time string:", err)
						return 0, err
					}

					OrderId := GenerateIdOrderList()

					go func() {
						now := time.Now()
						targetTime := time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute()-3, 0, 0, now.Location())
						duration := targetTime.Sub(now)

						timer := time.NewTimer(duration)

						<-timer.C

						for _, page := range ListCheckoutPage {
							page.AddOrder(productsOrder, OrderId)
							page.Window().Content().Refresh()
						}
					}()

					return OrderId, err
				})
				if err != nil {
					errorLabel.SetText("Error connect to websocket")
					errorLabel.Show()
				}
				break
			}
			GetAndSetAllData()
			SettingPage(MainWindow, LoginPage)
		} else {
			errorLabel.SetText("Error data")
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

	MainWindow.SetTitle("Login Page")
	MainWindow.SetContent(loginPage)
}

func GetAndSetAllData() {
	GetJson("http://localhost/api/bar/getAllProducts", &ProductList)
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

//66775588
//123456
