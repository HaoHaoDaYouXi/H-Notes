# 工厂（Factory）

工厂设计模式主要有3中形式：
- [简单工厂](factory.md#jdgc)
- [工厂方法](factory.md#gcff)
- [抽象工厂](factory.md#cxgc)

## <a id="jdgc">简单工厂（Simple Factory）</a>

### 问题

在创建一个对象时不向客户暴露内部细节，并提供一个创建对象的通用接口。

### 效果

简单工厂就是把实例化的操作单独放到一个类中，这个类成为简单工厂类，
让简单工厂类来决定具体用哪个子类进行实例化。

这样做能把客户类和具体子类的实现解耦，客户类不再需要知道有哪些子类以及应当实例化哪个子类。

客户类往往有多个，如果不使用简单工厂，那么所有的客户类都要知道所有子类的细节。
而且一旦子类发生改变，例如增加子类，那么所有的客户类都要进行修改。

### 解决方案

单词接口类，和几个实现类
```java
public interface Word {}

public class A implements Word {}

public class B implements Word {}

public class C implements Word {}
```

`SimpleFactory`简单工厂实现
```java
public class SimpleFactory {
    public Word generate(int type) {
        if (type == 1) {
            return new A();
        } else if (type == 2) {
            return new B();
        }// ...
        return new A();
    }
}
```

客户端调用
```java
public class Client {
    public static void main(String[] args) {
        SimpleFactory simpleFactory = new SimpleFactory();
        Word word = simpleFactory.generate(1);
        // ...
    }
}
```

## <a id="gcff">工厂方法（Factory Method）</a>

### 问题

定义了一个创建对象的接口，由子类决定要实例化哪个类。工厂方法把实例化操作推迟到子类。

### 效果

在简单工厂中，创建对象的是另一个类，而在工厂方法中，是由子类来创建对象。

`Factory`有一个`toDo()`方法，这个方法需要用到一个对象，
这个对象由`factoryMethod()`方法创建。

该方法是抽象的，需要由子类去实现。

### 解决方案

```java
public abstract class Factory {
    abstract public Word factoryMethod();
    
    public void toDo() {
        Word word = factoryMethod();
        // ...
    }
}

public class A extends Factory {
    public Word factoryMethod() {
        return new A();
    }
}

public class B extends Factory {
    public Word factoryMethod() {
        return new B();
    }
}

public class C extends Factory {
    public Word factoryMethod() {
        return new C();
    }
}
```

## <a id="cxgc">抽象工厂（Abstract Factory）</a>

### 问题

提供一个接口，用于创建相关的对象家族。

### 效果

抽象工厂模式创建的是对象家族，也就是很多对象而不是一个对象，并且这些对象是相关的，也就是说必须一起创建出来。
而工厂方法模式只是用于创建一个对象，和抽象工厂模式很不同。

### 解决方案

```java
public class AbstractA {}

public class AbstractB {}

public class WordA1 extends AbstractA {}

public class WordA2 extends AbstractA {}

public class WordB1 extends AbstractB {}

public class WordB2 extends AbstractB {}

public abstract class AbstractFactory {
    abstract AbstractA createA();
    abstract AbstractB createB();
}

public class Factory1 extends AbstractFactory {
    AbstractA createA() {
        return new WordA1();
    }

    AbstractB createB() {
        return new WordB1();
    }
}

public class Factory2 extends AbstractFactory {
    AbstractA createA() {
        return new WordA2();
    }

    AbstractB createB() {
        return new WordB2();
    }
}

public class Client {
    public static void main(String[] args) {
        AbstractFactory abstractFactory = new Factory1();
        AbstractA a = abstractFactory.createA();
        AbstractB b = abstractFactory.createB();
        // ...
    }
}
```

抽象工厂模式用到了工厂方法模式来创建单一对象，`AbstractFactory`中的`createA()`和`createB()`方法都是让子类来实现，这两个方法单独来看就是在创建一个对象，这符合工厂方法模式的定义。

创建对象的家族是指`Factory1`家族有`WordA1`和`WordB1`，`Factory2`家族有`WordA2`和`WordB2`。

抽象工厂使用了组合，组合了`AbstractFactory`，而工厂方法模式使用了继承。

----
