package auth_test

import (
	"context"
	"testing"

	"github.com/prashunchitkr/nepse-go/internal/auth"
)

func TestWasmHelper(t *testing.T) {
	ctx := context.Background()
	wasm, err := auth.NewWasmHelper(ctx)
	if err != nil {
		t.Errorf("cannot init wasm: %v", err)
	}
	wasm.Close()
}

func TestCDXCall(t *testing.T) {
	want := uint64(26)
	ctx := context.Background()

	wasm, err := auth.NewWasmHelper(ctx)
	if err != nil {
		t.Errorf("cannot init wasm: %v", err)
	}

	res, err := wasm.CDX(ctx, 1, 2, 3, 4, 5)
	if err != nil {
		t.Errorf("error calling cdx: %v", err)
	}

	if res != want {
		t.Errorf("Expected: %d, Recieved: %d", want, res)
	}

	wasm.Close()
}
