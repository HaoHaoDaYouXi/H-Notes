# 利用 logrotate 管理日志

1. 默认centos系统安装自带logrotate，如果未安装执行以下命令
    ```
    yum -y install logrotate
    ```
2. 软件包说明
    ```
    rpm -ql logrotate
    /etc/logrotate.conf  # 主配置文件
    /etc/logrotate.d   # 配置目录
    ```
3. 使用 logrotate 切割日志(nginx版本)
    ```
    vim /etc/logrotate.d/nginx 
    
    /var/log/nginx/*.log /var/log/nginx/*/*/*.log {   # 可以指定多个路径,此处为nginx存储日志的地方；
    create 0640 nginx root   #指定新文件权限，属主属组
    daily                    #指定周期为每天,除了daily，还可以选monthly-月，weekly-周，yearly-年
    rotate 10                #保留日志文件个数，保留10天
    missingok                #如果日志文件丢失，不要显示错误
    notifempty               #当日志文件为空时，不进行轮转 
    compress                 #通过gzip 压缩转储以后的日志   nocompress-不压缩
    delaycompress            #当前转储的日志文件到下一次转储时才压缩 
    sharedscripts            #所有的文件切割之后只执行一次下面脚本
    postrotate               #执行的指令
    /bin/kill -USR1 `cat /run/nginx.pid 2>/dev/null` 2>/dev/null || true
    endscript
    }
   
    在/etc/logrotate.d 编辑文件，下面是部分语法内容
    xxxxxx xxxx {   # 可以指定多个路径 空格分割
    create 0640 nginx root        #指定新文件权限，属主属组
    daily                         #指定周期为每天,除了daily，还可以选monthly-月，weekly-周，yearly-年
    rotate 10                     #保留日志文件个数，保留10天
    size +100M                    #超过100M时分割，单位K,M,G，优先级高于daily
    missingok                     #如果日志不存在，不提示错误，nomissingok-提示错误，默认值
    notifempty                    #日志为空时不进行切换，默认为ifempty 
    compress                      #压缩日志文件 nocompress-不压缩
    delaycompress                 #延迟压缩（当次切割不压缩，下次切割再压缩上一个文件）
    sharedscripts                 #所有的文件切割之后只执行一次下面脚本
    prerotate [命令] endscript     #指定切割前进行命令操作
    postrotate [命令] endscript    #指定切割后进行命令操作
    }
    ```
   
4. 保存好配置文件后，测试效果
   ```
   logrotate -vf /etc/logrotate.d/nginx 
   ```
