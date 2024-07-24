# Dubbo

`Dubbo`是阿里巴巴开源的基于`Java`的高性能`RPC`分布式服务框架，现已成为`Apache`基金会孵化项目。

地址：https://cn.dubbo.apache.org/zh-cn/

## <a id="lct">服务注册与发现的流程图</a>

下图来自`Dubbo`官网

![dubbo_architecture.png](img/dubbo_architecture.png)

## <a id="jdjs">Dubbo节点角色</a>

| 节点        | 角色说明                |
|-----------|---------------------|
| Provider  | 暴露服务的服务提供方          |
| Consumer  | 调用远程服务的服务消费方        |
| Registry  | 服务注册与发现的注册中心        |
| Monitor   | 统计服务的调用次数和调用时间的监控中心 |
| Container | 服务运行容器              |

## <a id="fwzl">`Dubbo` 服务治理</a>
`Dubbo`服务治理是一种服务管理和协调的解决方案，它主要是为分布式系统提供服务管理、服务调度、服务监控、服务负载均衡等功能。

`Dubbo`服务治理可以有效地管理和调度分布式系统中的服务，通过提供丰富的管理工具可以方便地实现服务的监控、调度和负载均衡等功能。
在分布式系统中，`Dubbo`服务治理可以提供一种方式，让不同的应用程序通过调用远程服务实现互联互通。

以下是一个简单的`Dubbo`服务治理的时序图，展示了`Dubbo`服务注册、发现和调用的过程：

![dubbo_fwzl.png](img/dubbo_fwzl.png)

在这个时序图中，`Client`是服务的消费者，`Registry`是服务注册中心，`Provider1`和`Provider2`是服务的提供者。

整个过程分为三个步骤：
- 服务发现：`Client`向`Registry`发起服务发现请求，`Registry`返回可用的服务列表。
- 服务调用：
  - `Client`向`Provider1`发起服务调用请求，`Provider1`返回结果。
  - `Client`向`Provider2`发起服务调用请求，`Provider2`返回结果。
- 结果返回：`Provider1`和`Provider2`返回结果给`Client`。

`Dubbo`服务治理的重要性在于，它可以帮助开发人员管理和协调不同的服务和组件，并确保服务的可用性和可靠性。
通过`Dubbo`服务治理，开发团队可以通过一个单一的入口管理所有服务，这对于大规模分布式服务的管理非常重要。

## <a id="zcxy">`Dubbo`支持的协议</a>

- `dubbo`(推荐使用的协议)：单一长连接和`NIO`异步通讯，适合大并发小数据量的服务调用，以及消费者远大于提供者。传输协议`TCP`，异步，`Hessian`序列化
  - 优点：支持异步通信，性能较高。
  - 缺点：只能在`Java`环境下使用。

- `hessian`：集成`Hessian`服务，基于`HTTP`通讯，采用`Servlet`暴露服务，`Dubbo`内嵌`Jetty`作为服务器时默认实现，提供与`Hessian`服务互操作。
  多个短连接，同步`HTTP`传输，`Hessian`序列化，传入参数较大，提供者大于消费者，提供者压力较大，可传文件
  - 优点：采用二进制序列化，传输效率高。
  - 缺点：只能在`Java`环境下使用

- `rmi`：采用`JDK`标准的`RMI`协议实现，传输参数和返回参数对象需要实现`Serializable`接口，
  使用`java`标准序列化机制，使用阻塞式短连接，传输数据包大小混合，消费者和提供者个数差不多，可传文件，传输协议`TCP`。
  多个短连接，`TCP`协议传输，同步传输，适用常规的远程服务调用和`rmi`互操作。
  在依赖低版本的`Common-Collections`包，`java`序列化存在安全漏洞
  - 优点：使用`JDK`标准的`RMI`协议，易于使用。
  - 缺点：只能在`Java`环境下使用。

- `http`：基于`Http`表单提交的远程调用协议，使用`Spring`的`HttpInvoke`实现。多个短连接，传输协议`HTTP`，传入参数大小混合，
  提供者个数多于消费者，需要给应用程序和浏览器`JS`调用。
  - 优点：支持跨语言调用，使用方便。
  - 缺点：传输效率相对较低。

- `webservice`：基于`WebService`的远程调用协议，集成`CXF`实现，提供和原生`WebService`的互操作。
  多个短连接，基于`HTTP`传输，同步传输，适用系统集成和跨语言调用
  - 优点：采用SOAP协议，支持跨语言调用。
  - 缺点：传输效率相对较低。

- `GRPC`：`GRPC`是谷歌开源的基于`HTTP2`的通信协议，支持多种编程语言，包括`C++`，`Java`，`Python`，`Go`等，适用于各种语言环境下的服务调用。
  - 优点：使用`HTTP2`协议，显著降低带宽消耗和提高性能。
  - 缺点：尚未提供连接池，基于`HTTP2`，绝大部多数`HTTP Server`、`Nginx`都尚不支持。

- `memcache`：基于`memcached`实现的`RPC`协议

- `redis`：基于`redis`实现的`RPC`协议
