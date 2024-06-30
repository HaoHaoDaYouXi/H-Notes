# `Spring`加载过程

### 刷新前的预处理：`PrepareRefresh()`

- `initPropertySources()`初始化一些属性设置
- `getEnvironment().validateRequiredProperties();` 检验属性的合法等
- `earlyApplicationEvents = new LinkedHashSet<ApplicationEvent>;` 保存容器中的一些早期时间

### 获取`BeanFactory`：`obtainFreshBeanFactory();`

- `refreshBeanFactory();` 刷新【创建】`BeanFactory`
- `getBeanFactory();` 返回刚才`GenericApplicationContext`创建的`BeanFactory`对象
- 将创建的`BeanFactory`对象【`DefaultListableBeanFactory`】返回

### `BeanFactory`的预备准备工作（`BeanFactory`进行一些设置）：`prepareBeanFactory(beanFactory);`

- 设置`BeanFactory`的类加载器、支持表达式解析器
- 添加部分`BeanPostProcessor`【`ApplicationContextAwareProcessor`】
- 设置忽略的自动装配的接口 `EnvironmentAware`、`EmbeddedValueResolverAware`。。。
- 注册可以解析的自动装配，我们能直接在任何组件中自动注入：`BeanFactory`、`ResourceLoader`、`ApplicationEventPublisher`、`ApplicationContext`
- 添加`BeanPostProcessor`,【`ApplicationListenerDetector`】
- 添加编译时的`AspectJ`支持
- 给`BeanFactory`中注册一些能用的组件：`environment`【`ConfigurableEnvironment`】、`SystemProperties`【`Map<String, Object>`】、　　              `systemEnvironment`【`Map<String, Object>`】

### `BeanFactory`准备工作完成后进行的后置处理工作：`postProcessBeanFactory(beanFactory);`

- 子类通过这个方法在`BeanFactory`创建并预准备完成后做的进一步设置

### 以上是BeanFactory的创建和预准备工作

----

### 执行`BeanFactoryPostProcessor`：`invokeBeanFactoryPostProcessor(beanFactory);`

`BeanFactoryPostProcessor`：`BeanFactory`的后置处理器。在`BeanFactory`标准初始化之后执行的。

两个接口：`BeanFactoryPostProcessor`、`BeanDefinitionRegistryPostProcessor`

- 执行`BeanFactoryPostProcessor`的方法：
  - 先执行`BeanDefinitionRegistryPostProcessor` 
    - 获取所有的`BeanDefinitionRegistryPostProcessor`
    - 先执行实现了`PriorityOrdered`优先级接口的`BeanDefinitionRegistryPostProcessor`
      - 执行方法`postProcessor.postProcessBeanDefinitionRegistry(registry)`
    - 再执行实现了`Ordered`顺序接口的`BeanDefinitionRegistryPostProcessor`
      - 执行方法`postProcessor.postProcessBeanDefinitionRegistry(registry)`
    - 最后执行没有实现任何优先级或者是顺序接口的`BeanDefinitionRegistryPostProcessor`
      - 执行方法`postProcessor.postProcessBeanDefinitionRegistry(registry)`
  - 再执行`BeanFactoryPostProcessor`的方法
    - 获取所有的`BeanFactoryPostProcessor`
    - 先执行实现了`PriorityOrdered`优先级接口的`BeanFactoryPostProcessor`
      - 执行方法`postProcessor.postProcessBeanFactory(registry)`
    - 再执行实现了`Ordered`顺序接口的`BeanFactoryPostProcessor`
      - 执行方法`postProcessor.postProcessBeanFactory(registry)`
    - 最后执行没有实现任何优先级或者是顺序接口的`BeanFactoryPostProcessor`
      - 执行方法`postProcessor.postProcessBeanFactory(registry)`

### 注册`BeanPostProcessor`（`Bean`的后置处理器）：`registerBeanPostProcessor(beanFactory);`
   
不同接口类型的`BeanPostProcessor`：在`Bean`创建前后的执行时机是不一样的

- `BeanPostProcessor`
- `DestructionAwareBeanPostProcessor`
- `InstantiationAwareBeanPostProcessor`
- `SmartInstantiationAwareBeanPostProcessor`
- `MergedBeanDefinitionPostProcessor`

1) 获取所有的`BeanPostProcessor`；后置处理器都默认可以通过`PriorityOrdered`、`Ordered`接口来指定优先级

2) 先注册`PriorityOrdered`优先级接口的`BeanPostProcessor`

   把每一个`BeanPostProcessor`添加到`BeanFactory`中
   
   `beanFactory.addBeanPostProcessor(postProcessor)`

3) 再注册`Ordered`优先级接口的`BeanPostProcessor`

4) 然后再注册没有任何优先级接口的`BeanPostProcessor`

5) 最终注册`MergedBeanDefinitionPostProcessor`

6) 注册一个`ApplicationListenerDetector`：再`Bean`创建完成后检查是否是`ApplicationListener`，如果是则执行

   `applicationContext.addApplicationListener((ApplicationListener<?>) bean)`

### 初始化`MessageSource`组件（做国际化功能；消息绑定，消息解析）：`InitMessageSource();`

1) 获取`BeanFactory`

2) 看容器中是否有`id`为`messageSource`，类型是`MessageSource`的组件

   如果有就赋值给`messageSource`，如果没有就自己创建一个`DelegatingMessageSource`;
   
   `MessageSource`: 取出国际化配置文件中某个`key`的值；能按照区域信息获取；

3) 把创建好的`messageSource`注册到容器中，以后获取国际化配置文件的时候，可以自动注入`MessageSource`，然后可以再调用它的`getMessage`方法　　　　
   `beanFactory.registerSingleton(MESSAGE_SOURCE_BEAN_NAME, this.messageSource)`

### 初始化事件派发器：`initApplicationEventMulticaster();`

1) 获取`BeanFactory`

2) 从`BeanFactory`中获取`applicationEventMulticaster`的`ApplicationEventMulticaster`

3) 如果上一步没有配置，那就会自己创建一个`SimpleApplicationEventMulticaster`，然后将创建的`ApplicationEventMulticaster`组件添加到`BeanFactory`中，以后其他组件可以直接注入

### 留给子容器(子类)：`onRefresh();`

子类重写这个方法，在容器刷新的时候可以自定义逻辑

### 将项目中所有ApplicationListener注册进容器中：`registerListeners();`

1) 从容器中拿到所有的`ApplicationListener`

2) 将每个监听器添加到事件派发器中`getApplicationEventMulticaster().addApplicationListenerBean(listenerBeanName)`

3) 派发之前步骤产生的事件；






----
