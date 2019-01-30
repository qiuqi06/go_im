package process

import (
	"common"
	"encoding/json"
	"model"
)

type SmsProcess struct {
}

func init() {

}
func (this *SmsProcess) SendGroupMsg(content string) (err error) {
	var msg common.Message
	msg.Type = "group"
	msg.Data = content

	tf := &common.Transfer{
		Conn: model.CurUser.Conn,
	}
	data, err := json.Marshal(msg)
	tf.WritePkg(data)
	return
}
