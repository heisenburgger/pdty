package domain

import "time"

type Project struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type Task struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Status    string     `json:"status"`
	Notes     *string    `json:"notes,omitempty"`
	ProjectID int        `json:"-"`
	Scheduled *time.Time `json:"scheduled,omitempty"`
	Deadline  *time.Time `json:"deadline,omitempty"`
}
