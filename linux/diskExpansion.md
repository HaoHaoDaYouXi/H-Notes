# 磁盘扩容

服务器磁盘扩容有很多中情况，主要根据系统、磁盘分区、文件系统等等对应不同的操作

这里简单介绍下`centos7`系统下`LVM`管理分区的情况下，如何创建新分区和扩展分区

## 创建新分区

**使用`fdisk`工具进行分区**

操作步骤
- 确定磁盘名称：`fdisk -l`，确定是哪个磁盘要扩容，假设为：`/dev/sda`
- 输入命令：`fisk /dev/sda`，操作分区，可以根据[说明](#fdisk命令说明)划分分区
- 输入`n`（表示增加分区），回车
- 输入`p`（创建主分区），回车
- 输入`分区数字`（默认会根据已有分区+1，使用默认即可，如：已经有了`sda1`和`sda2`分区，那么输入`3`，默认即`3`），回车
- 输入分区的`start`值，默认即可，回车
- 输入分区的`end`值，（使用默认，或者根据需要分配大小。默认即当前最大值），回车
- 这时可以看到分区信息，输入`p`，新的分区`/dev/sda3`后面`Id`写的是`83`和`Linux`，还需要修改成`8e`，表示`Linux LVM`分区
- 输入`t`，再输入`3`(最新分区号，默认就是3)，最后输入`8e`，回车
- 输入`w`保存，新分区创建好了，**这个时候记得重启，不然后续操作无法识别到这个分区，如果想不重启可以执行`partprobe`，进行分区重读**

分区创建后可以通过`fdisk -l`查看分区情况

### `fdisk`命令说明
```
a 设置可引导标记
b 编辑 bsd 磁盘标签
c 设置 DOS 操作系统兼容标记
d 删除一个分区
l 显示已知的文件系统类型。82 为 Linux swap 分区，83 为 Linux 分区
m 显示帮助菜单
n 新建分区
o 建立空白 DOS 分区表
p 显示分区列表
q 不保存退出
s 新建空白 SUN 磁盘标签
t 改变一个分区的系统 ID
u 改变显示记录单位
v 验证分区表
w 保存退出
x 附加功能（仅专家）
```

## 扩展分区

磁盘创建好后，需要对它格式化创建并扩容到具体的文件系统

**这里介绍下`pv`、`vg`、`lv`**

- `pv`（Phsical Volume：物理卷）：`pv`是`vg`的组成部分，由分区构成，多块盘的时候，可以把一块盘格式化成一个主分区，然后用这个分区做成一个`pv`，只有一块盘的时候，可以这块盘的某一个分区做成一个`pv`，实际上一个`pv`就一个分区。
- `vg`（Volume Group：卷组）：有若干个`pv`组成，作用就是将`pv`组成到以前，然后再重新划分空间。
- `lv`（Logical Volume：逻辑卷）：`lv`就是从`vg`中划分出来的卷，`lv`的使用要比`pv`灵活的多，可以在空间不够的情况下，增加空间。

可以说`pv`是硬盘，`vg`是管理硬盘的操作系统，`lv`是操作系统分出来的各个分区。

套娃逻辑：`pv`->`vg`->`lv`-> 文件系统使用(挂载到某个目录)，硬盘或分区做成`pv`，然后将`pv`或多个`pv`建立`vg`，`vg`上建立`lv`

现在开始扩容
- 创建物理卷`pv`：`pvcreate /dev/sda3`
- 可以使用`pvdisplay`查看物理卷信息
- 查看卷组：`vgdisplay`
- 添加到已有的`vg`中：`vgextend centos /dev/sda3`
- 再次查看卷组：`vgdisplay`，可以看到大小增加了
- 查看当前逻辑卷：`lvdisplay`
- 假设扩容`/dev/centos/root`目录`20G`的大小：`lvextend -L +20G /dev/centos/root`，如果是要全部的扩容，可以`lvextend -l +100%FREE /dev/centos/root`
- 刷新文件系统，不同的文件系统使用不同的命令，很多命令都可以看到文件系统类型，如：`df -T`、`lsblk -f`
  - `xfs`格式：`xfs_growfs /dev/centos/root`
  - `ext4`格式：`resize2fs /dev/centos/root`
- 刷新完成后，可以使用`df -h`查看文件系统大小是否增加了

## `pv`、`vg`、`lv`命令

- `pv`
  - `pvcreate`：根据物理盘,创建`pv`，如：`pvcreate /dev/sdb1`
  - `pvscan`：查询目前系统里的`pv`
  - `pvdisplay`：显示`pv`的状态
  - `pvremove`：将`pv`属性移除，如：`pvremove /dev/sdb1`
  - `pvresize`：刷新

- `vg`
  - `vgcreate`：创建`vg`，如：`vgcreate testvg /dev/sdb /dev/sdc`，创建`testvg`的卷组并将`sdb`和`sdc`物理卷加入其中
  - `vgscan`：查找当前系统里面的`vg`
  - `vgdisplay`：显示当前系统`vg`的状态
  - `vgextend`：给`vg`添加额外的`pv`
  - `vgreduce`：在`vg`内删除`pv`
  - `vgchange`：设置`vg`是否是启动状态
  - `vgremove`：删除一个`vg`，如：`vgremove testvg`

- `lv`
  - `lvcreate`：创建`lv`，如：`lvcreate -l +100%FREE -n testlv testvg`，创建`testlv`，大小为`100%FREE`，属于`testvg`
  - `lvscan`：查询当前系统的`lv`
  - `lvdisplay`：显示`lv`的属性
  - `lvextend`：给`lv`添加容量
  - `lvredurce`：给`lv`减少容量
  - `lvremove`：删除一个`lv`，如：`lvremove /dev/mapper/testvg-data`
  - `lvresize`：对`lv`大小的容量进行调整


----
