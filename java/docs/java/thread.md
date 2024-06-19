# 线程

## <div id="xc_jc">线程和进程</div>
### 进程
进程是程序的⼀次执⾏过程，是系统运⾏程序的基本单位，因此进程是动态的。系统运⾏⼀个程序即
是⼀个进程从创建，运⾏到消亡的过程。

在`Java`中，当我们启动`main`函数时其实就是启动了⼀个`JVM`的进程，⽽`main`函数所在的线程就
是这个进程中的⼀个线程，也称主线程。
### 线程
线程与进程相似，但线程是⼀个⽐进程更⼩的执⾏单位。⼀个进程在其执⾏的过程中可以产⽣多个线程。
与进程不同的是同类的多个线程共享进程的堆和⽅法区资源，但每个线程有⾃⼰的程序计数器、虚拟机栈和本地⽅法栈，
所以系统在产⽣⼀个线程，或是在各个线程之间作切换⼯作时，负担要⽐进程⼩得多，也正因为如此，线程也被称为轻量级进程。

`Java`程序天⽣就是多线程程序，我们可以通过`JMX`来看看⼀个普通的`Java`程序有哪些线程
```
// 获取 Java 线程管理 MXBean
ThreadMXBean threadMXBean = ManagementFactory.getThreadMXBean();
// 不需要获取同步的 monitor 和 synchronizer 信息，仅获取线程和线程堆栈信息
ThreadInfo[] threadInfos = threadMXBean.dumpAllThreads(false, false);
// 遍历线程信息，仅打印线程 ID 和线程名称信息
for (ThreadInfo threadInfo : threadInfos) {
    System.out.println("[" + threadInfo.getThreadId() + "] " +
    threadInfo.getThreadName());
}
```
程序输出
```
[5] Attach Listener //添加事件
[4] Signal Dispatcher // 分发处理给 JVM 信号的线程
[3] Finalizer //调⽤对象 finalize ⽅法的线程
[2] Reference Handler //清除 reference 线程
[1] main //main 线程,程序⼊⼝
```
从上⾯的输出内容可以看出：**⼀个 Java 程序的运⾏是 main 线程和多个其他线程同时运⾏。**

## <div id="xc_smzq">线程的生命周期</div>

## <div id="xc_zt">线程的状态</div>
