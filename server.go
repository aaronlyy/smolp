package smolp

type Server struct {
	Host string
	Port int
}

func NewServer(protocol *Protocol) (*Server, error) {
	return &Server{}, nil
}

func (srv *Server) Listen(host string, port int) error {
	return nil
}

// srv, err := smolp.newServer()
// if err := srv.Listen("localhost", 3000); err != nil {
// 	log.Fatalf("error: %v", err)
// }
