package server

import (
	"net"
	"net/http"
)

type Server struct {
	config       *Config
	httpListener net.Listener
	httpServer   *http.Server
}
