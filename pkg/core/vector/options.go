package vector

// Option provides a function definition to set options
type Option func(*BotOptions)

type BotOptions struct {
	Target string
	Token  string
}

// WithTarget sets the ip of the vector robot.
func WithTarget(s string) Option {
	return func(o *BotOptions) {
		o.Target = s
	}
}

// WithToken set the token for the vector robot.
func WithToken(s string) Option {
	return func(o *BotOptions) {
		o.Token = s
	}
}
