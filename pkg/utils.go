package pkg

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"net/http"
)

var ProductList []Product

var CheckListProducts []*widget.Check

var OrderListId []uint64

var App fyne.App

var Client *http.Client

//var Conn *websocket.Conn

type Product struct {
	Id    string `json:"Id"`
	Image string `json:"Image"`
	Name  string `json:"Name"`
}

type Order struct {
	Id       uint64
	Products []Product
}

func GetProducts() []Product {
	var products []Product
	for i, obj := range CheckListProducts {
		if obj.Checked {
			products = append(products, ProductList[i])
		}
	}
	return products
}

//type Page interface {
//	Window() fyne.Window
//	SetWindowContent()
//}

//type Message struct {
//	Type string
//	Msg  string
//}

//func receiver() {
//	for {
//		_, p, err := Conn.ReadMessage()
//		fmt.Println(p)
//		if err != nil {
//			err = Conn.Close()
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//			fmt.Println("err read message")
//			return
//		}
//
//		var m Message
//
//		err = json.Unmarshal(p, &m)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//
//		switch m.Type {
//		case "msg":
//			fmt.Println(m)
//		case "createOrder":
//			fmt.Println("create order", m)
//		}
//	}
//}

//func ConnectToWebsocketServer() {
//	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/wsChat?roomId=bar?sessionId=", nil)
//	if err != nil {
//		fmt.Println("websocket dial error: ", err)
//	}
//
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	err = conn.WriteJSON(Message{Type: "msg", Msg: "Hello, world!"})
//	if err != nil {
//		fmt.Println("write error:", err)
//	}
//	Conn = conn
//	go receiver()
//}
