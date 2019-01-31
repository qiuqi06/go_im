package model

import "net"

type User struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}
type CurrentUser struct {
	User
	Conn net.Conn
}

func (this *CurrentUser) Fc()  {

}
