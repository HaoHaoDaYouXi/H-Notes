# 单例（Singleton）

## 问题
确保一个类只有一个实例，并提供该实例的全局访问点。

## 效果
使用私有的构造函数、静态变量各一个，及一个公有的静态函数来实现。

私有构造函数保证了不能通过构造函数来创建对象实例，只能通过公有静态函数返回唯一的私有静态变量。

## 解决方案

### 懒汉式 线程不安全
```java
public class Singleton {

    private static Singleton singleton;

    private Singleton() {}

    public static Singleton getInstance() {
        if (singleton == null) {
            singleton = new Singleton();
        }
        return singleton;
    }
}
```
私有静态变量`singleton`被延迟实例化，这样如果没有用到该类，那么就不会实例化，从而节约资源。

这个实现在多线程环境下是不安全的，因为多个线程能够同时执行`if (singleton == null)`判断，会导致实例化多次`singleton`。

### 懒汉式 线程安全
为了线程安全，需要对`getInstance()`方法加锁，保证在一个时间点只能有一个线程能够进入该方法，避免了实例化多次。

因为加锁了，所以该方法会存在性能问题，不建议使用。
```java
public static synchronized Singleton getInstance() {
    if (singleton == null) {
        singleton = new Singleton();
    }
    return singleton;
}
```

### 饿汉式 线程安全

----
