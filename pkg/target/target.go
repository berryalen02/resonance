package target

import (
	"fmt"
	"resonance/pkg/port"
)

type Targets struct {
	Ip    string
	Range string      //存放端口范围
	Port  []port.Port //端口切片利于扫描
	//Port int
	SacnMode int //功能模块参数
}

func (h Targets) String() string {
	return fmt.Sprintf("%v:%v", h.Ip, h.Port)
}
