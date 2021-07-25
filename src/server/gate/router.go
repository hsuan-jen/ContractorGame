package gate

import (
	"server/game"
	"server/login"
	"server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.C2S_Login{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_Ping{}, game.ChanRPC)
}
