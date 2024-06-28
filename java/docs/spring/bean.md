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







----
