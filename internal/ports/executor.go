package ports

import (
	"context"

	"github.com/acme/actio/internal/domain"
)

type Executor interface {
	Execute(ctx context.Context, action *domain.Action, env map[string]string) (*domain.ExecutionResult, error)
	HealthCheck(ctx context.Context) error
	Type() string
	Version() string
	Location() string
}
