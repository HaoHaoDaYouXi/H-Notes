# redis安装

## yum安装

- 源依赖于epel源，因此需要先安装epel源

  ```
  yum -y install epel-release
  ```
  
- 安装redis

  ```
  yum -y install redis
  ```

- 启动redis服务

  ```
  systemctl start redis
  ```

- Redis常见命令

  ```
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