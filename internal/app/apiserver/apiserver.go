package apiserver

type APIServer struct{}

func NewAPIServer() *APIServer {
	return &APIServer{}
}

func (server *APIServer) StartAPIServer() error {
	return nil
}
