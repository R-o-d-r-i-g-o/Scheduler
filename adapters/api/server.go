package api

import (
	"scheduler/common/configs"

	"github.com/gin-gonic/gin"
)

type Server struct {
	host, port   string
	serverEngine *gin.Engine
}

func CreateServer() *Server {
	server := NewServer()

	server.SetPort(configs.Server.PORT)
	server.SetHost(configs.Server.HOST)
	return server
}

func NewServer() *Server {
	return &Server{
		host:         "",
		port:         "",
		serverEngine: gin.Default(),
	}
}

// ---------------------------------< Methods >--------------------------------- \\

func (s Server) GetServerEngine() *gin.Engine {
	return s.serverEngine
}

func (s *Server) SetPort(port string) {
	s.port = port
}

func (s *Server) SetHost(host string) {
	s.host = host
}
