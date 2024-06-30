# `Spring`加载过程

1. `PrepareRefresh()`刷新前的预处理

   `initPropertySources()`初始化一些属性设置
    
   `getEnvironment().validateRequiredProperties();` 检验属性的合法等
    
   `earlyApplicationEvents = new LinkedHashSet<ApplicationEvent>;` 保存容器中的一些早期时间

2. `obtainFreshBeanFactory();` 获取`BeanFactory`

   `refreshBeanFactory();` 刷新【创建】`BeanFactory`

   `getBeanFactory();` 返回刚才`GenericApplicationContext`创建的`BeanFactory`对象
   　　
   将创建的`BeanFactory`对象【`DefaultListableBeanFactory`】返回

3. `prepareBeanFactory(beanFactory);` `BeanFactory`的预备准备工作（`BeanFactory`进行一些设置）

   设置`BeanFactory`的类加载器、支持表达式解析器
　　添加部分`BeanPostProcessor`【`ApplicationContextAwareProcessor`】
　　设置忽略的自动装配的接口 `EnvironmentAware`、`EmbeddedValueResolverAware`。。。
　　注册可以解析的自动装配，我们能直接在任何组件中自动注入：`BeanFactory`、`ResourceLoader`、`ApplicationEventPublisher`、`ApplicationContext`
　　添加`BeanPostProcessor`,【`ApplicationListenerDetector`】
　　添加编译时的`AspectJ`支持
　　给`BeanFactory`中注册一些能用的组件：`environment`【`ConfigurableEnvironment`】、`SystemProperties`【`Map<String, Object>`】、　　              `systemEnvironment`【`Map<String, Object>`】

4. `postProcessBeanFactory(beanFactory);` `BeanFactory`准备工作完成后进行的后置处理工作：
   子类通过这个方法在BeanFactory创建并预准备完成后做的进一步设置

----

### 以上是BeanFactory的创建和预准备工作

----
