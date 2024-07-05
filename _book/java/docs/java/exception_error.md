# Exception && Error
# Exception
## java.lang.Exception
根异常。用以描述应用程序希望捕获的情况。
## java.lang.IllegalMonitorStateException
违法的监控状态异常。当某个线程试图等待一个自己并不拥有的对象（O）的监控器或者通知其他线程等待该对象（O）的监控器时，抛出该异常。

## java.lang.IllegalStateException
违法的状态异常。当在Java环境和应用尚未处于某个方法的合法调用状态，而调用了该方法时，抛出该异常。

## java.lang.IllegalThreadStateException
违法的线程状态异常。当县城尚未处于某个方法的合法调用状态，而调用了该方法时，抛出异常。

## java.lang.InstantiationException
实例化异常。当试图通过newInstance()方法创建某个类的实例，而该类是一个抽象类或接口时，抛出该异常。

## java.lang.InterruptedException
被中止异常。当某个线程处于长时间的等待、休眠或其他暂停状态，而此时其他的线程通过Thread的interrupt方法终止该线程时抛出该异常。

## java.lang.NullPointerException
空指针异常。简单地说就是调用了未经初始化的对象或者是不存在的对象。

## java.lang.ClassCastException
类型转换异常。通常发生在试图将一个对象强制转换成它不是实际类型的类型时。这个异常表明程序在运行时遇到了类型不兼容的强制转换。
假设有类A和B（A不是B的父类或子类），O是A的实例，那么当强制将O构造为类B的实例时抛出该异常。

## java.lang.ArithmeticException
算术异常。譬如：整数除零等。

## java.lang.ArrayIndexOutOfBoundsException
数组下标越界异常。通常在尝试访问数组中不存在的索引位置时抛出。

## java.lang.IllegalArgumentException
方法的参数错误异常。通常在方法接收到一个不合法的参数时抛出。

## java.lang.IllegalAccessException
访问异常。当我们试图动态地访问或修改一个类的成员，而该成员不是公开可访问的，就有可能遭遇这个异常。

## java.lang.ArrayStoreException
数组存储异常。当向数组中存放非数组声明类型对象时抛出。

## java.lang.ClassNotFoundException
找不到类异常。当应用试图根据字符串形式的类名构造类，而在遍历CLASSPAH之后找不到对应名称的class文件时，抛出该异常。

## java.lang.CloneNotSupportedException
不支持克隆异常。当没有实现Cloneable接口或者不支持克隆方法时,调用其clone()方法则抛出该异常。

## java.lang.EnumConstantNotPresentException
枚举常量不存在异常。当应用试图通过名称和枚举类型访问一个枚举对象，但该枚举对象并不包含常量时，抛出该异常。

~~~
数组负下标异常：NegativeArrayException

违背安全原则异常：SecturityException

文件已结束异常：EOFException

文件未找到异常：FileNotFoundException

字符串转换为数字异常：NumberFormatException

操作数据库异常：SQLException

输入输出异常：IOException

方法未找到异常：NoSuchMethodException

....
~~~

# Error
## java.lang.Error
错误。是所有错误的基类，用于标识严重的程序运行问题。这些问题通常描述一些不应被应用程序捕获的反常情况。

## java.lang.ThreadDeath
线程结束。当调用Thread类的stop方法时抛出该错误，用于指示线程结束。

## java.lang.OutOfMemoryError
内存不足错误。当可用内存不足以让Java虚拟机分配给一个对象时抛出该错误。

## java.lang.StackOverflowError
堆栈溢出错误。当一个应用递归调用的层次太深而导致堆栈溢出时抛出该错误。

## java.lang.NoSuchMethodError
方法不存在错误。当应用试图调用某类的某个方法，而该类的定义中没有该方法的定义时抛出该错误。

## java.lang.NoSuchFieldError
域不存在错误。当应用试图访问或者修改某类的某个域，而该类的定义中没有该域的定义时抛出该错误。

## java.lang.NoClassDefFoundError
未找到类定义错误。当Java虚拟机或者类装载器试图实例化某个类，而找不到该类的定义时抛出该错误。

## java.lang.AbstractMethodError
抽象方法错误。当应用试图调用抽象方法时抛出。

## java.lang.AssertionError
断言错。用来指示一个断言失败的情况。

## java.lang.ClassCircularityError
类循环依赖错误。在初始化一个类时，若检测到类之间循环依赖则抛出该异常。

## java.lang.ClassFormatError
类格式错误。当Java虚拟机试图从一个文件中读取Java类，而检测到该文件的内容不符合类的有效格式时抛出。

## java.lang.ExceptionInInitializerError
初始化程序错误。当执行一个类的静态初始化程序的过程中，发生了异常时抛出。静态初始化程序是指直接包含于类中的static语句段。

## java.lang.IllegalAccessError
违法访问错误。当一个应用试图访问、修改某个类的域（Field）或者调用其方法，但是又违反域或方法的可见性声明，则抛出该异常。

## java.lang.IncompatibleClassChangeError
不兼容的类变化错误。当正在执行的方法所依赖的类定义发生了不兼容的改变时，抛出该异常。一般在修改了应用中的某些类的声明定义而没有对整个应用重新编译而直接运行的情况下，容易引发该错误。

## java.lang.InstantiationError
实例化错误。当一个应用试图通过Java的new操作符构造一个抽象类或者接口时抛出该异常.

## java.lang.InternalError
内部错误。用于指示Java虚拟机发生了内部错误。

## java.lang.LinkageError
链接错误。该错误及其所有子类指示某个类依赖于另外一些类，在该类编译之后，被依赖的类改变了其类定义而没有重新编译所有的类，进而引发错误的情况。

## java.lang.UnknownError
未知错误。用于指示Java虚拟机发生了未知严重错误的情况。

## java.lang.UnsatisfiedLinkError
未满足的链接错误。当Java虚拟机未找到某个类的声明为native方法的本机语言定义时抛出。

## java.lang.UnsupportedClassVersionError
不支持的类版本错误。当Java虚拟机试图从读取某个类文件，但是发现该文件的主、次版本号不被当前Java虚拟机支持的时候，抛出该错误。

## java.lang.VerifyError
验证错误。当验证器检测到某个类文件中存在内部不兼容或者安全问题时抛出该错误。

## java.lang.VirtualMachineError
虚拟机错误。用于指示虚拟机被破坏或者继续执行操作所需的资源不足的情况。
