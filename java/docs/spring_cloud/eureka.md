# Eureka
`Eureka`是`Netflix`开发的服务发现组件，本身是一个基于`REST`的服务。
`Spring Cloud`将它集成在其子项目`spring-cloud-netflix`中， 以实现`Spring Cloud`的服务发现功能。
由于基于`REST`服务，这个服务一定会有心跳检测、健康检查和客户端缓存等机制。

`Eureka`包括两个端：
- `Eureka Server`：注册中心服务端，用于维护和管理注册服务列表。
- `Eureka Client`：注册中心客户端，向注册中心注册服务的应用都可以叫做`Eureka Client`（包括`Eureka Server`本身）。

## <a id="server">Eureka Server</a>
- 依赖
  - `spring-cloud-starter-netflix-eureka-server` `Eureka`服务端的标识，标志着此服务是做为注册中心
- 配置（`application.properties`）
  - ```yaml
    server:
      port: 8080
    spring:
      application:
        name: eureka-server # 服务名称
    eureka:
        client:
          register-with-eureka: false # 自身不做为服务注册到注册中心
          fetch-registry: false # 从注册表拉取信息
          serviceUrl:
            defaultZone: 'http://localhost:8080/eureka/' # 服务注册地址   
    ```
- 运行服务`eureka-server`
- 运行成功后访问`localhost:8080`，会显示`eureka`提供的服务页面

## <a id="client">Eureka Client</a>
- 依赖
  - `spring-cloud-starter-netflix-eureka-client` `eureka`客户端所需依赖。
  - `spring-boot-starter-web`等等其他`web`项目需要的依赖
- 配置（`application.properties`）
  - ```yaml
    server:
      port: 8081
    spring:
      application:
        name: eureka-client-8081
    eureka:
        client:
          serviceUrl:
            defaultZone: 'http://localhost:8080/eureka/'
    ```
- 运行服务`eureka-client-8081`
- 运行成功后访问`localhost:8080`，`eureka`服务页面里面就会有`eureka-client-8081`服务
