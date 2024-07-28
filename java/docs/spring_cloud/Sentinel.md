# Sentinel

`Sentinel`是阿里巴巴开源的一个用于保护微服务架构的流量控制组件，
主要以流量为切入点，从流量控制、熔断降级、系统负载保护等多个维度来帮助您提升微服务架构的稳定性。

`Sentinel`的设计目标是简单易用，可以无缝集成到`Spring Cloud`、`Dubbo`等微服务框架中，同时也支持独立使用。
它适用于各种微服务架构，无论是单体应用还是复杂的分布式系统，都能提供有效的流量管理和保护策略。

代码仓库：https://github.com/alibaba/Sentinel
文档地址：https://sentinelguard.io/zh-cn/docs/introduction.html

## `Sentinel`代码模块

- `sentinel-adapter`：适配器模块，主要实现了对一些常见框架的适配
- `sentinel-benchmark`：基准测试模块，对核心代码的精确性提供基准测试
- `sentinel-cluster`：集群流控制模块，默认实现。
  - `sentinel-cluster-common-default`：用于集群传输和功能的通用模块 
  - `sentinel-cluster-client-default`：使用`Netty`作为底层传输库的默认集群客户端模块 
  - `sentinel-cluster-server-default`：默认集群服务器模块
- `sentinel-core`：核心模块，限流、降级、系统保护等都在这里实现
- `sentinel-dashboard`：控制台模块，可以对连接上的`sentinel`客户端实现可视化的管理
- `sentinel-demo`：示例模块，可参考怎么使用`sentinel`进行限流、降级等
- `sentinel-extension`：扩展模块，主要对`DataSource`进行了部分扩展实现
- `sentinel-logging`：日志模块，提供了日志输出的接口，以及一些默认实现
- `sentinel-transport`：传输模块，提供了基本的监控服务端和客户端的API接口，以及一些基于不同库的实现

## `Sentinel`提供了以下核心功能

- 流量控制: `Sentinel`可以基于资源、`URL`或者自定义规则进行限流，支持`QPS`和线程数两种限流方式，以及冷启动和预热机制。
- 熔断降级: 当下游服务出现异常或者响应时间过长时，`Sentinel`可以进行熔断，快速返回错误信息，避免雪崩效应。
- 系统保护: `Sentinel`可以根据系统的负载情况，比如线程池使用率、`CPU`使用率、入口`QPS`等指标，对整个系统进行保护，避免系统过载。
- 集群限流: `Sentinel`支持跨应用的集群限流，可以确保整个集群的稳定运行。
- 动态规则管理: `Sentinel`支持动态调整规则，无需重启应用即可生效，方便运维和管理。
- `Dashboard`: `Sentinel`提供了一个可视化的控制台，可以实时监控服务的运行状态和流量情况，并且可以动态调整规则。
