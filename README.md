# resonance
一款不断成长的扫描器。Your loyal attack assistant.

# 介绍
resonance扫描器的定位，是一个代理池实时更新、高效并发的攻击面信息收集扫描器。
将会具有的功能：
- 实时更新代理池
- 可导入字典、目标列表
- 攻击面资产收集
- 后门程序扫描（webshell探测）
- 目标社工信息收集

----
# 功能
目前实现了并发端口扫描，支持nmap格式，多目标输入（逗号间隔）

---
# 用法

```GO
  _____
 |  __ \
 | |__) | ___  ___   ___   _ __    __ _  _ __    ___  ___
 |  _  / / _ \/ __| / _ \ | '_ \  / _` || '_ \  / __|/ _ \
 | | \ \|  __/\__ \| (_) || | | || (_| || | | || (__|  __/
 |_|  \_\\___||___/ \___/ |_| |_| \__,_||_| |_| \___|\___|

NAME:
   resonance - Your loyal attack assistant

USAGE:
   resonance [global options] command [command options] [arguments...]

VERSION:
   1.0

DESCRIPTION:
   https://github.com/berryalen02/resonance

AUTHOR:
   berryalen02

COMMANDS:
   portscan  start a port scan
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

扫描目标常见端口
```GO
resonance.exe portscan -i host
```

指定端口扫描
```GO
resonance.exe portscan -i host -p 1,2,3,5-1999
```

设定超时、协程数 扫描
```Go
resonance.exe portscan -i host -t 1 -c 2000
```

设定扫描等级
```GO
resonance.exe portscan -i host -l (0-4)
```

运行效果
```GO
  _____
 |  __ \
 | |__) | ___  ___   ___   _ __    __ _  _ __    ___  ___ 
 |  _  / / _ \/ __| / _ \ | '_ \  / _` || '_ \  / __|/ _ \
 | | \ \|  __/\__ \| (_) || | | || (_| || | | || (__|  __/
 |_|  \_\\___||___/ \___/ |_| |_| \__,_||_| |_| \___|\___|

start to scan ports....
┌──(resonance)-[portscan]
|---------------------------------------------------------
|ip:xx.xx.xx.xx
|ports: [80 22 21 443 888 1224 8080 8065]
|---------------------------------------------------------
└─portscan nums:65226
└─average port scan time:3.00903061s
```

# 更新日志
0.1 IPV4全端口扫描实现
0.2 优化任务调度

1.0
- 加端口扫描模式、等级，可以控制超时和协程数量；
- 优化并发超时机制，检测扫描平均时间；
- 优化程序结构；
- 美化UI；