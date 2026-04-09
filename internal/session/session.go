// Package session
package session

import (
	"sync"

	"github.com/prashunchitkr/nepse-go/internal/models"
)

type Manager struct {
	prove       *models.Prove
	salterToken string
	accessToken string
	sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{}
}

func (s *Manager) IsExpired() bool {
	return false
}

func (s *Manager) SaveTokens(accessToken string) {
	s.Lock()
	defer s.Unlock()
	s.accessToken = accessToken
}

func (s *Manager) GetTokens() string {
	s.RLock()
	defer s.RUnlock()
	return s.accessToken
}

func (s *Manager) SetFromProve(prove *models.Prove, salterToken string) {
	s.Lock()
	defer s.Unlock()
	if prove == nil {
		s.prove = nil
		s.salterToken = ""
		return
	}
	cp := new(models.Prove)
	*cp = *prove
	s.prove = cp
	s.salterToken = salterToken
	s.accessToken = prove.AccessToken
}

func (s *Manager) SalterAuthorization() (token string, ok bool) {
	s.RLock()
	defer s.RUnlock()
	if s.prove == nil || s.salterToken == "" {
		return "", false
	}
	return s.salterToken, true
}
