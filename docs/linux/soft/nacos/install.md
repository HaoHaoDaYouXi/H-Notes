# nacos 安装
需要安装JDK

1、下载并上传[安装包](https://github.com/alibaba/nacos/releases)

下载需要的版本并上传到服务器
~~~
wget https://github.com/alibaba/nacos/releases/download/2.2.0.1/nacos-server-2.2.0.1.tar.gz
~~~

2、解压安装包

（1）解压安装包：
~~~
tar -zxvf nacos-server-2.2.0.1.tar.gz -C /usr/local/nacos
~~~

（2）删除安装包（此步不执行也可以）
~~~
rm -rf nacos-server-2.2.0.1.tar.gz
~~~

3、配置Mysql并启动

（1）创建mysql库并运行sql脚本

（2）更改nacos配置文件
~~~
vim /usr/local/nacos/conf/application.properties
spring.datasource.platform=mysql
db.num=1
db.url.0=jdbc:mysql://192.168.1.201:61307/nacos?characterEncoding=utf8&connectTimeout=1000&socketTimeout=3000&autoReconnect=true&useUnicode=true&useSSL=false&serverTimezone=UTC
db.user.0=nacos
db.password.0=nacos
nacos.core.auth.plugin.nacos.token.secret.key=VGhpc0lzTXlDdXN0b21TZWNyZXRLZXkwMTIzNDU2Nzg=
~~~
（3）后台单机启动(可跳过，直接配置开机自启)
~~~
sudo sh startup.sh -m standalone &
~~~

4、设置开机自启动

（1）创建并编辑nacos.service文件
~~~
vim /lib/systemd/system/nacos.service
~~~
（2）nacos.service文件
~~~
[Unit]
Description=nacos
After=network.target

[Service]
Type=forking
ExecStart=/usr/local/nacos/bin/startup.sh -m standalone
ExecReload=/usr/local/nacos/bin/shutdown.sh
ExecStop=/usr/local/nacos/bin/shutdown.sh
PrivateTmp=true

[Install]
WantedBy=multi-user.target
~~~
四、执行以下命令
~~~
重新加载所有service服务：systemctl daemon-reload
开机启动nacos.service：systemctl enable nacos.service
查看该service是否开机启用：systemctl is-enabled nacos.service
启动该服务：systemctl start nacos.service
查看该服务状态：systemctl status nacos.service

~~~
