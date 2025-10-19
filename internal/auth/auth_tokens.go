package auth

import (
	"context"

	apitypes "github.com/prashunchitkr/nepse-go/internal/types"
)

func (a *AuthHandler) decodeAccessToken(ctx context.Context, prove apitypes.Prove) (string, error) {
	idx1, err := a.wasm.CDX(ctx, prove.Salt1, prove.Salt2, prove.Salt3, prove.Salt4, prove.Salt5)
	if err != nil {
		return "", err
	}

	idx2, err := a.wasm.RDX(ctx, prove.Salt1, prove.Salt2, prove.Salt3, prove.Salt4, prove.Salt5)
	if err != nil {
		return "", err
	}

	idx3, err := a.wasm.BDX(ctx, prove.Salt1, prove.Salt2, prove.Salt3, prove.Salt4, prove.Salt5)
	if err != nil {
		return "", err
	}

	idx4, err := a.wasm.NDX(ctx, prove.Salt1, prove.Salt2, prove.Salt3, prove.Salt4, prove.Salt5)
	if err != nil {
		return "", err
	}

	idx5, err := a.wasm.MDX(ctx, prove.Salt1, prove.Salt2, prove.Salt3, prove.Salt4, prove.Salt5)
	if err != nil {
		return "", err
	}

	decodedToken := prove.AccessToken[:idx1] +
		prove.AccessToken[idx1+1:idx2] +
		prove.AccessToken[idx2+1:idx3] +
		prove.AccessToken[idx3+1:idx4] +
		prove.AccessToken[idx4+1:idx5] +
		prove.AccessToken[idx5+1:]

	return decodedToken, nil
}

func (a *AuthHandler) decodeRefreshToken(ctx context.Context, prove apitypes.Prove) (string, error) {
	idx1, err := a.wasm.CDX(ctx, prove.Salt2, prove.Salt1, prove.Salt3, prove.Salt5, prove.Salt4)
	if err != nil {
		return "", err
	}

	idx2, err := a.wasm.RDX(ctx, prove.Salt2, prove.Salt1, prove.Salt3, prove.Salt5, prove.Salt4)
	if err != nil {
		return "", err
	}

	idx3, err := a.wasm.BDX(ctx, prove.Salt2, prove.Salt1, prove.Salt4, prove.Salt3, prove.Salt5)
	if err != nil {
		return "", err
	}

	idx4, err := a.wasm.NDX(ctx, prove.Salt2, prove.Salt1, prove.Salt4, prove.Salt3, prove.Salt5)
	if err != nil {
		return "", err
	}

	idx5, err := a.wasm.MDX(ctx, prove.Salt2, prove.Salt1, prove.Salt4, prove.Salt3, prove.Salt5)
	if err != nil {
		return "", err
	}

	decodedToken := prove.RefreshToken[:idx1] +
		prove.RefreshToken[idx1+1:idx2] +
		prove.RefreshToken[idx2+1:idx3] +
		prove.RefreshToken[idx3+1:idx4] +
		prove.RefreshToken[idx4+1:idx5] +
		prove.RefreshToken[idx5+1:]

	return decodedToken, nil
}
