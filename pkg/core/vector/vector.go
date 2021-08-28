package vector

import (
	"context"
	"fmt"
	"log"

	"github.com/digital-dream-labs/hugh/grpc/client"
	"github.com/digital-dream-labs/vector-go-sdk/pkg/vector"
	"github.com/digital-dream-labs/vector-go-sdk/pkg/vectorpb"
)

type Vector struct {
	client  *vector.Vector
	options BotOptions
}

// Needs further investigation
const (
	sessionID  = "id01"
	clientName = "id02"
)

func New(options *BotOptions) *Vector {

	v := &Vector{
		options: *options,
	}

	err := v.Authenticate()
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func (v *Vector) ConnectionStatus() string {
	status, _ := v.client.Conn.CheckCloudConnection(
		context.Background(),
		&vectorpb.CheckCloudRequest{},
	)
	log.Printf("\n\n\nConnection status: %s", status.Status)
	return status.StatusMessage
}

func (v *Vector) refreshClient() (err error) {

	c, err := vector.New(
		vector.WithTarget(
			fmt.Sprintf("%s:443", v.options.Target),
		),
		vector.WithToken(v.options.Token),
	)

	if err != nil {
		return
	}

	v.client = c
	return
}

func (v *Vector) setBotOptions(opts ...Option) (err error) {
	for _, opt := range opts {
		opt(&v.options)
	}

	err = v.refreshClient()
	if err != nil {
		return
	}
	return
}

func (v *Vector) Authenticate() error {

	log.Println("\n\n\ntest")
	if v.options.Token != "" {
		return v.refreshClient()
	}

	c, err := client.New(
		client.WithTarget(
			fmt.Sprintf("%s:443", v.options.Target),
		),
		client.WithInsecureSkipVerify(),
	)

	if err != nil {
		return err
	}
	if err := c.Connect(); err != nil {
		return err
	}

	vc := vectorpb.NewExternalInterfaceClient(c.Conn())

	login, err := vc.UserAuthentication(context.Background(),
		&vectorpb.UserAuthenticationRequest{
			UserSessionId: []byte(sessionID),
			ClientName:    []byte(clientName),
		},
	)

	if err != nil {
		return err
	}

	log.Println(login.ClientTokenGuid)

	v.setBotOptions(WithToken(string(login.ClientTokenGuid)))

	return nil
}

func (v *Vector) AssumeBehavior() {
	ctx := context.Background()
	start := make(chan bool)
	stop := make(chan bool)

	log.Printf("Bot: %#v", v.options)

	go func() {
		err := v.client.BehaviorControl(ctx, start, stop)
		if err != nil {
			log.Printf("Bot: %#v", v.options)
			log.Fatal(err)
		}
	}()
}
