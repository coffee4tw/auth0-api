package api

import (
	"net/http"
	"net/url"

	"github.com/rcrowley/go-tigertonic"
)

func Routes(c *Context) *tigertonic.TrieServeMux {
	mux := tigertonic.NewTrieServeMux()

	// Health
	mux.Handle("GET", "/health", wrap(c.GetHealth, "GET-health"))

	// Privileged endpoint
	mux.Handle("GET", "/user", authWrap(c, c.GetUserInfo, "GET-user"))

	return mux
}

// GetHeath check the health of the server.
func (c *Context) GetHealth(u *url.URL, h http.Header) (int, http.Header, string, error) {
	return http.StatusOK, nil, "Alive", nil
}
