# Resilience4j

`Resilience4j`是一个轻量级的故障容忍库，用于实现常见的容错模式，如断路器、重试、缓存、隔离、限流和超时。
它旨在解决分布式系统中的常见问题，如网络延迟、服务不可用或过载，从而提高应用程序的弹性和可靠性。

代码仓库：https://github.com/resilience4j/resilience4j

文档地址：https://resilience4j.readme.io/

## `Resilience4j`功能
- 断路器 (Circuit Breaker): 当依赖的服务出现故障时，断路器可以快速失败并避免进一步的调用，直到服务恢复。
- 重试 (Retry): 允许在检测到特定类型的错误时自动重试操作，以处理瞬态故障。
- 缓存 (Cache): 通过缓存结果来减少对远程服务的请求，提高响应速度。
- 隔离 (Bulkhead): 将不同的服务调用隔离在不同的线程池或信号量中，防止一个服务的故障影响其他服务。
- 限流 (RateLimiter): 控制对依赖服务的请求速率，以防止过载。
- 超时 (TimeLimiter): 设置操作的最大执行时间，超过这个时间将自动失败，以避免长时间等待。

`Resilience4j`的设计目标之一是低开销和易于集成，它不依赖于任何框架或容器，可以独立地添加到任何`Java`应用程序中。

## `Resilience4j`和`Hystrix`的不同

`Hystrix`使用`HystrixCommand`来调用外部的系统，而`R4j`提供了⼀些⾼阶函数，例如断路器、限流器、隔离机制等，这些函数作为装饰器对函数式接⼝、`lambda`表达式、函数引用进⾏装饰。

此外，R4j库还提供了失败重试和缓存调用结果的装饰器。

你可以在函数式接⼝、`lambda`表达式、函数引用上叠加地使用⼀个或多个装饰器，这意味着隔离机制、限流器、重试机制等能够进⾏组合使用。

这么做的优点在于，你可以根据需要选择特定的装饰器。任何被装饰的方法都可以同步或异步执⾏，异步执⾏可以采用`CompletableFuture`或`RxJava`。

当有很多超过规定响应时间的请求时，在远程系统没有响应和引发异常之前，断路器将会开启。

当Hystrix处于半开状态时，`Hystrix`根据只执⾏⼀次请求的结果来决定是否关闭断路器。
而`R4j`允许执⾏可配置次数的请求，将请求的结果和配置的阈值进⾏比较来决定是否关闭断路器。

`R4j`提供了⾃定义的`Reactor`和`Rx Java`操作符对断路器、隔离机制、限流器中任何的反应式类型进⾏装饰。

`Hystrix`和`R4j`都发出⼀个事件流，系统可以对发出的事件进⾏监听，得到相关的执⾏结果和延迟的时间统计数据都是⼗分有用的。


## `Resilience4j`组件

要使用`Resilience4j`，不需要引入所有依赖，只需要选择你需要的

`Resilience4j`提供了以下的核心模块和拓展模块：

| 组件名称                        | 功能              |
|-----------------------------|-----------------|
| resilience4j-circuitbreaker | 	熔断器            |
| resilience4j-ratelimiter	   | 限流器             |
| resilience4j-bulkhead       | 	隔离器--依赖隔离&负载保护 |
| resilience4j-retry          | 	重试、同步&异步       |
| resilience4j-cache	         | 缓存              |
| resilience4j-timelimiter	   | 超时处理            |
| resilience4j-all            | 所有              |

## 舱壁（Bulkhead-隔离）

`Resilience4j`提供了两种舱壁模式的实现，可用于限制并发执行的次数：
- `SemaphoreBulkhead`（信号量舱壁，默认），基于Java并发库中的Semaphore实现。
- `FixedThreadPoolBulkhead`（固定线程池舱壁），它使用一个有界队列和一个固定线程池。

`SemaphoreBulkhead`应该在各种线程和`I` / `O`模型上都能很好地工作。
它基于信号量，与`Hystrix`不同，它不提供“影子”线程池选项。取决于客户端，以确保正确的线程池大小将与舱壁配置保持一致。

### 信号量舱壁（SemaphoreBulkhead）

当信号量存在剩余时进入系统的请求会直接获取信号量并开始业务处理。

当信号量全被占用时，接下来的请求将会进入阻塞状态，`SemaphoreBulkhead`提供了一个阻塞计时器，如果阻塞状态的请求在阻塞计时内无法获取到信号量则系统会拒绝这些请求。

若请求在阻塞计时内获取到了信号量，那将直接获取信号量并执行相应的业务处理。

### 固定线程池舱壁（FixedThreadPoolBulkhead）
`FixedThreadPoolBulkhead`的功能与`SemaphoreBulkhead`一样也是用于限制并发执行的次数的，但是二者的实现原理存在差别而且表现效果也存在细微的差别。

`FixedThreadPoolBulkhead`使用一个固定线程池和一个等待队列来实现舱壁。

当线程池中存在空闲时，则此时进入系统的请求将直接进入线程池开启新线程或使用空闲线程来处理请求。

当线程池无空闲时接下来的请求将进入等待队列，若等待队列仍然无剩余空间时接下来的请求将直接被拒绝。

在队列中的请求等待线程池出现空闲时，将进入线程池进行业务处理。

可以看到FixedThreadPoolBulkhead和SemaphoreBulkhead一个明显的差别是FixedThreadPoolBulkhead没有阻塞的概念，而SemaphoreBulkhead没有一个队列容量的限制。

## 使用

使用上主要就是引用对应功能包、配置对应功能参数、使用对应注解就可以

还可以直接引用全部包
```xml
<dependency>
    <groupId>io.github.resilience4j</groupId>
    <artifactId>resilience4j-all</artifactId>
</dependency>
```

## 断路器（Circuit Breaker）

### `yml`配置

```yaml
resilience4j:
  circuitbreaker:
    configs:
      default:
        failureRateThreshold: 30 #失败请求百分比，超过这个比例，CircuitBreaker变为OPEN状态 
        slidingWindowSize: 10 #滑动窗⼝的⼤小，配置COUNT_BASED,表示10个请求，配置TIME_BASED表示10秒 
        minimumNumberOfCalls: 5 #最小请求个数，只有在滑动窗⼝内，请求个数达到这个个数，才会触发CircuitBreader对于断路器的判断 
        slidingWindowType: TIME_BASED #滑动窗⼝的类型 
        permittedNumberOfCallsInHalfOpenState: 3 #当CircuitBreaker处于HALF_OPEN状态的时候，允许通过的请求个数 
        automaticTransitionFromOpenToHalfOpenEnabled: true #设置true，表示⾃动从OPEN变成HALF_OPEN，即使没有请求过来 
        waitDurationInOpenState: 2s #从OPEN到HALF_OPEN状态需要等待的时间 
        recordExceptions: #异常名单 
          - java.lang.Exception
    instances:
      backendA:
        baseConfig: default #熔断器backendA，继承默认配置default 
      backendB:
        failureRateThreshold: 50
        slowCallDurationThreshold: 2s #慢调用时间阈值，⾼于这个阈值的呼叫视为慢调用，并增加慢调用比例。 
        slowCallRateThreshold: 30 #慢调用百分比阈值，断路器把调用时间⼤于 slowCallDurationThreshold，视为慢调用，当慢调用比例⼤于阈值，断路器打开，并进⾏服务降级
        slidingWindowSize: 10
        slidingWindowType: TIME_BASED
        minimumNumberOfCalls: 2
        permittedNumberOfCallsInHalfOpenState: 2
        waitDurationInOpenState: 2s #从OPEN到HALF_OPEN状态需要等待的时间 
```

### 注解

```java
@GetMapping("/getTest") 
@CircuitBreaker(name = "backendC", fallbackMethod = "fallback") 
public Response<Object> getById(@PathVariable("id") Long id) {
    Thread.sleep(10000L); // 阻塞10秒
    return Response.OK(id);
}

public Response<Object> fallback(Integer id) {
    return Response.ERROR("error："+id);
}
```
- name：对应配置里面的instances的名称，backendA、backendB，如果匹配不到会使用默认的，例如：backendC没有会使用backendA配置
- fallbackMethod：当断路器打开时，会调用该方法，返回值和原方法一致

### 测试过程

服务⽆法调用，所有请求报错，这时第⼀次并发发送20次请求，触发异常比例熔断，
断路器进入打开状态，2s后（waitDurationInOpenState: 2s），
断路器⾃动进入半开状态（automaticTransitionFromOpenToHalfOpenEnabled: true），
再次发送请求断路器处于半开状态，允许3次请求通过（permittedNumberOfCallsInHalfOpenState: 3），

注意此时控制台打印3次⽇志信息，说明半开状态，进入了3次请求调用，接着断路器继续进入打开状态。

慢比例调用熔断测试，修改代码，使用`backendB`熔断器。

第⼀次发送并发发送了20个请求，触发了慢比例熔断，但是因为没有配置（automaticTransitionFromOpenToHalfOpenEnabled: true），
⽆法⾃动从打开状态转为半开状态，需要浏览器中执⾏⼀次请求，这时，断路器才能从打开状态进入半开状态，接下来进入半开状态，
根据配置，允许2次请求在半开状态通过（permittedNumberOfCallsInHalfOpenState: 2）。

## 限流器（Rate Limiter）

### `yml`配置

```yaml
resilience4j: 
    ratelimiter: 
        configs: 
            default: 
                limitRefreshPeriod: 1s # 限流器每隔1s刷新⼀次，将允许处理的最⼤请求重置为2 
                limitForPeriod: 2 #在⼀个刷新周期内，允许执⾏的最⼤请求数 
                timeoutDuration: 5 # 线程等待权限的默认等待时间 
        instances: 
            backendA: 
                baseConfig: default 
            backendB: 
                limitRefreshPeriod: 1s 
                limitForPeriod: 5
                timeoutDuration: 5 
```

### 注解

```java
@GetMapping("/getTest") 
@RateLimiter(name = "backendA", fallbackMethod = "fallback") 
public Response<Object> getById(@PathVariable("id") Long id) {
    Thread.sleep(10000L); // 阻塞10秒
    return Response.OK(id);
}            
```

### 测试过程

因为在⼀个刷新周期1s（limitRefreshPeriod: 1s），允许执⾏的最⼤请求数为 2（limitForPeriod: 2），等待令牌时间5s（timeoutDuration: 5）。

并发发送20个请求后，只有2个请求拿到令牌执⾏，另外2个请求等5秒后拿到令牌，其他16个请求直接降级。

## 隔离（Bulkhead）

### SemaphoreBulkhead

#### `yml`配置

```yaml
resilience4j: 
    bulkhead: 
        configs: 
            default: 
                maxConcurrentCalls: 5 # 隔离允许并发线程执⾏的最⼤数量 
                maxWaitDuration: 20ms # 当达到并发调⽤数量时，新的线程的阻塞时间 
        instances: 
            backendA: 
                baseConfig: default 
            backendB: 
                maxConcurrentCalls: 10
                maxWaitDuration: 10ms
```

#### 注解

```java
@GetMapping("/getTest")
@Bulkhead(name = "backendA", fallbackMethod = "fallback", type = Bulkhead.Type.SEMAPHORE)
public Response<Object> getById(@PathVariable("id") Long id) {
    Thread.sleep(10000L); // 阻塞10秒
    return Response.OK(id);
}            
```
type默认为Bulkhead.Type.SEMAPHORE，表示信号量隔离。

#### 测试过程

因为并发线程数为5（maxConcurrentCalls: 5），只有5个线程进入执⾏，其他请求降直接降级。

### FixedThreadPoolBulkhead

#### `yml`配置

```yaml
resilience4j: 
    thread-pool-bulkhead: 
        configs: 
            default:
                queueCapacity: 2 # 队列容量 
                coreThreadPoolSize: 2 # 核⼼线程池⼤小 
                maxThreadPoolSize: 4 # 最⼤线程池⼤小 
        instances: 
            backendA: 
                baseConfig: default 
            backendB:
                queueCapacity: 1
                coreThreadPoolSize: 1
                maxThreadPoolSize: 1 
```

#### 注解

```java
@GetMapping("/getTest")
public Response<Object> getTest(@PathVariable("id") Long id) {
    return Response.OK(get(id));
}

@Bulkhead(name = "backendA", type = Bulkhead.Type.THREADPOOL)
public CompletableFuture<Object> get(Long id) {
    Thread.sleep(10000L); // 阻塞10秒
    return CompletableFuture.supplyAsync(() -> id);
}
```
`FixedThreadPoolBulkhead`只对`CompletableFuture`方法有效，所以必须返回`CompletableFuture`类型。

#### 测试过程

4个请求进入线程执⾏（maxThreadPoolSize: 4 ），2个请求（queueCapacity: 2）进入有界队列等待，等待10秒后有线程执⾏结束，队列中的线程开始执⾏。

---
