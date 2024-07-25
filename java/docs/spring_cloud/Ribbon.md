# Ribbon

`Ribbon`是一个在微服务架构中使用的客户端负载均衡器。它提供了一套规则来选择在多个服务实例可用时应使用哪一个实例。
它是早期的微服务架构中的负载均衡组件。

后因为停止维护，`Spring Cloud`2020之后移除了`Ribbon`，直接使用`Spring Cloud LoadBalancer`作为客户端负载均衡组件.

## Ribbon重要接口

| 接口	                | 作用	                  | 默认值                                                                             |
|--------------------|----------------------|---------------------------------------------------------------------------------|
| IClientConfig	     | 读取配置	                | DefaultclientConfigImpl                                                         |
| IRule	             | 负载均衡规则，选择实例	         | ZoneAvoidanceRule                                                               |
| IPing	             | 筛选掉ping不通的实例	        | 默认采用DummyPing实现，该检查策略是一个特殊的实现，实际上它并不会检查实例是否可用，而是始终返回true，默认认为所有服务实例都是可用的.       |
| ServerList<Server> | 	交给Ribbon的实例列表	      | Ribbon: ConfigurationBasedServerList</br> Spring Cloud Alibaba: NacosServerList |
| ServerListFilter	  | 过滤掉不符合条件的实例          | 	ZonePreferenceServerListFilter                                                 |
| ILoadBalancer      | 	Ribbon的入口	          | ZoneAwareLoadBalancer                                                           |
| ServerListUpdater  | 	更新交给Ribbon的List的策略	 | PollingServerListUpdater                                                        |

## Ribbon负载均衡规则

负载均衡的规则都定义在IRule接口中，而IRule有很多不同的实现类

| 规则名称	                      | 特点                                                                                                                                                                                                                                                             |
|----------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| RoundRobinRule	            | 简单轮询服务列表来选择服务器。它是Ribbon默认的负载均衡规则。                                                                                                                                                                                                                              |
| AvailabilityFilteringRule	 | 对以下两种服务器进行忽略： <br/>（1）在默认情况下，这台服务器如果3次连接失败，这台服务器就会被设置为“短路”状态。短路状态将持续30秒，如果再次连接失败，短路的持续时间就会几何级地增加。<br/>（2）并发数过高的服务器。如果一个服务器的并发连接数过高，配置了AvailabilityFilteringRule规则的客户端也会将其忽略。并发连接数的上限，可以由客户端的<clientName>.<clientConfigNameSpace>.ActiveConnectionsLimit属性进行配置。 |
| WeightedResponseTimeRule	  | 为每一个服务器赋予一个权重值。服务器响应时间越长，这个服务器的权重就越小。这个规则会随机选择服务器，这个权重值会影响服务器的选择。                                                                                                                                                                                              |
| ZoneAvoidanceRule（默认是这个）   | 	以区域可用的服务器为基础进行服务器的选择。使用Zone对服务器进行分类，这个Zone可以理解为一个机房、一个机架等。而后再对Zone内的多个服务做轮询。                                                                                                                                                                                  |
| BestAvailableRule	         | 忽略那些短路的服务器，并选择并发数较低的服务器。                                                                                                                                                                                                                                       |
| RandomRule	                | 随机选择一个可用的服务器。                                                                                                                                                                                                                                                  |
| RetryRule                  | 	重试机制的选择逻辑                                                                                                                                                                                                                                                     |

## <a id="zdy">自定义负载均衡策略</a>

通过定义`IRule`实现可以修改负载均衡规则

### 类配置方式
```java
public class RibbonConfiguration {
    @Bean
    public IRule ribbonRule(){
        //随机选择
        return new RandomRule();
    }
}

/**
 * 指定配置
 **/
@Configuration
@RibbonClient(name = "test",configuration = RibbonConfiguration.class)
public class TestRibbonConfiguration {
    @Bean
    @LoadBalanced
    public RestTemplate restTemplate(){
        return new RestTemplate();
    }
}
```

### 配置文件方式
```yaml
spring:
  application:
    name: testService
    
testService: # 给某个微服务配置负载均衡规则，这里是 testService 服务
  ribbon:
    NFLoadBalancerRuleClassName: com.netflix.loadbalancer.RandomRule # 负载均衡规则
```

### 优先级高低
类配置方式 > 配置文件方式

### 全局配置

`RibbonClient`改为`RibbonClients`，`configuration`改为`defaultConfiguration`
```java

/**
 * 指定配置
 **/
@Configuration
@RibbonClients(defaultConfiguration = RibbonConfiguration.class)
public class TestRibbonConfiguration {
    @Bean
    @LoadBalanced
    public RestTemplate restTemplate(){
        return new RestTemplate();
    }
}
```

## <a id="jejz">饥饿加载</a>

`Ribbon`默认是采用懒加载，即第一次访问时才会去创建`LoadBalanceClient`，请求时间会很长。

而饥饿加载则会在项目启动时创建，降低第一次访问的耗时，通过下面配置开启饥饿加载：
```yaml
ribbon:
  eager-load:
    enabled: true # 开启
    clients: testService # 配置 testService 使用饥饿加载，多个使用逗号分隔
```


---
