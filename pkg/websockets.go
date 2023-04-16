package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var Conn *websocket.Conn

type Message struct {
	Type string
	Msg  string
}

func receiver(callBackOnCreateOrder func(message Message) error) {
	for {
		_, p, err := Conn.ReadMessage()
		if err != nil {
			err = Conn.Close()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("err read message")
			return
		}

		var m Message

		err = json.Unmarshal(p, &m)
		if err != nil {
			err = Conn.WriteJSON(Message{Type: "", Msg: "Error"})
			if err != nil {
				fmt.Println("write error:", err)
			}
			fmt.Println(err)
			return
		}

		switch m.Type {
		case "createOrder":
			if err := callBackOnCreateOrder(m); err != nil {
				err = Conn.WriteJSON(Message{Type: "createOrder", Msg: "Error"})
				if err != nil {
					fmt.Println("write error:", err)
				}
			}

			err = Conn.WriteJSON(Message{Type: "createOrder", Msg: "Created"})
			if err != nil {
				fmt.Println("write error:", err)
			}
		}
	}
}

func ConnectToWebsocketServer(cookie *http.Cookie, callBackOnCreateOrder func(message Message) error) {
	req, err := http.NewRequest("GET", "ws://localhost:8080/wsChat?roomId=bar", nil)

	if err != nil {
		panic(err)
	}
	req.AddCookie(cookie)

	conn, _, err := websocket.DefaultDialer.Dial(req.URL.String(), req.Header)
	if err != nil {
		fmt.Println("websocket dial error: ", err)
	}

	Conn = conn
	go receiver(callBackOnCreateOrder)
}
