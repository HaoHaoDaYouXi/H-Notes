# Spring Data JPA

`Spring Data JPA`，事务是计算机应用中不可或缺的组件模型，它保证了用户操作的原子性 (`Atomicity`)、一致性(`Consistency`)、隔离性 (`Isolation`) 和持久性 (`Durability`)。

> 原子性（`Atomicity`）
>> 原子性，是指事务包含的所有操作，要么全部成功，要么全部失败回滚。因此，事务的操作如果成功就必须要完全持久化到数据库，如果操作失败则不能对数据库有任何影响。

> 一致性（`Consistency`）
>> 一致性，是指事务必须使数据库从一个一致性状态变换到另一个一致性状态，也就是说一个事务执行之前和执行之后都必须处于一致性状态。</br>
>> 例如：假设A和B两者的钱加起来一共是100，那么不管A和B之间如何转账，转几次账，事务结束后两个用户的钱相加起来应该还是100，这即是事务的一致性。

> 隔离性（`Isolation`）
>> 隔离性，是当多个用户并发访问数据库时，比如操作同一张表时，数据库为每一个用户开启的事务，不能被其他事务的操作所干扰，多个并发事务之间要相互隔离。</br>
>> 即要达到这么一种效果：对于任意两个并发的事务T1和T2，在事务T1看来，T2要么在T1开始之前就已经结束，要么在T1结束之后才开始，这样每个事务都感觉不到有其他事务在并发地执行。

> 持久性（`Durability`）
>> 持久性，是指一个事务一旦被提交了，那么对数据库中的数据的改变就是永久性的，即便是在数据库系统遇到故障的情况下也不会丢失提交事务的操作。</br>
>> 例如我们在使用JDBC操作数据库时，在提交事务方法后，提示用户事务操作完成，当我们程序执行完成直到看到提示后，就可以认定事务以及正确提交，即使这时候数据库出现了问题，也必须要将我们的事务完全执行完成，否则就会造成我们看到提示事务处理完毕，但是数据库因为故障而没有执行事务的重大错误。

## <div id="cbxw">Spring事物的传播行为</div>
事务的传播行为说的是，当多个事务同时存在的时候，`Spring`如何处理这些事务的行为。
- `PROPAGATION_REQUIRED`：如果当前没有事务，就创建一个新事务，如果当前存在事务，就加入该事务，`Spring`默认事务级别。
  - 执行`ServiceA.methodA`的时候，`ServiceA.methodA`已经起了事务，这时调用`ServiceB.methodB`，</br>
    `ServiceB.methodB`看到自己已经运行在`ServiceA.methodA`的事务内部，就不再起新的事务。</br>
    而假如`ServiceA.methodA`运行的时候发现自己没有在事务中，他就会为自己分配一个事务。</br>
    这样，在`ServiceA.methodA`或者在`ServiceB.methodB`内的任何地方出现异常，事务都被回滚。</br>
    即使`ServiceB.methodB`的事务已经被提交，但是`ServiceA.methodA`在接下来`error`要回滚，`ServiceB.methodB`也要回滚。
- `PROPAGATION_SUPPORTS`：支持当前事务，如果当前存在事务，就加入该事务，如果当前不存在事务，就以非事务执行。
- `PROPAGATION_MANDATORY`：支持当前事务，如果当前存在事务，就加入该事务，如果当前不存在事务，就抛出异常。
- `PROPAGATION_REQUIRES_NEW`：创建新事务，无论当前存不存在事务，都创建新事务。
  - `ServiceA.methodA`的事务级别为`PROPAGATION_REQUIRED`，`ServiceB.methodB`的事务级别为`PROPAGATION_REQUIRES_NEW`，</br>
    那么当执行到`ServiceB.methodB`的时候，`ServiceA.methodA`所在的事务就会挂起，`ServiceB.methodB`会起一个新的事务，</br>
    等待`ServiceB.methodB`的事务完成以后，`A`才继续执行。他与`PROPAGATION_REQUIRED`的事务区别在于事务的回滚程度了。</br>
    因为`ServiceB.methodB`是新起一个事务，那么就是存在两个不同的事务。</br>
    如果`ServiceB.methodB`已经提交，那么`ServiceA.methodA`失败回滚，`ServiceB.methodB`是不会回滚的。</br>
    如果`ServiceB.methodB`失败回滚，如果他抛出的异常被`ServiceA.methodA`捕获，`ServiceA.methodA`事务仍然可能提交。
- `PROPAGATION_NOT_SUPPORTED`：以非事务方式执行操作，如果当前存在事务，就把当前事务挂起。
- `PROPAGATION_NEVER`：以非事务方式执行，如果当前存在事务，则抛出异常。
- `PROPAGATION_NESTED`：如果当前存在事务，则在嵌套事务内执行。如果当前没有事务，则按`PROPAGATION_REQUIRED`属性执行。
  - 与`PROPAGATION_REQUIRES_NEW`的区别是`NESTED`的事务和他的父事务是相依的，它的提交是要等父事务一块提交。也就是说，如果父事务最后回滚，它也要回滚。

## <div id="gljb">事物的隔离级别</div>
在介绍数据库提供的各种隔离级别之前，先看看如果不考虑事务的隔离性，会发生的几种问题：
> 脏读
>> 脏读是指在一个事务处理过程里读取了另一个未提交的事务中的数据。</br>
>> 当一个事务正在多次修改某个数据，而在这个事务中这多次的修改都还未提交，这时一个并发的事务来访问该数据，就会造成两个事务得到的数据不一致。

> 不可重复读
>> 不可重复读是指在对于数据库中的某个数据，一个事务范围内多次查询却返回了不同的数据值，这是由于在查询间隔，被另一个事务修改并提交了。</br>
>> 不可重复读和脏读的区别是，脏读是某一事务读取了另一个事务未提交的脏数据，而不可重复读则是读取了前一事务提交的数据。</br>
>> 在某些情况下，不可重复读并不是问题，比如我们多次查询某个数据当然以最后查询得到的结果为主。但在另一些情况下就有可能发生问题，例如对于同一个数据A和B依次查询就可能不同，A和B就可能打起来了……

> 幻读(虚读)
>> 幻读是事务非独立执行时发生的一种现象。</br>
>> 如事务T1对一个表中所有的行的某个数据项做了从“1”修改为“2”的操作，这时事务T2又对这个表中插入了一行数据项，而这个数据项的数值还是为“1”并且提交给数据库。</br>
>> 而操作事务T1的用户如果再查看刚刚修改的数据，会发现还有一行没有修改，其实这行是从事务T2中添加的，就好像产生幻觉一样，这就是发生了幻读。</br>
>> 幻读和不可重复读都是读取了另一条已经提交的事务（这点就脏读不同），所不同的是不可重复读查询的都是同一个数据项，而幻读针对的是一批数据整体（比如数据的个数）

关于事务的隔离性数据库提供了多种隔离级别，`MySQL`数据库为我们提供的四种隔离级别：
- `Serializable`(串行化)：可避免脏读、不可重复读、幻读的发生。
- `Repeatable read`(可重复读)：可避免脏读、不可重复读的发生。
- `Read committed`(读已提交)：可避免脏读的发生。
- `Read uncommitted`(读未提交)：最低级别，任何情况都无法保证。

以上四种隔离级别最高的是`Serializable`级别，最低的是`Read uncommitted`级别，当然级别越高，执行效率就越低。
像`Serializable`这样的级别，就是以锁表的方式(类似于`Java`多线程中的锁)使得其他的线程只能在锁外等待，所以平时选用何种隔离级别应该根据实际情况。

**`MySQL`数据库中默认的隔离级别为`Repeatable read`(可重复读)。**

在`MySQL`数据库中，支持上面四种隔离级别，默认的为`Repeatable read`(可重复读)；而在`Oracle`数据库中，只支持`Serializable`(串行化)级别和`Read committed`(读已提交)这两种级别，其中默认的为`Read committed`级别。

查看`mysql`数据库事务的默认隔离级别的`SQL`：
```sql
SELECT @@tx_isolation;
```

## <div id="glsw">`Spring`管理事务</div>

`Spring`事务管理主要包括3个接口：

- `PlatformTransactionManager`：事务管理器，主要用于平台相关事务的管理。主要包括三个方法：
  - `commit`：事务提交。
  - `rollback`：事务回滚。
  - `getTransaction`：获取事务状态。
- `TransactionDefinition`：事务定义信息，用来定义事务相关属性，给事务管理器`PlatformTransactionManager`使用这个接口有下面四个主要方法：
  - `getIsolationLevel`：获取隔离级别。
  - `getPropagationBehavior`：获取传播行为。
  - `getTimeout`获取超时时间。
  - `isReadOnly`：是否只读（保存、更新、删除时属性变为`false`--可读写，查询时为`true`--只读）事务管理器能够根据这个返回值进行优化，这些事务的配置信息，都可以通过配置文件进行配置。
- `TransactionStatus`：事务具体运行状态，事务管理过程中，每个时间点事务的状态信息。例如：
  - `hasSavepoint()`：返回这个事务内部是否包含一个保存点。
  - `isCompleted()`：返回该事务是否已完成，也就是说，是否已经提交或回滚。
  - `isNewTransaction()`：判断当前事务是否是一个新事务。

## <div id="sxfs">事物的实现方式</div>

### <div id="bdsw">本地事务</div>
紧密依赖于底层资源管理器(例如数据库连接)，事务处理局限在当前事务资源内。</br>
此种事务处理方式不存在对应用服务器的依赖，因而部署灵活却无法支持多数据源的分布式事务。</br>

在数据库连接中使用本地事务示例如下：
```java
public void transferAccount() {  
    Connection conn = null;  
    Statement stmt = null;  
    try{
        conn = getDataSource().getConnection();  
        // 将自动提交设置为 false，若设置为 true 则数据库将会把每一次数据更新认定为一个事务并自动提交
        conn.setAutoCommit(false);
        stmt = conn.createStatement();  
        // 将 A 账户中的金额减少 100  
        stmt.execute("update u_account set amount = amount - 100 where account_id = 'A'");
        // 将 B 账户中的金额增加 100  
        stmt.execute("update u_account set amount = amount + 100 where account_id = 'B'");
        // 提交事务
        conn.commit();
        // 事务提交：转账的两步操作同时成功
    } catch(SQLException sqle){     
        // 发生异常，回滚在本事务中的操做
        conn.rollback();
        // 事务回滚：转账的两步操作完全撤销
        stmt.close();  
        conn.close();  
    }  
}
```

### <div id="fbssw">分布式事务</div>
`Java`事务编程接口(`JTA`：`Java Transaction API`)和`Java`事务服务 (`JTS`：`Java Transaction Service`)为`J2EE`平台提供了分布式事务服务。

分布式事务(`Distributed Transaction`)包括事务管理器(`Transaction Manager`)和一个或多个支持`XA`协议的资源管理器 (`Resource Manager`)。

我们可以将资源管理器看做任意类型的持久化数据存储；事务管理器承担着所有事务参与单元的协调与控制。
```java
public void transferAccount() {  
    UserTransaction userTx = null;  
    Connection connA,connB = null; 
    Statement stmtA,stmtB = null;   
    try{  
        // 获得 Transaction 管理对象
        userTx = (UserTransaction)getContext().lookup("java:comp/UserTransaction");
        // 从数据库 A 中取得数据库连接
        connA = getDataSourceA().getConnection();
        // 从数据库 B 中取得数据库连接
        connB = getDataSourceB().getConnection();
        // 启动事务
        userTx.begin();   
        stmtA = connA.createStatement();// 将 A 账户中的金额减少 100  
        stmtA.execute("update u_account set amount = amount - 100 where account_id = 'A'");
        // 将 B 账户中的金额增加 100  
        stmtB = connB.createStatement();
        stmtB.execute("update u_account set amount = amount + 100 where account_id = 'B'");
        // 提交事务   
        userTx.commit();
        // 事务提交：转账的两步操作同时成功（数据库 A 和数据库 B 中的数据被同时更新）
    } catch(SQLException sqlE){  
        // 发生异常，回滚在本事务中的操纵
        userTx.rollback();// 事务回滚：数据库 A 和数据库 B 中的数据更新被同时撤销
    } catch(Exception ne){
        // 发生异常，回滚在本事务中的操纵
        userTx.rollback();// 事务回滚：数据库 A 和数据库 B 中的数据更新被同时撤销
    }  
}
```

----
