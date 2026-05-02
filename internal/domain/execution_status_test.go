package domain

import "testing"

func TestExecutionStatusTransitions(t *testing.T) {
	if !StatusPending.CanTransitionTo(StatusRunning) {
		t.Fatal("expected pending -> running to be valid")
	}

	if !StatusPending.CanTransitionTo(StatusCanceled) {
		t.Fatal("expected pending -> canceled to be valid")
	}

	if StatusPending.CanTransitionTo(StatusSucceeded) {
		t.Fatal("expected pending -> succeeded to be invalid")
	}

	if !StatusRunning.CanTransitionTo(StatusSucceeded) {
		t.Fatal("expected running -> succeeded to be valid")
	}

	if !StatusRunning.CanTransitionTo(StatusFailed) {
		t.Fatal("expected running -> failed to be valid")
	}

	if !StatusRunning.CanTransitionTo(StatusCanceled) {
		t.Fatal("expected running -> canceled to be valid")
	}

	if StatusSucceeded.CanTransitionTo(StatusRunning) {
		t.Fatal("expected succeeded -> running to be invalid")
	}

	if StatusFailed.CanTransitionTo(StatusCanceled) {
		t.Fatal("expected failed -> canceled to be invalid")
	}
}

func TestExecutionStatusIsTerminal(t *testing.T) {
	if !StatusSucceeded.IsTerminal() {
		t.Fatal("expected succeeded to be terminal")
	}
	if !StatusFailed.IsTerminal() {
		t.Fatal("expected failed to be terminal")
	}
	if !StatusCanceled.IsTerminal() {
		t.Fatal("expected canceled to be terminal")
	}
	if StatusPending.IsTerminal() {
		t.Fatal("expected pending to be non-terminal")
	}
}
