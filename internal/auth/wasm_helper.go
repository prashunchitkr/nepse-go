package auth

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed assets/css.wasm
var wasmFile []byte

type WasmHelper struct {
	rt     wazero.Runtime
	module api.Module
}

func NewWasmHelper(ctx context.Context) (*WasmHelper, error) {
	runtime := wazero.NewRuntime(ctx)
	wasi_snapshot_preview1.MustInstantiate(ctx, runtime)

	mod, err := runtime.InstantiateWithConfig(ctx, wasmFile, wazero.NewModuleConfig().WithSysNanosleep().WithStartFunctions("_initialize"))
	if err != nil {
		return nil, fmt.Errorf("error instantiating wasm: %v", err)
	}

	return &WasmHelper{
		rt:     runtime,
		module: mod,
	}, nil
}

func (w *WasmHelper) CDX(ctx context.Context, salt1, salt2, salt3, salt4, salt5 uint64) (uint64, error) {
	results, err := w.module.ExportedFunction("cdx").Call(ctx, salt1, salt2, salt3, salt4, salt5)
	if err != nil {
		return 0, fmt.Errorf("error calling rdx: %v", err)
	}

	return results[0], nil
}

func (w *WasmHelper) RDX(ctx context.Context, salt1, salt2, salt3, salt4, salt5 uint64) (uint64, error) {
	results, err := w.module.ExportedFunction("rdx").Call(ctx, salt1, salt2, salt3, salt4, salt5)
	if err != nil {
		return 0, fmt.Errorf("error calling rdx: %v", err)
	}

	return results[0], nil
}

func (w *WasmHelper) BDX(ctx context.Context, salt1, salt2, salt3, salt4, salt5 uint64) (uint64, error) {
	results, err := w.module.ExportedFunction("bdx").Call(ctx, salt1, salt2, salt3, salt4, salt5)
	if err != nil {
		return 0, fmt.Errorf("error calling rdx: %v", err)
	}

	return results[0], nil
}

func (w *WasmHelper) NDX(ctx context.Context, salt1, salt2, salt3, salt4, salt5 uint64) (uint64, error) {
	results, err := w.module.ExportedFunction("ndx").Call(ctx, salt1, salt2, salt3, salt4, salt5)
	if err != nil {
		return 0, fmt.Errorf("error calling rdx: %v", err)
	}

	return results[0], nil
}

func (w *WasmHelper) MDX(ctx context.Context, salt1, salt2, salt3, salt4, salt5 uint64) (uint64, error) {
	results, err := w.module.ExportedFunction("mdx").Call(ctx, salt1, salt2, salt3, salt4, salt5)
	if err != nil {
		return 0, fmt.Errorf("error calling rdx: %v", err)
	}

	return results[0], nil
}

func (w *WasmHelper) Close() error {
	ctx := context.Background()

	var errs []error
	if err := w.module.Close(ctx); err != nil {
		errs = append(errs, fmt.Errorf("failed to close modile: %w", err))
	}

	if err := w.rt.Close(ctx); err != nil {
		errs = append(errs, fmt.Errorf("failed to close runtime: %w", err))
	}

	if len(errs) > 0 {
		return fmt.Errorf("multiple errors during close: %v", errs)
	}

	return nil
}
