package ports

import "github.com/acme/actio/internal/domain"

type ConfigLoader interface {
	LoadActionPackage(path string) (*domain.ActionPackage, error)
	ParseActionPackage(data []byte) (*domain.ActionPackage, error)
}
