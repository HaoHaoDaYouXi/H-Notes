# rocketmq

## 安装和使用
- [下载地址](https://rocketmq.apache.org/zh/download/)
- windows
  - 配置 环境变量 ROCKETMQ_HOME
  - 部署完成后进入bin目录，修改runserver.cmd的JAVA_OPT，大小可自己调整(初始值太大会导致分配内存不足而出错)
  - 进入到rocketMQ的bin目录下，然后执行下面的命令：
    ~~~
    #启动RocketMQ的注册中心
    start mqnamesrv.cmd
    #执行成功会弹出提示框，不要关闭
    #启动broker
    start mqbroker.cmd -n localhost:9876 -c ../conf/broker.conf autoCreateTopicEnable=true
    #执行成功会弹出提示框，不要关闭
    ~~~
~~~
当遇到在windows 下 出现无法启动broker的情况时，考虑是否为Rocketmq Broker异常关闭导致
解决方案：
尝试将C盘 C:\Users\用户 的store文件夹删除，再进行启动
~~~
windows启动命令：[start_mq.bat](start_mq.bat)

### 端口说明
我们在安装rocketmq后，要开放的端口一般有4个：9876，10911，10912，10909
#### 1. 首先说9876
>这个是nameserver中的端口,不做过多解释，链接nameserver就靠这个端口

#### 2. 剩下的3个都是RocketMQ-Broker中的端口
#### listenPort
>listenPort参数是broker的监听端口号，是remotingServer服务组件使用，作为对Producer和Consumer提供服务的端口号，默认为10911，可以通过配置文件修改。

打开broker-x.conf，修改或增加listenPort参数：
~~~
#Broker 对外服务的监听端口
listenPort=10911
~~~
#### fastListenPort
>fastListenPort参数是fastRemotingServer服务组件使用，默认为listenPort - 2，可以通过配置文件修改。

打开broker-x.conf，修改或增加fastListenPort参数
~~~
#主要用于slave同步master
fastListenPort=10909
~~~

#### haListenPort
>haListenPort参数是HAService服务组件使用，用于Broker的主从同步，默认为listenPort - 1，可以通过配置文件修改。

打开broker-x.conf，修改或增加haListenPort参数：
~~~
#haService中使用
haListenPort=10912
~~~~
#### remotingServer和fastRemotingServer的区别：
Broker端：

>remotingServer可以处理客户端所有请求，如：生产者发送消息的请求，消费者拉取消息的请求。</br>
>fastRemotingServer功能基本与remotingServer相同，唯一不同的是不可以处理消费者拉取消息的请求。</br>
>Broker在向NameServer注册时，只会上报remotingServer监听的listenPort端口。</br>

客户端：

>默认情况下，生产者发送消息是请求fastRemotingServer，我们也可以通过配置让其请求remotingServer；消费者拉取消息只能请求remotingServer。
