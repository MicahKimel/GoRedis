package data

import (
	"encoding/json"
	"io"
)

type User struct {
	username string `json:"username"`
	password string `json:"password"`
	phone string	`json:"phone"`
	email string	`json:"email"`
}

type Users []*User

func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}