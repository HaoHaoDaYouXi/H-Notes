# 外观（Facade）

## 问题

提供统一的接口，用来访问子系统中的一群接口，从而让子系统更容易使用。

## 效果

现在智能家具大家应该都接触过，但是智能设备有很多，回家时这设备要开启，那个设备要开启，智能话场景可以通过一个指令把很多设备统一开启，这就是外观模式。

## 解决方案
模拟一个场景，回家时，打开灯、空调、电视，继续观看。
```java
public class SubSystem {
    public void turnOnTheLights() {
        System.out.println("开灯");
    }
    
    public void turnOnTheAirConditioner() {
        System.out.println("开空调");
    }

    public void turnOnTheTV() {
        System.out.println("开电视");
    }

    public void keepWatching(){
        System.out.println("继续观看");
    }
}

public class Facade {
    private SubSystem subSystem = new SubSystem();

    public void wentHome() {
        subSystem.turnOnTheLights();
        subSystem.turnOnTheAirConditioner();
        subSystem.turnOnTheTV();
        subSystem.keepWatching();
    }
}
public class Client {
    public static void main(String[] args) {
        Facade facade = new Facade();
        facade.wentHome();
    }
}
```
设计上客户端所需要交互的对象应当尽可能少。








----
