// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package models

import (
	"database/sql"
)

type User struct {
	ID        int32        `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
