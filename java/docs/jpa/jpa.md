# JPA

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

## <div id="cbxw">事物的传播行为</div>
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






----
