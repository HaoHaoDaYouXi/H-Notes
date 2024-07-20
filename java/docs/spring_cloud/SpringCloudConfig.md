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

## 配置使用

### 搭建服务端

`Maven`依赖
```xml
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-config-server</artifactId>
</dependency>
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
浏览器访问git文件：http://localhost:9000/application-dev.yml

### 修改客户端程序

`Maven`依赖
```xml
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-config</artifactId>
</dependency>
```

`application.yml`改为`bootstrap.yml`
```yaml
spring:
  cloud:
    config:
      name: application # 应用名称，对应git中配置文件名称的前半部分
      profile: dev # 开发环境
      label: master # git中的分支
      uri: http://localhost:9000 # config-server的请求地址
```

启动，可成功访问相关服务

### 手动刷新

已经在客户端取到了配置中心的值，但当修改GitHub上面的值时，服务端（`Config Server`）能实时获取最新的值，
但客户端（`Config Client`）读的是缓存，无法实时获取最新值。

`Spring Cloud`已经为我们解决了这个问题，那就是客户端使用`post`去触发`refresh`，获取最新数据，需要依赖`springboot-starter-actuator`

`Maven`依赖
```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-actuator</artifactId>
</dependency>
```

对应的`controller`类加上`@RefreshScope`

```java
@RefreshScope
@RestController
@RequestMapping("/test")
public class TestController {
    @Resource
    private TestService productService;

    @Value("${spring.cloud.client.ip-address}") // spring cloud 自动的获取当前应用的ip地址
    private String ip;
  
    @Value("${server.port}")
    private String port;
  
    @Value("${test}")
    private String test;
  
    @GetMapping(value = "/test")
    public String test() {
        return "访问的服务地址：" + ip + ":" + port+"，[test]："+test;
    }
}
```

客户端配置文件添加端点：
```yaml
management:
  endpoints:
    web:
      exposure:
        include: refresh
```

发送`post`请求：http://localhost:9001/actuator/refresh

## 配置中心的高可用

上面的客户端都是直接调用配置中心的服务端来获取配置文件信息。
这样存在一个问题，客户端和服务端的耦合性太高

如果服务端要做集群，客户端只能通过原始的方式来路由，服务端端改变IP地址的时候，
客户端也需要修改配置，不符合`Spring Cloud`服务治理的理念。

`Spring Cloud`提供了这样的解决方案，我们只需要将服务端端当做一个服务注册到`eureka`，
客户端去`eureka`获取配置中心服务端的服务既可。

### 服务端改造：
