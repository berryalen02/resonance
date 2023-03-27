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
