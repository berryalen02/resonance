package scanner

import (
	"fmt"
	"net"
	"time"
)

func TCPConnect(ip net.IP, port int) (net.Conn, error) {
	conn, err := net.DialTimeout("tcp", fmt.Sprint("%v:%v", ip.String(), port), 1*time.Second)
	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()
	return conn, err
}
