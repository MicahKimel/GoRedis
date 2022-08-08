package data

import (
	"encoding/json"
	"io"
)

// swagger:parameters User
type User struct {
	// Username
	// in: body
	Username string `json:"username"`
	// password
	// in: body
	Password string `json:"password"`
	// phone
	// in: body
	Phone string `json:"phone"`
	// email
	// in: body
	Email string `json:"email"`
}

// swagger:model Users
type Users []*User

func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

// swagger:parameters Group
type Group struct {
	// name
	// in: body
	Name string `json:"name"`
	// groupid
	// in: body
	Groupid string `json:"groupid"`
}

// swagger:model Groups
type Groups []*Group

func (u *Groups) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *Group) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}
