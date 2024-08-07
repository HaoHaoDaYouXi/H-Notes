# 代理

`Java`中的代理机制是一种设计模式，也是一种编程技术，它允许你在不修改目标对象的情况下为其添加新的功能。

`Java`代理有两种主要形式：[`静态代理`](#静态代理)和[`动态代理`](#动态代理)。

## 静态代理

静态代理是在编译时期就已经定义好的代理类。
这种代理类通常是由程序员手动创建的，它实现了与目标对象相同的接口。

例：假设有一个`Fruit`接口，表示各种水果
```java
public interface Fruit {
    String getTaste();
}
```
我们可以为具体的水果如`Apple`创建一个实现
```java
public class Apple implements Fruit {
    @Override
    public String getTaste() {
        return "Sweet";
    }
}
```
然后创建一个静态代理类`AppleProxy`
```java
public class AppleProxy implements Fruit {
    private final Fruit apple;

    public AppleProxy(Fruit apple) {
        this.apple = apple;
    }

    @Override
    public String getTaste() {
        // 执行一些前置操作
        System.out.println("Before getting taste...");

        // 调用实际对象的方法
        String taste = apple.getTaste();

        // 执行一些后置操作
        System.out.println("After getting taste...");
        return taste;
    }
}
```
这样就可以使用代理对象：
```java
public static void main(String[] args) {
    Fruit apple = new Apple();
    Fruit appleProxy = new AppleProxy(apple);
    System.out.println(appleProxy.getTaste());
}
```

## 动态代理

动态代理则是在运行时动态创建的代理对象。

Java提供了两种动态代理的方式：基于`接口`的代理和基于`类`的代理。

### 基于接口的动态代理

动态代理主要依赖于`java.lang.reflect.Proxy`类和`java.lang.reflect.InvocationHandler`接口。

例：使用`Proxy`和`InvocationHandler`创建动态代理对象
```java
public class DynamicProxyDemo {
    public static void main(String[] args) {
        Fruit fruit = new Apple();
        Fruit proxyFruit = (Fruit) Proxy.newProxyInstance(
                fruit.getClass().getClassLoader(),
                fruit.getClass().getInterfaces(),
                new FruitInvocationHandler(fruit)
        );
        System.out.println(proxyFruit.getTaste());
    }
}

class FruitInvocationHandler implements InvocationHandler {
    private final Fruit fruit;

    public FruitInvocationHandler(Fruit fruit) {
        this.fruit = fruit;
    }

    @Override
    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
        System.out.println("Before method call");
        Object result = method.invoke(fruit, args);
        System.out.println("After method call");
        return result;
    }
}
```
例子中，`FruitInvocationHandler`实现了`InvocationHandler`接口，并在`invoke`方法中处理了方法调用前后的逻辑。

## 总结

代理机制可以在不修改现有代码的基础上增加新的功能，这对于`AOP`（面向切面编程）和拦截器等模式非常有用。
静态代理需要手动编写代理类，而动态代理则可以在运行时自动创建代理对象，因此更为灵活。

# <a id="cglib">`CGlib`</a>

`CGLIB`（Code Generation Library）是一个强大的、高性能的代码生成库，它可以在运行时动态地生成类。
`CGLIB`主要用于实现动态代理，尤其适用于那些没有实现接口的类。

在`Java`中，如果一个类没有实现任何接口，那么就不能直接使用`JDK`的动态代理机制来创建代理对象。
这时`CGLIB`就派上了用场。

## 工作原理

`CGLIB`通过字节码技术为一个现有的类创建子类，并在子类中采用方法拦截的技术来实现代理。
当我们在运行时创建一个代理对象时，`CGLIB`会生成一个新的类，这个新类继承自原始类，并覆盖了所有非最终（`non-final`）方法。
在覆盖的方法中，`CGLIB`使用回调机制来执行额外的操作。

## 使用

Maven依赖
```xml
<!-- https://mvnrepository.com/artifact/cglib/cglib -->
<dependency>
    <groupId>cglib</groupId>
    <artifactId>cglib</artifactId>
    <version>3.3.0</version>
</dependency>
```

例：假设有个简单的类`Calculator`，它没有实现任何接口
```java
public class Calculator {
    public int add(int a, int b) {
        return a + b;
    }

    public int multiply(int a, int b) {
        return a * b;
    }
}
```

使用`CGLIB`来创建一个动态代理对象
```java
public class CglibProxyExample {
    public static void main(String[] args) {
        Calculator calculator = new Calculator();
        Calculator proxyCalculator = (Calculator) createProxy(calculator);

        System.out.println(proxyCalculator.add(1, 2));
        System.out.println(proxyCalculator.multiply(3, 4));
    }

    private static Object createProxy(Calculator calculator) {
        Enhancer enhancer = new Enhancer();
        enhancer.setSuperclass(calculator.getClass());
        enhancer.setCallback(new MethodInterceptor() {
            @Override
            public Object intercept(Object obj, Method method, Object[] args, MethodProxy proxy) throws Throwable {
                System.out.println("Before method call");
                Object result = proxy.invokeSuper(obj, args);
                System.out.println("After method call");
                return result;
            }
        });
        return enhancer.create();
    }
}
```
例子中，`createProxy`方法使用`CGLIB`创建了一个`Calculator`类的代理对象。
`Enhancer`类负责创建代理对象，`MethodInterceptor`接口定义了方法调用前后需要执行的逻辑。

## 总结

`CGLIB`是一个非常有用的工具，尤其在需要为没有实现接口的类创建动态代理时。
通过使用`CGLIB`，你可以轻松地为现有类添加新的功能，而无需修改原始类的代码。
这在很多框架中都非常常见，`Spring AOP`就使用了`CGLIB`来实现对未实现接口的类的动态代理。






----
