## jvm配置优化

关于jvm的配置优化，我们得先了解jvm的配置，比如常用的年起代、老年代、方法区这类的配置是什么，
又应该如何配置，项目遇到情况我们又应该怎么排查问题并调整对应的配置

## XX参数

### Boolean类型

~~~
公式：-XX:+ 或者-某个属性 1+ 表示开启，-表示关闭
Case：-XX:-PrintGCDetails：表示关闭了GC详情输出
~~~

### key-value类型

~~~
公式：-XX:属性key=属性value
不满意初始值，可以通过下列命令调整
case：如何：-XX:MetaspaceSize=21807104 设置Java元空间的值
~~~

### jvm常见配置项

> -XX:+PrintFlagsInitial
>> 查看初始默认值

> -XX:+PrintFlagsFinal
>> 表示修改以后，最终的值，如果有 := 表示修改过的， = 表示没有修改过的

> -Xms
>> -Xms等价于-XX:InitialHeapSize，初始大小内存，默认物理内存1/64。
  默认当空余堆内存大于70%时， JVM会减小heap的大小到-Xms指定的大小，可通过-XX:MaxHeapFreeRation=来指定这个比列。
  Server端JVM最好将-Xms和-Xmx设为相同值，避免每次垃圾回收完成后JVM重新分配内存；
  开发测试机JVM可以保留默认值。(例如：-Xms4g)

> -Xmx
>> -Xmx等价于-XX:MaxHeapSize，最大分配内存，默认为物理内存1/4
  默认当空余堆内存小于40%时，JVM会增大Heap到-Xmx指定的大小，可通过-XX:MinHeapFreeRation=来指定这个比列。
  最佳设值应该视物理内存大小及计算机内其他内存开销而定。(例如：-Xmx4g)

> -Xmn
>> -Xmn等价于-XX:MaxNewSize 新生代内存配置。
  整个堆大小=年轻代大小 + 年老代大小 + 持久代大小(相对于HotSpot 类型的虚拟机来说)。
  持久代一般固定大小为64m，所以增大年轻代后，将会减小年老代大小。
  此值对系统性能影响较大，Sun官方推荐配置为整个堆的3/8。(例如：-Xmn2g)

程序新创建的对象都是从年轻代分配内存，年轻代由Eden Space和两块相同大小的SurvivorSpace(通常又称S0和S1或From和To)构成，
可通过-Xmn参数来指定年轻代的大小，也可以通过-XX:SurvivorRation来调整Eden Space及SurvivorSpace的大小。

老年代用于存放经过多次新生代GC仍然存活的对象，例如缓存对象，新建的对象也有可能直接进入老年代，
主要有两种情况：
1. 大对象，可通过启动参数设置-XX:PretenureSizeThreshold=1024(单位为字节，默认为0)来代表超过多大时就不在新生代分配，而是直接在老年代分配。
2. 大的数组对象，且数组中无引用外部对象。老年代所占的内存大小为-Xmx对应的值减去-Xmn对应的值。如果在堆中没有内存完成实例分配，并且堆也无法再扩展时，将会抛出OutOfMemoryError异常。

> -Xss
>> Java每个线程的Stack大小。JDK5.0以后每个线程堆栈大小为1M，以前每个线程堆栈大小为256K。 
> 根据应用的线程所需内存大小进行调整。在相同物理内存下，减小这个值能生成更多的线程。 
> 但是操作系统对一个进程内的线程数还是有限制的，不能无限生成，经验值在3000~5000左右。(例如：-Xss1024K)

#### 方法区

方法区JDK8和之前的版本有所区别。

##### 永生代,在JDK8之前，JVM中存在着方法区的概念，也可以叫做永生代（Perm）
> -XX:PermSize
>> 初始内存大小。（例如：-XX:PermSize=64m）

> -XX:MaxPermSize
>> 最大内存大小。（例如：-XX:MaxPermSize=512m）

##### 元空间,从JDK8开始，JVM将原来存放元数据的永生代Perm换成了本地元空间Metaspace
> -XX:MetaspaceSize
>> 初始内存大小。（例如：-XX:MetaspaceSize=64m）

> -XX:MaxMetaspaceSize
>> 最大内存大小。（例如：-XX:MaxMetaspaceSize=512m）

<div style="color: #ff5050" >
注意：方法区的空间设置一定要多注意，如果项目运行一段时间后内存就崩了，可能就是设置大了导致内存无法释放
</div>

> -XX:+UseSerialGC
>> 串行（SerialGC）是jvm的默认GC方式，一般适用于小型应用和单处理器，算法比较简单，GC效率也较高，但可能会给应用带来停顿。

> -XX:+UseParallelGC
>> 并行（ParallelGC）是指多个线程并行执行GC，一般适用于多处理器系统中，可以提高GC的效率，但算法复杂，系统消耗较大。（配合使用：-XX:ParallelGCThreads=8，并行收集器的线程数，此值最好配置与处理器数目相等）

> -XX:+UseParNewGC
>> 设置年轻代为并行收集，JKD5.0以上，JVM会根据系统配置自行设置，所以无需设置此值。

> -XX:+UseParallelOldGC
>> 设置年老代为并行收集，JKD6.0出现的参数选项。

> -XX:+UseConcMarkSweepGC
>> 并发（ConcMarkSweepGC）是指GC运行时，对应用程序运行几乎没有影响（也会造成停顿，不过很小而已），GC和app两者的线程在并发执行，这样可以最大限度不影响app的运行。

> -XX:+UseCMSCompactAtFullCollection
>> 在Full GC的时候，对老年代进行压缩整理。因为CMS是不会移动内存的，因此非常容易产生内存碎片。因此增加这个参数就可以在FullGC后对内存进行压缩整理，消除内存碎片。当然这个操作也有一定缺点，就是会增加CPU开销与GC时间，所以可以通过-XX:CMSFullGCsBeforeCompaction=3 这个参数来控制多少次Full GC以后进行一次碎片整理。

> -XX:+CMSInitiatingOccupancyFraction=80
>> 代表老年代使用空间达到80%后，就进行Full GC。CMS收集器在进行垃圾收集时，和应用程序一起工作，所以，不能等到老年代几乎完全被填满了再进行收集，这样会影响并发的应用线程的空间使用，从而再次触发不必要的Full GC。

> -XX:+MaxTenuringThreshold=10
>> 垃圾的最大年龄，代表对象在Survivor区经过10次复制以后才进入老年代。如果设置为0，则年轻代对象不经过Survivor区，直接进入老年代。

----
