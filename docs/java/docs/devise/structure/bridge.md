# 桥接（Bridge）

## 问题

将抽象与实现分离开来，使它们可以独立变化。

## 效果

用桥接模式通过将实现和抽象放在两个不同的类层次中而使它们可以独立改变。
- Abstraction：定义抽象类的接口
- Implementor：定义实现类接口

## 解决方案

例子：我们现在有个手机，手机充电口是Type-C的，现在有两个充电器一个是小米的，一个是华为的，也都是Type-C的。

手机充电是Abstraction，充电器是Implementor。

给手机充电这个事和使用的充电器是分离开的，独立的。

```java
public abstract class Charger {
    /**
     * 充电
     */
    public abstract void charging();
    /**
     * 不充电
     */
    public abstract void unCharging();
}

public class XMCharger extends Charger {
    @Override
    public void charging() {
        System.out.println("小米充电器，充电");
    }
    @Override
    public void unCharging() {
        System.out.println("小米充电器，不充电");
    }
}

public class HWCharger extends Charger {
    @Override
    public void charging() {
        System.out.println("华为充电器，充电");
    }
    @Override
    public void unCharging() {
        System.out.println("华为充电器，不充电");
    }
}

public abstract class PhoneCharger {
    protected Charger charger;

    public PhoneCharger(Charger charger) {
        this.charger = charger;
    }
    /**
     * 手机充电
     */
    public abstract void charging();
    /**
     * 手机不充电
     */
    public abstract void unCharging();
}

public class PhoneChargerImpl extends PhoneCharger {

    public PhoneChargerXM(Charger charger) {
        super(charger);
    }

    @Override
    public void charging() {
        System.out.println("手机充电");
        charger.charging();
    }

    @Override
    public void unCharging() {
        System.out.println("手机不充电");
        charger.unCharging();
    }
}
```

调用：
```java
public class Client {
    public static void main(String[] args) {
        PhoneCharger phoneCharger = new PhoneChargerImpl(new XMCharger());
        phoneCharger.charging();
        phoneCharger.unCharging();
    }
}
```









----
