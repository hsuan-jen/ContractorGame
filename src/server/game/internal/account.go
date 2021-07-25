package internal

import (
	"fmt"
	"server/base"
	"server/msg"
	"time"

	"github.com/name5566/leaf/gate"
)

func login(recv *msg.C2S_Login, sender *gate.Agent) {
	fmt.Println(sender)
	//send := msg.S2C_Login
	//todo :登录验证
	mgodb.Get(
		base.DBTask{
			ObjID:      recv.UID,
			DB:         base.DBNAME,
			Collection: base.ACCOUNTSET,
			Key:        "UID",
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
					//	info.ID, _ = uidbuilder.GenerateUID("AccountSet")
					//创建账号
					mgodb.Set(base.DBTask{ObjID: info.UID, DB: base.DBNAME, Collection: base.ACCOUNTSET, Key: "UID", KeyV: recv.UID, Ret: info, Cb: nil})
					/* //创建用户
					userInfo := &base.User{}
					userInfo.AccountID = info.ID
					userInfo.AvatarURL = recv.AvatarURL
					userInfo.NickName  = recv.NickName
					userInfo.Sex       = recv.Sex
					userInfo.ObjID	   = uidbuilder.GenerateUID("PlayerSet") + 10000
					userInfo.CreateTime = time.Now().Unix()
					userInfo.LastLoginTime = time.Now().Unix()
					userInfo.Level = 1
					mgodb.set(
						base.DBTask{
							ObjID: userInfo.AccountID ,
							DB: base.DBNAME,
							Collection: base.PLAYERSET,
							Key: ,
						}
					) */

				}

				if info.Password != recv.Password {
					fmt.Errorf("login uid:[%s] with password:[%s] not match", recv.UID, recv.Password)
					//	send.Reason = msg.S2C_Login_E_Err_LoginInfoNotMatch
					return
				}

			},
		},
	)
}
