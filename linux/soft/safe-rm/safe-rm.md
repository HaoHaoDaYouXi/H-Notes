## linux的safe-rm安装
#### 简介: safe-rm可以保护特定的路径内容不会被意外执行的rm命令去删除

#### 最新的版本可能需要手动编译安装,需要注意的是,linux下使用需要使用linux去编译,不同版本的rust编译产出的文件不像java一样可以通用执行

- 官方地址`https://launchpad.net/safe-rm`
- 选择一个版本下载`https://launchpad.net/safe-rm/+download`
- 解压文件`tar -zxvf safe-rm-1.1.0.tar.gz `
- 在linux上准备rust环境`curl https://sh.rustup.rs -sSf | sh`
- 配置当前shell的rust环境变量`source "$HOME/.cargo/env"`
- 进入解压之后的safe-rm文件夹,查看makefile文件,rust编译命令为`cargo build --release`
- 如果编译卡在类似`Updating `ustc` index`准备依赖位置,需要修改cargo源
- 此文档的编译产出文件对应版本为safe-rm-1.1.0,linux可以直接使用
- 上传至需要配置的服务器,剪切到/usr/local/bin/rm `mv safe-rm /usr/local/bin/rm`
- 添加环境变量,使用上一步的rm替换掉linux本身的rm命令`export PATH=/usr/local/sbin:/usr/sbin:/usr/local/bin:/usr/bin:/bin:/root/bin:/usr/local/bin`
- 刷新环境变量`source /etc/profile`
- 配置safe-rm保护目录`vim /etc/safe-rm.conf`
  ```
    /etc
    /usr
    /usr/lib
    /var
    /data
    /usr/local/bin
  ```
