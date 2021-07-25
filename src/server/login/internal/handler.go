package internal

import (
	"fmt"
	"reflect"
	"server/base"
	"server/game"
	"server/msg"
	"time"

	"github.com/name5566/leaf/gate"
	"gopkg.in/mgo.v2/bson"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.C2S_Login{}, handleLoginAuth)
}

func handleLoginAuth(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.C2S_Login)
	// 消息的发送者
	sender := args[1].(gate.Agent)

	send := &msg.S2C_Login{}
	mgodb.Get(
		base.DBTask{
			ObjID:      recv.UID,
			DB:         base.DBNAME,
			Collection: base.ACCOUNTSET,
			Key:        "uid",
			KeyV:       recv.UID,
			Ret:        &base.AccountInfo{},
			Cb: func(param interface{}, err error) {
				info := param.(*base.AccountInfo)

				if info.UID == "" {
					info.UID = recv.UID
					info.Password = recv.Password
					info.RegisterTime = time.Now().Unix()
					info.Location = recv.Location
					info.Ban = 0
					info.ObjID = bson.NewObjectId().Hex()

					//创建账号
					mgodb.Set(base.DBTask{ObjID: info.UID, DB: base.DBNAME, Collection: base.ACCOUNTSET, Key: "uid", KeyV: recv.UID, Ret: info, Cb: nil})
				}

				defer func() {
					sender.WriteMsg(send)
				}()

				if info.Password != recv.Password {
					fmt.Errorf("login uid:[%s] with password:[%s] not match", recv.UID, recv.Password)
					send.Reason = msg.S2C_Login_E_Err_LoginInfoNotMatch
					return
				}

				sender.SetUserData(info)
				skeleton.AsynCall(game.ChanRPC, "LoginSuccess", sender, recv, func(err error) {
					if nil != err {
						fmt.Errorf("login uid:[%s] failed:[%s]", info.UID, err.Error())
						send.Reason = msg.S2C_Login_E_Err_Unknown
						return
					}
				})
			},
		},
	)
	return
}
