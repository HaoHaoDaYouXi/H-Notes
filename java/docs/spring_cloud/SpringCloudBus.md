# Spring Cloud Bus

在微服务架构中，通常会使用轻量级的消息代理来构建一个共用的消息主题来连接各个微服务实例，
它广播的消息会被所有在注册中心的微服务实例监听和消费，也称消息总线。

SpringCloud中也有对应的解决方案，SpringCloud Bus 将分布式的节点用轻量的消息代理连接起来，
可以很容易搭建消息总线，配合SpringCloud config 实现微服务应用配置信息的动态更新。

Spring Cloud Bus做配置更新的步骤:
- 提交代码触发post请求给bus/refresh
- server端接收到请求并发送给Spring Cloud Bus
- Spring Cloud bus接到消息并通知给其它客户端
- 其它客户端接收到通知，请求Server端获取最新配置
- 全部客户端均获取到最新的配置

## 消息总线整合配置中心：

引入依赖：
```xml
<dependencies>
    <dependency>
        <groupId>org.springframework.cloud</groupId>
        <artifactId>spring-cloud-starter-config</artifactId>
    </dependency>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-actuator</artifactId>
    </dependency>
    <!--消息总线的依赖-->
    <dependency>
        <groupId>org.springframework.cloud</groupId>
        <artifactId>spring-cloud-bus</artifactId>
    </dependency>
    <dependency>
        <groupId>org.springframework.cloud</groupId>
        <artifactId>spring-cloud-stream-binder-rabbit</artifactId>
    </dependency>
</dependencies>
```

服务端配置：
```yaml
server:
  port: 9000
spring:
  application:
    name: config-server-9000
  cloud:
    config:
      server:
        git:
          uri: https://gitlab.com/xxx.git
  rabbitmq:
    host: 127.0.0.1
    port: 5672
    username: guest
    password: guest
management:
  endpoints:
    web:
      exposure:
        include: bus-refresh
eureka:
  client:
    serviceUrl:
      defaultZone: http://127.0.0.1:8888/eureka/
  instance:
    preferIpAddress: true
    instance-id: ${spring.cloud.client.ip-address}:${server.port}
```

客户端配置：
```yaml
spring:
    cloud:
        config:
        name: client-service
        profile: dev
        label: master
        discovery:
            enabled: true
            service-id: config-server
eureka:
    client:
    service-url:
      defaultZone: http://127.0.0.1:8888/eureka/
    instance:
        prefer-ip-address: true
        instance-id: ${spring.cloud.client.ip-address}:${server.port}
```

`git`配置中添加`rabbitmq`的配置信息

重新启动对应的`eureka-server`，`config-server`，`client-service`。

配置信息刷新后，只需要向配置中心发送对应的请求，即可刷新每个客户端的配置

----
