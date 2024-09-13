# `Spring`的`IOC`

`IoC`（Inverse of Control:控制反转）是一种设计思想，而不是一个具体的技术实现。

`Spring`通过一个配置文件描述`Bean`及`Bean`之间的依赖关系，利用`Java`语言的反射功能实例化`Bean`并建立`Bean`之间的依赖关系。
`Spring`的`IoC`容器在完成这些底层工作的基础上，还提供了`Bean`实例缓存、生命周期管理、`Bean`实例代理、事件发布、资源装载等高级服务。

`Spring`的`IOC`有三种注入方式 ：构造器注入、`setter`方法注入、根据注解注入。
`IoC`让相互协作的组件保持松散的耦合，而AOP编程允许你把遍布于应用各层的功能分离出来
形成可重用的功能组件。

## <a id="ioc_rqsx">`IOC`容器实现</a>
`IoC`的实现原理就是工厂模式加反射机制。
```java
interface A {
    void a();
}

class B implements A {
    @Override
    public void a() {
        System.out.println("B.a()");
    }
}

class C implements A {
    @Override
    public void a() {
        System.out.println("C.a()");
    }
}

class Factory {
    public static A getInstance(String className) {
        A a=null;
        try {
            a=(A)Class.forName(className).newInstance();
        }catch (Exception e){
            e.printStackTrace();
        }
        return a;
    }

    public static void main(String[] args) {
        // com.B 是B类的包路径
        Factory.getInstance("com.B").a();
        Factory.getInstance("com.C").a();
    }
}
```


----
