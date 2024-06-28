# `Bean`
`Spring Bean`是由`Spring IoC`容器管理的对象实例，也是`Spring`框架的基本组件之一。

`Bean`可以是任何一个普通的`Java`对象，也可以是第三方库中的对象，例如`Hibernate SessionFactory`或`MyBatis SqlSessionFactory`。

`Spring Bean`的创建、组装和管理是由`Spring IoC`容器负责的。在容器中注册一个`Bean`后，容器负责创建`Bean`的实例、管理`Bean`的生命周期，以及处理`Bean`之间的依赖关系。通过 Spring 容器，可以实现对象之间的松耦合，便于测试、模块化开发、重用等。

在`Spring`中，`Bean`是通过配置文件或注解来定义的。
配置文件通常是`XML`或`Java`配置类，通过声明`Bean`的类名、作用域、依赖关系、属性值等信息来定义`Bean`。
注解则是通过在`Bean`类上添加特定的注解来定义`Bean`。
无论是`XML`配置文件还是注解，都需要被`Spring IoC`容器加载和解析，以创建`Bean`的实例并放入容器中。

## <div id="component">`@Component`</div>

在运行时，`Spring`会找到所有使用`@Component`或其派生类进行注释的类，并将它们用作`bean`定义。
查找带注释的类的过程称为组件扫描。

`@Conpontent`衍生物是`Spring`构造型注释，它们本身用`@Component`注释。

`@Component`衍生列表包括：
- `@Service`
- `@Repository`
- `@Controller`

注释之间的区别纯粹是信息性的。它们允许你根据通用职责轻松对`bean`进行分类。
你可以使用这些注释将`bean`类标记为特定应用程序层的成员，`Spring`框架会将它们全部视为`@Components`。

## <div id="bean_zyy">Bean的作用域</div>
`Spring`中`Bean`的作⽤域通常有：
- `singleton`：单例，这种`bean`范围是默认的，这种范围确保不管接受到多少个请求，每个容器中只有一个`bean`的实例，单例的模式由`BeanFactory`自身来维护。
- `prototype`：原形，范围与单例范围相反，每次获取都会创建⼀个新的`bean`实例，连续`getBean()`两次，是不同的`Bean`实例。
- `request`(`Web`应⽤使⽤)：请求，在请求`bean`范围内会每一个来自客户端的网络请求创建一个实例，在请求完成以后，`bean`会失效并被垃圾回收器回收
- `session`(`Web`应⽤使⽤)：会话，与`request`范围类似，确保每个`session`中有一个`bean`的实例，在`session`过期后，`bean`会随之失效。
- `application/global-session`(`Web`应⽤使⽤)：全局会话，在一个全局的`Http Session`中，容器会返回该`bean`的同一个实例，仅在使用`portlet context`时有效。
- `websocket`(`Web`应⽤使⽤)：网络通信，在`WebSocket`会话范围内会创建⼀个新的`bean`。

## <div id="bean_smzq">Bean的生命周期</div>

`Spring`上下文中的`Bean`生命周期如下：

- 实例化`Bean`：
  - 对于`BeanFactory`容器，当客户向容器请求一个尚未初始化的`bean`时，或初始化`bean`的时候需要注入另一个尚未初始化的依赖时，容器就会调用`createBean`进行实例化。
  - 对于`ApplicationContext`容器，当容器启动结束后，通过获取`BeanDefinition`对象中的信息，实例化所有的`bean`。
- 设置对象属性（依赖注入）：实例化后的对象被封装在`BeanWrapper`对象中，紧接着，`Spring`根据`BeanDefinition`中的信息以及通过`BeanWrapper`提供的设置属性的接口完成依赖注入。
- 处理`Aware`接口：
  - `Spring`会检测该对象是否实现了`xxxAware`接口，并将相关的`xxxAware`实例注入给`Bean`：
    - 如果这个`Bean`已经实现了`BeanNameAware`接口，会调用它实现的`setBeanName(String beanId)`方法，此处传递的就是`Spring`配置文件中`Bean`的`id`值；
    - 如果这个`Bean`已经实现了`BeanFactoryAware`接口，会调用它实现的`setBeanFactory()`方法，传递的是`Spring`工厂自身。
    - 如果这个`Bean`已经实现了`ApplicationContextAware`接口，会调用`setApplicationContext(ApplicationContext)`方法，传入`Spring`上下文；
- `BeanPostProcessor`：如果想对`Bean`进行一些自定义的处理，那么可以让`Bean`实现了`BeanPostProcessor`接口，那将会调用`postProcessBeforeInitialization(Object obj, String s)`方法。
- `InitializingBean`与`init-method`：
  - 如果`Bean`在`Spring`配置文件中配置了`init-method`属性，则会自动调用其配置的初始化方法。
  - 如果这个`Bean`实现了`BeanPostProcessor`接口，将会调用`postProcessAfterInitialization(Object obj, String s)`方法；
  由于这个方法是在`Bean`初始化结束时调用的，所以可以被应用于内存或缓存技术

以上几个步骤完成后，`Bean`就已经被正确创建了，之后就可以使用这个`Bean`了。

- `DisposableBean`：当`Bean`不再需要时，会经过清理阶段，如果`Bean`实现了`DisposableBean`这个接口，会调用其实现的`destroy()`方法；
- `destroy-method`：如果这个`Bean`的`Spring`配置中配置了`destroy-method`属性，会自动调用其配置的销毁方法。

## <div id="bean_zrfs">Bean注入的实现方式</div>
早期的开发基本是基于`xml`的配置，目前大部分都是基于注解的配置

### 基于`xml`注入`bean`

**构造器注入**

```java
/*带参数，方便利用构造器进行注入*/
public ADaoImpl(String msg){
    this.msg = msg;
}
```
xml配置
```xml
<bean id="ADaoImpl" class="com.ADaoImpl">
    <constructor-arg value="msg"></constructor-arg>
</bean>
```

**`setter`方法注入**

```java
public class AId {
    private int id;
    public int getId() { return id; }
    public void setId(int id) { this.id = id; }
} 
```
xml配置
```xml
<bean id="AId" class="com.AId">
    <property name="id" value="111"></property>
</bean>
```

**静态工厂注入**

静态工厂顾名思义，就是通过调用静态工厂的方法来获取自己需要的对象，为了让`spring`管理所有对象，
我们不能直接通过`工程类.静态方法()`来获取对象，而是依然通过`spring`注入的形式获取：

```java
public class DaoFactory { //静态工厂
    public static final FactoryDao getStaticFactoryDaoImpl(){
        return new StaticFacotryDaoImpl();
    }
}
public class SpringAction {
    private FactoryDao staticFactoryDao; //注入对象
    //注入对象的 set 方法
    public void setStaticFactoryDao(FactoryDao staticFactoryDao) {
        this.staticFactoryDao = staticFactoryDao;
    }
}
```
xml配置
```xml
<!--factory-method="getStaticFactoryDaoImpl"指定调用哪个工厂方法-->
<bean name="springAction" class=" SpringAction" >
    <!--使用静态工厂的方法注入对象,对应下面的配置文件-->
    <property name="staticFactoryDao" ref="staticFactoryDao"></property>
</bean>

<!--此处获取对象的方式是从工厂类中获取静态方法-->
<bean name="staticFactoryDao" class="DaoFactory" 
      factory-method="getStaticFactoryDaoImpl"></bean>
```

**实例工厂**

实例工厂的意思是获取对象实例的方法不是静态的，所以你需要首先`new`工厂类，再调用普通的实例方法：

```java
 public class DaoFactory { //实例工厂 
    public FactoryDao getFactoryDaoImpl(){
        return new FactoryDaoImpl();
    }
}
public class SpringAction {
    private FactoryDao factoryDao; //注入对象 
    public void setFactoryDao(FactoryDao factoryDao) {
        this.factoryDao = factoryDao;
    }
} 
```
xml配置
```xml
<bean name="springAction" class="SpringAction">
    <!--使用实例工厂的方法注入对象,对应下面的配置文件-->
    <property name="factoryDao" ref="factoryDao"></property>
</bean>

<!--此处获取对象的方式是从工厂类中获取实例方法-->
<bean name="daoFactory" class="com.DaoFactory"></bean>
<bean name="factoryDao" factory-bean="daoFactory"
      factory-method="getFactoryDaoImpl"></bean>
```

### 基于注解注入`bean`

#### **声明`bean`**

- `@Component`：通⽤的注解，可标注任意类为`Spring`组件。如果⼀个`Bean`不清楚属于哪一层，可以使⽤`@Component`注解标注。
- `@Repository`: 对应持久层即`Dao`层，主要⽤于数据库相关操作。
- `@Service`: 对应服务层，主要涉及⼀些复杂的逻辑，需要⽤到`Dao`层。
- `@Controller`: 对应`Spring MVC`控制层，主要⽤户接受⽤户请求并调⽤`Service`层返回数据给前端。

##### `@Component`和`@Bean`的区别
- `@Component`注解作⽤于类，⽽`@Bean`注解作⽤于⽅法。
- `@Component`通常是通过类路径扫描来⾃动侦测以及⾃动装配到`Spring`容器中（我们可以使⽤
`@ComponentScan`注解定义要扫描的路径从中找出标识了需要装配的类⾃动装配到`Spring`的`bean`容器中）。
`@Bean`注解通常是我们在标有该注解的⽅法中定义产⽣这个`bean`, `@Bean`告诉了`Spring`这是某个类的实例，当我需要⽤它的时候还给我。
- `@Bean`注解⽐`@Component`注解的⾃定义性更强，⽽且很多地⽅我们只能通过`@Bean`注解来注册`bean`。
⽐如当我们引⽤第三⽅库中的类需要装配到`Spring`容器时，则只能通过`@Bean`来实现。


#### **使用`Bean`**

`Spring`内置的`@Autowired`以及`JDK`内置的`@Resource`和`@Inject`都可以⽤于注⼊`Bean`。

一般都是使用`@Autowired`和`@Resource`

[`@Autowired`和`@Resource`的区别](spring.md#autowired_resource)

----
