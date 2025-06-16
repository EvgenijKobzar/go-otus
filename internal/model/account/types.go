package account

import "time"

type Entity struct {
	Id        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
