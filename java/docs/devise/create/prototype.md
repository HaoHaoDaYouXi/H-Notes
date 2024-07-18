# 原型模式（Prototype）

## 问题

使用原型实例指定要创建对象的类型，通过复制这个原型来创建新对象。

也就是`克隆`

## 效果

`Object`类中有一个`clone()`方法，它用于生成一个新的对象，
如果我们要调用这个方法，`java`要求我们的类必须先实现`Cloneable`接口，
此接口没有定义任何方法，但是不这么做的话，在`clone()`的时候，会抛出`CloneNotSupportedException`异常。
```java
protected native Object clone() throws CloneNotSupportedException;
```
`Java`的克隆是浅克隆，碰到对象引用的时候，克隆出来的对象和原对象中的引用将指向同一个对象。

通常实现深克隆的方法是将对象进行序列化，然后再进行反序列化。

## 解决方案

```java
public abstract class Prototype {
    abstract Prototype myClone();
}

public class ConcretePrototype extends Prototype {

    private String str;

    public ConcretePrototype(String str) {
        this.str = str;
    }

    @Override
    Prototype myClone() {
        return new ConcretePrototype(str);
    }

    @Override
    public String toString() {
        return str;
    }
}

public class Client {
    public static void main(String[] args) {
        Prototype prototype = new ConcretePrototype("abc");
        Prototype clone = prototype.myClone();
        System.out.println(clone.toString());
        // 输出结果：abc
    }
}
```

----
