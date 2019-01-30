package process

import (
	"net"
	"common"
	"encoding/json"
)

type UserProcess struct {
	Conn net.Conn
	UserId int
}

func (this *UserProcess) ServerProcessMsg(msg *common.Message) (err error) {
	switch msg.Type {
	case common.RegisterMsgType:
		this.ServerLogin(msg)
	case common.LoginMsgType:
	default:
		return
	}
	return
}
func (this *UserProcess) ServerLogin(message *common.Message) (err error) {
	var loginMsg common.LoginMsg
	err = json.Unmarshal([]byte(message.Data), &loginMsg)
	println(loginMsg)
    //todo去数据库获取
    userMgr.AddOnlineUser(this)

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

