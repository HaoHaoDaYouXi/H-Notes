# Apollo

`Apollo`配置中心是携程开源的一款分布式配置管理中心，是一种可以集中管理应用配置的服务，它能够让配置变得灵活且易于维护。

`Apollo`配置中心对于大型分布式系统来说是一个强大的工具，它简化了配置管理，提高了运维效率，降低了因配置错误导致的故障风险。

## `Apollo`配置中心的一些特点

- 集中化管理：`Apollo`允许你在中心化的平台上管理所有微服务的配置，无论它们部署在何处。这包括不同环境（如开发、测试、生产）和不同集群的配置。
- 实时推送：当配置发生变化时，`Apollo`能够实时地将更新推送到客户端应用，无需重启应用即可生效。
- 权限与流程治理：`Apollo`提供了完整的权限管理系统和流程控制，确保只有授权用户可以修改配置，并且可以进行版本控制和回滚。
- 灰度发布：支持配置的灰度发布，允许逐步推广新配置到部分实例，以减少风险。
- 多环境支持：能够区分并管理多个环境的配置，每个环境可以有不同的配置版本。
- 高可用性：`Apollo`设计为高可用系统，即使部分节点失败，也能保证配置的正常访问。
- 缓存机制：客户端会缓存配置，即使与`Apollo`服务端通信中断，应用仍能继续使用已缓存的配置。
- 易于集成：`Apollo`提供了`Java SDK`和其他语言的客户端库，便于微服务应用集成。
- 可扩展性：`Apollo`的设计考虑了系统的可扩展性，能够处理大规模的配置管理和高并发请求。


## 使用`Apollo`配置中心

- 部署`Apollo`服务端：`Apollo`服务端基于`Spring Boot`和`Spring Cloud`开发，可以打包成独立的可执行`jar`包直接运行。
- 集成`Apollo`客户端：在你的微服务应用中引入`Apollo`的客户端库，根据文档进行配置，以便应用可以从`Apollo`获取配置。
- 配置管理：在`Apollo`控制台上创建项目，定义环境，上传配置，并管理这些配置的版本和变更。
- 监控与维护：定期检查`Apollo`的健康状态，监控配置变更的日志，以及处理可能出现的问题。

### 部署`Apollo`服务端

#### 安装`jdk`和`mysql`
`Apollo`服务需要`Java`和`MySQL`的支持。安装完成后启动`mysql`

#### 下载并启动`Apollo`
- 创建`Apollo`所需的数据库和表。
    - `Apollo`的`GitHub`仓库中找到`SQL`脚本
    - 项目地址：https://github.com/ctripcorp/apollo.git
- 构建`Apollo`
  - 在`build.sh`中配置了数据库的连接
      - 修该数据库连接信息（`host`、用户、密码等配置信息）
      - 其他配置信息可自行修改。
- 初始化数据库
    - 在`script`下存在`apolloconfigdb.sql`和`apolloportaldb.sql`文件，将这2个文件导入到数据库中。
- 运行`Apollo`服务
    - 运行`Apollo`配置中心通常需要运行三个`jar`文件，分别对应`Apollo`的三个主要服务：
        - `apollo-configservice.jar`：`Apollo`的配置服务，为`Apollo`客户端提供配置信息，端口默认是`8080`
        - `apollo-adminservice.jar`：`Apollo`的管理服务，为`Apollo`管理界面提供后端服务，端口默认是`8090`
        - `apollo-portal.jar`：`Apollo`的门户服务，提供`Apollo`的管理界面，端口默认是`8070`

#### 访问`Apollo`服务

访问`http://localhost:8070`来查看`Apollo`的控制台。登录名密码默认为：`apollo/admin`

访问`http://localhost:8080`可查看`Apollo`的服务信息。

登录名和密码可以在`README.md`中查看到。

### `Apollo`中创建项目
在`Apollo`配置中心中，`部门`（Department）是一个组织结构的概念，用于管理和控制不同的项目和配置。
部门可以包含多个项目，每个项目可以有多个环境和集群。

- 创建`department`
  - 可以在`Admin`->`Tools-System`->`Configuration`下配置部门信息，修改完信息后点击下`Reset`按钮。
- 创建项目
- 添加配置参数

### Apollo客户端

#### `Maven`依赖
```xml
<!-- 在项目的pom.xml中添加Apollo客户端依赖 -->
<dependency>
    <groupId>com.ctrip.framework.apollo</groupId>
    <artifactId>apollo-client</artifactId>
</dependency>
```

#### 在`yml`中配置`Apollo`信息
```yaml
app:
  id: app-id
apollo:
  meta: http://localhost:8080
  bootstrap:
    enabled: true
    namespaces: application
```

#### 在代码中使用`@Value`注解获取配置：
```java
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

@Component
public class YourComponent {

    @Value("${some.key:default}")
    private String someKey;
 
    // ...
}
```

启动应用程序，`Apollo`客户端会自动从配置中心拉取配置并注入到应用中。


---- 
