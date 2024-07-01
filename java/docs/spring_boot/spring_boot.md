# `Spring Boot`

`Spring Boot` 是由`Pivotal`团队提供的全新框架，其设计目的是用来简化新`Spring`应用的初始搭建以及开发过程。
该框架使用了特定的方式来进行配置，从而使开发人员不再需要定义样板化的配置。
通过这种方式，`Spring Boot`致力于在蓬勃发展的快速应用开发领域(`rapid application development`)成为领导者。 

### <div id="td">其特点如下：</div>
- 创建独立的`Spring`应用程序
- 嵌入的`Tomcat`，无需部署`WAR`文件
- 简化`Maven`配置
- 自动配置`Spring`
- 提供生产就绪型功能，如指标，健康检查和外部配置
- 没有代码生成和对`XML`没有要求配置

### <div id="ydyypz">`Spring Boot`约定优于配置</div>
`Spring Boot Starter`、`Spring Boot Jpa`都是**约定优于配置**的一种体现。
都是通过**约定优于配置**的设计思路来设计的，`Spring Boot Starter`在启动的过程中会根据约定的信息对资源进行初始化；
`Spring Boot Jpa`通过约定的方式来自动生成`Sql`，避免大量无效代码编写。

### <div id="cshhjbl">`Spring Boot`初始化环境变量</div>
- 调用`prepareEnvironment`方法去设置环境变量
- `getOrCreateEnvironment`去初始化系统环境变量
- `configureEnvironment`去初始化命令行参数
- `environmentPrepared`当广播到来的时候调用`onApplicationEnvironmentPreparedEvent`方法
  - 去使用`postProcessEnvironment`方法`load yml`和`properties`变量

----


----
