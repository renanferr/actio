package domain

import "testing"

func TestDependencyValidateSelfDependency(t *testing.T) {
	d := &Dependency{
		FromActionID: NewActionID("billing", "bill"),
		ToActionID:   NewActionID("billing", "bill"),
	}

	if err := d.Validate(); err == nil {
		t.Fatal("expected self-dependency validation to fail")
	}
}

func TestDependencyValidateDistinctDependency(t *testing.T) {
	d := &Dependency{
		FromActionID: NewActionID("billing", "bill"),
		ToActionID:   NewActionID("billing", "invoice"),
	}

	if err := d.Validate(); err != nil {
		t.Fatal("expected distinct dependency validation to succeed")
	}
}
