package util

import (
	"fmt"
	"net"
	"os"
	"resonance/pkg/scanner"
	"strconv"
	"strings"

	"github.com/malfunkt/iprange"
	"github.com/urfave/cli/v2"
)

func IsRoot() bool {
	return os.Geteuid() == 0
}

func CheckRoot() {
	if !IsRoot() {
		fmt.Println("must run with root")
		os.Exit(0)
		//如果不是root就退出程序
	}
}

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

func Scan(cli *cli.Context) error {
	if cli.IsSet("iplist") {
		scanner.Scan.Host.Ip = net.ParseIP(cli.String("iplist"))
	} else {
		return fmt.Errorf("Invalid input parameter: %s", cli.IsSet("iplist"))
	}

	if cli.IsSet("port") {
		var err error
		scanner.Scan.Host.Port.Port, err = strconv.Atoi(cli.String("port"))
	} else {
		return fmt.Errorf("Invalid input parameter: %s", cli.String("port"))
	}

	if cli.IsSet("mode") {
		// var err error
		scanner.Scan.Mode = cli.String("mode")
	} else {
		return fmt.Errorf("Invalid input parameter: %s", cli.String("mode"))
	}

	if cli.IsSet("timeout") {
		scanner.Scan.Timeout = cli.Int("timeout")
	} else {
		return fmt.Errorf("Invalid input parameter: %s", cli.Int("timeout"))
	}

	if cli.IsSet("concurrency") {
		scanner.Scan.Concurrency = cli.Int("concurrency")
	} else {
		return fmt.Errorf("Invalid input parameter: %s", cli.Int("concurrency"))
	}

	// if strings.ToLower(scanner.Scan.Mode) == "" {
	// 	CheckRoot()
	// }
	// 这里留个判断扫描模式对不对，用白名单
	// 如果不对的话，就不能用root去运行

	ips, err := GetIpList(scanner.Host)
	ports, err := GetPorts(scanner.Port)
	tasks, n := scanner.GenerateTask(ips, ports)
	_ = n
	scanner.RunTask(tasks)
	scanner.PrintResult()
	return err
}
