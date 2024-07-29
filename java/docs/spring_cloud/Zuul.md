# Zuul

`Zuul`是`Netflix`开源的一个边缘服务框架，主要用于构建`API`网关。
它提供了一种路由请求到后端服务的方式，并且可以添加安全、监控和弹性等功能。
`Zuul`可以作为微服务架构中的一个核心组件，用于处理所有进入和离开系统的请求。
`Zuul`自动集成了`Ribbon`，天生就有负载均衡。

## 主要作用

- 路由：根据请求的`URL`、`HTTP`方法或其他自定义规则将请求转发给正确的微服务。
- 过滤器：可以在请求到达微服务之前或响应返回客户端之前执行一些操作，如身份验证、监控、负载均衡等。
- 容错：通过内置的`Hystrix`客户端实现断路器模式，提高系统的稳定性和响应性。
- 静态响应：在某些情况下，可以直接返回预定义的响应，而无需调用后端服务。

`Zuul`通常与`Spring Cloud`框架结合使用，以构建高度可扩展和健壮的微服务网关。

`Zuul`截止`Spring Cloud`的`H.SR12`版本之后就彻底从官网移除了，假如你这时候还想使用`Zuul`，需要注意版本，`SpringBoot`版本也需要注意，不可以高于`2.3.12.RELEASE`。

## 通信原理

`Zuul`是通过`Servlet`来实现的(`Servlet`会为每个请求绑创建一个线程，而线程上线文切换，内存消耗大)，
`Zuul`通过自定义的`ZuulServlet`（类似于`Spring MVC`的`DispatcherServlet`）来对请求进行控制(一系列过滤器处理Http请求)。

所有的`Request`都要经过`ZuulServlet`的处理，三个核心的方法`preRoute()`，`route()`，`postRoute()`，`Zuul`对`request`处理逻辑都在这三个方法里，`ZuulServlet`交给`ZuulRunner`去执行。
`ZuulRunner`直接将执行逻辑交由`FilterProcessor`处理，`ZuulServlet`、`ZuulRunner`、`FilterProcessor`都是单例。

`FilterProcessor`对`filter`的处理逻辑
- 根据`Type`获取所有输入该`Type`的`filter`，`List<ZuulFilter> list`。
- 遍历该`list`，执行每个`filter`的处理逻辑，`processZuulFilter(ZuulFilter filter)`。
- `RequestContext`对每个`filter`的执行状况进行记录，此处的执行状态主要包括其执行时间、以及执行成功或者失败，若失败则对异常封装后抛出。
- `Zuul`框架对每个`filter`的执行结果都没有太多的处理，上一个`filter`的执行结果没交给下一个将要执行的`filter`，仅记录执行状态，如果执行失败抛出异常并终止执行。

## 过滤器的功能

- 身份验证和安全性：确定每个资源的身份验证要求并拒绝不满足这些要求的请求。
- 洞察和监控：在边缘跟踪有意义的数据和统计数据，以便为我们提供准确的生产视图。
- 动态路由：根据需要动态地将请求路由到不同的后端群集。
- 压力测试：逐渐增加群集的流量以衡量性能。
- 负载分配：为每种类型的请求分配容量并删除超过限制的请求。
- 静态响应处理：直接在边缘构建一些响应，而不是将它们转发到内部集群。

## 生命周期(四类过滤)

- `PRE`：在请求被路由之前调用。可在集群中选择请求的微服务、认证鉴权，限流等。
- `ROUTING`：将请求路由到微服务。可构建发送给微服务的请求，并使用Apache HttpClient或 Ribbon请求微服务，现在也支持OKHTTP。
- `POST`：在路由到微服务以后执行。可在这种过滤中处理逻辑，如收集统计信息和指标、将响应从微服务发送给客户端等。
- `ERROR`：在其他阶段发生错误时执行该过滤器。可做全局异常处理。

## 使用

### 引用依赖
```xml
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-netflix-zuul</artifactId>
</dependency>
```

### 路由配置

`Zuul`在不添加配置的情况下，默认就是允许通过服务名称来调用其他服务的，`Zuul`也可以指定`url`来访问

通过`url`来访问，不通过注册中心来转发请求。
```yaml
zuul:
  routes:
    test:	# 自定义的路由名称
      path: /test/**  # 匹配路径
      url: http://127.0.0.1:8001  # 请求地址

```

#### 配置说明

```yaml
zuul:
  prefix: /haohaodayouxi # 代表的是所有的路由前缀
  ignored-services: "*"   # 关闭具体的服务名称访问，"*" 代表全部。
  ignoredPatterns: /**/system/**   # 指定哪些请求 不允许进行路由。
  routes: # 路由映射配置
    test:
      path: /test/**             # 匹配路径
      serviceId: test-service    # 注册中心的服务名称
      stripPrefix: false         # 针对单个路由是否要用前缀访问的设置，默认是true
```

### 注解启动

```java
@SpringBootApplication
@EnableZuulProxy // 开启网关功能
public class TestZuulApplication {
 
    public static void main(String[] args) {
        SpringApplication.run(TestZuulApplication.class, args);
    }

    /**
     * 配置动态路由规则
     * @return
     */
    @Bean
    public PatternServiceRouteMapper getPatternServiceRouteMapper() {
        return new PatternServiceRouteMapper("(?<name>^.+)", "${name}");
    }
}
```

### 自定义过滤器
```java
@Component
public class TokenFilter extends ZuulFilter {
 
    /**
     * 拦截类型,4种类型 pre route error post
     */
    @Override
    public String filterType() {
        //  FilterConstants.PRE_TYPE;
        //  FilterConstants.ROUTE_TYPE;
        //  FilterConstants.ERROR_TYPE;
        //  FilterConstants.POST_TYPE;
        return FilterConstants.PRE_TYPE;
    }
 
    /**
     * 该过滤器在所有过滤器的执行顺序值，值越小，越前面执行
     */
    @Override
    public int filterOrder() {
        return 0;
    }
 
    /**
     * 是否拦截
     */
    @Override
    public boolean shouldFilter() {
        RequestContext ctx = RequestContext.getCurrentContext();
        // RequestContext ctx = RequestContext.getCurrentContext();
        // ctx.getBoolean("isOk");
        HttpServletRequest request = ctx.getRequest();
        String requestURI = request.getRequestURI();
        // 排除拦截的url
        if (requestURI.equals("/test/token")) {
            return false;
        }
        return true;
    }
 
     /**
      * 过滤器具体的业务逻辑
      */
    @Override
    public Object run() throws ZuulException {
        RequestContext ctx = RequestContext.getCurrentContext();
        HttpServletRequest request = ctx.getRequest();
        String token = request.getParameter("token");
        // ctx.set("isOk",true); // 可以在上下文里面设置一个key，在下一次拦截时，就可以获取到
        if (null == token) {
            ctx.setResponseBody("token is null");
            ctx.setResponseStatusCode(400);
            ctx.setSendZuulResponse(false);
            return null;
        }
        if (!"123456".equals(token)) {
            ctx.setResponseBody("token is error");
            ctx.setResponseStatusCode(400);
            ctx.setSendZuulResponse(false);
            return null;
        }
        ctx.setSendZuulResponse(true);
        return null;
    }
}
```

### 关闭过滤器

```yaml
zuul:
  TokenFilter:
    pre:
      disable: true # 关闭前置过滤器
```

### 超时时间设置

如果使用`@EnableZuulProxy`，则可以使用代理路径上传文件，只要文件很小，它应该可以工作。

对于大文件接口访问慢，这时候需要设置超时时间
```yaml
hystrix:
  command:
    default:
      execution:
        isolation:
          thread:
            timeoutInMilliseconds: 60000
ribbon:
  ConnectTimeout: 3000
  ReadTimeout: 60000
```

如果想通过`Zuul`代理的请求，配置套接字超时和读取超时
```yaml
zuul:
  host:
    connect-timeout-millis: 40000
    socket-timeout-millis: 40000
    connection-request-timeout-millis: 40000
```

----
