package port

import (
	"fmt"

	"resonance/pkg/protocol"
)

type Port struct {
	Port     int
	Protocol protocol.Protocol
	TLS      bool
}

func (p Port) String() string {
	return fmt.Sprintf("%d-%s-%v", p.Port, p.Protocol.String(), p.TLS)
}

func (p Port) int() int {
	return p.Port
}

// func (p Port) StringPortonly() string {
// 	return fmt.Sprintf("%d-%d-%v", p.Port, p.Protocol, p.TLS)
// }
