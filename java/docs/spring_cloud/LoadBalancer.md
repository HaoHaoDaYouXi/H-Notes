# Spring Cloud LoadBalancer

`Spring Cloud LoadBalancer`是由`SpringCloud`官方提供的一个开源的、简单易用的客户端负载均衡器，
它包含在`SpringCloud-commons`中用它来替换了以前的Ribbon组件。
相比较于`Ribbon`，`Spring Cloud LoadBalancer`不仅能够支持`RestTemplate`，还支持`WebClient`（`WeClient`是`Spring Web Flux`中提供的功能，可以实现响应式异步请求）

官网：https://docs.spring.io/spring-cloud-commons/reference/spring-cloud-commons/loadbalancer.html

# <a id="nginxqb">和`Nginx`的区别</a>
- `Nginx`是服务器负载均衡，客户端所有请求都会交给`Nginx`，然后由`Nginx`实现转发请求，即负载均衡是由服务端实现的。
- `loadbalancer`本地负载均衡，在调用微服务接口时候，会在注册中心上获取注册信息服务列表之后缓存到`JVM`本地，从而在本地实现`RPC`远程服务调用技术。

负载均衡算法也是常见的：轮询、随机、最小活跃数、源地址哈希、一致性哈希。

## <a id="pzsy">配置使用</a>

如果是`Hoxton`之前的版本，默认负载均衡器为`Ribbon`，需要移除`Ribbon`引用和增加配置

对应版本：

| Spring Cloud Alibaba | 	Spring cloud	          | Spring Boot    |
|----------------------|-------------------------|----------------|
| 2.2.6.RELEASE	       | Spring Cloud Hoxton.SR9 | 	2.3.2.RELEASE |

如果不移除，也可以在`yml`中配置不使用`Ribbon`
```yaml
spring:
  cloud:
    loadbalancer:
      ribbon:
        enabled: false
```

移除`Ribbon`依赖，增加`loadBalance`依赖

```xml
<dependencies>
    <!--nacos-->
    <dependency>
        <groupId>com.alibaba.cloud</groupId>
        <artifactId>spring-cloud-starter-alibaba-nacos-discovery</artifactId>
        <exclusions>
            <!--排除ribbon-->
            <exclusion>
                <groupId>org.springframework.cloud</groupId>
                <artifactId>spring-cloud-starter-netflix-ribbon</artifactId>
            </exclusion>
        </exclusions>
    </dependency>
     
    <!--添加loadbalanncer依赖, 添加spring-cloud的依赖-->
    <dependency>
        <groupId>org.springframework.cloud</groupId>
        <artifactId>spring-cloud-starter-loadbalancer</artifactId>
    </dependency>
</dependencies>
```
默认情况下，如果同时拥有`RibbonLoadBalancerClient`和`BlockingLoadBalancerClient`，为了保持兼容性，将使用`RibbonLoadBalancerClient`。

配置文件的参数也调整了：https://docs.spring.io/spring-cloud-commons/reference/spring-cloud-commons/loadbalancer.html#configuring-individual-loadbalancerclients
```yaml
spring:
  cloud:
    loadbalancer:
      health-check:
        initial-delay: 1s
      clients:
        myclient:
          health-check:
            interval: 30s
```

## <a id="zdy">自定义负载均衡器</a>

自定义随机负载均衡器
```java
public class CustomRandomLoadBalancerClient implements ReactorServiceInstanceLoadBalancer {
 
    // 服务列表
    private ObjectProvider<ServiceInstanceListSupplier> serviceInstanceListSupplierProvider;
 
    public CustomRandomLoadBalancerClient(ObjectProvider<ServiceInstanceListSupplier> serviceInstanceListSupplierProvider) {
        this.serviceInstanceListSupplierProvider = serviceInstanceListSupplierProvider;
    }
 
    @Override
    public Mono<Response<ServiceInstance>> choose(Request request) {
        ServiceInstanceListSupplier supplier = serviceInstanceListSupplierProvider.getIfAvailable();
        return supplier.get().next().map(this::getInstanceResponse);
    }
 
    /**
     * 使用随机数获取服务
     * @param instances
     * @return
     */
    private Response<ServiceInstance> getInstanceResponse(List<ServiceInstance> instances) {
        if (instances.isEmpty()) {
            return new EmptyResponse();
        }
        // 随机算法
        int size = instances.size();
        Random random = new Random();
        ServiceInstance instance = instances.get(random.nextInt(size));
        return new DefaultResponse(instance);
    }
}
```

配置使用
```java
// 设置全局负载均衡器
@LoadBalancerClients(defaultConfiguration = {CustomRandomLoadBalancerClient.class})
// 指定具体服务用某个负载均衡
//@LoadBalancerClient(name = "test",configuration = CustomRandomLoadBalancerClient.class)
//@LoadBalancerClients(
//        value = {
//                @LoadBalancerClient(value = "test",configuration = CustomRandomLoadBalancerClient.class)
//        },defaultConfiguration = LoadBalancerClientConfiguration.class
//)
public class TestConfiguration {
    @Bean
    @LoadBalanced
    public RestTemplate restTemplate(){
        return new RestTemplate();
    }
}
```

这里说下`@LoadBalanced`注解的使用：它是与`@Bean`一起作用在`RestTemplate`上的
```java
@Bean
@LoadBalanced
public RestTemplate restTemplate(){
   return new RestTemplate();
}
```
是为了实现，在注入`RestTemplate`对象到`Spring IoC`容器的同时，启用`Spring`的负载均衡机制。

### `@LoadBalanced`注解的原理
在`LoadBalancerAutoConfiguration`初始化的过程中，创建拦截器`LoadBalancerInterceptor`，对请求进行拦截从而实现负载均衡。

`LoadBalancerInterceptor`拦截器在执行请求前调用其`intercept`方法，`intercept`负责负载均衡的实现

## <a id="csjz">重试机制</a>

```yaml
spring:
  cloud:
    loadbalancer:
      clients:
        # default 表示去全局配置，如要针对某个服务，写对应的服务名称
        default:
          retry:
            enbled: true
            # 是否有的的请求都重试，false表示只有GET请求才重试
            retryOnAllOperation: true
            # 同一个实例的重试次数，不包括第一次调用：比如第填写3 ，实际会调用4次
            maxRetriesOnSameServiceInstance: 3
            # 其他实例的重试次数，多节点情况下使用
            maxRetriesOnNextServiceInstance: 0
```

----
