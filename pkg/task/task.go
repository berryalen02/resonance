package task

import (
	"fmt"
	"net"
	"resonance/pkg/scanner"
	"strings"
	"sync"
)

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

	taskChan := make(chan map[string]int, scanner.Scanmode.Concurrency)

	for i := 0; i < scanner.Scanmode.Concurrency; i++ {
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
	for task := range taskChan {
		for ip, port := range task {
			if strings.ToLower(scanner.Scanmode.Protocol.String()) == "syn" {
				err := SaveResult(scanner.SynScan(net.ParseIP(ip), port))
				_ = err
			} else {
				err := SaveResult(scanner.TCPConnect(net.ParseIP(ip), port))
				_ = err
			}
			wg.Done()
		}
	}
}

func SaveResult(ip string, port int, err error) error {
	// fmt.Printf("ip:%v, port: %v, goruntineNum: %v\n", ip, port, runtime.NumGoroutine())
	if err != nil {
		return err
	}

	if port > 0 {
		v, ok := scanner.Scanmode.Result.Load(ip)
		if ok {
			ports, ok1 := v.([]int)
			if ok1 {
				ports = append(ports, port)
				scanner.Scanmode.Result.Store(ip, ports)
			}
		} else {
			ports := make([]int, 0)
			ports = append(ports, port)
			scanner.Scanmode.Result.Store(ip, ports)
		}
	}

	return err
}

func PrintResult() {
	scanner.Scanmode.Result.Range(func(key, value interface{}) bool {
		fmt.Printf("ip:%v\n", key)
		fmt.Printf("ports: %v\n", value)
		fmt.Println(strings.Repeat("-", 100))
		return true
	})
}
