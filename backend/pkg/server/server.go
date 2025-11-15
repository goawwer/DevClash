package server

import (
	"fmt"
	"net/http"
)

type Config struct {
	Host string `env:"HTTP_HOST"`
	Port int    `env:"HTTP_PORT"`
}

type Server struct {
	*http.Server
}

func New(handler http.Handler, c *Config) *Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", c.Host, c.Port),
		Handler: handler,
	}

	return &Server{
		srv,
	}
}
