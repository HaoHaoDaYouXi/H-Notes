# 享元（Flyweight）

## 问题

利用共享的方式来支持大量细粒度的对象，这些对象一部分内部状态是相同的。

## 效果

享元模式中可以共享的相同内容称为内部状态(享元角色)，而那些需要外部环境来设置的不能共享的内容称为外部状态(非享元角色)

一个系统有大量相同或者相似的对象，由于这类对象的大量使用，造成内存的大量耗费

对象的大部分状态都可以外部化，可以将这些外部状态传入对象中(细粒度对象)

使用享元模式需要维护一个存储享元对象的享元池，而这需要耗费资源，
因此，应当在多次重复使用享元对象时才值得使用享元模式。


## 解决方案

```java
public interface Flyweight {
    void check(String out);
}

public class FlyweightImpl implements Flyweight {
    private String in;

    public FlyweightImpl(String in) {
        this.in = in;
    }

    @Override
    public void check(String out) {
        System.out.println("内存地址: " + System.identityHashCode(this));
        System.out.println("内部: " + in);
        System.out.println("外部: " + out);
    }
}

public class FlyweightFactory {

    private Map<String, Flyweight> flyweightMap = new HashMap<>();

    Flyweight getFlyweight(String in) {
        if (!flyweights.containsKey(in)) {
            Flyweight flyweight = new FlyweightImpl(in);
            flyweightMap.put(in, flyweight);
        }
        return flyweightMap.get(in);
    }
}

public class Client {
    public static void main(String[] args) {
        FlyweightFactory factory = new FlyweightFactory();
        Flyweight f1 = factory.getFlyweight("qq");
        Flyweight f2 = factory.getFlyweight("qq");
        f1.doOperation("1");
        // Object address: 1043227783
        // IntrinsicState: qq
        // ExtrinsicState: 1
        f2.doOperation("2");
        // Object address: 1043227783
        // IntrinsicState: qq
        // ExtrinsicState: 2
        // f1=f2
    }
}
```









----
