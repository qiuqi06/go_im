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

var CurUser CurrentUser

func init() {
	CurUser = CurrentUser{
		Conn: nil,
		User: nil,
	}

}
