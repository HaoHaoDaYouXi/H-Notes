# 更新为阿里云的Yum源

步骤：

1. 以root用户打开/etc/yum.repos.d文件夹，找到CentOS-Base.repo

2. 加#注释掉原来的地址mirrorlist=http://mirrorlist.centos.org/?

3. 替换阿里云镜像源地址，也可以直接替换文件[CentOS-Base.repo](ALi_yum/CentOS-Base.repo)
    ```
    baseurl=http://mirrors.aliyun.com/centos/$releasever/os/$basearch/
    ```
   
4. 在终端输入yum clean all 回车执行

5. 在终端输入yum makecache 回车执行

6. 完成。

----
