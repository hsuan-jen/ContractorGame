package internal

import (
	"reflect"
	"server/msg"

	"github.com/name5566/leaf/gate"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.C2S_Ping{}, handlePong)

}
func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlePong(args []interface{}) {
	// 收到的 C2S_Ping 消息
	_ = args[0].(*msg.C2S_Ping)
	// 消息的发送者
	a := args[1].(gate.Agent)
	// 输出收到的消息的内容
	a.WriteMsg(&msg.S2C_Pong{})
}
