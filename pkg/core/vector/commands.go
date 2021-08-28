package vector

import (
	"context"

	"github.com/digital-dream-labs/vector-go-sdk/pkg/vectorpb"
)

func (v *Vector) Talk(ctx context.Context, message string) {
	_, _ = v.client.Conn.SayText(
		ctx,
		&vectorpb.SayTextRequest{
			Text:           message,
			UseVectorVoice: true,
			DurationScalar: 1.0,
		},
	)

}
