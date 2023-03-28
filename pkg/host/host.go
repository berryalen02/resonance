package host

import (
	"fmt"
	"net"
	"resonance/pkg/port"
)

type Host struct {
	Ip   net.IP
	Port port.Port
	//预留一个port接口
	//Port int
}

func (h Host) String() string {
	return fmt.Sprintf("%v:%v", h.Ip, h.Port)
}
