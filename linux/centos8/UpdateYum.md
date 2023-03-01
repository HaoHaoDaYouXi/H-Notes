# 更新为阿里云的Yum源

步骤：

1. 以root用户打开/etc/yum.repos.d文件夹，找到三个文件CentOS-AppStream.repo、CentOS-Base.repo、CentOS-Extras.repo

2. 加#注释掉原来的地址mirrorlist=http://mirrorlist.centos.org/?

3. 在注释掉的下一行分别加上阿里云镜像源地址，可以看到文件名与下面要替换的地址是对应的
    ```
    baseurl=https://mirrors.aliyun.com/centos/$releasever-stream/AppStream/$basearch/os/
    
    baseurl=https://mirrors.aliyun.com/centos/$releasever-stream/BaseOS/$basearch/os/
    
    baseurl=https://mirrors.aliyun.com/centos/$releasever-stream/extras/$basearch/os/
    ```
   
4. 在终端输入yum clean all 回车执行

5. 在终端输入yum makecache 回车执行

6. 完成。