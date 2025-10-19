// Package auth
package auth

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	apitypes "github.com/prashunchitkr/nepse-go/internal/types"
)

// Token Holds auth data
type Token struct {
	AccessToken  string
	RefreshToken string
	Expiry       time.Time
	DummyID      int
}

// AuthHandler manages token lifecycle
type AuthHandler struct {
	mu     sync.Mutex
	token  *Token
	wasm   *WasmHelper
	client resty.Client
}

func NewAuthHandler(client resty.Client, cssWasm *WasmHelper) *AuthHandler {
	return &AuthHandler{
		client: client,
		wasm:   cssWasm,
	}
}

func (a *AuthHandler) GetToken(ctx context.Context) (*Token, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	var prove apitypes.Prove
	if a.token != nil && time.Now().Before(a.token.Expiry) {
		log.Printf("[AuthHandler] Returning old token: %v\n", a.token)
		return a.token, nil
	}

	resp, err := a.client.R().
		SetContext(ctx).
		SetResult(&prove).
		Get("/authenticate/prove")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New("getting prove failed: " + resp.Status())
	}

	expiry := time.Now().Add(45 * time.Second)

	accessToken, err := a.decodeAccessToken(ctx, prove)
	if err != nil {
		return nil, err
	}

	refreshToken, err := a.decodeRefreshToken(ctx, prove)
	if err != nil {
		return nil, err
	}

	a.token = &Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expiry:       expiry,
		DummyID:      -1,
	}

	return a.token, nil
}
