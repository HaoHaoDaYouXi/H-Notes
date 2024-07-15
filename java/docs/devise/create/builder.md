# 建造者（Builder）

## 问题

封装一个对象的构造过程，并允许按步骤构造。

## 效果

由产品、抽象建造者、具体建造者、指挥者等`4`个要素构成

建造者（Builder）模式的主要角色如下
- 产品角色（Product）：它是包含多个组成部件的复杂对象，由具体建造者来创建其各个零部件。
- 抽象建造者（Builder）：它是一个包含创建产品各个子部件的抽象方法的接口，通常还包含一个返回复杂产品的方法。
- 具体建造者(Concrete Builder）：实现`Builder`接口，完成复杂产品的各个部件的具体创建方法。
- 指挥者（Director）：它调用建造者对象中的部件构造与装配方法完成复杂对象的创建，在指挥者中不涉及具体产品的信息。

建造者（Builder）模式和工厂模式的关注点不同：
- 建造者模式注重零部件的组装过程
- 工厂方法模式更注重零部件的创建过程
- 两者可以结合使用

## 解决方案
产品角色：包含多个组成部件的复杂对象。
```java
class Product {
    private String partA;
    private String partB;
    private String partC;
    public void setPartA(String partA) {
        this.partA = partA;
    }
    public void setPartB(String partB) {
        this.partB = partB;
    }
    public void setPartC(String partC) {
        this.partC = partC;
    }
    public void show() {
        // 显示产品的特性
    }
}
```

抽象建造者：包含创建产品各个子部件的抽象方法。
```java
abstract class Builder {
    // 创建产品对象
    protected Product product = new Product();
    public abstract void buildPartA();
    public abstract void buildPartB();
    public abstract void buildPartC();
    // 返回产品对象
    public Product getResult() {
        return product;
    }
}
```

具体建造者：实现了抽象建造者接口。
```java
public class ConcreteBuilder extends Builder {
    public void buildPartA() {
        product.setPartA("建造 PartA");
    }
    public void buildPartB() {
        product.setPartB("建造 PartB");
    }
    public void buildPartC() {
        product.setPartC("建造 PartC");
    }
}
```

指挥者：调用建造者中的方法完成复杂对象的创建。
```java
class Director {
    private Builder builder;
    public Director(Builder builder) {
        this.builder = builder;
    }
    // 产品构建与组装方法
    public Product construct() {
        builder.buildPartA();
        builder.buildPartB();
        builder.buildPartC();
        return builder.getResult();
    }
}

```

具体使用
```java
public class Client {
    public static void main(String[] args) {
        Builder builder = new ConcreteBuilder();
        Director director = new Director(builder);
        Product product = director.construct();
        product.show();
    }
}
```


----
