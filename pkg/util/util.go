package util

import (
	"fmt"
	"net"
	"os"
	"resonance/pkg/port"
	"resonance/pkg/protocol"
	"resonance/pkg/target"
	"strconv"
	"strings"
	"sync"

	"github.com/malfunkt/iprange"
	"github.com/urfave/cli/v2"
)

type ScanConfig struct {
	Targets     target.Targets
	ScanMode    ScanMode
	Timeout     int
	Concurrency int
	Result      *sync.Map
}

var Scanmode ScanConfig

type ScanMode int

const (
	Portscan ScanMode = iota
	Dirscan
)

type PortScanConfig struct {
	ScanMode string //tcp or syn
	Level    int
}

var Portscanconfig PortScanConfig

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

type Flag struct {
	Name    string
	Default interface{}
}

var flags = []Flag{
	{"iplist", ""},
	{"port", ""},
	{"tcp", false},
	{"syn", false},
	{"level", 2},
	{"timeout", 3000},
	{"concurrency", 1000},
}

// 解析域名和nmap格式IPv4地址区间
// 返回IP切片
func GetIpList(s string) ([]net.IP, error) {
	// 同时传入ipv4地址区间和域名
	input := strings.Split(s, ",")
	output := []net.IP{}
	var err error
	for _, s := range input {
		list, err := iprange.ParseList(s)
		if err == nil {
			// 如果成功，遍历范围切片
			for _, r := range list {
				output = append(output, r.Expand()...)
			}
		} else {
			// 如果失败，尝试将字符串解析为域名
			ips, err := net.LookupIP(s)
			if err != nil {
				fmt.Println("error:", err)
				continue
			}
			for _, ip := range ips {
				// 如果是IPv4地址，直接添加到输出切片
				if ipv4 := ip.To4(); ipv4 != nil {
					output = append(output, ipv4)
				}
			}
		}
	}
	return output, err
}

// 将Ports从[]port.Port结构体队列，转换成"20-588"格式字符串
func PortsUnsirializeToString(list []port.Port) string {
	first := list[0].Port
	final := list[len(list)-1].Port
	return fmt.Sprintf("%v-%v", first, final)
}

// 将Ports从[]port.Port结构体队列，转换成[]int队列，保留端口数值
func PortsUnsirializeToIntList(list []port.Port) []int {
	intlist := make([]int, 0)
	for _, r := range list {
		intlist = append(intlist, r.Port)
	}
	return intlist
}

// 将Ports从“20-588”的int队列，转换成Port结构体队列，加上TLS和Protocol
func PortsSerializeFromInt(ports []int, protocol protocol.Protocol, tls bool) []port.Port {
	SerializePorts := make([]port.Port, 0)
	for _, p := range ports {
		SerializePorts = append(SerializePorts, port.Port{Port: p, Protocol: protocol, TLS: tls})
	}
	return SerializePorts
}

// 将Ports从“20-588”的string，转换成Port结构体队列，加上TLS和Protocol
func PortsSerializeFromString(ports string, protocol protocol.Protocol, tls bool) []port.Port {
	SerializePorts := make([]port.Port, 0)

	first, err3 := strconv.ParseInt(ports[:1], 10, 0) // 10表示十进制，0表示自动推断位数
	final, err4 := strconv.ParseInt(ports[2:], 10, 0)
	if err3 != nil || err4 != nil {
		fmt.Printf("invalid ports range")
	}

	for i := first; i < final; i++ {
		SerializePorts = append(SerializePorts, port.Port{Port: int(i), Protocol: protocol, TLS: tls})
	}
	return SerializePorts
}

func GetPorts(portslist string) ([]int, error) {
	//期望传入portlist：20-588
	//定义好的端口接口string()：20-TCP-false
	//定义好了PortsUnsirializeToString()

	//默认为CommonPort，常见端口
	ports := make([]int, 0)
	if portslist == "" {
		return port.CommonPort, nil
	}
	ranges := strings.Split(portslist, ",")
	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if strings.Contains(r, "-") {
			parts := strings.Split(r, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid ports arguments: '%s'", r)
			}

			p1, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, fmt.Errorf("invalid port number: '%s'", r)
			}

			p2, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("invalid port number: '%s'", r)
			}

			if p1 > p2 {
				return nil, fmt.Errorf("invalid port : '%s'", r)
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

func TargetsInit(cli *cli.Context) {
	for _, flag := range flags {
		switch flag.Name {
		case "iplist":
			if cli.IsSet(flag.Name) {
				Scanmode.Targets.Ip = cli.String(flag.Name)
			}
		case "port":
			if cli.IsSet(flag.Name) && cli.IsSet("full") {
				Scanmode.Targets.Range = "1-65535"
			} else if cli.IsSet(flag.Name) {
				Scanmode.Targets.Range = cli.String(flag.Name)
			} else if cli.IsSet("full") {
				Scanmode.Targets.Range = "1-65535"
			}
		case "tcp":
			if cli.IsSet(flag.Name) && cli.IsSet("syn") {
				Portscanconfig.ScanMode = "syn"
			} else if cli.IsSet(flag.Name) {
				Portscanconfig.ScanMode = "tcp"
			} else if cli.IsSet("syn") {
				Portscanconfig.ScanMode = "syn"
			}
		case "syn":
			if cli.IsSet(flag.Name) && cli.IsSet("tcp") {
				Portscanconfig.ScanMode = "syn"
			} else if cli.IsSet(flag.Name) {
				Portscanconfig.ScanMode = "syn"
			} else if cli.IsSet("tcp") {
				Portscanconfig.ScanMode = "tcp"
			}
		case "level":
			if cli.IsSet(flag.Name) {
				Portscanconfig.Level = cli.Int(flag.Name)
				switch Portscanconfig.Level {
				case 0:
					Scanmode.Timeout = 10000
					Scanmode.Concurrency = 200
				case 1:
					Scanmode.Timeout = 5000
					Scanmode.Concurrency = 500
				case 2:
					Scanmode.Timeout = 3000
					Scanmode.Concurrency = 1000
				case 3:
					Scanmode.Timeout = 2000
					Scanmode.Concurrency = 2000
				case 4:
					Scanmode.Timeout = 1500
					Scanmode.Concurrency = 3000
				default:
				}
			}
		case "timeout":
			if cli.IsSet(flag.Name) {
				Scanmode.Timeout = cli.Int(flag.Name)
			} else {
				Scanmode.Timeout = flag.Default.(int)
			}
		case "concurrency":
			if cli.IsSet(flag.Name) {
				Scanmode.Concurrency = cli.Int(flag.Name)
			} else {
				Scanmode.Concurrency = flag.Default.(int)
			}
		}
	}

	if Scanmode.Targets.Ip == "" {
		fmt.Println("请指定要扫描的目标IP或IP列表文件")
		os.Exit(1)
	}
}

func StringToProtocol(s string) protocol.Protocol {
	switch s {
	case "TCP":
		return protocol.TCP
	case "UDP":
		return protocol.UDP
	case "ARP":
		return protocol.ARP
	default:
		panic("unknown protocol")
	}
}

func init() {
	Scanmode = ScanConfig{
		Timeout:     1,
		Concurrency: 1000,
		Result:      &sync.Map{},
	}
	Portscanconfig = PortScanConfig{ScanMode: "tcp"}
}
