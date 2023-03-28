// package main

// import (
// 	"github.com/urfave/cli/v2"
// )

//	func main() {
//		app:=&cli.App{
//			Name: "sacnner",
//			Authors: ,
//		}
//		Scan:=cli.Command
//	}
package main

import (
	"fmt"
	"net"
)

func main() {
	//srcIp, srcPort, err := localIPPort(net.ParseIP())
	dstAddrs, err := net.LookupIP("192.168.2.231")
	// if err != nil {
	// 	return "192.168.2.231", 0, err
	// }

	dstip := dstAddrs[0].To4()
	fmt.Printf("%v", dstip)
	fmt.Println("Hello, World!")
	_ = err
}
