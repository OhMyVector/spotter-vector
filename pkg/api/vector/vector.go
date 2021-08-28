package vector

import (
	"context"
	"fmt"
	"net/http"

	"github.com/OhMyVector/spotter-vector/pkg/core/vector"
	"github.com/labstack/echo/v4"
)

var (
	bot *vector.Vector
)

type Message struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

func Connect(c echo.Context) error {

	var body struct {
		Target string `json:"target"`
		Token  string `json:"token,omitempty"`
	}

	if err := c.Bind(&body); err != nil {
		return err
	}

	c.Logger().Info("Setting vector config")
	bot = vector.New(&vector.BotOptions{Target: body.Target, Token: body.Token})
	bot.AssumeBehavior()
	c.Logger().Info("Connected to Vector")

	return c.JSON(http.StatusOK, body)
}

func GetConnectionStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, bot.ConnectionStatus())
}

func HandleMessage(msg Message) error {
	switch msg.Type {
	case "talk":
		bot.Talk(context.Background(), msg.Data)
	default:
		return fmt.Errorf("invalid message type")
	}
	return nil
}
