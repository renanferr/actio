package config

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/acme/actio/internal/domain"
)

func TestParseActionPackageValid(t *testing.T) {
	yamlData := []byte(`apiVersion: actio.dev/v1alpha1
kind: ActionPackage
metadata:
  name: billing-suite
  version: 1.2.0
  description: End-of-month billing orchestration
  created: 2026-04-30T10:30:00Z
domain: billing
actions:
  import-invoices:
    image: ghcr.io/acme/billing/import:1.0.0
    command: ["python", "main.py"]
    args: ["--format", "csv"]
    timeout: 30m
`)

	loader := NewYAMLConfigLoader()
	pkg, err := loader.ParseActionPackage(yamlData)
	if err != nil {
		t.Fatalf("expected valid package, got %v", err)
	}

	if pkg.Metadata.Name != "billing-suite" {
		t.Fatalf("expected metadata.name %q, got %q", "billing-suite", pkg.Metadata.Name)
	}

	if pkg.Actions == nil || pkg.Actions["import-invoices"].Image != "ghcr.io/acme/billing/import:1.0.0" {
		t.Fatalf("expected action 'import-invoices' to be parsed")
	}

	if pkg.Actions["import-invoices"].ID != domain.NewActionID("billing", "import-invoices") {
		t.Fatalf("expected action ID to be set, got %q", pkg.Actions["import-invoices"].ID)
	}

	if pkg.Actions["import-invoices"].TimeoutDuration != 30*time.Minute {
		t.Fatalf("expected timeout duration to parse to 30m, got %v", pkg.Actions["import-invoices"].TimeoutDuration)
	}
}

func TestParseActionPackageInvalidYAML(t *testing.T) {
	yamlData := []byte(`apiVersion: actio.dev/v1alpha1
kind: ActionPackage
metadata:
  name: billing-suite
  version: 1.2.0
 domain: billing
`)

	loader := NewYAMLConfigLoader()
	_, err := loader.ParseActionPackage(yamlData)
	if err == nil {
		t.Fatal("expected YAML parse error")
	}
}

func TestParseActionPackageMissingRequiredField(t *testing.T) {
	yamlData := []byte(`apiVersion: actio.dev/v1alpha1
kind: ActionPackage
metadata:
  name: billing-suite
  version: 1.2.0
domain: billing
actions:
  import-invoices:
    command: ["python", "main.py"]
`)

	loader := NewYAMLConfigLoader()
	_, err := loader.ParseActionPackage(yamlData)
	if err == nil {
		t.Fatal("expected validation error for missing image")
	}

	if _, ok := err.(domain.ValidationError); !ok {
		t.Fatalf("expected ValidationError, got %T", err)
	}
}

func TestLoadActionPackageNotFound(t *testing.T) {
	loader := NewYAMLConfigLoader()
	path := filepath.Join(t.TempDir(), "missing.yaml")
	_, err := loader.LoadActionPackage(path)
	if err == nil {
		t.Fatal("expected file not found error")
	}

	if _, ok := err.(domain.FileNotFoundError); !ok {
		t.Fatalf("expected FileNotFoundError, got %T", err)
	}

	if !os.IsNotExist(err.(domain.FileNotFoundError).Err) {
		t.Fatalf("expected underlying error to be OS not exist, got %v", err)
	}
}
