package process

type UserMgr struct {
	//onLineUser [] *UserProcess
	onLineUsers map[int]*UserProcess
}
var userMgr *UserMgr
func init(){
 userMgr =&UserMgr{
		onLineUsers: make(map[int]*UserProcess,1024),
	}
}
func(this *UserMgr)AddOnlineUser(up *UserProcess) {
	//this.onLineUser = append(this.onLineUser, up)
	this.onLineUsers[up.UserId]=up
}

func (this *UserMgr) NotifyOtherOnlineUser(userId int)  {
	for id ,_ :=range this.onLineUsers {
		if(userId==id) {
			continue
		}
		//up.
	}


	return
}
func (this *UserMgr) NotifyOneOnline(userId int) {
	for id, _ := range this.onLineUsers {
		if (userId == id) {
			continue
		}
	}
}