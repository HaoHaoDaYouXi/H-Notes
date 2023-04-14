# linux 安装git-lab

## 1、配置yum源

~~~
vim /etc/yum.repos.d/gitlab-ce.repo
复制以下内容，然后 :wq! 保存退出

[gitlab-ce]
name=Gitlab CE Repository
baseurl=https://mirrors.tuna.tsinghua.edu.cn/gitlab-ce/yum/el$releasever/
gpgcheck=0
enabled=1
~~~

## 2、更新本地yum缓存
~~~
yum makecache
~~~

## 3、安装gitlab社区版
~~~
#自动安装最新版本并安装相关依赖
yum install gitlab-ce
也可以下载想要安装的gitlab版本，使用rpm命令安装

下载地址：Index of /gitlab-ce/yum/el7/ | 清华大学开源软件镜像站 | Tsinghua Open Source Mirror

#下载最新版本
wget https://mirrors.tuna.tsinghua.edu.cn/gitlab-ce/yum/el7/gitlab-ce-15.2.0-ce.0.el7.x86_64.rpm --no-check-certificate

#使用rpm安装需要手动解决依赖问题
#前置条件依赖policycoreutils-python、openssh-server设置开机自启
yum -y install policycoreutils-python openssh-server
systemctl enable sshd
systemctl start sshd

#还缺啥安装时会有提示，缺啥安装啥就行
rpm -Uvh gitlab-ce-15.2.0-ce.0.el7.x86_64.rpm

#安装成功后启动
gitlab-ctl reconfigure

gitlab-ctl restart

~~~
## 4、更改配置文件参数
~~~
默认安装完成后安装目录

gitlab组件日志路径：/var/log/gitlab
gitlab配置路径：/etc/gitlab/  路径下有gitlab.rb配置文件
应用代码和组件依赖程序：/opt/gitlab
各个组件存储路径： /var/opt/gitlab/
仓库默认存储路径   /var/opt/gitlab/git-data/repositories
版本文件备份路径：/var/opt/gitlab/backups/
nginx安装路径：/var/opt/gitlab/nginx/
redis安装路径：/var/opt/gitlab/redis
要更改配置文件我们需要修改/etc/gitlab/gitlab.rb

~~~

- 1)更改默认端口
~~~
#ip为你服务器ip port你想要修改的端口号
#修改如下：
external_url 'http://ip:port'
nginx['listen_https'] = false
nginx['listen_port'] = port
nignx['listen_address'] = ['*']
#同时还需要修改nginx配置文件
vim /var/opt/gitlab/nginx/conf/gitlab-http.conf

#修改如下
server {
listen *:port;
server_name ip
if ($http_host = ""){
set $http_host_with_default "ip:port";
}
}
~~~

- 2）配置邮箱
~~~
前置条件：需要安装postfix邮箱

#安装
yum install -y postfix
#设置开机自启
systemctl enable postfix
#启动
systemctl start postfix
#修改以下内容
gitlab_rails['gitlab_email_enabled'] = true
gitlab_rails['gitlab_email_from'] = '发信邮箱'
gitlab_rails['gitlab_email_display_name'] = 'xxx'

gitlab_rails['smtp_enable'] = true
gitlab_rails['smtp_address'] = "smtp.163.com"
gitlab_rails['smtp_port'] = 465
gitlab_rails['smtp_user_name'] = "发信邮箱"
gitlab_rails['smtp_password'] = "smtp客户端授权码"
gitlab_rails['smtp_domain'] = "163.com"
gitlab_rails['smtp_authentication'] = "login"
gitlab_rails['smtp_enable_starttls_auto'] = true
gitlab_rails['smtp_tls'] = true
gitlab_rails['smtp_openssl_verify_mode'] = 'none'


修改完成后重新加载配置文件

gitlab-ctl reconfigure
如果修改了邮箱配置，测试邮箱是否生效

[root@localhost gitlab]# gitlab-rails console
--------------------------------------------------------------------------------
Ruby:         ruby 2.7.5p203 (2021-11-24 revision f69aeb8314) [x86_64-linux]
GitLab:       15.2.0 (a876afc5fd8) FOSS
GitLab Shell: 14.9.0
PostgreSQL:   13.6
------------------------------------------------------------[ booted in 14.38s ]
Loading production environment (Rails 6.1.4.7)
irb(main):001:0> Notify.test_email('xxx@163.com','test','gitlab').deliver_now
Delivered mail 62e23a5839096_40bb4664558c8@localhost.localdomain.mail (1078.9ms)
=> #<Mail::Message:290380, Multipart: false, Headers: <Date: Thu, 28 Jul 2022 15:27:20 +0800>, <From: lick <lick0064@163.com>>, <Reply-To: xxx <noreply@xxx>>, <To: xxx@163.com>, <Message-ID: <62e23a5839096_40bb4664558c8@localhost.localdomain.mail>>, <Subject: test>, <Mime-Version: 1.0>, <Content-Type: text/html; charset=UTF-8>, <Content-Transfer-Encoding: 7bit>, <Auto-Submitted: auto-generated>, <X-Auto-Response-Suppress: All>>
~~~

- 3)修改root管理员密码
~~~
ps:初始密码可以查看

cat /etc/gitlab/initial_root_password
#登录控制台
gitlab-rails console
#查找切换账号
u=User.where(id:1).first
#修改密码
u.password='更改后的密码'
#再次确认密码
u.password='更改后的密码'
#保存
u.save!
~~~
- 4)性能优化
~~~
unicorn['worker_processes'] = 2                         #官方建议值为CPU核数+1（服务器只部署gitLab的情况下），可提高服务器响应速度，此参数最小值为2，设为1服务器可能卡死
unicorn['work_timeout'] = 60                            #设置超时时间
unicorn['worker_memory_limit_min'] = "200 * 1 << 20"    #减少最小内存
unicorn['worker_memory_limit_max'] = "300 * 1 << 20"    #减少最大内存
postgresql['shared_buffers'] = "128MB"                  #减少数据库缓存
postgresql['max_worker_processes'] = 6                  #减少数据库并发数
sidekiq['concurrency'] = 15                             #减少sidekiq并发数
由于 Gitlab 核心功能是代码托管，所以有些额外的组件比较浪费资源，所以可以考虑关闭。

prometheus['enable'] = false
prometheus['monitor_kubernetes'] = false
alertmanager['enable'] = false  
node_exporter['enable'] = false
redis_exporter['enable'] = false
postgres_exporter['enable'] = false
gitlab_exporter['probe_sidekiq'] = false
prometheus_monitoring['enable'] = false
grafana['enable'] = false  
以上就是修改配置文件！

~~~
## 5、重新启动
~~~
sudo gitlab-ctl reconfigure

sudo gitlab-ctl restart
~~~
## 6、其他命令
~~~
gitlab-ctl start #启动全部服务
gitlab-ctl restart#重启全部服务
gitlab-ctl stop #停止全部服务
gitlab-ctl restart nginx #重启单个服务，如重启nginx
gitlab-ctl status #查看服务状态
gitlab-ctl reconfigure #使配置文件生效
gitlab-ctl show-config #验证配置文件
gitlab-ctl uninstall #删除gitlab（保留数据）
gitlab-ctl cleanse #删除所有数据，从新开始
gitlab-ctl tail <service name>查看服务的日志
gitlab-ctl tail nginx  #如查看gitlab下nginx日志
gitlab-rails console  #进入控制台
gitlab-ctl help                  #查看gitlab帮助信息

~~~

