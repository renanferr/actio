package logger

import "fmt"

type Logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Debug(msg string)
}

type SimpleLogger struct {
	Verbose bool
}

func NewLogger(verbose bool) *SimpleLogger {
	return &SimpleLogger{Verbose: verbose}
}

func (l *SimpleLogger) Info(msg string) {
	fmt.Println("[INFO]", msg)
}

func (l *SimpleLogger) Warn(msg string) {
	fmt.Println("[WARN]", msg)
}

func (l *SimpleLogger) Error(msg string) {
	fmt.Println("[ERROR]", msg)
}

func (l *SimpleLogger) Debug(msg string) {
	if l.Verbose {
		fmt.Println("[DEBUG]", msg)
	}
}
