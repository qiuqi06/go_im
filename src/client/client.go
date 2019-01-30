package main

import (
	"common"
	"fmt"
	"net"
	"encoding/json"
	"unsafe"
)

func main() {
	var key int
	var loop = true
	for loop {
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			print("ss")
			loop = false
		case 2:
			print("ss")
			loop = false
		default:
			print("error")

		}
	}
	login(1,"qq")
	fmt.Print("ss")

}

func login(userId int, userPwd string) (err error) {
	print("err")
	conn, err := net.Dial("tcp", "localhost:8089")
	defer conn.Close()
	var msg common.Message
	msg.Type = common.LoginMsgType

	var  x unsafe.Pointer
	x = unsafe.Pointer(&msg)
	println(x)

	var loginMsg common.LoginMsg
	loginMsg.UserId = 1
	loginMsg.UserName = "qq"
	loginMsg.UserPwd = "123"

	bytes, err := json.Marshal(loginMsg)
	msg.Data=string(bytes)

	transfer := common.Transfer{
		Conn: conn,
	}
	dataSlice, err := json.Marshal(msg)
	transfer.WritePkg( dataSlice)

	return

}
