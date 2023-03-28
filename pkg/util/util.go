package util

import (
	"fmt"
	"net"
	"os"
	"resonance/pkg/protocol"
	"resonance/pkg/scanner"
	"resonance/pkg/task"
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
	//这里的portslist和定义的接口port对不上
	//期望传入portlist：20-588
	//定义好的端口接口string()：20-TCP-false
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
		scanner.Scanmode.Target.Ip = net.ParseIP(cli.String("iplist"))
	} else {
		return fmt.Errorf("Invalid input parameter: %s", cli.IsSet("iplist"))
	}

	if cli.IsSet("port") {
		var err error
		_ = err
		scanner.Scanmode.Target.Port.Port, err = strconv.Atoi(cli.String("port"))
	} else {
		return fmt.Errorf("Invalid input parameter: %s", cli.String("port"))
	}

	if cli.IsSet("mode") {
		// var err error
		scanner.Scanmode.Mode = protocol.Protocol(cli.Int("mode"))
	} else {
		return fmt.Errorf("Invalid input parameter: %s", cli.String("mode"))
	}

	if cli.IsSet("timeout") {
		scanner.Scanmode.Timeout = cli.Int("timeout")
	} else {
		return fmt.Errorf("Invalid input parameter: %s", cli.Int("timeout"))
	}

	if cli.IsSet("concurrency") {
		scanner.Scanmode.Concurrency = cli.Int("concurrency")
	} else {
		return fmt.Errorf("Invalid input parameter: %s", cli.Int("concurrency"))
	}

	// if strings.ToLower(scanner.Scanmode.Mode) == "" {
	// 	CheckRoot()
	// }
	// 这里留个判断扫描模式对不对，用白名单
	// 如果不对的话，就不能用root去运行

	// 以上判断

	ips, err := GetIpList(string(scanner.Scanmode.Target.Ip))
	ports, err := GetPorts(scanner.Scanmode.Target.Port.String())
	//需要中间加一层解析端口列表string格式
	tasks, n := task.GenerateTask(ips, ports)
	_ = n
	task.RunTask(tasks)
	task.PrintResult()
	return err
}
