package domain

import "testing"

func TestExecutionResultHelpers(t *testing.T) {
	r := &ExecutionResult{
		Status:   StatusSucceeded,
		ExitCode: 0,
	}

	if !r.IsSuccessful() {
		t.Fatal("expected execution result to be successful")
	}

	if r.IsFailed() {
		t.Fatal("expected execution result not to be failed")
	}

	if r.IsCanceled() {
		t.Fatal("expected execution result not to be canceled")
	}

	r.Status = StatusFailed
	r.ExitCode = 1

	if r.IsSuccessful() {
		t.Fatal("expected execution result not to be successful")
	}

	if !r.IsFailed() {
		t.Fatal("expected execution result to be failed")
	}

	if r.IsCanceled() {
		t.Fatal("expected execution result not to be canceled")
	}

	r.Status = StatusCanceled

	if !r.IsCanceled() {
		t.Fatal("expected execution result to be canceled")
	}
}
