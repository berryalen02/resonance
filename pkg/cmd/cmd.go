package cmd

import (
	"resonance/pkg/util"

	"github.com/urfave/cli/v2"
)

// func GetIpList(ips string) ([]net.IP, error) {
// 	iplist, err := iprange.ParseList(ips)
// 	if err != nil {
// 		return nil, err
// 	}
// 	list := iplist.Expand()
// 	return list, err
// }

// func GetPorts(portslist string) ([]int, error) {
// 	ports := make([]int, 0)
// 	if portslist == "" {
// 		return ports, nil
// 	}
// 	ranges := strings.Split(portslist, ",")
// 	for _, r := range ranges {
// 		r = strings.TrimSpace(r)
// 		if strings.Contains(r, "-") {
// 			parts := strings.Split(r, "-")
// 			if len(parts) != 2 {
// 				return nil, fmt.Errorf("error ports arguments: '%s'", r)
// 			}

// 			p1, err := strconv.Atoi(parts[0])
// 			if err != nil {
// 				return nil, fmt.Errorf("error port number: '%s'", r)
// 			}

// 			p2, err := strconv.Atoi(parts[0])
// 			if err != nil {
// 				return nil, fmt.Errorf("error port number: '%s'", r)
// 			}

// 			if p1 > p2 {
// 				return nil, fmt.Errorf("error port : '%s'", r)
// 			}

// 			for i := p1; i <= p2; i++ {
// 				ports = append(ports, i)
// 			}
// 		} else {
// 			port, err := strconv.Atoi(r)
// 			if err != nil {
// 				ports = append(ports, port)
// 			}
// 		}
// 	}
// 	return ports, nil
// }

var Scan = cli.Command{
	Name:        "Scan",
	Usage:       "start a port scan",
	Description: "start a port scan",
	Action:      util.Scan,
	Flags: []cli.Flag{
		stringFlag("iplist, i", "", "ip list"),
		stringFlag("port, p", "", "port"),
		stringFlag("mode, m", "", "scan mode"),
		intFlag("timeout,t", 3, "timeout"),
		intFlag("concurrency, c", 1000, "concurrency"),
	},
}

func stringFlag(name, value, usage string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func intFlag(name string, value int, usage string) *cli.IntFlag {
	return &cli.IntFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func boolFlag(name string, value bool, usage string) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}
