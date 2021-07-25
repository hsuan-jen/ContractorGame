package internal

import (
	"server/base"
	"server/conf"
	"server/db"

	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	mgodb    = new(db.Mongodb)
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton

	var err error
	mgodb, err = db.Dial(conf.Server.MgodbAddr, conf.Server.LoginMgoConnNum, skeleton)
	if nil == mgodb {
		log.Error("dial mongodb failed:", conf.Server.MgodbAddr, " ", err.Error())
	}

	mgodb.EnsureUniqueIndex(base.DBNAME, base.ACCOUNTSET, []string{"UID"})
}

func (m *Module) OnDestroy() {
	mgodb.Close()
}
