# Spring

`Spring`是⼀款开源的轻量级`Java`开发框架，旨在提⾼开发⼈员的开发效率以及系统的可维护性。

`Spring`框架指的是[`Spring Framework`](https://github.com/spring-projects/spring-framework)

项目地址：https://github.com/spring-projects/spring-framework

官网地址：https://spring.io/


## <div id="spring_td">Spring特点</div>
- 轻量：`Spring`是轻量的，基本的版本大约2MB。
- 控制反转：`Spring`通过控制反转实现了松散耦合，对象们给出它们的依赖，而不是创建或查找依赖的对象们。
- 面向切面的编程(`AOP`)：`Spring`支持面向切面的编程，并且把应用业务逻辑和系统服务分开。
- 容器：`Spring`包含并管理应用中对象的生命周期和配置。
- `MVC`框架：`Spring`的`WEB`框架是个精心设计的框架，是`Web`框架的一个很好的替代品。
- 事务管理：`Spring`提供一个持续的事务管理接口，可以扩展到上至本地事务下至全局事务（`JTA`，即`Java Transaction API`）。
- 异常处理：`Spring`提供方便的`API`把具体技术相关的异常（比如由`JDBC`，`Hibernate` or `JDO`抛出的）转化为一致的`unchecked`异常。

**最重要的一点使用的人多，生态完善！！！**


## <div id="spring_hxzj">Spring的核心组件</div>
目前`Spring`框架已集成了`20`多个模块。这些模块主要被分如下图所示的核心容器、数据访问、集成、`Web`、`AOP`（面向切面编程）、工具、消息和测试模块。

下图取自spring4.3.30版本的：https://docs.spring.io/spring-framework/docs/4.3.30.RELEASE/spring-framework-reference/htmlsingle/#overview-modules

![spring_framework_runtime_4.3.30.png](img/spring_framework_runtime_4.3.30.png)

spring5.2.25版本：https://docs.spring.io/spring-framework/docs/5.2.25.RELEASE/spring-framework-reference/

![spring_framework_doc_5.2.25.png](img/spring_framework_doc_5.2.25.png)

spring6.1.x版本：https://docs.spring.io/spring-framework/reference/6.1-SNAPSHOT/

![spring_framework_doc_6.1.10.png](img/spring_framework_doc_6.1.10.png)

### `Core Container`
`Spring`框架的核⼼模块，也可以说是基础模块，主要提供`IoC`依赖注⼊功能的⽀持，`Spring`其他所有的功能基本都需要依赖于该模块。
- `spring-core`：`Spring`框架基本的核⼼⼯具类。
- `spring-beans`：提供对`bean`的创建、配置和管理等功能的⽀持。
- `spring-context`：提供对国际化、事件传播、资源加载等功能的⽀持。
- `spring-expression`：提供对表达式语⾔（`Spring Expression Language`）`SpEL`的⽀持，只依赖于`core`模块，不依赖于其他模块，可以单独使⽤。

### `AOP`
- `spring-aspects`：该模块为与`AspectJ`的集成提供⽀持。
- `spring-aop`：提供了⾯向切⾯的编程实现。
- `spring-instrument`：提供了为`JVM`添加代理（`agent`）的功能。
  具体来讲，它为`Tomcat`提供了⼀个织⼊代理，能够为`Tomcat`传递类⽂件，就像这些⽂件是被类加载器加载的⼀样。

### `Data Access/Integration`
- `spring-jdbc`：提供了对数据库访问的抽象`JDBC`。不同的数据库都有⾃⼰独⽴的`API`⽤于操作数据库，⽽`ava`程序只需要和`JDBC API`交互，这样就屏蔽了数据库的影响。
- `spring-tx`：提供对事务的⽀持。
- `spring-orm`： 提供对`Hibernate`、`JPA`、`iBatis`等`ORM`框架的⽀持。
- `spring-oxm`：提供⼀个抽象层⽀撑`OXM`(`Object-to-XML-Mapping`)，例如：`JAXB`、`Castor`、`XMLBeans`、`JiBX`和`XStream`等。
- `spring-jms`: 消息服务。⾃`Spring Framework 4.1`以后，它还提供了对`spring-messaging`模块的继承。

### `Spring Web`
- `spring-web`：对`Web`功能的实现提供⼀些最基础的⽀持。
- `spring-webmvc`： 提供对`Spring MVC`的实现。
- `spring-websocket`： 提供了对`WebSocket`的⽀持，`WebSocket`可以让客户端和服务端进⾏双向通信。
- `spring-webflux`：提供对`WebFlux`的⽀持。`WebFlux`是`Spring Framework 5.0`中引⼊的新的响应式框架。与`Spring MVC`不同，它不需要`Servlet API`，是完全异步。

### `Messaging`
- `spring-messaging`是从`Spring4.0`开始新加⼊的⼀个模块，主要职责是为`Spring`框架集成⼀些基础的报⽂传送应⽤。

### Spring Test
- `Spring`团队提倡测试驱动开发（`TDD`）。有了控制反转 (`IoC`)的帮助，单元测试和集成测试变得更简单。
- `Spring`的测试模块对`JUnit`（单元测试框架）、`TestNG`（类似`JUnit`）、`Mockito`（主要⽤来`Mock`对象）、
  `PowerMock`（解决`Mockito`的问题⽐如⽆法模拟`final`，`static`，`private`⽅法）等等常⽤的测试框架⽀持的都比较好。


----
