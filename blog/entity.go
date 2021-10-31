package blog

import "time"

type Blog struct {
	id          int
	title       string
	thumbnail   string
	description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
