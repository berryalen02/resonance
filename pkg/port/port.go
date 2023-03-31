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

var CommonPort []int

func init() {
	CommonPort = []int{21, 22, 23, 25, 53, 80, 110, 135, 139, 143, 161, 389, 443, 445, 1433, 1521, 3306, 3389}
}
