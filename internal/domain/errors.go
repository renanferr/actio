package domain

type ParseError struct {
	Message string
}

type FileNotFoundError struct {
	Path string
	Err  error
}

type ValidationError struct {
	Message string
}

type DomainError struct {
	Message string
}

type ExecutionError struct {
	Message string
}

type CycleDetectedError struct {
	Message string
}

func (e ParseError) Error() string {
	return e.Message
}

func (e FileNotFoundError) Error() string {
	return e.Err.Error()
}

func (e FileNotFoundError) Unwrap() error {
	return e.Err
}

func (e ValidationError) Error() string {
	return e.Message
}

func (e DomainError) Error() string {
	return e.Message
}

func (e ExecutionError) Error() string {
	return e.Message
}

func (e CycleDetectedError) Error() string {
	return e.Message
}
