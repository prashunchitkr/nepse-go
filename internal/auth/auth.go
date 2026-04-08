// Package auth
package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/prashunchitkr/nepse-go/internal/session"
)

type Manager struct {
	session *session.Manager
}

func (a *Manager) Authenticate(ctx context.Context) error {
	return fmt.Errorf("not implemented")
}

func (a *Manager) Refresh(ctx context.Context) error {
	return fmt.Errorf("not implemented")
}

func (a *Manager) InjectHeaders(req *http.Request) error {
	return fmt.Errorf("not implemented")
}
