package common

import (
	"net"
	"fmt"
	"encoding/binary"
	"encoding/json"
)
type Transfer struct {
	Conn net.Conn
	Buf [8096]byte
}
func( this *Transfer) ReadPkg() (msg Message, err error) {
	//读头，长度
	n, err := this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("read package header error")
	}
	if n != 4 {
		print("error")
	}
	fmt.Printf("len %s", this.Buf[:4])
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	//读正真数据
	n2, err := this.Conn.Read(this.Buf[:pkgLen])
	if pkgLen != uint32(n2) || err != nil {
		fmt.Println("read body error")
		return
	}
	json.Unmarshal(this.Buf[:pkgLen], &msg)

	err = nil
	return
}

func(this *Transfer) WritePkg( data []byte) (err error) {
	var bytes [4]byte
	pkgLen := uint32(len(data))
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)
	n, err := this.Conn.Write(bytes[:])

	n, err = this.Conn.Write(data)
	if uint32(n) != pkgLen {
		fmt.Println("server write body error")
		return
	}
	return
}