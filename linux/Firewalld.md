# Firewalld
## 1、centos中firewalld与iptables
centos7以前的版本默认使用iptables服务进行管理防火墙规则。centos7以及其以上版本默认使用firewalld服务管理防火墙。所以在centos8中，就使用其默认的firewalld配置防火墙。

firewalld和iptables的[官方说明](https://access.redhat.com/documentation/zh-cn/red_hat_enterprise_linux/7/html/security_guide/sec-using_firewalls#sec-Comparison_of_Firewalld_to_system-config-firewall_and_iptables)可以参照

## 2、firewalld相关命令
~~~
#进程与状态相关
systemctl start firewalld.service            #启动防火墙
systemctl stop firewalld.service             #停止防火墙
systemctl status firewalld                   #查看防火墙状态
systemctl enable firewalld             #设置防火墙随系统启动
systemctl disable firewalld                #禁止防火墙随系统启动
firewall-cmd --state                         #查看防火墙状态
firewall-cmd --reload                        #更新防火墙规则
firewall-cmd --list-ports                    #查看所有打开的端口
firewall-cmd --list-services                 #查看所有允许的服务
firewall-cmd --get-services                  #获取所有支持的服务

#区域相关
firewall-cmd --list-all-zones                    #查看所有区域信息
firewall-cmd --get-active-zones                  #查看活动区域信息
firewall-cmd --set-default-zone=public           #设置public为默认区域
firewall-cmd --get-default-zone                  #查看默认区域信息


#接口相关
firewall-cmd --zone=public --add-interface=eth0  #将接口eth0加入区域public
firewall-cmd --zone=public --remove-interface=eth0       #从区域public中删除接口eth0
firewall-cmd --zone=default --change-interface=eth0      #修改接口eth0所属区域为default
firewall-cmd --get-zone-of-interface=eth0                #查看接口eth0所属区域

#端口控制
firewall-cmd --query-port=8080/tcp             # 查询端口是否开放
firewall-cmd --add-port=8080/tcp --permanent               #永久添加8080端口例外(全局)
firewall-cmd --remove-port=8800/tcp --permanent            #永久删除8080端口例外(全局)
firewall-cmd --add-port=65001-65010/tcp --permanent      #永久增加65001-65010例外(全局)
firewall-cmd  --zone=public --add-port=8080/tcp --permanent            #永久添加8080端口例外(区域public)
firewall-cmd  --zone=public --remove-port=8080/tcp --permanent         #永久删除8080端口例外(区域public)
firewall-cmd  --zone=public --add-port=65001-65010/tcp --permanent   #永久增加65001-65010例外(区域public)
~~~

## 最常用命令

~~~
firewalld-cmd --zone=public --add-ports=8080/tcp --permanent

命令解析

firwall-cmd：是Linux提供的操作firewall的一个工具（服务）命令
--zone #作用域
--add-port=8080/tcp #添加端口，格式为：端口/通讯协议 ；add表示添加，remove则对应移除
--permanent #永久生效，没有此参数重启后失效
~~~
