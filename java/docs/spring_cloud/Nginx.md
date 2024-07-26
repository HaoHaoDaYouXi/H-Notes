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



----
