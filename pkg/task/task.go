package task

import (
	"fmt"
	"net"
	"resonance/pkg/scanner"
	"resonance/pkg/util"
	"strings"
	"sync"
	"time"
)

type Counter struct {
	mu      sync.Mutex
	rwmu    sync.RWMutex
	counter int
}

type ScansTime struct {
	//mu sync.Mutex
	mu        sync.RWMutex
	timeTotal time.Duration
	scansNum  int
}

// 留作之后新增系统指纹等功能使用
// IP对应端口扫描时间，由于设计，不好添加
// type PortScanResult struct {
// 	IP   string
// 	Port []int
// }

var ScanTime ScansTime

func GenerateTask(ips []net.IP, ports []int) ([]map[string]int, int) {
	tasks := make([]map[string]int, 0)
	for _, ip := range ips {
		for _, port := range ports {
			task := map[string]int{ip.String(): port}
			tasks = append(tasks, task)
		}
	}
	return tasks, len(tasks)
}

func RunTask(tasks []map[string]int) {
	wg := &sync.WaitGroup{}

	taskChan := make(chan map[string]int, util.Scanmode.Concurrency)

	for i := 0; i < util.Scanmode.Concurrency; i++ {
		go Scan(taskChan, wg)
	}

	// 不断地往taskChan channel发送数据，直到channel阻塞
	for _, task := range tasks {
		wg.Add(1)
		taskChan <- task
	}

	close(taskChan)
	wg.Wait()
}

func Scan(taskChan chan map[string]int, wg *sync.WaitGroup) {
	// 任务调度配套扫描接口，插件式方法调用，可以实现后续的目录扫描，指纹识别等工作
	// 每个协程都从channel中读取数据后开始扫描并入库
	// 预设计数器，错误超过设定次数放弃当前IP
	var errcounter Counter
	// var ScansTime ScansTime
	for task := range taskChan {
		for ip, port := range task {
			func() {
				// 使用defer语句来确保wg.Done()在函数返回之前被调用
				defer wg.Done()

				// // 创建一个超时上下文
				// timeoutCtx, cancel := context.WithTimeout(ctx, 4*time.Second)
				// defer cancel()
				startTime := time.Now()
				var err error

				if util.Scanmode.ScanMode == 0 {
					err = PortScanMenu(ip, port, &startTime)
				}

				//报错、超时控制
				//但是这里写的是配合端口扫描的，不确定以后加功能比如目录扫描什么的会不会难受，封装一个scanner.PostScan函数来进行端口扫描的操作

				if err != nil {
					// 在修改errnumber之前获取互斥锁
					errcounter.mu.Lock()
					errcounter.counter++
					// 在修改errnumber之后释放互斥锁
					errcounter.mu.Unlock()
				}

				// 使用读写锁来保护errcounter.counter变量
				errcounter.rwmu.RLock()
				counter := errcounter.counter
				errcounter.rwmu.RUnlock()
				if counter > 100 {
					return
				}
				//fmt.Println("Current error counter:", counter)
			}()
		}
	}
}

func SavePortScanResult(ip string, port int, err error) error {
	// fmt.Printf("ip:%v, port: %v, goruntineNum: %v\n", ip, port, runtime.NumGoroutine())
	if err != nil {
		return err
	}
	// data:=make(map[string]PortScan)

	if port > 0 {
		v, ok := util.Scanmode.Result.Load(ip)
		if ok {
			ports, ok1 := v.([]int)
			// ports, ok1 := v.(PortScan)
			if ok1 {
				ports = append(ports, port)
				util.Scanmode.Result.Store(ip, ports)
			}
		} else {
			ports := make([]int, 0)
			ports = append(ports, port)
			util.Scanmode.Result.Store(ip, ports)
		}
	}

	return err
}

func PrintPortScanResult() {

	util.Scanmode.Result.Range(func(key, value interface{}) bool {
		fmt.Printf("|ip:%v\n", key)
		fmt.Printf("|ports: %v\n|", value)
		fmt.Println(strings.Repeat("-", 57))
		return true
	})
	//fmt.Println("%v", int(ScanTime.timeTotal))
	if ScanTime.scansNum != 0 {
		avetime := ScanTime.timeTotal / time.Duration(ScanTime.scansNum)
		fmt.Printf("└─")
		fmt.Printf("portscan nums:%v\n└─average port scan time:%v\n\n", ScanTime.scansNum, avetime)
	}
}

func PortScan() error {
	fmt.Printf("start to scan ports....\n")
	fmt.Printf("┌──(resonance)-[portscan]\n|")
	fmt.Println(strings.Repeat("-", 57))
	ips, err := util.GetIpList(util.Scanmode.Targets.Ip)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	ports, err := util.GetPorts(util.Scanmode.Targets.Range)
	//	fmt.Println("2") 测试
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	tasks, n := GenerateTask(ips, ports)
	_ = n
	RunTask(tasks)
	PrintPortScanResult()
	return err
}

// Portscan路由，判断SYN还是TCP，配合超时控制机制
func PortScanMenu(ip string, port int, start *time.Time) error {
	var err error
	if strings.ToLower(util.Portscanconfig.ScanMode) == "syn" {
		err = SavePortScanResult(scanner.SynScan(net.ParseIP(ip), port))
	} else {
		err = SavePortScanResult(scanner.TCPConnect(net.ParseIP(ip), port))
	}
	elapsedTime := time.Since(*start)
	//fmt.Printf("elapsedTime:%v,timeout:%v\n", elapsedTime, util.Scanmode.Timeout)
	ScanTime.mu.RLock()
	if float64(elapsedTime) <= 1.5*float64(time.Duration(util.Scanmode.Timeout)*time.Millisecond) {
		ScanTime.scansNum++
		ScanTime.timeTotal += elapsedTime
	}

	ScanTime.mu.RUnlock()
	return err
}
