package smolp

type HandlerFunc func(payload *Payload)

func (f HandlerFunc) Handle(payload *Payload) {
	f(payload)
}
