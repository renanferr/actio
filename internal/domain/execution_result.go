package domain

import "time"

type ExecutionResult struct {
	ExecutionID      ExecutionID
	ActionID         ActionID
	Status           ExecutionStatus
	ExitCode         int
	Stdout           string
	Duration         time.Duration
	StartedAt        time.Time
	CompletedAt      time.Time
	ExecutorType     string
	ExecutorLocation string
	ExecutorVersion  string
	Error            string
}

func (r *ExecutionResult) IsSuccessful() bool {
	return r.Status == StatusSucceeded && r.ExitCode == 0
}

func (r *ExecutionResult) IsFailed() bool {
	return r.Status == StatusFailed
}

func (r *ExecutionResult) IsCanceled() bool {
	return r.Status == StatusCanceled
}
