package todo

import "time"

type Todo struct {
	ID        int
	Content   string
	createdAt time.Time
	Done      bool
}
