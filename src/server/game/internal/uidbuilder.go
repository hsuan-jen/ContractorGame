package internal

import (
	"fmt"
	"server/base"

	"gopkg.in/mgo.v2"
)

type UidBuilder struct {
}

func (builder *UidBuilder) Init() {

	for _, seq := range base.Seqs {
		err := mgodb.GetSync(base.DBNAME, base.INCREMENTSET, "_id", seq, builder)
		if mgo.ErrNotFound == err {
			mgodb.EnsureCounter(base.DBNAME, base.INCREMENTSET, seq)
		} else if nil != err {
			fmt.Errorf("ensure counters [%s] error, %s", seq, err.Error())
			return
		}
	}

	/* playercount := int64(mgodb.GetTableCountSync(base.DBNAME, base.PLAYERSET))
	if builder.Seq <= playercount {
		builder.Seq = playercount + 1
		mgodb.SetSync(base.DBNAME, base.INCREMENTSET, "_id", "uid", builder)
	} */
}

//GenerateUID 获取增长id
func (builder *UidBuilder) GenerateUID(id string) (int64, error) {
	return mgodb.IncreSeq(base.DBNAME, base.INCREMENTSET, id, nil)
}
