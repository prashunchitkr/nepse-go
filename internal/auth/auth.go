// Package auth
package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

func (a *Manager) Authenticate(ctx context.Context) error { return a.prove(ctx) }
func (a *Manager) Refresh(ctx context.Context) error      { return a.prove(ctx) }

func (a *Manager) prove(ctx context.Context) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, a.baseURL+"/api/authenticate/prove", nil)
	resp, err := a.publicClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("prove: unexpected status %s", resp.Status)
	}

	var p models.Prove
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return fmt.Errorf("prove: decode: %w", err)
	}

	a.session.SetFromProve(&p, deriveSalterToken(&p))
	return nil
}

func (a *Manager) InjectHeaders(req *http.Request) error {
	// Double-checked locking to ensure we only prove once
	if _, ok := a.session.SalterAuthorization(); !ok {
		if err := a.prove(req.Context()); err != nil {
			return err
		}
	}

	token, ok := a.session.SalterAuthorization()
	if !ok {
		return fmt.Errorf("auth: session missing after prove")
	}

	req.Header.Set("Authorization", "Salter "+token)
	return nil
}

func deriveSalterToken(p *models.Prove) string {
	s1, s2, s3, s4, s5 := int32(p.Salt1), int32(p.Salt2), int32(p.Salt3), int32(p.Salt4), int32(p.Salt5)

	idxs := []int32{
		cdx(s1, s2, s3, s4, s5),
		rdx(s1, s2, s3, s4, s5),
		bdx(s1, s2, s3, s4, s5),
		ndx(s1, s2, s3, s4, s5),
		mdx(s1, s2, s3, s4, s5),
	}

	var res strings.Builder
	res.WriteString(p.AccessToken[:idxs[0]])
	for i := 0; i < len(idxs)-1; i++ {
		res.WriteString(p.AccessToken[idxs[i]+1 : idxs[i+1]])
	}
	return res.String() + p.AccessToken[idxs[4]+1:]
}
