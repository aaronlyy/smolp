package smolp

type Server struct {
	Host string
	Port int
}

func NewServer(protocol *Protocol, handler HandlerFunc) (*Server, error) {
	return &Server{}, nil
}

func (srv *Server) Listen(host string, port int) error {
	return nil
}
