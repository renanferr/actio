# Actio

Actio is a Go-based CLI application for defining, validating, and applying declarative action packages. It is built around a hexagonal architecture with a lightweight domain model, pluggable adapters for configuration loading, execution, and persistence, and a CLI for local developer workflows.

## What it does

- Parses `ActionPackage` YAML manifests
- Validates package schema and action definitions
- Supports action dependencies, timeouts, and environment configuration
- Provides a CLI for `init`, `validate`, `apply`, `list`, `inspect`, `run`, and `graph`
- Uses a modular port-and-adapter design for easy extension

## MVP Scope

- Local YAML configuration loader
- Domain-driven parsing and validation
- Docker executor support planned in adapters
- SQLite persistence support planned in the state adapter
- CLI entrypoint under `cmd/actio`

## Repository Structure

- `cmd/actio/` - CLI entrypoint and bootstrap
- `internal/adapters/` - adapter implementations for CLI, config loading, executors, and persistence
- `internal/application/` - application service layer and orchestration logic
- `internal/domain/` - core business entities and domain error types
- `internal/logger/` - simple logging utilities
- `internal/ports/` - interface definitions for config loader, executor, state store, and validator
- `specs/` - specification and markdown files (ignored by Git)

## Key Components

### Domain

- `internal/domain/action_package.go` - ActionPackage schema model
- `internal/domain/action.go` - action entity and ID helpers
- `internal/domain/errors.go` - typed domain errors

### Ports

- `internal/ports/config_loader.go` - config loader interface
- `internal/ports/executor.go` - executor interface
- `internal/ports/state_store.go` - persistence interface
- `internal/ports/validator.go` - validation interface

### Adapters

- `internal/adapters/config/loader.go` - YAML `ActionPackage` parser and validator
- `internal/adapters/cli/` - CLI command wiring and command implementations
- `internal/adapters/executor/` - executor implementations
- `internal/adapters/state/` - persistence adapters

## Getting Started

### Requirements

- Go 1.21
- Git
- Docker (for executor-related workflows, if implemented)

### Build

```sh
go build ./cmd/actio
```

### Test

```sh
go test ./...
```

## CLI Usage

From the repository root, run:

```sh
go run ./cmd/actio --help
```

Example commands:

- `actio init` - initialize project or example packages
- `actio validate <path>` - validate an ActionPackage YAML file
- `actio apply <path>` - apply an ActionPackage to the registry
- `actio list` - list registered actions
- `actio inspect <action-id>` - inspect a registered action
- `actio run <action-id>` - execute an action and its dependency graph
- `actio graph <action-id>` - render the dependency graph for an action

## Development Notes

- Use the port-and-adapter pattern to keep domain logic isolated
- Add new execution or persistence behavior by implementing the relevant port interfaces
- Domain validation should remain independent from CLI and adapter concerns
- `specs/` contains project specifications and is intentionally excluded from version control via `.gitignore`

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add or update tests
4. Run `go test ./...`
5. Submit a pull request

## License

This project is licensed under the Apache License, Version 2.0. See `LICENSE` for details.
