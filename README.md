# resonance
一款不断成长的综合扫描器。A growing comprehensive scanner.

# 介绍
resonance扫描器的定位，是一个轻量级、代理池实时更新、高效并发的综合扫描器。
将会具有的功能：
- 实时更新代理池
- 可导入字典、目标列表
- 弱口令爆破
- 目录扫描
- 资产探测
- 后门程序扫描（webshell探测）
- 漏洞扫描（或许）

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
   v0.2

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

# 更新日志
0.1 IPV4全端口扫描实现
0.2 
- 加端口扫描模式、等级，可以控制超时和协程数量；
- 优化并发超时机制，检测扫描平均时间；
- 优化程序结构；
- 美化UI




