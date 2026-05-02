package domain

import "time"

type ActionID string // Format: domain.action-name (globally unique)

type Action struct {
	ID        ActionID
	Domain    string
	Name      string
	Image     string
	Command   []string
	Args      []string
	Env       map[string]string
	Timeout   time.Duration
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewActionID(domain, name string) ActionID {
	return ActionID(domain + "." + name)
}

func (a ActionID) Domain() string {
	id := string(a)
	for i := 0; i < len(id); i++ {
		if id[i] == '.' {
			return id[:i]
		}
	}
	return ""
}

func (a ActionID) Name() string {
	id := string(a)
	for i := 0; i < len(id); i++ {
		if id[i] == '.' {
			return id[i+1:]
		}
	}
	return ""
}
