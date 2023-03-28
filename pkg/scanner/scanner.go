package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"

	"resonance/pkg/host"
	"resonance/pkg/protocol"
)

type Scanserver struct {
	Host        host.Host
	Mode        protocol.Protocol
	Timeout     int
	Concurrency int
	Result      *sync.Map
}

var Scan Scanserver

func TCPConnect(ip net.IP, port int) (net.Conn, error) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ip.String(), port), 1*time.Second)
	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()
	return conn, err
}

func init() {
	Scan = Scanserver{
		Mode:        0,
		Timeout:     1,
		Concurrency: 1000,
		Result:      &sync.Map{},
	}

}
