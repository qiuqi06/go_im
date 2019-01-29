package common

const (
	LoginMsgType="login"
	LoginResMsgType="res"
	RegisterMsgType="register"
)

type Message struct{
	Type string `json:"type"`
	Data string	`json:"data"`
}
type LoginMsg struct {
	UserId int `json:"userId"`
	UserName string	`json:"userName"`
	UserPwd string `json:"userPwd"`
}
type LoginResMsg struct {
	Code int
	Error string
}
