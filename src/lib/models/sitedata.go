package models

import (
	"github.com/fxtlabs/date"
)

type SiteData struct {
	Id           int       `json:"id"`
	Country      string    `json:"country"`
	Domain       string    `json:"domain"`
	Name         string    `json:"name"`
	Organization string    `json:"organization"`
	Address      string    `json:"address"`
	Postalcode   string    `json:"postalcode"`
	Email        string    `json:"email"`
	IPAddress    string    `json:"ipaddress"`
	Phone        string    `json:"phone"`
	Phone2       string    `json:"phone2"`
	Fax          string    `json:"fax"`
	Outcome      string    `json:"outcome"`
	CreatedDate  date.Date `json:"createddate"`
	ExpiresDate  date.Date `json:"expiresdate"`
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
