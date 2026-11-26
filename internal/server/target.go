package server

import (
	"net/http"
	"net/url"
)

type Target struct {
	targetURL *url.URL
	client    *http.Client
}
