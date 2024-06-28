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




----
