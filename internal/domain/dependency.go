package domain

import "fmt"

type Dependency struct {
	FromActionID ActionID // Action that must complete
	ToActionID   ActionID // Action that depends on it
}

func (d *Dependency) Validate() error {
	if d.FromActionID == d.ToActionID {
		return fmt.Errorf("action cannot depend on itself: %s", d.FromActionID)
	}
	return nil
}
