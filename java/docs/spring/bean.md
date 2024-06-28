# `Bean`
`Spring Bean`是由`Spring IoC`容器管理的对象实例，也是`Spring`框架的基本组件之一。

`Bean`可以是任何一个普通的`Java`对象，也可以是第三方库中的对象，例如`Hibernate SessionFactory`或`MyBatis SqlSessionFactory`。

`Spring Bean`的创建、组装和管理是由`Spring IoC`容器负责的。在容器中注册一个`Bean`后，容器负责创建`Bean`的实例、管理`Bean`的生命周期，以及处理`Bean`之间的依赖关系。通过 Spring 容器，可以实现对象之间的松耦合，便于测试、模块化开发、重用等。

在`Spring`中，`Bean`是通过配置文件或注解来定义的。
配置文件通常是`XML`或`Java`配置类，通过声明`Bean`的类名、作用域、依赖关系、属性值等信息来定义`Bean`。
注解则是通过在`Bean`类上添加特定的注解来定义`Bean`。
无论是`XML`配置文件还是注解，都需要被`Spring IoC`容器加载和解析，以创建`Bean`的实例并放入容器中。







----
