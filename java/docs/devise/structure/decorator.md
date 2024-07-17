# 装饰（Decorator）

## 问题

为对象动态添加功能。

## 效果

装饰者（Decorator）和具体组件（ConcreteComponent）都继承自组件（Component），具体组件的方法实现不需要依赖于其它对象，而装饰者组合了一个组件，这样它可以装饰其它装饰者或者具体组件。

所谓装饰，就是把这个装饰者套在被装饰者之上，从而动态扩展被装饰者的功能。

装饰者的方法有一部分是自己的，这属于它的功能，然后调用被装饰者的方法实现，从而也保留了被装饰者的功能。

具体组件应当是装饰层次的最低层，因为只有具体组件的方法实现不需要依赖于其它对象。

## 解决方案

这里我们以咖啡店举例：现在有黑咖啡、牛奶、冰块，冰美式=黑咖啡+冰块，摩卡=黑咖啡+牛奶+牛奶，拿铁=黑咖啡+牛奶+牛奶+牛奶

```java
public interface Coffee {
    int cost();
}
// 基调
public class BlackCoffee implements Coffee {
    @Override
    public int cost() {
        return 8;
    }
}
// 调料
public abstract class Condiment implements Coffee {
    protected Coffee coffee;
}

public class Ice extends Condiment {

    public Milk(Coffee coffee) {
        this.coffee = coffee;
    }

    @Override
    public double cost() {
        return 1 + coffee.cost();
    }
}

public class Milk extends Condiment {

    public Milk(Coffee coffee) {
        this.coffee = coffee;
    }

    @Override
    public double cost() {
        return 3 + coffee.cost();
    }
}
public class Client {
    public static void main(String[] args) {
        // 冰美式
        Coffee iceAmerican = new Ice(new BlackCoffee());
        // 摩卡
        Coffee mk = new Milk(new Milk(new BlackCoffee()));
        // 冰摩卡
        Coffee bmk = new Ice(new Milk(new Milk(new BlackCoffee())));
        // 拿铁
        Coffee nt = new Milk(new Milk(new Milk(new BlackCoffee())));
    }
}
```

设计上类应该对扩展开放，对修改关闭，这样添加新功能时不需要修改代码。
可以动态添加新的调料，而不需要去修改基调的代码。

不可能把所有的类设计成都满足这一原则，应当把该原则应用于最有可能发生改变的地方。

----
