package domain

import "time"

type ActionPackage struct {
	APIVersion string                `yaml:"apiVersion"`
	Kind       string                `yaml:"kind"`
	Metadata   ActionPackageMetadata `yaml:"metadata"`
	Domain     string                `yaml:"domain"`
	Actions    map[string]ActionSpec `yaml:"actions"`
	CreatedAt  time.Time
}

type ActionPackageMetadata struct {
	Name        string    `yaml:"name"`
	Version     string    `yaml:"version"`
	Description string    `yaml:"description"`
	Created     time.Time `yaml:"created"`
}

type ActionSpec struct {
	ID              ActionID              `yaml:"-"`
	Image           string                `yaml:"image"`
	Command         []string              `yaml:"command"`
	Args            []string              `yaml:"args"`
	Env             map[string]string     `yaml:"env"`
	DependsOn       []string              `yaml:"depends_on"`
	Timeout         string                `yaml:"timeout"`
	TimeoutDuration time.Duration         `yaml:"-"`
	Resources       ResourceSpec          `yaml:"resources"`
	Namespace       string                `yaml:"namespace"`
	ServiceAccount  string                `yaml:"serviceAccount"`
	Secrets         map[string]SecretSpec `yaml:"secrets"`
}

type ResourceSpec struct {
	Limits   ResourceQuantity `yaml:"limits"`
	Requests ResourceQuantity `yaml:"requests"`
}

type ResourceQuantity struct {
	CPU    string `yaml:"cpu"`
	Memory string `yaml:"memory"`
}

type SecretSpec struct {
	Name string `yaml:"name"`
	Key  string `yaml:"key"`
}
