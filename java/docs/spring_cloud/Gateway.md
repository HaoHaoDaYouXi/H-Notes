# Gateway

`Spring Cloud Gateway`是`Spring Cloud`生态系统中的一个全新项目，旨在为微服务架构提供一种简单有效的统一的`API`路由管理方式。
它是基于`Project Reactor`和`Spring Framework 5.0`的`WebFlux`构建的，提供了比`Zuul`更高的性能和更现代的非阻塞反应式编程模型。
它是为了替换`Zuul`而开发的网关服务，底层使用的是`Netty`。

`Spring Cloud Gateway`的主要功能包括：

- 动态路由：可以根据不同的条件（如路径、方法、头信息等）将请求路由到不同的微服务。
- 过滤器：类似于`Zuul`的过滤器，可以在请求和响应之间执行预定义的操作，如修改请求或响应、日志记录、认证授权等。
- 限流与熔断：可以通过集成`Resilience4j`或其他库来实现请求限流和熔断机制，增强系统的稳定性和可用性。
- 协议代理：可以代理`HTTP`、`WebSocket`等多种协议的请求，支持更广泛的应用场景。

`Spring Cloud Gateway`的设计目标是提供一个轻量级、高性能、易于扩展的`API`网关解决方案，适用于现代的微服务架构。

## Gateway 三个核心点

**`Route`（路由）**

路由是构建网关的基础模块，它由`ID`，目标`URI`，包括一些列的断言和过滤器组成，如果断言为`true`则匹配该路由

**`Predicate`（断言）**

参考的是`Java8`的`java.util.function.Predicate`，开发人员可以匹配`HTTP`请求中的所有内容（例如请求头或请求参数），请求与断言匹配则进行路由

**`Filter`（过滤）**

指的是`Spring`框架中`GateWayFilter`的实例，使用过滤器，可以在请求被路由前或者之后对请求进行修改。

三个核心点连起来：

当用户发出请求到达`Gateway`，`Gateway`会通过一些匹配条件，定位到真正的服务节点，并在这个转发过程前后，进行一些及细化控制。其中`Predicate`就是我们匹配的条件，而`Filter`可以理解为一个拦截器，有了这两个点，再加上目标`URI`，就可以实现一个具体的路由了。

`Gateway`核心的流程就是：`路由转发`+`执行过滤器链`

## `Route`（路由）

路由是`Gateway`中最基本的组件之一，表示一个具体的路由信息载体。

主要定义了下面的几个信息:

- `id`：路由标识、区别于其他`route`
- `uri`：路由指向的目的地`uri`，即客户端请求最终被转发到的微服务
- `order`：用于多个`route`之间的排序，数值越小排序越靠前，匹配优先级越高
- `predicate`：断言的作用是进行条件判断，只有断言都返回真，才会真正的执行路由
- `filter`：过滤器用于修改请求和响应信息

## 执行流程：

- `Gateway Client`向`Gateway Server`发送请求
- 请求首先会被`HttpWebHandlerAdapter`进行提取组装成网关上下文
- 然后网关的上下文会传递到`DispatcherHandler`，它负责将请求分发给`RoutePredicateHandlerMapping`
- `RoutePredicateHandlerMapping`负责路由查找，并根据路由断言判断路由是否可用
- 如果过断言成功，由`FilteringWebHandler`创建过滤器链并调用
- 请求会一次经过`PreFilter` >> 微服务 >> `PostFilter`的方法，最终返回响应

## 使用

### 引用依赖

```xml
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-gateway</artifactId>
</dependency>
```

### yml配置

```yml
spring:
    gateway:
      discovery:
        locator:
          enabled: true # 表明gateway开启服务注册和发现的功能，并且spring cloud gateway自动根据服务发现为每一个服务创建了一个router
          lower-case-service-id: true  # 服务名小写
      # 全局的跨域处理
      globalcors:
        add-to-simple-url-handler-mapping: true # 解决options请求被拦截问题
        corsConfigurations:
          '[/**]':
            allowedOrigins: # 允许哪些网站的跨域请求 "*" 表示全部
              - "http://localhost:9000"
            allowedMethods: # 允许的跨域ajax的请求方式 "*" 表示全部
              - "GET"
              - "POST"
              - "DELETE"
              - "PUT"
              - "OPTIONS"
            allowedHeaders: "*" # 允许在请求中携带的头信息
            allowCredentials: true # 是否允许携带cookie
            maxAge: 360000 # 这次跨域检测的有效期
      # 当service实例不存在时默认返回503，显示配置返回404
      loadbalancer:
        use404: true
      # 路由(如果使用动态路由方式，不要在配置文件中配置路由）
      routes:  # 路由
        - id: test # 路由ID，没有固定要求，但是要保证唯一，建议配合服务名
          uri: http://localhost:9000/test  # 转发地址可以是http://xxx.xx 直接转发，也可以是lb://服务名 通过nacos进行访问   lb=loadbalancer
          predicates: # 断言
            - Path=/lwz/** # 断言，路径相匹配进行路由
#            - After=yyyy-MM-DDTHH:mm:ss.sss+08:00[Asia/Shanghai] # 在这个时间之后的请求都能通过
#            - Cookie=token,[a-z]+ # 匹配Cookie的key和value（正则表达式）
#            - Header=X-Request-Id, \d+
#            - Method=GET  # 等等判断条件
#          filters: # 过滤器是路由转发请求时所经过的过滤逻辑，可用于修改请求、响应内容
#            - StripPrefix=1 # 去掉地址中的第一部分
#            - RequestTime=true # 自定义过滤器是否开启，可以写间称- RequestTime，或者全称- RequestTimeGatewayFilterFactory，具体看你的名称叫什么
```

### 全局过滤器

```java
public class AuthFilter implements GlobalFilter, Ordered {

    @Override
    public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
        log.debug("AuthFilter");
        ServerHttpRequest request = exchange.getRequest();
        ServerHttpRequest.Builder mutate = request.mutate();
        String url = request.getURI().getPath();
        String ignoreUrl="test";
        // 跳过不需要验证的路径
        if (url.contains(ignoreUrl)) {
            return chain.filter(exchange);
        }
        // 获取token
        String token = getToken(request);
        if (StringUtils.isEmpty(token)) {
             return "令牌不能为空";
        }
        return chain.filter(exchange.mutate().request(mutate.build()).build());
    }

    /**
     * 获取请求token
     */
    private String getToken(ServerHttpRequest request) {
        return request.getHeaders().getFirst("token");
    }
    
    @Override
    public int getOrder() {
        return 0;
    }
}
```

### 自定义过滤器

#### 通过GatewayFilter实现

这种需要配合`GatewayConfig`进行配置使用

```java
@Component
public class TestGatewayFilter implements GatewayFilter, Ordered {
    @Override
    public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
        // 获取请求路径
        String path = exchange.getRequest().getPath().toString();
        URI uri = exchange.getRequest().getURI();
        System.err.println(String.format("获取到请求路径：%s", uri.toString()));
        // 如果请求路径以 test 开头，则截取掉第一个路径段
        if (path.startsWith("/test")) {
            path = path.substring("/test".length());
        }
 
        // 创建新的请求对象，并将新路径设置为请求路径
        ServerHttpRequest newRequest = exchange.getRequest().mutate()
                .path(path)
                .build();
 
        // 使用新请求对象创建新的ServerWebExchange对象
        ServerWebExchange newExchange = exchange.mutate()
                .request(newRequest)
                .build();
        System.err.println(String.format("获取到新的请求路径：%s", newExchange.getRequest().getURI()));
        // 继续执行过滤器链
        return chain.filter(newExchange);
    }
 
    @Override
    public int getOrder() {
        return 0;
    }
}

@Configuration
public class GatewayConfig {
    @Value("${server.port}")
    private String port;

    @Bean
    public RouteLocator customerRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                .route(r -> r.path("/test/**")
                        .filters(f -> f.filter(new TestGatewayFilter()))
                        .uri("http://localhost:" + port)
                )
                .build();
    }
}
```


#### 继承AbstractGatewayFilterFactory

这种可以通过配置文件直接配置使用
```java
public class RequestTimeGatewayFilterFactory extends AbstractGatewayFilterFactory<RequestTimeGatewayFilterFactory.Config> {
    private static final String REQUEST_TIME_BEGIN = "requestTimeBegin";
    private static final String KEY = "withParams";

    @Override
    public List<String> shortcutFieldOrder() {
        return Collections.singletonList(KEY);
    }

    public RequestTimeGatewayFilterFactory() {
        super(Config.class);
    }

    @Override
    public GatewayFilter apply(Config config) {
        return (exchange, chain) -> {
            log.debug("RequestTimeGatewayFilterFactory");
            exchange.getAttributes().put(REQUEST_TIME_BEGIN, System.currentTimeMillis());
            return chain.filter(exchange).then(
                    Mono.fromRunnable(() -> {
                        Long startTime = exchange.getAttribute(REQUEST_TIME_BEGIN);
                        if (startTime != null) {
                            StringBuilder sb = new StringBuilder(exchange.getRequest().getURI().getRawPath())
                                    .append(": ")
                                    .append(System.currentTimeMillis() - startTime)
                                    .append("ms");
                            if (config.isWithParams()) {
                                sb.append(" params:").append(exchange.getRequest().getQueryParams());
                            }
                            log.debug(sb.toString());
                        }
                    })
            );
        };
    }
    
    public static class Config {
        private boolean withParams;

        public boolean isWithParams() {
            return withParams;
        }

        public void setWithParams(boolean withParams) {
            this.withParams = withParams;
        }

    }
}
```

----
