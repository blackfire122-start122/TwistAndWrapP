package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var Conn *websocket.Conn

type Message struct {
	Type string
	Data json.RawMessage `json:"Data"`
}

type CreatedOrderMessage struct {
	ProductsCreated []uint64 `json:"ProductsCreated"`
	Id              uint64   `json:"Id"`
}

type OrderGiveMessage struct {
	Id uint64 `json:"Id"`
}

type CreateOrderMessage struct {
	FoodIdCount map[uint64]uint8
	Time        string
}

func receiver(conn *websocket.Conn, cookie *http.Cookie, callBackOnCreateOrder func(message Message) (uint64, []uint64, error)) {
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {

			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println(err)
			}

			fmt.Println("err read message ", err)

			time.Sleep(2 * time.Second)
			if err := ConnectToWebsocketServer(cookie, callBackOnCreateOrder); err != nil {
				fmt.Println("error reconnect ", err)
			}
			break
		}

		var m Message
		err = json.Unmarshal(p, &m)

		if err != nil {
			err = conn.WriteJSON(Message{Type: "Error"})
			if err != nil {
				fmt.Println("write error:", err)
			}
			fmt.Println("error unmarshal", err)
			break
		}

		fmt.Println(m)

		switch m.Type {
		case "createOrder":
			if id, ProductsCreated, err := callBackOnCreateOrder(m); err != nil {
				err = conn.WriteJSON(Message{Type: "Error"})
				if err != nil {
					fmt.Println("write error:", err)
				}
			} else {
				data, err := json.Marshal(CreatedOrderMessage{Id: id, ProductsCreated: ProductsCreated})

				if err != nil {
					fmt.Println("marshal error: ", err)
				}

				err = conn.WriteJSON(Message{Type: "OrderCreated", Data: data})
				if err != nil {
					fmt.Println("write error:", err)
				}
			}
		}
	}
}

func ConnectToWebsocketServer(cookie *http.Cookie, callBackOnCreateOrder func(message Message) (uint64, []uint64, error)) error {
	req, err := http.NewRequest("GET", "ws://"+Host+"/websocket/wsChat?roomId="+IdBar, nil)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.AddCookie(cookie)

	conn, _, err := websocket.DefaultDialer.Dial(req.URL.String(), req.Header)
	if err != nil {
		fmt.Println("websocket dial error: ", err)
		return err
	}

	Conn = conn

	fmt.Println("Connected")
	go receiver(conn, cookie, callBackOnCreateOrder)
	go pingPong(conn, cookie, callBackOnCreateOrder)
	return nil
}

func pingPong(conn *websocket.Conn, cookie *http.Cookie, callBackOnCreateOrder func(message Message) (uint64, []uint64, error)) {
	ticker := time.NewTicker(30 * time.Second)

	defer func() {
		ticker.Stop()
		err := conn.Close()
		if err != nil {
			fmt.Println("close error", err)
		}
	}()

	for range ticker.C {
		if err := conn.WriteControl(websocket.PingMessage, []byte("ping"), time.Now().Add(5*time.Second)); err != nil {
			fmt.Println("Error sending ping message: " + err.Error())
			time.Sleep(2 * time.Second)
			if err := ConnectToWebsocketServer(cookie, callBackOnCreateOrder); err != nil {
				fmt.Println("error reconnect ", err)
			}
			return
		}
	}
}
