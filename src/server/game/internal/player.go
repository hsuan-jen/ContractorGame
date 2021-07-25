package internal

import (
	"server/base"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/util"
)

type Player struct {
	ObjID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" bson:"_id"`
	agent gate.Agent
	// 帐号id
	AccountID int64 `protobuf:"varint,3,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
	//等级
	Level int32 `protobuf:"varint,4,opt,name=Level,proto3" json:"Level,omitempty"`
	//头像url
	AvatarURL string `protobuf:"bytes,5,opt,name=AvatarURL,proto3" json:"AvatarURL,omitempty"`
	// 名字
	NickName string `protobuf:"bytes,6,opt,name=NickName,proto3" json:"NickName,omitempty"`
	// 性别
	Sex string `protobuf:"bytes,7,opt,name=Sex,proto3" json:"Sex,omitempty"`
	// 创建时间
	CreateTime int64 `protobuf:"varint,8,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	// 上次登录时间
	LastLoginTime int64 `protobuf:"varint,9,opt,name=LastLoginTime,proto3" json:"LastLoginTime,omitempty"`
	// 上次登出时间
	LastLogoutTime int64 `protobuf:"varint,10,opt,name=LastLogoutTime,proto3" json:"LastLogoutTime,omitempty"`
	//经验
	Exp int64 `protobuf:"varint,11,opt,name=Exp,proto3" json:"Exp,omitempty"`
}

func CreatePlayer() *Player {
	player := new(Player)
	//player.IntAttr = make([]int, INTATTR_MAX)
	//player.StrAttr = make([]string, STRATTR_MAX)
	return player
}

/* func (self *Player) GetIntAttr(index int) int {
	if index < 0 || index >= INTATTR_MAX {
		return 0
	}
	return self.IntAttr[index]
}
*/
/* func (self *Player) SetIntAttr(index, val int) {
	if index < 0 || index >= INTATTR_MAX {
		return
	}

	self.IntAttr[index] = val
} */

/* func (self *Player) GetStrAttr(index int) string {
	if index < 0 || index >= STRATTR_MAX {
		return ""
	}
	return self.StrAttr[index]
} */

/* func (self *Player) SetStrAttr(index int, val string) {
	if index < 0 || index >= STRATTR_MAX {
		return
	}

	self.StrAttr[index] = val
} */

func (self *Player) InitData(ObjID string, info *msg.C2S_Login) {
	self.ObjID = ObjID
	self.Sex = info.Sex
	self.AvatarURL = info.AvatarURL
	self.NickName = info.NickName
	nextSeq, _ := uidbuilder.GenerateUID("PlayerSet")
	self.AccountID = nextSeq

}

func (self *Player) CallClientFunc(ret int, cmd string, ans interface{}) {
	//errmsg := ""

	//message := &msg.RetMsg{ret, errmsg, cmd, ans}
	//self.agent.WriteMsg(message)
}

/* func (self *Player) Test(req *msg.TestReq) {

} */

//登陆
func (self *Player) OnLogin() {
}

func (self *Player) OnLogout() {
	self.Save()

	TimerManager.RmvAllTimer(self)
}

//保存玩家数据
func (self *Player) Save() {
	mgodb.Set(
		base.DBTask{
			ObjID:      self.ObjID,
			DB:         base.DBNAME,
			Collection: base.PLAYERSET,
			Key:        "_id",
			KeyV:       self.ObjID,
			Ret:        util.DeepClone(self),
			Cb: func(param interface{}, err error) {
				if nil != err {
					log.Error("save playerdata failed:", self.AccountID)
				}
			},
		},
	)
}

//同步保存玩家数据
func (self *Player) SaveSync() {
	if nil != mgodb.SetSync(base.DBNAME, base.PLAYERSET, "AccountID", self.AccountID, self) {
		log.Error("save playerdata failed:", self.AccountID)
	}
}

func (self *Player) OnAfterLogin() {

	send := &msg.S2C_LoginInfo{
		ID:        self.AccountID,
		NickName:  self.NickName,
		AvatarURL: self.AvatarURL,
	}
	self.agent.WriteMsg(send)
}
