package ports

import "github.com/acme/actio/internal/domain"

type Validator interface {
	ValidateActionPackage(pkg *domain.ActionPackage) ValidationResult
	ValidateDependencies(actions []*domain.Action, deps []*domain.Dependency) ValidationResult
}

type ValidationResult struct {
	Errors   []string
	Warnings []string
	Valid    bool
}
