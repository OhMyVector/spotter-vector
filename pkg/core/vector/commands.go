package vector

import (
	"context"

	"github.com/digital-dream-labs/vector-go-sdk/pkg/vectorpb"
	"github.com/ohmyvector/spotter-vector/pkg/core/constants"
)

func turnInPlace(ctx context.Context, v *Vector, angleRad float32) {
	_, _ = v.client.Conn.TurnInPlace(ctx, &vectorpb.TurnInPlaceRequest{
		AngleRad: angleRad,
		IdTag:    constants.VECTOR_ID_TAG,
	})
}

func (v *Vector) TurnLeft(ctx context.Context) {
	turnInPlace(ctx, v, constants.TURN_LEFT_ANGLE_RAD)
}

func (v *Vector) TurnRight(ctx context.Context) {
	turnInPlace(ctx, v, constants.TURN_RIGHT_ANGLE_RAD)
}

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

func (v *Vector) DriveStraight(ctx context.Context) {
	_, _ = v.client.Conn.DriveStraight(ctx, &vectorpb.DriveStraightRequest{
		SpeedMmps: constants.SPEED_MMPS,
		IdTag:     constants.VECTOR_ID_TAG,
	})
}
