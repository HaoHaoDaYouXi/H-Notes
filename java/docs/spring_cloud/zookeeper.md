# Zookeeper
`Zookeeper`是一个分布式协调服务，可用于服务发现，分布式锁，分布式领导选举，配置管理等。

`Zookeeper`提供了一个类似于`Linux`文件系统的树形结构
（可认为是轻量级的内存文件系统，但只适合存少量信息，完全不适合存储大量文件或者大文件），
同时提供了对于每个节点的监控与通知机制。

据节点提交的反馈进行下一步合理操作。最终，将简单易用的接口和性能高效、功能稳定的系统提供给用户。

**分布式应用程序可以基于`Zookeeper`实现**
- 数据发布/订阅
- 负载均衡
- 命名服务
- 分布式协调/通知
- 集群管理
- `Master`选举
- 分布式锁和分布式队列
- 等

**`Zookeeper`保证了如下分布式一致性特性：**
- 顺序一致性
- 原子性
- 单一视图
- 可靠性
- 实时性（最终一致性）

客户端的读请求可以被集群中的任意一台机器处理，如果读请求在节点上注册了监听器，这个监听器也是由所连接的`zookeeper`机器来处理。
对于写请求，这些请求会同时发给其他`zookeeper`机器并且达成一致后，请求才会返回成功。
因此，随着`zookeeper`的集群机器增多，读请求的吞吐会提高但是写请求的吞吐会下降。

有序性是`zookeeper`中非常重要的一个特性，所有的更新都是全局有序的，每个更新都有一个唯一的时间戳，
这个时间戳称为`zxid`（Zookeeper Transaction Id）。
而读请求只会相对于更新有序，也就是读请求的返回结果中会带有这个`zookeeper`最新的`zxid`。

## Zookeeper 文件系统
`Zookeeper`提供一个多层级的节点命名空间（节点称为`znode`）。与文件系统不同的是，这些节点都可
以设置关联的数据，而文件系统中只有文件节点可以存放数据而目录节点不行。
`Zookeeper`为了保证高吞吐和低延迟，在内存中维护了这个树状的目录结构，这种特性使得`Zookeeper`
不能用于存放大量的数据，每个节点的存放数据上限为`1M`。

## Zookeeper 通知机制
`Zookeeper`允许客户端对服务端的某个`znode`注册一个`watcher`监听事件，当服务端的一些指定事件触发了这个`watcher`，
服务端会向指定客户端发送一个事件通知来实现分布式的通知功能，然后客户端根据`watcher`通知状态和事件类型做出业务上的改变。

**大致分为三个步骤：**

- 客户端注册`watcher`
  - 调用`getData`、`getChildren`、`exist`三个API ，传入`watcher`对象。 
  - 标记请求`request`，封装`watcher`到`WatchRegistration`。 
  - 封装成`Packet`对象，发服务端发送`request`。 
  - 收到服务端响应后，将`watcher`注册到`ZKWatcherManager`中进行管理。
  - 请求返回，完成注册。
- 服务端处理`watcher`
  - 服务端接收`watcher`并存储
  - `watcher`触发，调用`process`方法来触发`watcher`。
- 客户端回调`watcher`
  - 客户端`SendThread`线程接收事件通知，交由`EventThread`线程回调`watcher`。 
  - 客户端的`watcher`机制同样是一次性的，一旦被触发后，该`watcher`就失效了。

**总结**：客户端会对某个`znode`建立一个`watcher`事件，当该`znode`发生变化时，
这些客户端会收到`Zookeeper`的通知，然后客户端可以根据`znode`变化来做出业务上的改变等。

### Zookeeper 通知机制的特点

- 一次性触发数据发生改变时，一个`watcher event`会被发送到客户端，但是客户端只会收到一次这样的信息。
- `watcher event`异步发送`watcher`的通知事件从服务端发送到客户端是异步的，这就存在一个问题，
  - 不同的客户端和服务器之间通过`socket`进行通信，由于网络延迟或其他因素导致客户端在不通的时刻监听到事件，
  - 由于`Zookeeper`本身提供了`ordering guarantee`，即客户端监听事件后，才会感知它所监视`znode`发生了变化。
  - 所以我们使用`Zookeeper`不能期望能够监控到节点每次的变化。
  - `Zookeeper`只能保证最终的一致性，而无法保证强一致性。
- 数据监视`Zookeeper`有数据监视和子数据监视`getData()`、`exists()`设置数据监视，`getchildren()`设置了子节点监视。
- 注册`watcher`，`getData`、`exists`、`getChildren`
- 触发`watcher`，`create`、`delete`、`setData`
- `setData()`会触发`znode`上设置的`data watch`（如果`set`成功的话）。
  - 一个成功的`create()`操作会触发被创建的`znode`上的数据`watch`，以及其父节点上的`child watch`。
  - 一个成功的`delete()`操作将会同时触发一个`znode`的`data watch`和`child watch`（因为这样就没有子节点了），同时也会触发其父节点的`child watch`。
- 当一个客户端连接到一个新的服务器上时，`watch`将会被以任意会话事件触发。
  - 当与一个服务器失去连接的时候，是无法接收到`watch`的。而当客户端重新连接时，如果需要的话，所有先前注册过的`watch`，都会被重新注册。通常这是完全透明的。
  - 有在一个特殊情况下，`watch`可能会丢失：对于一个未创建 的`znode`的`exist watch`，如果在客户端断开连接期间被创建了，
    并且随后在客户端连接上之前又删除了，这种情况下，这个`watch`事件可能会被丢失。
- `watch`是轻量级的，其实就是本地`JVM`的`Callback`，服务器端只是存了是否有设置了`watcher`的布尔类型。

## Zookeeper节点ZNode和相关属性

`ZNode`有两种类型 ：
- 持久的（PERSISTENT）：客户端和服务器端断开连接后，创建的节点不删除（默认）。
- 短暂的（EPHEMERAL）：客户端和服务器端断开连接后，创建的节点自己删除。

`ZNode`有四种形式：
- PERSISTENT-持久节点
  - 客户端与`Zookeeper`断开连接后，除非手动删除，否则节点一直存在于`Zookeeper`上
- PERSISTENT_SEQUENTIAL-持久顺序节点
  - 基本特性同持久节点，只是增加了顺序属性，节点名后边会追加一个由父节点维护的自增整型数字。
- EPHEMERAL-临时节点
  - 临时节点的生命周期与客户端会话绑定，一旦客户端会话失效（客户端与`Zookeeper`连接断开不一定会话失效），
    那么这个客户端创建的所有临时节点都会被移除。
- EPHEMERAL_SEQUENTIAL-临时顺序节点
  - 基本特性同临时节点，增加了顺序属性，节点名后边会追加一个由父节点维护的自增整型数字。

创建`ZNode`时设置顺序标识，`ZNode`名称后会附加一个值，顺序号是一个单调递增的计数器，由父节点维护。

**节点属性**

`znode`节点不仅可以存储数据，还有一些其他特别的属性。

| 节点属性           | 注解                                                    |
|----------------|-------------------------------------------------------|
| cZxid          | 该数据节点被创建时的事务Id                                        |
| ctime          | 该数据节点创建时间                                             |
| mZxid          | 该数据节点被修改时最新的事物Id                                      |
| mtime          | 该数据节点最后修改时间                                           |
| pZxid          | 当前节点的父级节点事务Id                                         |
| cversion       | 子节点版本号(子节点修改次数，每修改一次值+1递增)                            |
| dataVersion    | 当前节点版本号(每修改一次值+1递增)                                   |
| aclVersion     | 当前节点acl版本号(节点被修改acl权限，每修改一次值+1递增)                     |
| ephemeralOwner | 临时节点标示，当前节点如果是临时节点，则存储的创建者的会话id(sessionId)，如果不是，那么值=0 |
| dataLength     | 当前节点所存储的数据长度                                          |
| numChildren    | 当前节点下子节点的个数                                           |

## Zookeeper 集群中的角色
`Zookeeper`集群是一个基于主从复制的高可用集群，每个服务器承担如下三种角色中的一种

- `Leader`
  - 一个`Zookeeper`集群同一时间只会有一个实际工作的`Leader`，它会发起并维护与各`Follower`及`Observer`间的心跳。
  - 所有的写操作必须要通过`Leader`完成再由`Leader`将写操作广播给其它服务器。
    只要有超过半数节点（不包括`observer`节点）写入成功，该写请求就会被提交（类`2PC`协议）。
- `Follower`
  - 一个`Zookeeper`集群可能同时存在多个`Follower`，它会响应`Leader`的心跳，
  - `Follower`可直接处理并返回客户端的读请求，同时会将写请求转发给`Leader`处理，并且负责在`Leader`处理写请求时对请求进行投票。

- `Observer`
  - 与`Follower`类似，但是无投票权。
  - `Zookeeper`需保证高可用和强一致性，为了支持更多的客户端，需要增加更多`Server`
    - `Server`增多，投票阶段延迟增大，影响性能
  - 引入`Observer`，`Observer`不参与投票
    - `Observers`接受客户端的连接，并将写请求转发给`leader`节点
    - 加入更多`Observer`节点，提高伸缩性，同时不影响吞吐率。

## Zookeeper集群中Server工作状态
- `LOOKING`
  - 寻找`Leader`状态；当服务器处于该状态时，它会认为当前集群中没有`Leader`，因此需要进入`Leader`选举状态
- `FOLLOWING`
  - 跟随者状态；表明当前服务器角色是`Follower`
- `LEADING`
  - 领导者状态；表明当前服务器角色是`Leader`
- `OBSERVING`
  - 观察者状态；表明当前服务器角色是`Observer`

## ZooKeeper集群中服务器之间通信
`Leader`服务器会和每一个`Follower/Observer`服务器都建立`TCP`连接，
同时为每个`Follower/Observer`都创建一个叫做`LearnerHandler`的实体。

`LearnerHandler`主要负责`Leader`和`Follower/Observer`之间的网络通讯，包括数据同步，请求转发和`proposal`提议的投票等。

`Leader`服务器保存了所有`Follower/Observer`的`LearnerHandler`。

## ZAB 协议

### 事务编号`Zxid`（事务请求计数器 + epoch）
在`ZAB`(`ZooKeeper Atomic Broadcast`，`ZooKeeper`原子消息广播协议）协议的事务编号`Zxid`设计中，
`Zxid`是一个`64`位的数字，其中低`32`位是一个简单的单调递增的计数器，针对客户端每一个事务请求，计数器加`1`。
而高`32`位则代表`Leader`周期`epoch`的编号，每个当选产生一个新的`Leader`服务器，
就会从这个`Leader`服务器上取出其本地日志中最大事务的`Zxid`，并从中读取`epoch`值，然后加`1`，
以此作为新的`epoch`，并将低`32`位从`0`开始计数。

`Zxid`（`Transaction id`）类似于`RDBMS`中的事务`ID`，用于标识一次更新操作的`Proposal`（提议）`ID`。
为了保证顺序性，该`id`必须单调递增。

### epoch
`epoch`：可以理解为当前集群所处的年代或者周期，每个`Leader`就像皇帝，都有自己的年号，
所以每次改朝换代，`leader`变更之后，都会在前一个年代的基础上加`1`。
这样就算旧的`Leader`崩溃恢复之后，也没有人听他的了，因为`Follower`只听从当前年代的`Leader`的命令。

### Zab 协议有两种模式-恢复模式（选主）、广播模式（同步）
`Zab`协议有两种模式，它们分别是`恢复模式（选主）`和`广播模式（同步）`。
当服务启动或者在领导者崩溃后，`Zab`就进入了恢复模式，当领导者被选举出来，
且大多数`Server`完成了和`Leader`的状态同步以后，恢复模式就结束了。
状态同步保证了`Leader`和`Server`具有相同的系统状态。

### ZAB协议4阶段
- `Leader election`（选举阶段-选出准`Leader`）
  - 节点在一开始都处于选举阶段，只要有一个节点得到超半数节点的票数，它就可以当选准`Leader`。
  - 只有到达广播阶段（`broadcast`）准`leader`才会成为真正的`leader`。
  - 这一阶段的目的是就是为了选出一个准`leader`，然后进入下一个阶段。
- `Discovery`（发现阶段-接受提议、生成`epoch`、接受`epoch`）
  - 在这个阶段，`Followers`跟准`Leader`进行通信，同步`Followers`最近接收的事务提议。
  - 这个一阶段的主要目的是发现当前大多数节点接收的最新提议，并且准`Leader`生成新的`epoch`，
    让`Followers`接受，更新它们的`accepted Epoch`
  - 一个`Follower`只会连接一个`Leader`，如果有一个节点`f`认为另一个`Follower p`是`Leader`，
    `f`在尝试连接`p`时会被拒绝，`f`被拒绝之后，就会进入重新选举阶段。
- `Synchronization`（同步阶段-同步`Follower`副本）
  - 同步阶段主要是利用`Leader`前一阶段获得的最新提议历史，同步集群中所有的副本。
  - 只有当大多数节点都同步完成，准`Leader`才会成为真正的`Leader`。
  - `Follower`只会接收`Zxid`比自己的`lastZxid`大的提议。
- `Broadcast`（广播阶段-`Leader`消息广播）
  - 到了这个阶段，`Zookeeper`集群才能正式对外提供事务服务，并且`Leader`可以进行消息广播。
  - 同时如果有新的节点加入，还需要对新节点进行同步。

`ZAB`提交事务并不像`2PC`一样需要全部`Follower`都`ACK`，只需要得到超过半数的节点的`ACK`就可以了。

### `ZAB`协议`JAVA`实现（FLE-发现阶段和同步合并为`Recovery Phase`（恢复阶段））

协议的`Java`版本实现跟上面的定义有些不同，选举阶段使用的是`Fast Leader Election`（FLE），它包含了选举的发现职责。

因为`FLE`会选举拥有最新提议历史的节点作为`Leader`，这样就省去了发现最新提议的步骤。

实际的实现将`发现阶段`和`同步`合并为`Recovery Phase`（恢复阶段）。

所以，`ZAB`的实现只有三个阶段：`Fast Leader Election`、`Recovery Phase`、`Broadcast Phase`。


#

----
