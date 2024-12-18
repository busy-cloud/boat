package broker

import (
	"bytes"
	"github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

type Hook struct {
	mqtt.HookBase
}

func (h *Hook) ID() string {
	return "broker"
}
func (h *Hook) Provides(b byte) bool {
	//高效吗？
	return bytes.Contains([]byte{
		mqtt.OnConnectAuthenticate,
		mqtt.OnACLCheck,
		mqtt.OnDisconnect,
		mqtt.OnSubscribed,
		mqtt.OnUnsubscribed,
	}, []byte{b})
}

func (h *Hook) OnConnectAuthenticate(cl *mqtt.Client, pk packets.Packet) bool {
	//cl.Net.Listener todo websocket 直接鉴权通过

	//TODO 使用

	return true
}

func (h *Hook) OnACLCheck(cl *mqtt.Client, topic string, write bool) bool {
	//只允许发送属性事件

	return true
}

func (h *Hook) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	//执行unsubscribe

}

func (h *Hook) OnSubscribed(cl *mqtt.Client, pk packets.Packet, reasonCodes []byte) {

}

func (h *Hook) OnUnsubscribed(cl *mqtt.Client, pk packets.Packet) {

}
