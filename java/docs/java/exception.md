## 异常

### 异常分类
在`Java`中，所有的异常都有一个共同的祖先`java.lang`包中的`Throwable`类。
`Throwable`类有两个重要的⼦类，分为`Error`和`Exception`

- `Error`
  - `Error`类是指`Java`运行时系统的内部错误和资源耗尽错误。属于程序⽆法处理的错误，我们没办法通过`catch`来进行捕获不建议通过`catch`捕获。
  例如`Java`虚拟机运行错误(`Virtual MachineError`)、虚拟机内存不够错误(`OutOfMemoryError`)、类定义错误(`NoClassDefFoundError`)等。
  应用程序不会抛出该类对象。如果出现了这样的错误，除了告知用户，剩下的就是尽力使程序安全的终止。

- `Exception`（`RuntimeException`、`CheckedException`）
  - `Exception`又有两个分支，一个是运行时异常`RuntimeException`，一个是`CheckedException`。
    - `RuntimeException`如：`NullPointerException`、`ClassCastException`；
    - `CheckedException`，如`I/O`错误导致的`IOException`、`SQLException`。
    `RuntimeException`是那些可能在`Java`虚拟机正常运行期间抛出的异常的超类。
    如果出现`RuntimeException`，那么一定是代码错误.

### 检查异常 `CheckedException`：
`Java`代码在编译过程中，如果受检查异常没有被`try catch`或者`throws`关键字处理的话，就没办法通过编译。
除了`RuntimeException`及其⼦类以外，其他的`Exception`类及其⼦类都属于受检查异常。
常⻅的受检查异常有：`IO`相关的异常、`ClassNotFoundException`、`SQLException`...。

### 不受检查异常 `UncheckedException`：
`Java`代码在编译过程中，我们即使不处理不受检查异常也可以正常通过编译。
`RuntimeException`及其⼦类都统称为⾮受检查异常

常⻅的[Exception && Error](exception_error.md):
- `NullPointerException`(空指针异常)
- `IllegalArgumentException`(参数异常比如方法入参类型错误)
- `NumberFormatException`(字符串转换为数字格式异常，`IllegalArgumentException`的⼦类)
- `ArrayIndexOutOfBoundsException`(数组越界异常)
- `ClassCastException`(类型转换异常)
- `ArithmeticException`(算术异常)
- `SecurityException`(安全异常比如权限不够)
- `UnsupportedOperationException`(不⽀持的操作异常比如重复创建同一用户)

### `Throwable`类常用方法有哪些？
- `String getMessage()`: 返回异常发生时的简要描述
- `String toString()`: 返回异常发生时的详细信息
- `String getLocalizedMessage()`: 返回异常对象的本地化信息。使用`Throwable`的⼦类覆盖这个方
法，可以生成本地化信息。如果⼦类没有覆盖该方法，则该方法返回的信息与`getMessage()`返
回的结果相同
- `void printStackTrace()`: 在控制台上打印`Throwable`对象封装的异常信息

## 异常的处理方式
#### 遇到问题不进行具体处理，而是继续抛给调用者 （throw,throws）
抛出异常有三种形式，一是`throw`,一个`throws`，还有一种系统自动抛异常。

### `try catch finally`捕获异常针对性处理方式
- `try`块：用于捕获异常。其后可接零个或多个`catch`块，如果没有`catch`块，则必须跟一个`finally`块。
- `catch`块：用于处理`try`捕获到的异常。
- `finally`块：⽆论是否捕获或处理异常，`finally`块⾥的语句都会被执行。当在`try`块或`catch`块中遇到`return`语句时，`finally`语句块将在方法返回之前被执行。
  
注意：不要在`finally`语句块中使用`return!`当`try`语句和`finally`语句中都有`return`语句时，`try`语句块中的`return`语句会被忽略。这是因为`try`语句中的`return`返回值会先被暂存在一个本地变量中，
当执行到`finally`语句中的`return`之后，这个本地变量的值就变为了`finally`语句中的`return`返回值。

#### `finally`中的代码一定会执行吗？
不一定的，在某些情况下，`finally`中的代码不会被执行。
就比如说`finally`之前虚拟机被终止运行的话，`finally`中的代码就不会被执行。
在以下 2 种特殊情况下，`finally`块的代码也不会被执行：
1. 程序所在的线程死亡。
2. 关闭`CPU`。

### 如何使用`try-with-resources`代替`try-catch-finally`？
面对必须要关闭的资源，我们总是应该优先使用`try-with-resources`而不是`try-finally`。随之产生的代码更简短，更清晰，产生的异常对我们也更有用。
`try-with-resources`语句让我们更容易编写必须要关闭的资源的代码，若采用`try-finally`则几乎做不到这点。

`Java`中类似于`InputStream`、`OutputStream`等的资源都需要我们调用`close()`方法来手动关闭，一般情况下我们都是通过`try-catch-finally`语句来实现这个需求

如下：
~~~java
public static void test() {
  File file=new File("/test.txt");
  InputStream is=null;
  try {
    is=Files.newInputStream(file.toPath());
    System.out.println("todo");
  } catch (IOException e) {
    e.printStackTrace();
  }finally {
    IOUtils.closeQuietly(is);
  }
}
~~~

使用`try-with-resources`语句改造上面的代码:
~~~java
public static void test() {
  File file=new File("/test.txt");
  InputStream is=null;
  try (InputStream is=Files.newInputStream(file.toPath())) {
    System.out.println("todo");
  } catch (IOException e) {
    e.printStackTrace();
  }
}
~~~
要注意的是：`try-with-resources`如果要声明多个资源，则需要分号`;`隔开。

### `throw`和`throws`的区别
#### 位置不同
- `throws`用在函数上，后面跟的是异常类，可以跟多个；而`throw`用在函数内，后面跟的是异常对象。

#### 功能不同：
- `throws`用来声明异常，让调用者只知道该功能可能出现的问题，可以给出预先的处理方式；
  `throw`抛出具体的问题对象，执行到`throw`，功能就已经结束了，跳转到调用者，并将具体的问题对象抛给调用者。
  也就是说`throw`语句独立存在时，下面不要定义其他语句，因为执行不到。
- `throws`表示出现异常的一种可能性，并不一定会发生这些异常；`throw`则是抛出了异常，执行`throw`则一定抛出了某种异常对象。

两者都是消极处理异常的方式，只是抛出或者可能抛出异常，但是不会由函数去处理异常，真正的处理异常由函数的上层调用处理。
