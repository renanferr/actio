package domain

import "time"

type Domain struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
}
