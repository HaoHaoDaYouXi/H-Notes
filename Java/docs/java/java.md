## Java 基础

### JVM
Java 虚拟机（JVM）是运⾏ Java 字节码的虚拟机。JVM 有针对不同系统的特定实现（Windows，
Linux，macOS），⽬的是使⽤相同的字节码，它们都会给出相同的结果。字节码和不同系统的
JVM 实现是 Java 语⾔“⼀次编译，随处可以运⾏”的关键所在。
JVM 并不是只有⼀种！只要满⾜ JVM 规范，每个公司、组织或者个⼈都可以开发⾃⼰的专属
JVM。 也就是说我们平时接触到的 HotSpot VM 仅仅是是 JVM 规范的⼀种实现⽽已。
除了我们平时最常⽤的 HotSpot VM 外，还有 J9 VM、Zing VM、JRockit VM 等 JVM 。维基百科上
就有常⻅ JVM 的对⽐：Comparison of Java virtual machines ，感兴趣的可以去看看。并且，你可
以在 Java SE Specifications 上找到各个版本的 JDK 对应的 JVM 规范。

### JDK 和 JRE

JDK 是 Java Development Kit 缩写，它是功能⻬全的 Java SDK。它拥有 JRE 所拥有的⼀切，还有
编译器（javac）和⼯具（如 javadoc 和 jdb）。它能够创建和编译程序。
JRE 是 Java 运⾏时环境。它是运⾏已编译 Java 程序所需的所有内容的集合，包括 Java 虚拟机
（JVM），Java 类库，java 命令和其他的⼀些基础构件。但是，它不能⽤于创建新程序。
如果你只是为了运⾏⼀下 Java 程序的话，那么你只需要安装 JRE 就可以了。如果你需要进⾏⼀些
Java 编程⽅⾯的⼯作，那么你就需要安装 JDK 了。但是，这不是绝对的。有时，即使您不打算在计
算机上进⾏任何 Java 开发，仍然需要安装 JDK。例如，如果要使⽤ JSP 部署 Web 应⽤程序，那么
从技术上讲，您只是在应⽤程序服务器中运⾏ Java 程序。那你为什么需要 JDK 呢？因为应⽤程序
服务器会将 JSP 转换为 Java servlet，并且需要使⽤ JDK 来编译 servlet。

### <div id="java_ts">java的特色</div>
1. 简单易学、有丰富的类库
2. 面向对象（Java最重要的特性，让程序耦合度更低，内聚性更高）
3. 与平台无关性（JVM是Java跨平台使用的根本）
4. 可靠安全
5. 支持多线程
