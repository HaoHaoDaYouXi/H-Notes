# Nginx

`Nginx`是一个`Web`服务器和反向代理服务器，用于`HTTP`、`HTTPS`、`SMTP`、`IMAP`等协议。

下载地址：https://nginx.org/en/download.html

## <a id="zxdlfxdl">正向代理和反向代理</a>

- 正向代理
    - 一个人发送一个请求直接就到达了目标的服务器
- 反方代理
    - 请求统一被`Nginx`接收，`Nginx`反向代理服务器接收到之后，按照一定的规则分发给后端的业务处理服务器进行处理

反向代理服务器可以隐藏源服务器的存在和特征，充当互联网和`Web`服务器之间的中间层，保证安全性。

## <a id="masterworker">`Master`和`Worker`进程</a>

- `Master`进程：读取及评估配置和维持
- `Worker`进程：处理请求

## <a id="fzjh">负载均衡</a>

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

## <a id="xzfw">限制访问</a>

如果要限制IP访问的可以使用`if`指令，或者`geo`和`map`模块。
```nginx
if  ($remote_addr = 192.168.0.105) {  
    return 403;  
}
```

其他限制访问的都可以使用`if`指令和`map`模块，只需要调整参数和判断就可以
```nginx
if ($http_user_agent ~ Chrome) {   
  return 403;  
}
```

## <a id="zz">正则</a>

| 优先级 | 匹配符 | 说明       |
|-----|-----|----------|
| 1   | =   | 精准匹配     |
| 2   | ^~  | 匹配开头     |
| 3   | ~   | 区分大小写的匹配 |
| 4   | ~*  | 不区分大小写   |
| 5   | !~  | ~的取反     |
| 6   | !~* | ~*的取反    |
| 7   | /   | 通用匹配     |

## <a id="cybl">常用变量</a>

- $host: 请求的主机头
- $remote_addr: 客户端IP地址
- $remote_port: 客户端端口号
- $remote_user: 已经经过Auth Basic Module验证的用户名
- $http_referer: 请求引用地址
- $http_user_agent: 客户端代理信息(UA)
- $http_x_forwarded_for: 相当于网络访问路径
- $body_bytes_sent: 页面传送的字节数
- $time_local: 服务器时间
- $request: 客户端请求
- $request_uri: 请求的URI,带参数, 不包含主机名
- $request_filename: 请求的文件路径
- $request_method: 请求的方法，如GET、POST
- $args: 客户端请求中的参数
- $query_string: 等同于$args, 客户端请求的参数
- $nginx_version: 当前nginx版本
- $status: 服务器响应状态码
- $server_addr: 服务器地址
- $server_port: 请求到达的服务器端口号
- $server_protocol: 请求的协议版本
- $content_type: HTTP请求信息里的Content-Type字段
- $content_length: HTTP请求信息里的Content-Length字段
- $uri: 请求中的当前URI(不带请求参数，参数位于$args)
- $document_root: 当前请求在root指令中指定的值
- $document_uri: 与$uri相同

## <a id="cdx">重定向</a>

- return
  - return code;
  - return code http://xxx.xxx/$request_uri;
  - return http://xxx.xxx/$request_uri;
- rewrite
  - rewrite ^/$ http://xxx.xxx permanent;
  - last: 停止处理后续rewrite指令集，然后对当前重写的新URI在rewrite指令集上重新查找
  - break: 停止处理后续rewrite指令集，并不在重新查找,但是当前location内剩余非rewrite语句和location外的非rewrite语句可以执行
  - redirect: 如果replacement不是以http:// 或https://开始，返回302临时重定向
  - permanent: 返回301永久重定向

## <a id="xl">限流</a>

- 限制访问频率（正常、突发）
- 限制并发连接数

### 限制访问频率

限制一个用户发送的请求，`Nginx`多久接收一个请求。

`Nginx`中使用`ngx_http_limit_req_module`模块来限制的访问频率，限制的原理实质是基于漏桶算法原理来实现的。

使用`limit_req_zone`命令及`limit_req`命令限制单个`IP`的请求处理频率。
```nginx
http {
    # 定义限流维度
    limit_req_zone $binary_remote_addr zone=test:10m rate=1r/s;
  
    server{
        location /test {
            # 使用限流规则
            limit_req zone=test burst=5 nodelay;
            proxy_pass http://test_service;
        }
    }
}
```
`1r/s`代表`1`秒一个请求，`1r/m`一分钟接收一个请求， 如果`Nginx`这时还有别人的请求没有处理完，`Nginx`就会拒绝处理该用户请求。

- `limit_req_zone`
  - 第一个参数：`$binary_remote_addr`表示通过`remote_addr`这个标识来做限制，`binary_`的目的是缩写内存占用量，是限制同一客户端`ip`地址。
  - 第二个参数：`zone=test:10m`表示生成一个大小为`10M`，名字为`test`的内存区域，用来存储访问的频次信息。
  - 第三个参数：`rate=1r/s`表示允许相同标识的客户端的访问频次，这里限制的是每秒`1`次，还可以有比如`30r/m`的。
- `limit_req`
  - 第一个参数：`zone=test` 设置使用哪个配置区域来做限制，与`limit_req_zone`里的`name`对应。
  - 第二个参数：`burst=5`，重点说明一下这个配置，burst爆发的意思，这个配置的意思是设置一个大小为`5`的缓冲区当有大量请求（爆发）过来时，超过了访问频次限制的请求可以先放到这个缓冲区内。
  - 第三个参数：`nodelay`，如果设置，超过访问频次而且缓冲区也满了的时候就会直接返回`503`，如果没有设置，则所有请求会等待排队。

正常情况下`limit_req`需要第一个参数，如果要限制突发情况，`limit_req`就需要配置第二、三个参数

### 限制并发连接数

`Nginx`中的`ngx_http_limit_conn_module`模块提供了限制并发连接数的功能，
可以使用`limit_conn_zone`指令以及`limit_conn`执行进行配置。

```nginx
http {
   limit_conn_zone $binary_remote_addr zone=testIp:10m;
   limit_conn_zone $server_name zone=testName:10m;
  
   server {
       location / {
           limit_conn testIp 10;
           limit_conn testName 100;
           rewrite / https://haohaodayouxi.github.io/MyNotes/ permanent;
       }
   }
}
```

- `limit_conn_zone`
  - 第一个参数：`$binary_remote_addr`表示通过`remote_addr`这个标识来做限制，`binary_`的目的是缩写内存占用量，是限制同一客户端`ip`地址。
  - 第二个参数：`zone=testIp:10m`表示生成一个大小为`10M`，名字为`testIp`的内存区域，用来存储访问的频次信息。
- `limit_conn`
  - 第一个参数：`testIp`，表示使用哪个配置区域来做限制，与`limit_conn_zone`里的`name`对应。
  - 第二个参数：`10`，表示允许同一客户端`ip`地址同时连接的个数。

可以配置多个`limit_conn`指令。

上面配置了单个IP同时并发连接数最多只能10个连接，并且设置了整个虚拟服务器同时最大并发数最多只能100个链接。

只有当请求的`header`被服务器处理后，虚拟服务器的连接数才会计数。

## <a id="djfl">动静分离</a>

`Nginx`是目前最热门的`Web`容器，网站优化的重要点在于静态化网站，网站静态化的关键点则是是动静分离，
动静分离是让动态网站里的动态网页根据一定规则把不变的资源和经常变的资源区分开来，
动静资源做好了拆分以后，我们则根据静态资源的特点将其做缓存操作。

让静态的资源只走静态资源服务器，动态的走动态的服务器`Nginx`的静态处理能力很强，但是动态处理能力不足，因此，在企业中常用动静分离技术。

对于静态资源比如图片，`js`，`css`等文件，我们则在反向代理服务器`Nginx`中进行缓存。
这样浏览器在请求一个静态资源时，代理服务器`Nginx`就可以直接处理，无需将请求转发给后端服务器。

若用户请求的动态文件，比如导出文件，下载文件等等则转发给服务器处理，从而实现动静分离。

这也是反向代理服务器的一个重要的作用。

简单来讲，不需要根据业务功能动态变动的文件通过`Nginx`直接访问，减少网络带宽，提高网站性能。

```nginx
location /static/ {
    alias /static/;
    expires 30d;
    add_header Cache-Control "public";
    try_files $uri $uri/ =404;
}
```
静态文件位于`/static/`目录下，通过`location /static/`匹配。

对于静态文件，设置了缓存时间为`30`天，并添加了`Cache-Control`头，以指示客户端和代理服务器缓存文件。

所有其他请求都被转发到服务器，以处理应用的动态内容。

## <a id="pzwjsm">配置文件说明</a>


----
