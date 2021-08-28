package websocket

import (
	"encoding/json"
	"net/http"

	"github.com/OhMyVector/spotter-vector/pkg/api/vector"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func reader(c echo.Context, conn *websocket.Conn) error {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
			return err
		}
		c.Logger().Info(string(p))

		var msg vector.Message
		err = json.Unmarshal(p, &msg)
		if err != nil {
			return err
		}

		err = vector.HandleMessage(msg)
		if err != nil {
			return err
		}

		if err := conn.WriteMessage(messageType, p); err != nil {
			c.Logger().Error(err)
			return err
		}
	}
}

func Connection(c echo.Context) (err error) {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error(err)
		return
	}
	defer ws.Close()

	c.Logger().Info("Client Connected")

	err = ws.WriteMessage(1, []byte("Hi Spotter!"))
	if err != nil {
		c.Logger().Error(err)
		return
	}

	err = reader(c, ws)
	if err != nil {
		c.Logger().Error(err)
		return
	}
	return
}
