package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"

	"resonance/pkg/protocol"
	"resonance/pkg/target"
)

type Scanserver struct {
	Target      target.Target
	Mode        protocol.Protocol
	Timeout     int
	Concurrency int
	Result      *sync.Map
}

var Scanmode Scanserver

func TCPConnect(ip net.IP, port int) (string, int, error) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ip, port), time.Duration(Scanmode.Timeout)*time.Second)

	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()

	return ip.String(), port, err
}

func SynScan() {

}

func init() {
	Scanmode = Scanserver{
		Mode:        0,
		Timeout:     1,
		Concurrency: 1000,
		Result:      &sync.Map{},
	}
}
