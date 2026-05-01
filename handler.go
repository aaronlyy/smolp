package smolp

type HandlerFunc func(payload *Payload) error

func (f HandlerFunc) Handle(payload *Payload) error {
	return f(payload)
}
