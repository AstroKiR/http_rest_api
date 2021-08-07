package apiserver

type APIServer struct {
	config *Config
}

func NewAPIServer(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

func (server *APIServer) StartAPIServer() error {
	return nil
}
