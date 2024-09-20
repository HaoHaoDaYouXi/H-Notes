# DockerFile 参数配置

## 常见指令
- [ARG](#ARG)               # 构建参数，在构建镜像时，可以设置参数，在构建镜像过程中，参数可以动态修改
- [FROM](#FROM)             # 基础镜像，当前新镜像是基于哪个镜像的
- [LABEL](#LABEL)           # 为镜像打标签，标签可以添加到镜像中，在镜像构建过程中，可以添加标签，在镜像运行过程中，可以查看标签
- [MAINTAINER](#MAINTAINER)         # 镜像维护者的姓名混合邮箱地址
- [RUN](#RUN)               # 容器构建时需要运行的命令
- [EXPOSE](#EXPOSE)         # 当前容器对外保留出的端口
- [WORKDIR](#WORKDIR)       # 指定在创建容器后，终端默认登录的进来工作目录，一个落脚点
- [ENV](#ENV)               # 用来在构建镜像过程中设置环境变量
- [USER](#USER)             # 设置运行容器时使用的用户名或UID。
- [ADD](#ADD)               # 将宿主机目录下的文件拷贝进镜像且ADD命令会自动处理URL和解压tar压缩包
- [COPY](#COPY)             # 类似ADD，拷贝文件和目录到镜像中
- [VOLUME](#VOLUME)         # 容器数据卷，用于数据保存和持久化工作
- [CMD](#CMD)               # 指定一个容器启动时要运行的命令，dockerFile中可以有多个CMD指令，但只有最后一个生效！
- [ENTRYPOINT](#ENTRYPOINT) # 指定一个容器启动时要运行的命令！和CMD一样
- [ONBUILD](#ONBUILD)       # 当构建一个被继承的DockerFile时运行命令，父镜像在被子镜像继承后，父镜像的ONBUILD被触发

文档：https://docs.docker.com/reference/dockerfile/

## 指令示例

### ARG

#### 格式

```
ARG <name>[=<default value>]
```

#### 示例

```dockerfile
ARG username=default_user
```

### FROM

#### 格式

```
FROM <image>[:<tag>] 或 FROM <image>@<digest>

<image>		# 指定base image的名称
<tag>		# base image的标签，省略时默认latest
<digest>	# 时镜像的哈希码；使用哈希码会更安全一点
```

#### 示例

```dockerfile
FROM alpine:latest
```

### LABEL

#### 格式

```
LABEL <key>=<value>
```

#### 示例

```dockerfile
LABEL author="haohaodayouxi <2601183227@qq.com>"
LABEL version="1.0"
LABEL description="This is a simple alpine image"
```

### MAINTAINER

#### 格式

```
MAINTAINER <name>
```

#### 示例

```dockerfile
MAINTAINER haohaodayouxi <2601183227@qq.com>
```

### RUN

#### 格式

```
RUN <command> 或
RUN [“<executable>”,”<param1>”,”<param2>”]

第一种格式中<command>通常是一个shell命令，且以”/bin/sh -c“来运行，这意味着此进程在容器中的PID不为1，不能接收Unix信号。
因此，当使用docker stop <container>命令停止容器时，此进程接收不到SIGTERM信号

第二种语法格式中的参数是一个JSON格式数组，其中<executable>为要运行的命令，后面的<paramN>为传递给命令的选项或参数；
然而，此种格式指定的命令不会以”/bin/sh -c”来发起，因此常见的shell操作如变量替换以及通配符替换将不会进行；
```

#### 示例

```dockerfile
RUN apk update && apk add curl
```

### EXPOSE

#### 格式

```
EXPOSE <port> [<port>/<protocol>]

<port>	    # 暴露的端口号，如80
<protocol>	# 传输层协议，可为tcp或udp，默认为tcp
```

#### 示例

```dockerfile
EXPOSE 80/tcp
```

### WORKDIR

#### 格式

```
WORKDIR <path>

<path>	# 工作目录，必须是一个绝对路径

# WORKDIR可以多次出现，其路径也可以为相对路径；相对路径是对此前一个WORKDIER指令指定的路径；
# WORKDIR也可以调用由ENV指定定义的变量；
```

#### 示例

```dockerfile
# 设置工作目录为 /data1
WORKDIR /data1

# 在 /data1 目录下创建一个子目录 folder1
RUN mkdir folder1

# 设置工作目录为 /data2
WORKDIR /data2

# 在 /data2 目录下创建一个子目录 folder2
RUN mkdir folder2

# 设置工作目录为 /data2/folder2
WORKDIR folder2

# 在 /data2/folder2 目录下创建一个子目录 folder3
RUN mkdir folder3
```

### ENV

#### 格式

```
ENV <key>=<value> [<key>=<value>, ...] 或 ENV <key> <value>

第一种格式中，<key>=<value>为键值对，可以定义多个键值对，键值对之间用空格隔开，如果<value>中包含空格，可以用“\“进行转义，也可以通过对<value>加引号进行标识；反斜线也可用于续行；
第二种格式中，<key>为键，<value>为值，只能定义一个键值对；
```

#### 示例

```dockerfile
ENV username haohaodayouxi

ENV test=test doc='doc' \
    test2='test 2'
```

### USER

#### 格式

```
USER <user>[:<group>]
```

#### 示例

```dockerfile
USER root
```

### ADD

#### 格式

```
ADD <src> <dest> 

如果为URL且不以/结尾，则指定的文件将被下载并直接被创建为，如果以/结尾，则文件名URL指定的文件将被直接下载并保存为/
如果是一个本地文件系统上的压缩格式为tar文件，将被展开为一个目录，其行为类似于“tar -x“命令；然而，通过URL获取到的tar文件将不会自动展开
如果有多个，或其简介或直接使用了通配符，则必须是一个以/结尾的目录路径；如果不以/结尾，则其被视作一个普通文件，的内容将被直接写入到
```

#### 示例

```dockerfile
ADD ./test.html /data/test/

ADD https://nginx.org/download/nginx-1.26.2.tar.gz /usr/local/src/
```

### COPY

#### 格式

```
COPY <src> <dest>

拷贝文件到容器，源文件地址与容器同目录
若地址是目录，则其内部文件或子目录会被递归复制；但目录自身不会被复制
如果指定了多个，或在中使用了通配符，则必须是一个目录，且必须以/结尾
如果事先不存在，他将会被自动创建，这包括父目录路径
```

#### 示例

```dockerfile
COPY ./test.html /data/test/
```

### VOLUME

#### 格式

```
VOLUME <path>

如果挂载点目录路径下此前的文件存在，docker run命令会在卷挂载完成后将此前的所有文件复制到新挂载的卷中
```

#### 示例

```dockerfile
VOLUME /data
```

### CMD

#### 格式

```
CMD <command> 或 CMD [“<executable>”,”<param1>”,”<param2>”]
```

#### 示例

```dockerfile
CMD java -jar /app/app.jar

CMD ["java","-jar","/app/app.jar"]
```

### ENTRYPOINT

#### 格式

```
ENTRYPOINT <command> 或 ENTRYPOINT [“<executable>”,”<param1>”,”<param2>”]

与CMD不同的是，有ENTRYPOINT启动的程序不会被docker run命令指定的参数覆盖，而且，这些命令参数会被当做参数传递给ENTRYPOTINT指定的程序

docker run命令的—entrypoint选项的参数可覆盖ENTRYPOINT指令指定的程序
Dockerfile文件中也可以存在多个ENTRYPOINT指令，但仅有最后一个会生效
```

#### 示例

```dockerfile
ENTRYPOINT java -jar /app/app.jar
```

### ONBUILD

#### 格式

```
ONBUILD <INSTRUCTION>

ONBUILD不能自我嵌套，且不会触发FROM和MAINTAINER指令；
```

#### 示例

```dockerfile
ONBUILD RUN mkdir /data
```


----
