

### dd 基本概念


### 文件概念
- 挂载
- 分区
- 镜像 iso 


### 这些信息是什么意思？

dd if=/dev/zero of=loobakfile.img bs=1G count=1


loobakfile.img: Linux rev 1.0 ext4 filesystem data, UUID=1d73f75e-c00b-4eec-b287-3ee65d3b84dc (extents) (large files) (│
huge files)


mke2fs 1.42.13 (17-May-2015)                                                                                           │
Discarding device blocks: done                                                                                         │
Creating filesystem with 262144 4k blocks and 65536 inodes                                                             │
Filesystem UUID: 1d73f75e-c00b-4eec-b287-3ee65d3b84dc                                                                  │
Superblock backups stored on blocks:                                                                                   │
        32768, 98304, 163840, 229376                                                                                   │
                                                                                                                       │
Allocating group tables: done                                                                                          │
Writing inode tables: done                                                                                             │
Creating journal (8192 blocks): done                                                                                   │
Writing superblocks and filesystem accounting information: done

- 快捷挂载方法，无需手动连接任何设备。但是在内部这个环回文件连接到了一个名为 `/dev/loop1` 或 `loop2` 的设备上。