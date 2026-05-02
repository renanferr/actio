package domain

type ExecutionStatus string

const (
	StatusPending   ExecutionStatus = "pending"
	StatusRunning   ExecutionStatus = "running"
	StatusSucceeded ExecutionStatus = "succeeded"
	StatusFailed    ExecutionStatus = "failed"
	StatusCanceled  ExecutionStatus = "canceled"
)

func (s ExecutionStatus) IsTerminal() bool {
	return s == StatusSucceeded || s == StatusFailed || s == StatusCanceled
}

func (s ExecutionStatus) CanTransitionTo(next ExecutionStatus) bool {
	if s.IsTerminal() {
		return false
	}

	validTransitions := map[ExecutionStatus][]ExecutionStatus{
		StatusPending: {StatusRunning, StatusCanceled},
		StatusRunning: {StatusSucceeded, StatusFailed, StatusCanceled},
	}

	for _, candidate := range validTransitions[s] {
		if candidate == next {
			return true
		}
	}

	return false
}
