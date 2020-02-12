package main

func main() {
	srv := CreateServer()
	srv.RunServer()
}

type Server struct {
	myFavField string
}

func CreateServer() *Server {
	// this is just a demo
	return nil
}

func (s *Server) RunServer() {
	// this is just a demo
}
