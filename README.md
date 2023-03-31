# resonance
一款不断成长的综合扫描器。A growing comprehensive scanner.

# 介绍
resonance扫描器的定位，是一个轻量级、代理池实时更新、高效并发的综合扫描器。
希望具有的功能：
- 实时更新代理池
- 可导入字典、目标列表
- 弱口令爆破
- 目录扫描
- 端口扫描
- 资产探测
- 系统安全信息扫描检测
- 后门程序扫描

----
# 功能
目前实现了并发端口扫描，支持nmap格式，多目标输入（逗号间隔）

```GO
NAME:
   resonance - start a scan server

USAGE:
   resonance [global options] command [command options] [arguments...]

VERSION:
   2023.3.28

DESCRIPTION:
   A growing comprehensive scanner.

AUTHOR:
   oink <wx11211@hotmail.com>

COMMANDS:
   portscan  start a port scan
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

