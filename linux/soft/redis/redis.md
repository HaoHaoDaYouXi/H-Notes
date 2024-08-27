# redis安装

## yum安装

- 源依赖于epel源，因此需要先安装epel源

  ```shell
  yum -y install epel-release
  ```

- 安装redis

  ```shell
  yum -y install redis
  ```

- 启动redis服务

  ```shell
  systemctl start redis
  ```

- Redis常见命令

  ```shell
  systemctl status redis   查看服务状态
  systemctl stop redis     停止服务
  systemctl restart redis  重启服务
  ps -ef | grep redis      查看reids服务信息
  systemctl enable redis   redis开机启动
  ```

- 设置redis 远程连接和密码，重启后生效

  ```
  vim /etc/redis.config
  注释 #bind 127.0.0.1
  修改protected-mode no
  修改 daemonize yes
  修改 requirepass 123456
  ```
  若出现更改后启动失败可以查看日志文件（`/var/log/redis/redis.log`），若是端口权限问题，可查看`SELinux`策略相关问题</br>
  - `getenforce`：查看`SELinux`状态，`Enforcing`强制模式，`Permissive`宽容模式，`Disabled`为关闭
  - `grubby --update-kernel ALL --args selinux=0;`：关闭，需要重启生效

## 编译安装

- 安装编译环境

  ```shell
  yum -y install gcc gcc-c++ libstdc++-devel
  ```

- 下载最新的稳定版本的redis包

  ```shell
  wget http://download.redis.io/redis-stable.tar.gz
  ```

- 解压，并进入

  ```shell
  tar -zxvf redis-stable.tar.gz && cd redis-stable
  ```

- 编译安装

  ```shell
  make install PREFIX=/usr/local/redis
  ```

- 修改配置文件

  ```
  cp redis.conf /usr/local/redis/bin 
  vim redis.conf
  bind 127.0.0.1 -::1  >>>   bind 0.0.0.0 -::1      # 允许外网访问
  daemonize no         >>>   daemonize yes          # 开启守护线程
  requirepass 12345                                 # 设置密码为123456
  ```
  若出现更改后启动失败可以查看日志文件（`/var/log/redis/redis.log`），若是端口权限问题，可查看`SELinux`策略相关问题</br>
  - `getenforce`：查看`SELinux`状态，`Enforcing`强制模式，`Permissive`宽容模式，`Disabled`为关闭
  - `grubby --update-kernel ALL --args selinux=0;`：关闭，需要重启生效

- 通过systemctl管理redis服务，配置文件

  ```
  vim /lib/systemd/system/redis.service
  Unit]
  Description=redis
  After=network.target remote-fs.target nss-lookup.target
  
  [Service]
  Type=forking
  ExecStart=/usr/local/redis/bin/redis-server /usr/local/redis/bin/redis.conf
  ExecReload=/bin/kill -s HUP $MAINPID
  ExecStop=/usr/local/redis/bin/redis-cli -p 6379 -a 123456 shutdown
  PrivateTmp=true
  
  [Install]
  WantedBy=multi-user.target
  # 重新加载配置文件
  systemctl daemon-reload
  ```

- systemctl管理redis

  ```shell
  # 服务启动
  systemctl start redis    
  # 服务停止
  systemctl stop redis
  # 服务重启
  systemctl restart redis
  # 查看状态
  systemctl status redis
  # 开机启动
  systemctl enable redis
  ```

----
