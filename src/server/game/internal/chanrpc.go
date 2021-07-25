package internal

import (
	"fmt"
	"server/base"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("LoginSuccess", rpcLoginSuccess)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	log.Debug("new agent")
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	log.Debug("close agent")

	userdata := a.UserData()
	if nil == userdata {
		return
	}

	info := userdata.(*base.AccountInfo)
	player := PlayerManager.Get(info.ObjID)
	if nil != player {
		player.OnLogout()
	}
	PlayerManager.DelPlayer(info.ObjID)
}

func rpcLoginSuccess(args []interface{}) {
	a := args[0].(gate.Agent)
	recv := args[1].(*msg.C2S_Login)
	userdata := a.UserData()
	info := userdata.(*base.AccountInfo)

	//判断玩家重复登陆
	player := PlayerManager.Get(info.ObjID)
	if nil != player {
		player.agent.Close()
		player.agent = a
		return
	}

	mgodb.Get(
		base.DBTask{
			ObjID:      info.ObjID,
			DB:         base.DBNAME,
			Collection: base.PLAYERSET,
			Key:        "_id",
			KeyV:       info.ObjID,
			Ret:        CreatePlayer(),
			Cb: func(param interface{}, err error) {
				player := param.(*Player)
				player.agent = a

				if player.AccountID == 0 { //保存新玩家数据
					player.InitData(info.ObjID, recv)
					player.Save()
				}

				player.OnLogin()
				PlayerManager.AddPlayer(player)
				fmt.Println(player)
				//下发用户数据
				player.OnAfterLogin()
			},
		},
	)
}
