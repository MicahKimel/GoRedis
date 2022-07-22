package data

import (
	"encoding/json"
	"io"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
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

type Group struct {
	Name string `json:"name"`
}

type Groups []*Group

func (u *Groups) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *Group) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}
