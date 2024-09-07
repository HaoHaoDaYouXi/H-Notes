# 虚拟线程

虚拟线程是 Java 19 新增的特性，现在 Java 21 已确定保留，它允许在 Java 中使用线程，而无需显式地创建线程对象。

虚拟线程与普通线程相比，具有更少的开销，更少的资源消耗，更少的上下文切换，更少的内存占用等优势。
虚拟线程也支持协程，并且可以与普通线程一起使用。

## 使用

线程的创建方式与普通线程基本一致

```java
Thread thread = Thread.ofVirtual().name("test").start(new Task());
```

线程池

```
普通：Executors.newFixedThreadPool(10);
虚拟：Executors.newVirtualThreadPerTaskExecutor();
```

----
