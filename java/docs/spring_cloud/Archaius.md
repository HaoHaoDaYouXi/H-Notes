# Archaius

Archaius是一个提供获取可以运行时动态改变的属性的API的Java类库，它主要实现为`Apache Commons Configuration`库的扩展。

## Archaius特性

### 动态属性
- 可以使用简洁的代码去动态获取类型特定的属性
- 可以对某个属性发生修改之后创建callback
```java
DynamicIntProperty prop = DynamicPropertyFactory.getInstance().getIntProperty("myProperty", DEFAULT_VALUE);
// prop.get() may change value at runtime
myMethod(prop.get());

// 创建callback
prop.addCallback(new Runnable() {
  public void run() {
      // ...
  }
});
```















----
