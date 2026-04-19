// Package transport
package transport

import (
	"errors"
	"io"
	"net/http"

	"github.com/prashunchitkr/nepse-go/internal/auth"
)

const (
	userAgent = "Mozilla/5.0 (X11; Linux x86_64; rv:149.0) Gecko/20100101 Firefox/149.0"
	referer   = "https://nepalstock.com.np/"
	host      = "www.nepalstock.com.np"
	accept    = "application/json"
)

type Transport interface {
	RoundTrip(req *http.Request) (*http.Response, error)
}

type nepsePublicTransport struct {
	base http.RoundTripper
}

func NewNepsePublicTransport(roundTripper http.RoundTripper) Transport {
	return &nepsePublicTransport{
		base: roundTripper,
	}
}

func (t *nepsePublicTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Host", host)
	req.Header.Set("Referer", referer)
	req.Header.Set("Accept", accept)

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
	do := func(r *http.Request) (*http.Response, error) {
		if err := t.auth.InjectHeaders(r); err != nil {
			return nil, err
		}

		r.Header.Set("User-Agent", userAgent)
		r.Header.Set("Referer", referer)
		r.Header.Set("Host", host)
		r.Header.Set("Accept", accept)

		return t.base.RoundTrip(r)
	}

	resp, err := do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusUnauthorized {
		return resp, nil
	}

	_, _ = io.Copy(io.Discard, resp.Body)
	resp.Body.Close()

	if err := t.auth.Refresh(req.Context()); err != nil {
		return nil, err
	}

	req2 := req.Clone(req.Context())
	if req2 == nil {
		return nil, errors.New("transport: request clone failed")
	}

	resp2, err := do(req2)
	if err != nil {
		return nil, err
	}
	if resp2.StatusCode == http.StatusUnauthorized {
		resp2.Body.Close()
		return nil, errors.New("unauthorized after prove refresh")
	}
	return resp2, nil
}
