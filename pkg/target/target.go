package target

import (
	"fmt"
	"net"
	"resonance/pkg/port"
)

type Target struct {
	Ip   net.IP
	Port port.Port
	//预留一个port接口
	//Port int
}

func (h Target) String() string {
	return fmt.Sprintf("%v:%v", h.Ip, h.Port)
}
