package target

import (
	"fmt"
	"net"
	"resonance/pkg/port"
)

type Target struct {
	Ip   net.IP
	Port port.Port
	//Port int
	SacnMode int //功能模块参数
}

func (h Target) String() string {
	return fmt.Sprintf("%v:%v", h.Ip, h.Port)
}
