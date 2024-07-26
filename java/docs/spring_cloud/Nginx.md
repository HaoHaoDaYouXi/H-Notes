# Nginx

`Nginx`是一个`Web`服务器和反向代理服务器，用于`HTTP`、`HTTPS`、`SMTP`、`IMAP`等协议。

下载地址：https://nginx.org/en/download.html

## <a id="zxdlfxdl>正向代理和反向代理</a>

- 正向代理
    - 一个人发送一个请求直接就到达了目标的服务器
- 反方代理
    - 请求统一被`Nginx`接收，`Nginx`反向代理服务器接收到之后，按照一定的规则分发给后端的业务处理服务器进行处理

反向代理服务器可以隐藏源服务器的存在和特征，充当互联网和`Web`服务器之间的中间层，保证安全性。

## `Nginx`服务器上的`Master`和`Worker`进程

- `Master`进程：读取及评估配置和维持
- `Worker`进程：处理请求

## <a id="fzjh>Nginx负载均衡</a>

Nginx负载均衡支持：
- [轮询（默认方式）](Nginx.md#lx)
- [`weight`（权重方式）](Nginx.md#weight)
- [`ip_hash`（依据`ip`分配方式）](Nginx.md#iphash)
- [`least_conn`（依据最少连接方式）](Nginx.md#leastconn)
- [`url_hash`（依据`URL`分配方式，使用第三方插件）](Nginx.md#urlhash)
- [`fair`（依据响应时间方式，使用第三方插件）](Nginx.md#fair)


### <a id="lx">轮询（默认方式）</a>

每个请求按时间顺序逐一分配到不同的后端服务器，如果后端某个服务器宕机，能自动剔除故障系统。
```nginx
upstream test_server {
    server 192.168.0.100:9000;
    server 192.168.0.100:9001;
    server 192.168.0.100:9002;
}
```

### <a id="lx">`weight`（权重方式）</a>

`weight`的值越大分配到的访问概率越高，主要用于每台服务器性能不均衡的情况下。
其次是为在主从的情况下设置不同的权值，达到合理有效的地利用主机资源。
```nginx
upstream test_server {
    server 192.168.0.100:9000 weight=1;
    server 192.168.0.100:9001 weight=2;
    server 192.168.0.100:9002 weight=2;
}
```
权重越高，在被访问的概率越大，5个请求，第一个为1个，第二个为2个，第三个为2个

### <a id="iphash">`ip_hash`（依据`ip`分配方式）</a>

每个请求按访问`IP`的哈希结果分配，使来自同一个IP的访客固定访问一台后端服务器，可以有效解决动态网页存在的`session`共享问题
```nginx
upstream test_server {
    ip_hash;
    server 192.168.0.100:9000;
    server 192.168.0.100:9001;
    server 192.168.0.100:9002;
}
```

### <a id="leastconn">`least_conn`（依据最少连接方式）</a>

最少连接，把请求转发给连接数较少的后端服务器。
轮询算法是把请求平均的转发给各个后端，使它们的负载大致相同

但是，有些请求占用的时间很长，会导致其所在的后端负载较高。
这种情况下，`least_conn`这种方式就可以达到更好的负载均衡效果。
```nginx
upstream test_server {
    least_conn;
    server 192.168.0.100:9000;
    server 192.168.0.100:9001;
    server 192.168.0.100:9002;
}
```

### <a id="urlhash">`url_hash`（依据`URL`分配方式，使用第三方插件）</a>

必须安装`Nginx`的`hash`软件包

按访问`url`的`hash`结果来分配请求，使每个`url`定向到同一个后端服务器，可以进一步提高后端缓存服务器的效率。

使用`Hash`后，不能使用`weight`参数，因为`hash`是根据`url`进行分配的。
```nginx
upstream test_server {
    hash $request_uri;
    hash_method crc32;
    server 192.168.0.100:9000;
    server 192.168.0.100:9001;
    server 192.168.0.100:9002;
}
```
`hash_method`是使用的hash算法，`crc32`是一种校验数值的算法


### <a id="fair">`fair`（依据响应时间方式，使用第三方插件）</a>

必须安装`upstream_fair`模块。

对比`weight`、`ip_hash`更加智能的负载均衡算法，`fair`算法可以根据页面大小和加载时间长短智能地进行负载均衡，响应时间短的优先分配。
```nginx
upstream test_server {
    fair;
    server 192.168.0.100:9000;
    server 192.168.0.100:9001;
    server 192.168.0.100:9002;
}
```
哪个服务器的响应速度快，就将请求分配到那个服务器上。
