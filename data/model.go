package data

import (
	"encoding/json"
	"io"
)

// swagger:model User
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

// swagger:model Group
type Group struct {
	// group name
	// in: body
	Name string `json:"name"`
	// group id
	// in: body
	Groupid string `json: Groupid`
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
