package main

import (
	"net"
	"common"
	"encoding/json"
)

type UserProcessor struct {
	Conn net.Conn
}

func (this *UserProcessor) serverProcessMsg(msg *common.Message) (err error) {
	switch msg.Type {
	case common.RegisterMsgType:
		this.serverLogin(msg)
	case common.LoginMsgType:
	default:
		return
	}
	return
}
func (this *UserProcessor) serverLogin(message *common.Message) (err error) {
	var loginMsg common.LoginMsg
	err = json.Unmarshal([]byte(message.Data), &loginMsg)
	println(loginMsg)

	var loginRes common.LoginResMsg
	loginRes.Code = 100
	loginRes.Error = "err"
	data, err := json.Marshal(loginRes)
	transfer := common.Transfer{
		Conn: this.Conn,
	}
	transfer.WritePkg(data)
	return
}
