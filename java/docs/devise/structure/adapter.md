# 适配器（Adapter）

## 问题

把一个类接口转换成另一个用户需要的接口。

## 效果

适配器和代理，它们很相似

适配器让原本接口不兼容的类可以合作无间，对象适配器使用组合, 类适配器使用多重继承。

适配器模式做的就是，有一个接口需要实现，但是我们现成的对象都不满足，需要加一层适配器来进行适配。

适配器模式总体来说分三种：默认适配器模式、对象适配器模式、类适配器模式。

## 解决方案

### 默认适配器模式

我们用`Apache commons-io`包中的`FileAlterationListener`做例子，
此接口定义了很多的方法，用于对文件或文件夹进行监控，一旦发生了对应的操作，就会触发相应的方法。
```java
public interface FileAlterationListener {
    void onStart(final FileAlterationObserver observer);
    void onDirectoryCreate(final File directory);
    void onDirectoryChange(final File directory);
    void onDirectoryDelete(final File directory);
    void onFileCreate(final File file);
    void onFileChange(final File file);
    void onFileDelete(final File file);
    void onStop(final FileAlterationObserver observer);
}
```

此接口的一大问题是抽象方法太多了，如果我们要用这个接口，意味着我们要实现每一个抽象方法，
如果我们只是想要监控文件夹中的文件创建和文件删除事件，可是我们还是不得不实现所有的方法，这不是我们想要的。

所以，我们需要下面的一个适配器，它用于实现上面的接口，但是所有的方法都是空方法，
这样，我们就可以转而定义自己的类来继承下面这个类即可。
```java
public class FileAlterationListenerAdaptor implements FileAlterationListener {

    public void onStart(final FileAlterationObserver observer) {}

    public void onDirectoryCreate(final File directory) {}

    public void onDirectoryChange(final File directory) {}

    public void onDirectoryDelete(final File directory) {}

    public void onFileCreate(final File file) {}

    public void onFileChange(final File file) {}

    public void onFileDelete(final File file) {}

    public void onStop(final FileAlterationObserver observer) {}
}
```

比如我们可以定义以下类，我们仅仅需要实现我们想实现的方法就可以了：
```java
public class FileMonitor extends FileAlterationListenerAdaptor {
    public void onFileCreate(final File file) {
        // 文件创建
        doSomething();
    }

    public void onFileDelete(final File file) {
        // 文件删除
        doSomething();
    }
}
```
这是最简单的一种适配器模式。

### 对象适配器模式

我们现在模拟一个例子：因为猪肉干加工和牛肉干加工一样，现在牛肉干的需求大，所有猪肉干流水线需要适配加工牛肉干。

```java
public interface Beef {
    void beefJerky(); // 牛肉干
}

public interface Pork {
    void porkJerky(); // 猪肉干
}

public class BeefJerky implements Beef {
    public void beefJerky() {
        System.out.println("牛肉干");
    }
}
```

牛肉的接口有`beefJerky()`，猪肉厂如果要加工牛肉，因为没有牛肉干的方法，就需要适配了：
```java
public class PorkAdapter implements Pork {
    Beef beef;
    
    public CockAdapter(Beef beef) {
        this.beef = beef;
    }

    // 猪肉干流水线加工牛肉干
    @Override
    public void porkJerky() {
        // 加工牛肉
        beef.beefJerky();
    }
}
```
使用：
```java
public static void main(String[] args) {
    // 牛肉
    Beef beef = new BeefJerky();
    // 到猪肉干流水线进行加工
    Pork pork = new PorkAdapter(beef);
    // ...
}
```
适配器实际就是我们现在有一个A，但我们需要B，这个时候就需要定义一个适配器，需要A来适配B，当作B来使用。

类适配和对象适配有一些不同：
- 类适配采用继承，对象适配采用组合
- 类适配属于静态实现，对象适配属于组合的动态实现，对象适配需要多实例化一个对象

一般对象适配用得比较多。

----
