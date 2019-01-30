package process

import (
	"common"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) Process() (err error) {
	defer this.Conn.Close()
	for {
		transfer := common.Transfer{
			Conn: this.Conn,
		}
		msg, err := transfer.ReadPkg()
		if err != nil {
			if err == io.EOF {
				println("对方关闭连接")
				return
			} else {
				println("readPackage err=", err)
				return
			}
		}
		userProcessor := UserProcess{
			Conn: this.Conn,
		}
		go userProcessor.ServerProcessMsg(&msg)
	}
}
