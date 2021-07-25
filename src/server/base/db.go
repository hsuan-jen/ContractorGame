package base

import "gopkg.in/mgo.v2/bson"

const (
	DBNAME = "ContractorGame"

	ACCOUNTSET   = "AccountSet" //账号集合
	PLAYERSET    = "PlayerSet"  //玩家集合
	INCREMENTSET = "Increment"  //增量集合
)

const (
	UserIdTimes   = 1000  //真人用户Id倍数
	UserIdOffset  = 0     //真人用户Id偏移,也可用服务器id
	RobotIdTimes  = 10000 //机器人用户Id倍数
	RobotIdOffset = 1     //机器人用户Id偏移
)

// 定义库表的递增序列
var Seqs = []string{
	PLAYERSET,
}

type DBTask struct {
	ObjID      string
	DB         string
	Collection string
	Key        string
	KeyV       interface{}
	Ret        interface{}
	Cb         func(interface{}, error)
}

type DBSearch struct {
	ObjID      string
	DB         string
	Collection string
	M          bson.M
	Limit      int
	Skip       int
	Ret        interface{}
	Cb         func(interface{}, error)
}

//AccountInfo 账号信息
type AccountInfo struct {
	ObjID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" bson:"_id"`
	// 客户端传来的唯一id,如微信unionID/steamID
	UID string `protobuf:"bytes,2,opt,name=UID,proto3" json:"UID,omitempty"`
	// 密码
	Password string `protobuf:"bytes,9,opt,name=password,proto3" json:"password,omitempty"`
	// 注册时间
	RegisterTime int64 `protobuf:"varint,5,opt,name=RegisterTime,proto3" json:"RegisterTime,omitempty"`
	//登陆地理位置
	Location string `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	// 封号标记,1为被封
	Ban int32 `protobuf:"varint,12,opt,name=Ban,proto3" json:"Ban,omitempty"`
}

//BsonObjectID 转换一个Bson
func BsonObjectID(s string) bson.ObjectId {
	if s == "" {
		return bson.NewObjectId()
	}

	if bson.IsObjectIdHex(s) {
		return bson.ObjectIdHex(s)
	}

	return bson.ObjectId(s)
}
