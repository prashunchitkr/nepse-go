// Package transport
package transport

import (
	"errors"
	"net/http"

	"github.com/prashunchitkr/nepse-go/internal/auth"
)

type Transport interface {
	RoundTrip(req *http.Request) (*http.Response, error)
}

// Struct for public API calls such as /api/authenticate/prove
type nepsePublicTransport struct {
	base http.RoundTripper
}

func NewNepsePublicTransport(roundTripper http.RoundTripper) Transport {
	return &nepsePublicTransport{
		base: roundTripper,
	}
}

func (t *nepsePublicTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Referer", "https://nepalstock.com.np")

	resp, err := t.base.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

type nepseAuthTransport struct {
	base http.RoundTripper
	auth *auth.Manager
}

func NewNepseAuthTransport(roundTripper http.RoundTripper, authManager *auth.Manager) Transport {
	return &nepseAuthTransport{
		base: roundTripper,
		auth: authManager,
	}
}

func (t *nepseAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.auth.InjectHeaders(req)

	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Referer", "https://nepalstock.com.np")

	resp, err := t.base.RoundTrip(req)

	if resp.StatusCode == http.StatusUnauthorized {
		// trigger re-auth
		return nil, errors.New("Unauthorized")
	}

	return resp, err
}
