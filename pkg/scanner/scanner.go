package scanner

import (
	"fmt"
	"net"
	"resonance/pkg/util"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func TCPConnect(ip net.IP, port int) (string, int, error) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ip, port), time.Duration(util.Scanmode.Timeout)*time.Millisecond)

	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()

	return ip.String(), port, err
}

// 基于目标IP设置当前IP和端口
// 思路：向目标IP发UDP包（一来一回即可）
// 连接成功，调用localAddr()返回net.Addr显示源IP和端口
func localIPPort(dstip net.IP) (net.IP, int, error) {
	serverAddr, err := net.ResolveUDPAddr("udp", dstip.String()+":54321")
	if err != nil {
		return nil, 0, err
	}
	if con, err := net.DialUDP("udp", nil, serverAddr); err == nil {
		if udpaddr, ok := con.LocalAddr().(*net.UDPAddr); ok {
			return udpaddr.IP, udpaddr.Port, nil
		}
	}
	return nil, -1, err
}

func SynScan(dstIp net.IP, Port int) (string, int, error) {
	srcIp, srcPort, err := localIPPort(dstIp)
	if err != nil {
		panic(err)
	} else {
		_ = err
	}

	dstIp = dstIp.To4()

	dstport := layers.TCPPort(Port)
	srcport := layers.TCPPort(srcPort)

	ip := &layers.IPv4{
		SrcIP:    srcIp,
		DstIP:    dstIp,
		Protocol: layers.IPProtocolTCP,
	}
	//设置ip帧
	tcp := &layers.TCP{
		SrcPort: srcport,
		DstPort: dstport,
		SYN:     true,
	}
	//设置tcp帧，SYN为1
	err = tcp.SetNetworkLayerForChecksum(ip)

	if err != nil {
		panic(err)
	} else {
		_ = err
	}

	buf := gopacket.NewSerializeBuffer()
	//gopacket发送数据包准备工作，必要的创建缓冲区存放数据流

	opts := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}
	//两个校验机制
	//可以不设置为true，不赋值默认为false

	if err := gopacket.SerializeLayers(buf, opts, tcp); err != nil {
		return dstIp.String(), 0, err
	}
	//序列化layers，把tcp层写入NewSerializeBuffer()创建的缓冲区buf

	conn, err := net.ListenPacket("ip4:tcp", "0.0.0.0")
	if err != nil {
		return dstIp.String(), 0, err
	}
	defer conn.Close()

	if _, err := conn.WriteTo(buf.Bytes(), &net.IPAddr{IP: dstIp}); err != nil {
		return dstIp.String(), 0, err
	}
	// WriteTo writes a packet with payload p to addr.
	// WriteTo can be made to time out and return an Error after a
	// fixed time limit; see SetDeadline and SetWriteDeadline.
	// On packet-oriented connections, write timeouts are rare.

	// 设置连接时间
	if err := conn.SetDeadline(time.Now().Add(time.Duration(util.Scanmode.Timeout) * time.Millisecond)); err != nil {
		return dstIp.String(), 0, err
	}

	// 以上都是前置准备
	// 下面才是扫描的重头戏

	for {
		b := make([]byte, 4096)
		n, addr, err := conn.ReadFrom(b)
		if err != nil {
			return dstIp.String(), 0, err
		} else if addr.String() == dstIp.String() {
			// Decode a packet
			packet := gopacket.NewPacket(b[:n], layers.LayerTypeTCP, gopacket.Default)
			// Get the TCP layer from this packet
			// 将TCP层数据导入b切片中

			if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
				tcp, _ := tcpLayer.(*layers.TCP)

				if tcp.DstPort == srcport {
					//响应包目的端口==发送包源端口
					//确认是对应包
					if tcp.SYN && tcp.ACK {
						// log.Printf("%v:%d is OPEN\n", dstIp, dstport)
						return dstIp.String(), Port, err
					} else {
						return dstIp.String(), 0, err
					}
				}
			}
		}
	}
}
