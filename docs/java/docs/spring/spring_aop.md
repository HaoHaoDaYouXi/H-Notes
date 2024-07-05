# `Spring`的`AOP`

`AOP`（Aspect-Oriented Programming，面向切面编程）能够将那些与业务无关，
却为业务模块所共同调用的逻辑或责任（例如事务处理、日志管理、权限控制等）封装起来，
便于减少系统的重复代码，降低模块间的耦合度，并有利于未来的可扩展性和可维护性。

`Spring AOP`是基于动态代理的，如果要代理的对象实现了某个接口，那么`Spring AOP`就会使用`JDK`动态代理去创建代理对象；
而对于没有实现接口的对象，就无法使用`JDK`动态代理，转而使用`CGlib`动态代理生成一个被代理对象的子类来作为代理。

也可以使⽤`AspectJ`，`Spring AOP`已经集成了`AspectJ`，`AspectJ`是`Java`⽣态系统中最完整的`AOP`框架了。

## <a id="aophxgn">`AOP`核心概念</a>
- 切面（`aspect`）：类是对物体特征的抽象，切面就是对横切关注点的抽象
- 横切关注点：对哪些方法进行拦截，拦截后怎么处理，这些关注点称之为横切关注点。
- 连接点（`joinpoint`）：被拦截到的点，因为`Spring`只支持方法类型的连接点，所以在`Spring`中连接点指的就是被拦截到的方法，实际上连接点还可以是字段或者构造器。
- 切入点（`pointcut`）：对连接点进行拦截的定义
- 通知（`advice`）：所谓通知指的就是指拦截到连接点之后要执行的代码，通知分为前置、后置、异常、最终、环绕通知五类。
- 目标对象：代理的目标对象
- 织入（`weave`）：将切面应用到目标对象并导致代理对象创建的过程
- 引入（`introduction`）：在不修改代码的前提下，引入可以在运行期为类动态地添加一些方法或字段。

## <a id="aopdlfs">`AOP`代理方式</a>
`Spring`提供了两种方式来生成代理对象: `JDKProxy`和`Cglib`，具体使用哪种方式生成由
`AopProxyFactory`根据`AdvisedSupport`对象的配置来决定。默认的策略是如果目标类是接口，
则使用`JDK`动态代理技术，否则使用`Cglib`来生成代理。
- `JDK`动态接口代理
  - `JDK`动态代理主要涉及到`java.lang.reflect`包中的两个类：`Proxy`和`InvocationHandler`。
      - `InvocationHandler`是一个接口，通过实现该接口定义横切逻辑，并通过反射机制调用目标类的代码，动态将横切逻辑和业务逻辑编制在一起。
      - `Proxy`利用`InvocationHandler`动态创建一个符合某一接口的实例，生成目标类的代理对象。
- `CGLib`动态代理
  - `CGLib`全称为`Code Generation Library`，是一个强大的高性能，高质量的代码生成类库，
     可以在运行期扩展`Java`类与实现`Java`接口，`CGLib`封装了`asm`，可以再运行期动态生成新的`class`。
  - 和`JDK`动态代理相比较：`JDK`创建代理有一个限制，就是只能为接口创建代理实例，而对于没有通过接口定义业务方法的类，则可以通过`CGLib`创建动态代理。

## <a id="aopsx">`AOP`实现</a>
```java
@Aspect
public class TransactionDemo {
    @Pointcut(value="execution(* com.yangxin.core.service.*.*.*(..))")
    public void point(){
    }
    
    @Before(value="point()")
    public void before(){
        System.out.println("transaction begin");
    }
    
    @AfterReturning(value = "point()")
    public void after(){
        System.out.println("transaction commit");
    }
    
    @Around("point()")
    public void around(ProceedingJoinPoint joinPoint) throws Throwable {
        System.out.println("transaction begin");
        joinPoint.proceed();
        System.out.println("transaction commit");
    }
}
```

## <a id="aspectJ">`AspectJ`</a>

### <a id="aspectJ_spring">`Spring AOP`和`AspectJ AOP`</a>
`Spring AOP`是属于运行时增强，而`AspectJ`是编译时增强。
`Spring AOP`基于代理(`Proxying`)，而`AspectJ`基于字节码操作(`Bytecode Manipulation`)。
`Spring AOP`已经集成了`AspectJ`，`AspectJ`算得上是`Java`生态系统中最完整的`AOP`框架。
`AspectJ`相比于`Spring AOP`功能更加强大，但是`Spring AOP`相对来说更简单。
如果我们的切面比较少，那么两者性能差异不大。当切面太多的话，最好选择`AspectJ`，它比`Spring AOP`快很多。

### <a id="aspectJ_tzlx">`AspectJ`定义的通知类型</a>
- `Before`(前置通知)：⽬标对象的⽅法调⽤之前触发
- `After`(后置通知)：⽬标对象的⽅法调⽤之后触发
- `AfterReturning`(返回通知)：⽬标对象的⽅法调⽤完成，在返回结果值之后触发
- `AfterThrowing`(异常通知) ：⽬标对象的⽅法运⾏中抛出或触发异常后触发。
  - `AfterReturning`和`AfterThrowing`两者互斥。
  - 如果⽅法调⽤成功⽆异常，则会有返回值；如果⽅法抛出了异常，则不会有返回值。
- `Around`：(环绕通知)编程式控制⽬标对象的⽅法调⽤。
  - 环绕通知是所有通知类型中可操作范围最⼤的⼀种，因为它可以直接拿到⽬标对象，以及要执⾏的⽅法
  - 所以环绕通知可以任意的在⽬标对象的⽅法调⽤前后搞事，甚⾄不调⽤⽬标对象的⽅法

### <a id="dgqmdzxsx">多个切面的执行顺序</a>

**通常使⽤`@Order`注解直接定义切⾯顺序**

```java
// 值越⼩优先级越⾼
@Order(3)
@Aspect
@Component
public class AClassAspect implements Ordered {
    
}
```

**实现`Ordered`接⼝重写`getOrder`⽅法。**

```java
@Aspect
@Component
public class AClassAspect implements Ordered {
    
    @Override
    public int getOrder() {
        // 返回值越⼩优先级越⾼
        return 1;
    }
}
```


----
