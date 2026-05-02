package config

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/acme/actio/internal/domain"
	"github.com/acme/actio/internal/ports"
	"gopkg.in/yaml.v3"
)

const (
	expectedAPIVersion = "actio.dev/v1alpha1"
	expectedKind       = "ActionPackage"
)

var (
	namePattern       = regexp.MustCompile(`^[a-z0-9-]{1,63}$`)
	semverPattern     = regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+$`)
	qualifiedActionID = regexp.MustCompile(`^[a-z0-9-]{1,63}\.[a-z0-9-]{1,63}$`)
)

type YAMLConfigLoader struct{}

func NewYAMLConfigLoader() ports.ConfigLoader {
	return &YAMLConfigLoader{}
}

func (l *YAMLConfigLoader) LoadActionPackage(path string) (*domain.ActionPackage, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, domain.FileNotFoundError{Path: path, Err: err}
		}
		return nil, err
	}

	pkg, err := l.ParseActionPackage(data)
	if err != nil {
		return nil, err
	}

	return pkg, nil
}

func (l *YAMLConfigLoader) ParseActionPackage(data []byte) (*domain.ActionPackage, error) {
	var pkg domain.ActionPackage
	if err := yaml.Unmarshal(data, &pkg); err != nil {
		return nil, domain.ParseError{Message: fmt.Sprintf("invalid YAML: %s", err.Error())}
	}

	if err := validateActionPackage(&pkg); err != nil {
		return nil, err
	}

	for actionName, action := range pkg.Actions {
		action.ID = domain.NewActionID(pkg.Domain, actionName)
		if action.Timeout != "" {
			dur, _ := time.ParseDuration(action.Timeout)
			action.TimeoutDuration = dur
		}
		pkg.Actions[actionName] = action
	}

	return &pkg, nil
}

func validateActionPackage(pkg *domain.ActionPackage) error {
	if pkg.APIVersion != expectedAPIVersion {
		return domain.ValidationError{Message: fmt.Sprintf("apiVersion must be %s", expectedAPIVersion)}
	}

	if pkg.Kind != expectedKind {
		return domain.ValidationError{Message: fmt.Sprintf("kind must be %s", expectedKind)}
	}

	if pkg.Metadata.Name == "" {
		return domain.ValidationError{Message: "metadata.name is required"}
	}

	if !namePattern.MatchString(pkg.Metadata.Name) {
		return domain.ValidationError{Message: "metadata.name must be lowercase alphanumeric with hyphens and no more than 63 characters"}
	}

	if pkg.Metadata.Version == "" {
		return domain.ValidationError{Message: "metadata.version is required"}
	}

	if !semverPattern.MatchString(pkg.Metadata.Version) {
		return domain.ValidationError{Message: "metadata.version must be semantic version MAJOR.MINOR.PATCH"}
	}

	if pkg.Domain == "" {
		return domain.ValidationError{Message: "domain is required"}
	}

	if !namePattern.MatchString(pkg.Domain) {
		return domain.ValidationError{Message: "domain must be lowercase alphanumeric with hyphens and no more than 63 characters"}
	}

	if pkg.Actions == nil {
		pkg.Actions = map[string]domain.ActionSpec{}
	}

	for actionName, action := range pkg.Actions {
		if actionName == "" {
			return domain.ValidationError{Message: "action name must not be empty"}
		}

		if !namePattern.MatchString(actionName) {
			return domain.ValidationError{Message: fmt.Sprintf("action name %q must be lowercase alphanumeric with hyphens and no more than 63 characters", actionName)}
		}

		if action.Image == "" {
			return domain.ValidationError{Message: fmt.Sprintf("actions.%s.image is required", actionName)}
		}

		if action.Command != nil && len(action.Command) == 0 {
			return domain.ValidationError{Message: fmt.Sprintf("actions.%s.command must be a non-empty array", actionName)}
		}

		if action.Args != nil && len(action.Args) == 0 {
			return domain.ValidationError{Message: fmt.Sprintf("actions.%s.args must be a non-empty array", actionName)}
		}

		if action.DependsOn != nil && len(action.DependsOn) == 0 {
			return domain.ValidationError{Message: fmt.Sprintf("actions.%s.depends_on must be a non-empty array", actionName)}
		}

		for _, dependency := range action.DependsOn {
			if !qualifiedActionID.MatchString(dependency) {
				return domain.ValidationError{Message: fmt.Sprintf("actions.%s.depends_on contains invalid action id %q", actionName, dependency)}
			}
			if dependency == fmt.Sprintf("%s.%s", pkg.Domain, actionName) {
				return domain.ValidationError{Message: fmt.Sprintf("actions.%s.depends_on must not reference itself", actionName)}
			}
		}

		if action.Timeout != "" {
			duration, err := time.ParseDuration(action.Timeout)
			if err != nil {
				return domain.ValidationError{Message: fmt.Sprintf("actions.%s.timeout is invalid: %s", actionName, err.Error())}
			}
			if duration <= 0 {
				return domain.ValidationError{Message: fmt.Sprintf("actions.%s.timeout must be positive", actionName)}
			}
		}
	}

	return nil
}
