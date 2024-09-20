# docker创建 jdk21 镜像

```dockerfile
# 基础镜像，镜像从小到大依次为 alpine、ubuntu、debian、centos。推荐alpine构建镜像，减小构建出的镜像文件大小。
FROM alpine:latest

# 维护者
MAINTAINER haohaodayouxi

# 安装JDK21 
RUN apk update \
    && apk -U upgrade \
    && apk add --no-cache \
    openjdk21 --repository=https://dl-cdn.alpinelinux.org/alpine/v$(cut -d. -f1,2 /etc/alpine-release)/community/
```

构建：
```shell
docker build -t jdk21:21.0.4 .
```

----
