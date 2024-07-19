# SpringCloudConfig
`Spring Cloud Config`项目是一个解决分布式系统的配置管理方案，其不依赖于注册中心，是一个独立的配置中心，支持多种存储配置信息形式，
目前主要有`jdbc`、`value`、`native`、`svn`、`git`，其中默认是`git`。

因为采用`native`方式必然决定了每次配置完相关文件后一定要重启`Spring Cloud Config`，所以一般不会采用此方案，

在实际操作中主要的实现方案有以下四种：
- `Spring Cloud Config`结合`Git`实现配置中心方案
- `Spring Cloud Config`结合关系型数据库实现配置中心方案
- `Spring Cloud Config`结合非关系型数据库实现配置中心方案
- `Spring Cloud Config`与`Apollo`配置结合实现界面化配置中心方案

`Spring Cloud Config`服务端特性：
- `HTTP`，为外部配置提供基于资源的`API`（键值对，或者等价的`YAML`内容）
- 属性值的加密和解密（对称加密和非对称加密）
- 通过使用`@EnableConfigServer`在`Spring boot`应用中非常简单的嵌入。
- `Config`客户端的特性（特指`Spring`应用）
- 绑定`Config`服务端，并使用远程的属性源初始化`Spring`环境。
- 属性值的加密和解密（对称加密和非对称加密）

它包含了`Client`和`Server`两个部分
- `server`提供配置文件的存储、以接口的形式将配置文件的内容提供出去
- `client`通过接口获取数据、并依据此数据初始化自己的应用
  Config Server是一个可横向扩展、集中式的配置服务器，它用于集中管理应用程序各个环境下的配置，默认使用Git存储配置文件内容，也可以使用SVN存储，或者是本地文件存储。

## 搭建服务端

Maven配置
```xml 
<dependencies>
    <dependency>
        <groupId>org.springframework.cloud</groupId>
        <artifactId>spring-cloud-config-server</artifactId>
    </dependency>
</dependencies>
```

添加启动类，通过`@EnableConfigServer`注解开启注册中心服务端功能
```java
@SpringBootApplication
@EnableConfigServer
public class ConfigApplication {
  public static void main(String[] args) {
      SpringApplication.run(ConfigApplication.class,args);
  }
}
```

配置`application.yml`
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
          uri: https://gitlab.com/xxx.git # git地址
```
浏览器访问git文件:http://localhost:9000/application-dev.yml
