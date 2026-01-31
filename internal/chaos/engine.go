package chaos

import (
	"context"
	"errors"
	"math/rand"
	"time"
)

var ErrInjectedFailure = errors.New("chaos: injected failure")

type Engine struct {
	store Store
	rand  *rand.Rand
}

func NewEngine(store Store) *Engine {
	return &Engine{
		store: store,
		rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (e *Engine) Apply(ctx context.Context, path, method string) error {
	cfg, ok := e.store.Find(path, method)
	if !ok {
		return nil
	}

	// Delay
	if cfg.MaxDelayMs > 0 {
		delay := e.rand.Intn(cfg.MaxDelayMs-cfg.MinDelayMs+1) + cfg.MinDelayMs
		select {
		case <-time.After(time.Duration(delay) * time.Millisecond):
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	// Error injection
	if e.rand.Float64() < cfg.ErrorRate {
		return ErrInjectedFailure
	}

	return nil
}
