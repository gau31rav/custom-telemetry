package model

import "github.com/google/uuid"

type Order struct {
	ID    uuid.UUID `json:"id"`
	Items []Item    `json:"items"`
	User  User      `josn:"user"`
}

type User struct {
	ID    int    `json:"ID"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
