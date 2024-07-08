# [SpringCloud](https://spring.io/projects/spring-cloud)

Spring Cloud 为开发人员提供了一整套的开发框架

具体可以查看：https://spring.io/projects/spring-cloud

开源地址：https://github.com/spring-cloud

里面有我们熟知的：服务注册与发现、路由、熔断器、配置中心、消息总线等等项目

## SpringBoot和SpringCloud的区别

`Spring Boot`和`Spring Cloud`在多个方面存在显著差异：

- 作用：
  - `Spring Cloud`是一个综合管理框架，用于给微服务提供一个综合管理框架。
  - `Spring Boot`主要的作用是为微服务开发提供一种快速的方式，简化配置文件，提高工作效率。
- 使用方式：
  - `Spring Cloud`必须在`Spring Boot`使用的前提下才能使用。
  - `Spring Boot`可以单独使用，
- 创作初衷：
  - `Spring Cloud`的设计目的是为了管理同一项目中的各项微服务。
  - `Spring Boot`的设计目的是为了在微服务开发过程中可以简化配置文件，提高工作效率，
- 目的：
  - `Spring Cloud`的目标是建立一个`有生态系统`的框架，这个框架涵盖了微服务的各个方面，
  - `Spring Boot`的目标是简化`Spring`应用的初始搭建以及开发过程。
- 集成性：
  - `Spring Cloud`集成了所有的服务治理组件，比如`Eureka`、`OpenFeign`、`Ribbon`等。
  - `Spring Boot`都可以与这些组件一起使用，但并不是必须的。
- 扩展性：
  - `Spring Cloud`是基于`Netflix`的`Eureka`、`Ribbon`、`Hystrix`等组件实现的，这些组件都提供了可扩展的`API`，允许开发者根据需要进行定制。
  - `Spring Boot`则没有这样的组件。
- 复杂性：
  - `Spring Cloud`的功能更丰富，因此相对更复杂。
  - `Spring Boot`则更加简单，更易于上手。
- 社区支持：
  - 尽管两者都得到了广泛的社区支持，但在某些方面，`Spring Boot`可能更受欢迎，因为它简化了开发过程并提供了许多实用的功能。
- 安全性：
  - `Spring Cloud`在安全性方面提供了很多组件，例如`Spring Cloud Security`，这使得它更适合处理敏感数据和需要高度安全性的应用。
- 部署和运维：
  - `Spring Cloud`集成了所有的服务治理组件，因此在部署和运维方面更加方便。
  - `Spring Boot`则需要开发者自行解决这些问题。


----
