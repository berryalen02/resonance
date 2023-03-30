package protocol

type Protocol int

const (
	TCP Protocol = iota
	UDP
	ARP
)

func (p Protocol) String() string {
	switch p {
	case TCP:
		return "TCP"
	case UDP:
		return "UDP"
	case ARP:
		return "ARP"
	default:
		panic("unknow protocol")
	}
}

// func (p Protocol) int() int {
// 	switch p {
// 	case TCP:
// 		return int(TCP)
// 	case UDP:
// 		return int(UDP)
// 	case ARP:
// 		return int(ARP)
// 	default:
// 		panic("unknow protocol")
// 	}
// }
