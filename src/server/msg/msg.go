package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

// Processor 使用 Protobuf 消息处理器
var Processor = protobuf.NewProcessor()

func init() {

	Processor.Register(&C2S_Ping{})
	Processor.Register(&S2C_Pong{})
	Processor.Register(&C2S_Login{})
	Processor.Register(&S2C_Login{})
	Processor.Register(&S2C_LoginInfo{})

}
