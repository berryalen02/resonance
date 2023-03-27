package shell

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/malfunkt/iprange"
)

func GetIpList(ips string) ([]net.IP, error) {
	iplist, err := iprange.ParseList(ips)
	if err != nil {
		return nil, err
	}
	list := iplist.Expand()
	return list, err
}

func GetPorts(portslist string) ([]int, error) {
	ports := make([]int, 0)
	if portslist == "" {
		return ports, nil
	}
	ranges := strings.Split(portslist, ",")
	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if strings.Contains(r, "-") {
			parts := strings.Split(r, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("error ports arguments: '%s'", r)
			}

			p1, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, fmt.Errorf("error port number: '%s'", r)
			}

			p2, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, fmt.Errorf("error port number: '%s'", r)
			}

			if p1 > p2 {
				return nil, fmt.Errorf("error port : '%s'", r)
			}

			for i := p1; i <= p2; i++ {
				ports = append(ports, i)
			}
		} else {
			port, err := strconv.Atoi(r)
			if err != nil {
				ports = append(ports, port)
			}
		}
	}
	return ports, nil
}
