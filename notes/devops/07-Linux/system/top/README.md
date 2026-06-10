### 操作
1-top -l 1 -s 0 | grep PhysMem 查看物理内存的使用
2- cat  /proc/cpuinfo ，查看cpu 逆序
3.显示进程参数
4. top -c
5. shift + p/m
6.高亮 排序列, 按x键
7.高亮 变化进程行, 按b键
8.显示cpu每核的运行状态,按1键
9.退出, 按q键


### 官方文件阅读


### 参考
top : https://www.booleanworld.com/guide-linux-top-command/

1. https://apple.stackexchange.com/questions/4286/is-there-a-mac-os-x-terminal-version-of-the-free-command-in-linux-systems
2. top 命令：https://www.cnblogs.com/peida/archive/2012/12/24/2831353.html
3. 如何使用使用 top 命令：https://diabloneo.github.io/2019/08/29/How-to-use-top-command/
4. 30 个实例 https://blog.csdn.net/zhusongziye/article/details/80356989
5. 指标参考 https://zshell.cc/2017/03/16/linux-process--top_%E5%91%BD%E4%BB%A4%E4%B8%8E%E6%93%8D%E4%BD%9C%E5%8F%8A%E5%85%B6%E5%8F%82%E6%95%B0%E6%8C%87%E6%A0%87%E6%80%BB%E7%BB%93/
6. http://www.4e00.com/blog/linux/2018/06/06/linux-commands-top.html
7. https://learnku.com/articles/30384 命令详细解析

### TOP 命令数据结构

* PID     = Process Id                                                                                                                                                                │
* USER    = Effective User Name                                                                                                                                                       │
* PR      = Priority                                                                                                                                                                  │
* NI      = Nice Value                                                                                                                                                                │
* VIRT    = Virtual Image (KiB)                                                                                                                                                       │
* RES     = Resident Size (KiB)                                                                                                                                                       │
* SHR     = Shared Memory (KiB)                                                                                                                                                       │
* S       = Process Status                                                                                                                                                            │
* %CPU    = CPU Usage                                                                                                                                                                 │
* %MEM    = Memory Usage (RES)                                                                                                                                                        │
* TIME+   = CPU Time, hundredths                                                                                                                                                      │
* COMMAND = Command Name/Line                                                                                                                                                         │
  PPID    = Parent Process pid                                                                                                                                                        │
  UID     = Effective User Id                                                                                                                                                         │
  RUID    = Real User Id                                                                                                                                                              │
  RUSER   = Real User Name                                                                                                                                                            │
  SUID    = Saved User Id                                                                                                                                                             │
  SUSER   = Saved User Name                                                                                                                                                           │
  GID     = Group Id                                                                                                                                                                  │
  GROUP   = Group Name                                                                                                                                                                │
  PGRP    = Process Group Id                                                                                                                                                          │
  TTY     = Controlling Tty                                                                                                                                                           │
  TPGID   = Tty Process Grp Id                                                                                                                                                        │
  SID     = Session Id                                                                                                                                                                │
  nTH     = Number of Threads                                                                                                                                                         │
  P       = Last Used Cpu (SMP)                                                                                                                                                       │
  TIME    = CPU Time                                                                                                                                                                  │
  SWAP    = Swapped Size (KiB)                                                                                                                                                        │
  CODE    = Code Size (KiB)                                                                                                                                                           │
  DATA    = Data+Stack (KiB)                                                                                                                                                          │
  nMaj    = Major Page Faults                                                                                                                                                         │
  nMin    = Minor Page Faults                                                                                                                                                         │
  nDRT    = Dirty Pages Count                                                                                                                                                         │
  WCHAN   = Sleeping in Function                                                                                                                                                      │
  Flags   = Task Flags <sched.h>                                                                                                                                                      │
  CGROUPS = Control Groups                                                                                                                                                            │
  SUPGIDS = Supp Groups IDs                                                                                                                                                           │
  SUPGRPS = Supp Groups Names                                                                                                                                                         │
  TGID    = Thread Group Id                                                                                                                                                           │
  ENVIRON = Environment vars                                                                                                                                                          │
  vMj     = Major Faults delta                                                                                                                                                        │
  vMn     = Minor Faults delta                                                                                                                                                        │
  USED    = Res+Swap Size (KiB)                                                                                                                                                       │
  nsIPC   = IPC namespace Inode                                                                                                                                                       │
  nsMNT   = MNT namespace Inode                                                                                                                                                       │
  nsNET   = NET namespace Inode                                                                                                                                                       │
  nsPID   = PID namespace Inode                                                                                                                                                       │
  nsUSER  = USER namespace Inode                                                                                                                                                      │
  nsUTS   = UTS namespace Inode