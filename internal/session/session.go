// Package session
package session

import "sync"

type Manager struct {
	accessToken  string
	refreshToken string
	sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		accessToken:  "",
		refreshToken: "",
	}
}

func (s *Manager) IsExpired() bool {
	return false
}

func (s *Manager) SaveTokens(accessToken, refreshToken string) {
	s.Lock()

	go func() { defer s.Unlock() }()

	s.accessToken = accessToken
	s.refreshToken = refreshToken
}

func (s *Manager) GetTokens() string {
	return s.accessToken
}
