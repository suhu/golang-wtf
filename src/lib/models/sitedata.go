package models

import (
	"github.com/fxtlabs/date"
)

type SiteData struct {
	Id           int
	Country      string
	Domain       string
	Name         string
	Organization string
	Address      string
	Postalcode   string
	Email        string
	IPAddress    string
	Phone        string
	Phone2       string
	Fax          string
	Outcome      string
	CreatedDate  date.Date
	ExpiresDate  date.Date
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Exception struct {
	Message string `json:"message"`
}

type JwtToken struct {
	Token string `json:"token"`
}
