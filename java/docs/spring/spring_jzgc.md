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

### 初始化所有剩下的单实例`Bean`：`finishBeanFactoryInitialization(beanFactory);`

- `beanFactory.preInstantiateSingletons();`初始化所有剩下的单实例`Bean`
  - 获取容器中的所有`bean`, 依次进行初始化和创建对象
  - 获取`Bean`的定义信息；`RootBeanDefinition`
  - `Bean`不是抽象的，是单实例的，且不是懒加载的，
    - 判断是不是`FactoryBean`；是否是实现`FactoryBean`接口的`Bean`
    - 如果不是`FactoryBean`；使用`getBean(beanName)`创建对象
      - `getBean(beanName)` -> `ioc.getBean();`
      - `doGetBean(name, null, null, false)`
      - 先获取缓存中保存的单实例`Bean`。如果能获取到，说明这个`Bean`之前被创建过（所有创建过的单实例`Bean`都会被缓存起来）从`singletonObjects`中获取
      - 缓存中获取不到，开始`Bean`的创建对象流程；
      - 标记当前`Bean`已经被创建
      - 获取`Bean`的定义信息
      - 获取当前`Bean`依赖的其它`Bean`；如果有，还是按照`getBean()`把依赖的`Bean`先创建出来
      - 启动单实例`Bean`的创建流程
        - `createBean(beanName, mbd, args);`
        - `Object bean = resolveBeforeInstantiation(beanName, mbdToUse);` 
          - 让`BeanPostProcessor`先拦截返回代理对象；
          - `InstantiationAwareBeanPostProcessor`提前执行
          - 先触发：`postProcessBeforeInstantiation();`
          - 如果有返回值；再触发`postProcessAfterInstantiation()`
        - 如果前面的`InstantiationAwareBeanPostProcessor`没有返回代理对象；调用`Object beanInstance = doCreateBean(beanName, mbdToUse, args)`创建`Bean`
          - 创建`Bean`实例，`createBeanInstance(beanName, mbd, args)`利用工厂方法或者对象的构造器创建出`Bean`实例
          - `applyMergedBeanDefinitionPostProcessors(mbd, beanType, beanName)`
            调用`MergedBeanDefinitionPostProcessor``的postProcessMergedBeanDefinition(mbd, beanType, beanName)`
          - 给`Bean`属性赋值，调用`populateBean(beanName, mbd, instanceWrapper)`
            - 赋值之前：
              - 拿到`InstantiationAwareBeanPostProcessor`后置处理器
                - 执行`postProcessAfterInstantiation()`
              - 拿到`InstantiationAwareBeanPostProcessor`后置处理器
                - 执行`postProcessPropertyValues()`
              - 应用`Bean`属性的值：为属性利用`setter`方法进行赋值（反射）
                - `applyPropertyValues(beanName, mbd, bw, pvs)`
          - 初始化`Bean`；`initializeBean(beanName, exposedObject, mbd);`
            - 执行`Aware`接口方法：`invokeAwareMethods(beanName, bean);`执行`xxxAware`接口的方法
              - `BeanNameAware`、`BeanClassLoaderAware`、`BeanFactoryAware`
            - 执行后置处理器初始化之前：`applyBeanPostProcessorsBeforeInitialization(wrappedBean, beanName)`
              - `BeanPostProcessor.postProcessBeforeInitialization()`
            - 执行初始化方法：`invokeInitMethods(beanName, wrappedBean, mbd)`
              - 是否是`InitializingBean`接口的实现：执行接口规定的初始化
              - 是否自定义初始化方法
            - 执行后置处理器初始化之后：`applyBeanPostProcessorsAfterInitialization`
              - `BeanPostProcessor.postProcessAfterInitialization()`
            - 注册`Bean`的销毁方法
          - 将创建的`Bean`添加到缓存中 - `singletonObjects`（`Map`对象）

`IOC`容器就是这些`Map`；很多的`Map`里保存了单实例`Bean`，环境信息、、、

所有`Bean`都利用`getBean`创建完成以后；再来检查所有`Bean`是否是`SmartInitializingSingleton`接口的实现类，

如果是，就执行`afterSingletonsInstantiated();`

### 完成`BeanFactory`的初始化创建工作；`IOC`容器就创建完成：`finishRefresh();`

- `initLifecycleProcessor();` 初始化和生命周期相关的后置处理器；`LifecycleProcessor`
  - 默认从容器中找是否有`lifecycleProcessor`的组件(`LifecycleProcessor`)
  - 如果没有，创建/使用默认的生命周期组件`new DefaultLifecycleProcessor();`再加入到容器中；
  - 写一个`LifecycleProcessor`的实现类，可以在`BeanFactory`的下面两个方法刷新和关闭前后进行拦截调用
    - `onRefresh()`
    - `onClose()`
- `getLifecycleProcessor().onRefresh();`
  - 拿到前面定义的生命周期处理器（`BeanFactory`）；`回调.onRefresh()`;
- `publishEvent(new ContextRefreshedEvent(this));` 发布容器刷新完成时间；
- `liveBeansView.registerApplicationContext();`

## 总结

- `spring`容器在启动的时候，先回保存所有注册进来的`Bean`的定义信息
  - `xml`注册`bean`：`<bean>`
  - 注解注册`Bean`：`@Service`、`@Repository`、`@Component`、`@Bean`、`xxx`
- `Spring`容器会在合适的时机创建这些`Bean`
  - 用到这个`bean`的时候，利用`getBean`创建`Bean`，创建好以后保存在容器中。
  - 统一创建剩下的所有`bean`的时候：`finishBeanFactoryInitialization();`
- 后置处理器：
  - 每一个`bean`创建完成，都会使用各种后置处理器处理，来增强`bean`的功能；
  - 例如：`AutoWiredAnnotationBeanPostProcessor`: 处理自动注入功能
  - `AnnotationAwareAspectJAutoProxyCreator`: 来做`AOP`功能
- 事件驱动模型：
  - `ApplicationListener`: 事件监听
  - `ApplicationEventMulticaster`: 事件派发

----
