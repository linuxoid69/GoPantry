package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Port       int
	ListenAddr string
}

func NewServer(port int, listenAddr string) *Server {
	return &Server{
		Port:       port,
		ListenAddr: listenAddr,
	}
}

func (s *Server) Run() error {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})
	r.POST("/api/movies/add", HandlerAddMovie)

	return r.Run(s.ListenAddr + ":" + strconv.Itoa(s.Port))
}
