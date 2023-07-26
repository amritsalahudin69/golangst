package model

import "time"

type Campus struct {
	ID         uint       `json:"id"`
	Name       string     `json:"name"`
	Address    string     `json:"address"`
	Rektor     string     `json:"rektor"`
	Akreditasi string     `json:"akreditasi"`
	Email      string     `json:"email"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeleteAt   *time.Time `json:"delete_at"`
}

func (m Campus) GetTableName() string {
	return `campus`
}
