package domain

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

type ExecutionID string

type Execution struct {
	ID          ExecutionID
	ActionID    ActionID
	Status      ExecutionStatus
	StartedAt   *time.Time
	CompletedAt *time.Time
	ExitCode    *int
	Stdout      string
	CreatedAt   time.Time
}

func NewExecution(actionID ActionID) *Execution {
	return &Execution{
		ID:        ExecutionID(generateUniqueID()),
		ActionID:  actionID,
		Status:    StatusPending,
		CreatedAt: time.Now().UTC(),
	}
}

func (e *Execution) Start() error {
	if !e.Status.CanTransitionTo(StatusRunning) {
		return fmt.Errorf("cannot start execution in status %s", e.Status)
	}
	if e.StartedAt != nil {
		return fmt.Errorf("execution already started")
	}
	now := time.Now().UTC()
	e.Status = StatusRunning
	e.StartedAt = &now
	return nil
}

func (e *Execution) Complete(status ExecutionStatus, exitCode int, output string) error {
	if !e.Status.CanTransitionTo(status) {
		return fmt.Errorf("invalid transition: %s -> %s", e.Status, status)
	}
	code := exitCode
	e.Status = status
	e.ExitCode = &code
	e.Stdout = output
	now := time.Now().UTC()
	e.CompletedAt = &now
	return nil
}

func generateUniqueID() string {
	buf := make([]byte, 12)
	if _, err := rand.Read(buf); err != nil {
		return fmt.Sprintf("exec-%d", time.Now().UnixNano())
	}
	return "exec-" + hex.EncodeToString(buf)
}
