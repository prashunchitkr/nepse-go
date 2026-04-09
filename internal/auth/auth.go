// Package auth
package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/prashunchitkr/nepse-go/internal/models"
	"github.com/prashunchitkr/nepse-go/internal/session"
)

type Manager struct {
	session      *session.Manager
	publicClient *http.Client
	baseURL      string
	mu           sync.Mutex
}

func NewManager(session *session.Manager, publicClient *http.Client, baseURL string) *Manager {
	return &Manager{
		session:      session,
		publicClient: publicClient,
		baseURL:      baseURL,
	}
}

// Authenticate forces a new prove exchange.
func (a *Manager) Authenticate(ctx context.Context) error {
	return a.Prove(ctx)
}

// Prove forces a new prove exchange.
func (a *Manager) Prove(ctx context.Context) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.proveUnlocked(ctx)
}

// Refresh re-runs prove (e.g. after 401).
func (a *Manager) Refresh(ctx context.Context) error {
	return a.Prove(ctx)
}

func (a *Manager) ensureProved(ctx context.Context) error {
	if _, ok := a.session.SalterAuthorization(); ok {
		return nil
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.session.SalterAuthorization(); ok {
		return nil
	}
	return a.proveUnlocked(ctx)
}

func (a *Manager) proveUnlocked(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.baseURL+"/api/authenticate/prove", http.NoBody)
	if err != nil {
		return err
	}

	resp, err := a.publicClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("prove: unexpected status %s", resp.Status)
	}

	var prove models.Prove

	if err := json.NewDecoder(resp.Body).Decode(&prove); err != nil {
		return fmt.Errorf("prove: decode body: %w", err)
	}

	salter := deriveSalterToken(&prove)
	a.session.SetFromProve(&prove, salter)
	return nil
}

func (a *Manager) InjectHeaders(req *http.Request) error {
	if err := a.ensureProved(req.Context()); err != nil {
		return err
	}
	token, ok := a.session.SalterAuthorization()
	if !ok {
		return fmt.Errorf("prove: session missing salter after successful prove")
	}
	req.Header.Set("Authorization", "Salter "+token)
	return nil
}

func deriveSalterToken(p *models.Prove) string {
	idx1 := cdx(int32(p.Salt1), int32(p.Salt2), int32(p.Salt3), int32(p.Salt4), int32(p.Salt4))
	idx2 := rdx(int32(p.Salt1), int32(p.Salt2), int32(p.Salt3), int32(p.Salt4), int32(p.Salt4))
	idx3 := bdx(int32(p.Salt1), int32(p.Salt2), int32(p.Salt3), int32(p.Salt4), int32(p.Salt4))
	idx4 := ndx(int32(p.Salt1), int32(p.Salt2), int32(p.Salt3), int32(p.Salt4), int32(p.Salt4))
	idx5 := mdx(int32(p.Salt1), int32(p.Salt2), int32(p.Salt3), int32(p.Salt4), int32(p.Salt4))

	return p.AccessToken[:idx1] +
		p.AccessToken[idx1+1:idx2] +
		p.AccessToken[idx2+1:idx3] +
		p.AccessToken[idx3+1:idx4] +
		p.AccessToken[idx4+1:idx5] +
		p.AccessToken[idx5+1:]
}
