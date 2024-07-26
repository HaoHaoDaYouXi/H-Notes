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

详细的可以查看官方文档：https://nginx.org/en/docs/varindex.html

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

详细的可查看官方文档：https://nginx.org/en/docs/dirindex.html

```nginx
#定义Nginx运行的用户和用户组 
user root root;

#nginx进程数，建议设置为等于CPU总核心数。
worker_processes 8;
 
#全局错误日志定义类型，[ debug | info | notice | warn | error | crit ]
error_log /usr/local/nginx/logs/error.log info;

#进程pid文件
pid /usr/local/nginx/logs/nginx.pid;

#指定进程可以打开的最大描述符：数目
#工作模式与连接数上限
#这个指令是指当一个nginx进程打开的最多文件描述符数目，理论值应该是最多打开文件数（ulimit -n）与nginx进程数相除，但是nginx分配请求并不是那么均匀，所以最好与ulimit -n 的值保持一致。
#现在在linux 2.6内核下开启文件打开数为65535，worker_rlimit_nofile就相应应该填写65535。
#这是因为nginx调度时分配请求到进程并不是那么的均衡，所以假如填写10240，总并发量达到3-4万时就有进程可能超过10240了，这时会返回502错误。
worker_rlimit_nofile 65535;

events {
    #参考事件模型，use [ kqueue | rtsig | epoll | /dev/poll | select | poll ]; epoll模型
    #是Linux 2.6以上版本内核中的高性能网络I/O模型，linux建议epoll，如果跑在FreeBSD上面，就用kqueue模型。
    #补充说明：
    #与apache相类，nginx针对不同的操作系统，有不同的事件模型
    #A）标准事件模型
    #Select、poll属于标准事件模型，如果当前系统不存在更有效的方法，nginx会选择select或poll
    #B）高效事件模型
    #Kqueue：使用于FreeBSD 4.1+, OpenBSD 2.9+, NetBSD 2.0 和 MacOS X.使用双处理器的MacOS X系统使用kqueue可能会造成内核崩溃。
    #Epoll：使用于Linux内核2.6版本及以后的系统。
    #/dev/poll：使用于Solaris 7 11/99+，HP/UX 11.22+ (eventport)，IRIX 6.5.15+ 和 Tru64 UNIX 5.1A+。
    #Eventport：使用于Solaris 10。 为了防止出现内核崩溃的问题， 有必要安装安全补丁。
    use epoll;

    #单个进程最大连接数（最大连接数=连接数*进程数）
    #根据硬件调整，和前面工作进程配合起来用，尽量大，但是别把cpu跑到100%就行。每个进程允许的最多连接数，理论上每台nginx服务器的最大连接数为。
    worker_connections 65535;

    #keepalive超时时间。
    keepalive_timeout 60;

    #客户端请求头部的缓冲区大小。这个可以根据你的系统分页大小来设置，一般一个请求头的大小不会超过1k，不过由于一般系统分页都要大于1k，所以这里设置为分页大小。
    #分页大小可以用命令getconf PAGESIZE 取得。
    #[root@web001 ~]# getconf PAGESIZE
    #4096
    #但也有client_header_buffer_size超过4k的情况，但是client_header_buffer_size该值必须设置为“系统分页大小”的整倍数。
    client_header_buffer_size 4k;

    #这个将为打开文件指定缓存，默认是没有启用的，max指定缓存数量，建议和打开文件数一致，inactive是指经过多长时间文件没被请求后删除缓存。
    open_file_cache max=65535 inactive=60s;

    #这个是指多长时间检查一次缓存的有效信息。
    #语法:open_file_cache_valid time 默认值:open_file_cache_valid 60 使用字段:http, server, location 这个指令指定了何时需要检查open_file_cache中缓存项目的有效信息.
    open_file_cache_valid 80s;

    #open_file_cache指令中的inactive参数时间内文件的最少使用次数，如果超过这个数字，文件描述符一直是在缓存中打开的，如上例，如果有一个文件在inactive时间内一次没被使用，它将被移除。
    #语法:open_file_cache_min_uses number 默认值:open_file_cache_min_uses 1 使用字段:http, server, location  这个指令指定了在open_file_cache指令无效的参数中一定的时间范围内可以使用的最小文件数,如果使用更大的值,文件描述符在cache中总是打开状态.
    open_file_cache_min_uses 1;
    
    #语法:open_file_cache_errors on | off 默认值:open_file_cache_errors off 使用字段:http, server, location 这个指令指定是否在搜索一个文件时记录cache错误.
    open_file_cache_errors on;
}

#设定http服务器，利用它的反向代理功能提供负载均衡支持
http {
    #文件扩展名与文件类型映射表
    include mime.types;

    #默认文件类型
    default_type application/octet-stream;

    #默认编码
    charset utf-8;

    #服务器名字的hash表大小
    #保存服务器名字的hash表是由指令server_names_hash_max_size 和server_names_hash_bucket_size所控制的。参数hash bucket size总是等于hash表的大小，并且是一路处理器缓存大小的倍数。在减少了在内存中的存取次数后，使在处理器中加速查找hash表键值成为可能。如果hash bucket size等于一路处理器缓存的大小，那么在查找键的时候，最坏的情况下在内存中查找的次数为2。第一次是确定存储单元的地址，第二次是在存储单元中查找键 值。因此，如果Nginx给出需要增大hash max size 或 hash bucket size的提示，那么首要的是增大前一个参数的大小.
    server_names_hash_bucket_size 128;

    #客户端请求头部的缓冲区大小。这个可以根据你的系统分页大小来设置，一般一个请求的头部大小不会超过1k，不过由于一般系统分页都要大于1k，所以这里设置为分页大小。分页大小可以用命令getconf PAGESIZE取得。
    client_header_buffer_size 32k;

    #客户请求头缓冲大小。nginx默认会用client_header_buffer_size这个buffer来读取header值，如果header过大，它会使用large_client_header_buffers来读取。
    large_client_header_buffers 4 64k;

    #设定通过nginx上传文件的大小
    client_max_body_size 8m;

    #开启高效文件传输模式，sendfile指令指定nginx是否调用sendfile函数来输出文件，对于普通应用设为 on，如果用来进行下载等应用磁盘IO重负载应用，可设置为off，以平衡磁盘与网络I/O处理速度，降低系统的负载。注意：如果图片显示不正常把这个改成off。
    #sendfile指令指定 nginx 是否调用sendfile 函数（zero copy 方式）来输出文件，对于普通应用，必须设为on。如果用来进行下载等应用磁盘IO重负载应用，可设置为off，以平衡磁盘与网络IO处理速度，降低系统uptime。
    sendfile on;

    #开启目录列表访问，合适下载服务器，默认关闭。
    autoindex on;

    #此选项允许或禁止使用socke的TCP_CORK的选项，此选项仅在使用sendfile的时候使用
    tcp_nopush on;
     
    tcp_nodelay on;

    #长连接超时时间，单位是秒
    keepalive_timeout 120;

    #FastCGI相关参数是为了改善网站的性能：减少资源占用，提高访问速度。下面参数看字面意思都能理解。
    fastcgi_connect_timeout 300;
    fastcgi_send_timeout 300;
    fastcgi_read_timeout 300;
    fastcgi_buffer_size 64k;
    fastcgi_buffers 4 64k;
    fastcgi_busy_buffers_size 128k;
    fastcgi_temp_file_write_size 128k;

    #gzip模块设置
    gzip on; #开启gzip压缩输出
    gzip_min_length 1k;    #最小压缩文件大小
    gzip_buffers 4 16k;    #压缩缓冲区
    gzip_http_version 1.0;    #压缩版本（默认1.1，前端如果是squid2.5请使用1.0）
    gzip_comp_level 2;    #压缩等级
    gzip_types text/plain application/x-javascript text/css application/xml;    #压缩类型，默认就已经包含textml，所以下面就不用再写了，写上去也不会有问题，但是会有一个warn。
    gzip_vary on;

    #限流和限制访问时使用
    #limit_req_zone $binary_remote_addr zone=test:10m rate=1r/s;
    #limit_conn_zone $binary_remote_addr zone=testIp:10m;
    #limit_conn_zone $server_name zone=testName:10m;
  
    #日志格式设定
    #$remote_addr与$http_x_forwarded_for用以记录客户端的ip地址；
    #$remote_user：用来记录客户端用户名称；
    #$time_local： 用来记录访问时间与时区；
    #$request： 用来记录请求的url与http协议；
    #$status： 用来记录请求状态；成功是200，
    #$body_bytes_sent ：记录发送给客户端文件主体内容大小；
    #$http_referer：用来记录从那个页面链接访问过来的；
    #$http_user_agent：记录客户浏览器的相关信息；
    #通常web服务器放在反向代理的后面，这样就不能获取到客户的IP地址了，通过$remote_add拿到的IP地址是反向代理服务器的iP地址。反向代理服务器在转发请求的http头信息中，可以增加x_forwarded_for信息，用以记录原有客户端的IP地址和原来客户端的请求的服务器地址。
    log_format access '$remote_addr - $remote_user [$time_local] "$request" '
    '$status $body_bytes_sent "$http_referer" '
    '"$http_user_agent" $http_x_forwarded_for';
  
    #定义本虚拟主机的访问日志
    access_log  /usr/local/nginx/logs/host.access.log  main;
  
    # 负载均衡，自行根据需要配置，可参考上面的，一般情况都是使用权重的配置
    upstream test_server {
        server 192.168.0.100:9000 weight=1;
        server 192.168.0.100:9001 weight=2;
        server 192.168.0.100:9002 weight=2;
    }
  
    #虚拟主机的配置
    server {
        #监听端口
        listen 80;
        #监听IPv6端口
        listen [::]:80;
        #listen 80 ssl; ssl配置可以不是443端口的
        #域名可以有多个，用空格隔开
        server_name test.com www.test.com;
        #日志文件
        access_log /var/log/nginx/test/80/access.log main;
        error_log /var/log/nginx/test/80/error.log;

        #如果是SSL访问，则需要配置ssl参数
        #ssl证书文件
        ssl_certificate conf.d/ssl/test.com.pem;
        ssl_certificate_key conf.d/ssl/test.com.key;
        #ssl会话缓存
        ssl_session_timeout 5m;
        #ssl加密算法
        ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
        #ssl加密算法的集合
        ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:HIGH:!aNULL:!MD5:!RC4:!DHE:!MD5;
        #优先使用服务器端的加密算法
        ssl_prefer_server_ciphers on;
      
        #对 "/" 启用反向代理
        location / {
            proxy_pass http://127.0.0.1:9000;
            proxy_redirect off;
            proxy_set_header X-Real-IP $remote_addr;
             
            #后端的Web服务器可以通过X-Forwarded-For获取用户真实IP
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
             
            #以下是一些反向代理的配置，可选。
            proxy_set_header Host $host;

            #允许客户端请求的最大单文件字节数
            client_max_body_size 10m;

            #缓冲区代理缓冲用户端请求的最大字节数，
            #如果把它设置为比较大的数值，例如256k，那么，无论使用firefox还是IE浏览器，来提交任意小于256k的图片，都很正常。如果注释该指令，使用默认的client_body_buffer_size设置，也就是操作系统页面大小的两倍，8k或者16k，问题就出现了。
            #无论使用firefox4.0还是IE8.0，提交一个比较大，200k左右的图片，都返回500 Internal Server Error错误
            client_body_buffer_size 128k;

            #表示使nginx阻止HTTP应答代码为400或者更高的应答。
            proxy_intercept_errors on;

            #后端服务器连接的超时时间_发起握手等候响应超时时间
            #nginx跟后端服务器连接超时时间(代理连接超时)
            proxy_connect_timeout 90;

            #后端服务器数据回传时间(代理发送超时)
            #后端服务器数据回传时间_就是在规定时间之内后端服务器必须传完所有的数据
            proxy_send_timeout 90;

            #连接成功后，后端服务器响应时间(代理接收超时)
            #连接成功后_等候后端服务器响应时间_其实已经进入后端的排队之中等候处理（也可以说是后端服务器处理请求的时间）
            proxy_read_timeout 90;

            #设置代理服务器（nginx）保存用户头信息的缓冲区大小
            #设置从被代理服务器读取的第一部分应答的缓冲区大小，通常情况下这部分应答中包含一个小的应答头，默认情况下这个值的大小为指令proxy_buffers中指定的一个缓冲区的大小，不过可以将其设置为更小
            proxy_buffer_size 4k;

            #proxy_buffers缓冲区，网页平均在32k以下的设置
            #设置用于读取应答（来自被代理服务器）的缓冲区数目和大小，默认情况也为分页大小，根据操作系统的不同可能是4k或者8k
            proxy_buffers 4 32k;

            #高负荷下缓冲大小（proxy_buffers*2）
            proxy_busy_buffers_size 64k;

            #设置在写入proxy_temp_path时数据的大小，预防一个工作进程在传递文件时阻塞太长
            #设定缓存文件夹大小，大于这个值，将从upstream服务器传
            proxy_temp_file_write_size 64k;
        }
      
        #所有静态文件由nginx直接读取不经过tomcat或resin
        location ~ .*.(htm|html|gif|jpg|jpeg|png|bmp|swf|ioc|rar|zip|txt|flv|mid|doc|ppt|pdf|xls|mp3|wma)$ {
            expires 15d; 
        }
         
        location ~ .*.(js|css)?$ {
            expires 1h;
        }

        #设定查看Nginx状态的地址
        location /NginxStatus {
          stub_status on;
          access_log on;
          auth_basic "NginxStatus";
          auth_basic_user_file confpasswd;
          #htpasswd文件的内容可以用apache提供的htpasswd工具来产生。
        }
      
        #本地动静分离反向代理配置
        #所有jsp的页面均交由tomcat或resin处理
        location ~ .(jsp|jspx|do)?$ {
          proxy_set_header Host $host;
          proxy_set_header X-Real-IP $remote_addr;
          proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
          proxy_pass http://127.0.0.1:8080;
        }
    }
}
```

----
