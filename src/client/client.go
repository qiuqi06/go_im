package main

import (
	"common"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
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
	if err != nil {

	}
	var msg common.Message
	msg.Type = common.LoginMsgType

	var loginMsg common.LoginMsg
	loginMsg.UserId = 1
	loginMsg.UserName = "qq"
	loginMsg.UserPwd = "123"

	data, _ := json.Marshal(loginMsg)
	msg.Data = string(data)

	var bytes [4]byte

	pkgLen := uint32(len(data))

	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)

	n, _ := conn.Write(bytes[:])
	println(n)

	print(conn)
	return nil
}
