# 更新为静态IP

1. ifconfig，查看网卡，确定网卡
2. vim /etc/sysconfig/network-scripts/ifcfg-网卡名
   </br>(以防万一先备份

## 配置文件修改内容
1. 将 BOOTPROTO = dhcp   改成  BOOTPROTO = static
   </br>也就是将动态ip，改成静态ip
2. 新增4行内容：
   ```
   IPADDR="192.168.1.200"          # 静态IP地址
   NETMASK="255.255.255.0"         # 子网掩码
   GATEWAY="192.168.1.1"         # 网关地址(路由器的或提供网络的ip)
   DNS1="223.5.5.5"            # DNS服务器(可以写223.5.5.5或提供网络的ip)
   ```
3. 所有配置说明
   ```
   #类型
   TYPE=Ethernet
   PROXY_METHOD=none
   BROWSER_ONLY=no
   #是否启动DPCH：none为禁止使用；static是使用静态ip；DPCH为使用DPCH服务
   #如果要设定多网口绑定bond，必须为none
   BOOTPROTO=static
   # 设置的静态IP地址
   IPADDR="192.168.1.100"        
   # 子网掩码
   NETMASK="255.255.255.0"      
    # 网关地址   
   GATEWAY="192.168.1.1"      
   # DNS服务器  
   DNS1="223.5.5.5" 
   #default route  是否把这个网卡设置为ipv4默认路由         
   DEFROUTE=yes
   #如果ipv4设置失败则禁用设备
   IPV4_FAILURE_FATAL=no
   #是否使用ipv6
   IPV6INIT=yes
   #ipv6自动配置
   IPV6_AUTOCONF=yes
   #是否把这个网卡设置为ipv6默认路由
   IPV6_DEFROUTE=yes
   #如果ipv6设置失败则禁用设备
   IPV6_FAILURE_FATAL=no
   IPV6_ADDR_GEN_MODE=stable-privacy
   #网络连接的名字
   NAME="ens160"
   #随机的唯一标识
   UUID="05ee6118-eae9-404e-a1af-a20fdd2ccbd6"
   #网卡名称
   DEVICE="ens160"
   #启动或重启是否启动该设备
   ONBOOT=yes
   ```
   
4. 重启网络服务
   ```
   ip地址修改完毕之后，需要重启网络服务，执行如下指令：
   systemctl restart network
   ```
