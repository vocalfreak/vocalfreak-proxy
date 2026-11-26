package server

import (
	"net/url"
	"time"
)

type Config struct {
	Bind            string
	HttpPort        int
	Upstream        *url.URL
	ResponseTimeout time.Duration
	Debug           bool
}
