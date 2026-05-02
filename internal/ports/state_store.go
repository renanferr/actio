package ports

import (
	"context"

	"github.com/acme/actio/internal/domain"
)

type StateStore interface {
	// Domain operations
	CreateDomain(ctx context.Context, name, description string) (*domain.Domain, error)
	GetDomain(ctx context.Context, name string) (*domain.Domain, error)
	ListDomains(ctx context.Context) ([]*domain.Domain, error)

	// Action operations
	CreateAction(ctx context.Context, action *domain.Action) error
	UpdateAction(ctx context.Context, action *domain.Action) error
	GetAction(ctx context.Context, actionID string) (*domain.Action, error)
	ListActionsByDomain(ctx context.Context, domainName string) ([]*domain.Action, error)
	ActionExists(ctx context.Context, actionID string) (bool, error)

	// Dependency operations
	SetDependencies(ctx context.Context, actionID string, deps []*domain.Dependency) error
	GetDependencies(ctx context.Context, actionID string) ([]*domain.Dependency, error)
	GetDependentsOf(ctx context.Context, actionID string) ([]*domain.Action, error)

	// Execution operations
	CreateExecution(ctx context.Context, exec *domain.Execution) error
	UpdateExecution(ctx context.Context, exec *domain.Execution) error
	GetExecution(ctx context.Context, executionID string) (*domain.Execution, error)
	ListExecutionsByAction(ctx context.Context, actionID string, limit int) ([]*domain.Execution, error)

	// Transaction support
	BeginTx(ctx context.Context) (Transaction, error)

	// Shutdown
	Close() error
}

type Transaction interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error

	CreateDomain(ctx context.Context, name, description string) (*domain.Domain, error)
	GetDomain(ctx context.Context, name string) (*domain.Domain, error)
	ListDomains(ctx context.Context) ([]*domain.Domain, error)

	CreateAction(ctx context.Context, action *domain.Action) error
	UpdateAction(ctx context.Context, action *domain.Action) error
	GetAction(ctx context.Context, actionID string) (*domain.Action, error)
	ListActionsByDomain(ctx context.Context, domainName string) ([]*domain.Action, error)
	ActionExists(ctx context.Context, actionID string) (bool, error)

	SetDependencies(ctx context.Context, actionID string, deps []*domain.Dependency) error
	GetDependencies(ctx context.Context, actionID string) ([]*domain.Dependency, error)
	GetDependentsOf(ctx context.Context, actionID string) ([]*domain.Action, error)

	CreateExecution(ctx context.Context, exec *domain.Execution) error
	UpdateExecution(ctx context.Context, exec *domain.Execution) error
	GetExecution(ctx context.Context, executionID string) (*domain.Execution, error)
	ListExecutionsByAction(ctx context.Context, actionID string, limit int) ([]*domain.Execution, error)
}
