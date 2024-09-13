# MySQL

`MySQL`是一种关系型数据库，在`Java`企业级开发中⾮常常用，因为`MySQL`是开源免费的，并且方便扩展。

## <a id="tx">特性</a>

- 开源与成本效益：`MySQL`是开源的，这意味着它可以免费使用，降低了软件许可的成本。
- 跨平台兼容性：`MySQL`可在多种操作系统上运行，包括`Windows`、`Linux`和`macOS`。
- 高性能：`MySQL`通过优化的存储引擎和内存管理提供了高效率的数据处理能力。
- 安全性：`MySQL`提供了强大的安全功能，如用户权限管理和加密数据传输。
- 可扩展性：`MySQL`支持从小型数据库到大型集群的多种部署方式。
- `ACID`遵守：事务处理遵循`ACID`（原子性、一致性、隔离性、持久性）原则，确保数据完整性。
- `SQL`标准支持：`MySQL`支持`SQL92`和`SQL99`标准，以及一些扩展功能。
- 存储引擎：`MySQL`支持多种存储引擎，如`InnoDB`、`MyISAM`和`MEMORY`，每种引擎具有不同的特性和用途。
- 复制和高可用性：`MySQL`支持主从复制和群集配置，以提高可用性和容错性。
- 社区与文档：`MySQL`拥有庞大的开发者社区和详尽的官方文档，易于学习和解决问题。

## <a id="gcbf">构成部分</a>

- 连接器：身份认证和权限相关(登录`MySQL`的时候)。
- 查询缓存：执行查询语句的时候，会先查询缓存（`MySQL 8.0`版本后移除，因为这个功能不太实用）。
- 分析器：没有命中缓存的话，`SQL`语句就会经过分析器，分析器说⽩了就是要先看你的`SQL`语句要⼲嘛，再检查你的`SQL`语句语法是否正确。
- 优化器：按照`MySQL`认为最优的方案去执行。
- 执行器：执行语句，然后从存储引擎返回数据。 执行语句之前会先判断是否有权限，如果没有权限的话，就会报错。
- 插件式存储引擎：主要负责数据的存储和读取，采用的是插件式架构，⽀持`InnoDB`、`MyISAM`、`Memory`等多种存储引擎。

## <a id="sfs">三范式</a>

- 第一范式：每个列都不可以再拆分。
- 第二范式：非主键列完全依赖于主键，而不能是依赖于主键的一部分。
- 第三范式：非主键列只依赖于主键，不依赖于其他非主键。

在设计数据库结构的时候，要尽量遵守三范式，如果不遵守，必须有足够的理由，比如性能。
事实上经常会为了性能而妥协数据库的设计。

## <a id="ccyq">存储引擎</a>

可以通过`select version()`命令查看你的`MySQL`版本。

`MySQL`⽀持多种存储引擎，可以通过`show engines`命令来查看`MySQL`⽀持的所有存储引擎。

也可以通过`show variables like '%storage_engine%'`命令直接查看`MySQL`当前默认的存储引擎。

如果你只想查看数据库中某个表使用的存储引擎的话，可以使用`show table status from db_name where name='table_name'`命令。

`MySQL 5.5.5`之前，`MyISAM`是`MySQL`的默认存储引擎。`5.5.5`版本之后，`InnoDB`是`MySQL`的默认存储引擎。
所有的存储引擎中只有`InnoDB`是事务性存储引擎，只有`InnoDB`⽀持事务。

## <a id="gljb">隔离级别</a>

`MySQL`数据库为我们提供的四种隔离级别：

- `Serializable`(串行化)：可避免脏读、不可重复读、幻读的发生。
- `Repeatable read`(可重复读)：可避免脏读、不可重复读的发生。
- `Read committed`(读已提交)：可避免脏读的发生。
- `Read uncommitted`(读未提交)：最低级别，任何情况都无法保证。

`InnoDB`存储引擎的默认⽀持的隔离级别是`Repeatable read`（可重复读）。

可以通过`SELECT @@tx_isolation`命令来查看，
`MySQL 8.0`该命令改为`SELECT @@transaction_isolation`

## <a id="myisaminnodb">`MyISAM`和`InnoDB`的区别</a>

| 区别                                    | MyISAM                           | Innodb                |
|---------------------------------------|----------------------------------|-----------------------|
| 文件格式                                  | 数据和索引是分别存储的数据`.MYD`，索引`.MYI`     | 数据和索引是集中存储的，`.ibd`    |
| 文件能否移动                                | 能，一张表就对应`.frm`、`MYD`、`MYI`3个文件   | 否，因为关联的还有`data`下的其它文件 |
| 记录存储顺序                                | 按记录插入顺序保存                        | 按主键大小有序插入             |
| 空间碎片(删除记录并`flush table`表名 之后，表文件大小不变) | 产生，定时整理，使用命令`optimize table`表名实现 | 不产生                   |
| 事务                                    | 不支持                              | 支持                    |
| 外键                                    | 不支持                              | 支持                    |
| 全文索引                                  | 支持                               | 不支持                   |
| 锁支持(锁是避免资源争用的一个机制，MySQL锁对用户几乎是透明的)    | 表级锁                              | 行级锁、表级锁，锁定粒度小并发能力高    |
| MVCC(多版本并发控制)                         | 不支持                              | 支持                    |

## <a id="bjsyjshjs">表级锁、页级锁和行级锁</a>

### 表级锁

`MySQL`中锁定粒度最⼤的一种锁，是针对⾮索引字段加的锁，对当前操作的整张表加锁，实现简单，资源消耗也比较少，加锁快，不会出现死锁。
其锁定粒度最⼤，触发锁冲突的概率最⾼，并发度最低，`MyISAM`和`InnoDB`引擎都⽀持表级锁。

### 页级锁

开销和加锁时间界于表锁和行锁之间，会出现死锁，锁定粒度界于表锁和行锁之间，并发度一般

### 行级锁

`MySQL`中锁定粒度最小的一种锁，是针对索引字段加的锁，只针对当前操作的行记录进行加锁。行级锁能⼤⼤减少数据库操作的冲突。
其加锁粒度最小，并发度⾼，但加锁的开销也最⼤，加锁慢，会出现死锁。

`InnoDB`的行锁是针对索引字段加的锁，表级锁是针对⾮索引字段加的锁。

当我们执行`UPDATE`、`DELETE`语句时，如果`WHERE`条件中字段没有命中唯一索引或者索引失效的话，
就会导致扫描全表对表中的所有行记录进行加锁，一定要多注意

不过，很多时候即使用了索引也有可能会⾛全表扫描，这是因为`MySQL`优化器的原因。

## <a id="gxspts">共享锁和排他锁</a>

不论是表级锁还是行级锁，都存在共享锁（`Share Lock`，`S`锁）和排他锁（`Exclusive Lock`，`X`锁）

- 共享锁（`S`锁）：又称读锁，事务在读取记录的时候获取共享锁，允许多个事务同时获取（锁兼容）。
- 排他锁（`X`锁）：又称写锁/独占锁，事务在修改记录的时候获取排他锁，不允许多个事务同时获取。
  如果一个记录已经被加了排他锁，那其他事务不能再对这条事务加任何类型的锁（锁不兼容）。

排他锁与任何的锁都不兼容，共享锁仅和共享锁兼容。

`MVCC`的存在，对于一般的`SELECT`语句，`InnoDB`不会加任何锁。

可以通过以下语句显式加共享锁或排他锁
```sql
-- 共享锁
SELECT ... LOCK IN SHARE MODE;
-- 排他锁
SELECT ... FOR UPDATE;
```

## <a id="yxs">意向锁</a>

意向锁的可以快速判断是否可以对某个表使用表锁

意向锁是表级锁，共有两种

- 意向共享锁（`Intention Shared Lock`，`IS`锁）
  - 事务有意向对表中的某些加共享锁（`S`锁），加共享锁前必须先取得该表的`IS`锁。

- 意向排他锁（`Intention Exclusive Lock`，`IX`锁）
  - 事务有意向对表中的某些记录加排他锁（`X`锁），加排他锁之前必须先取得该表的`IX`锁。

意向锁是由数据引擎维护的，⽆法手动操作意向锁，在为数据行加共享锁、排他锁之前，`InnoDB`会先获取该数据行所在在数据表的对应意向锁。

意向锁之间是互相兼容的。

意向锁和表级的共享锁和排他锁互斥，行级的共享锁和排他锁不互斥。

## <a id="hs">`InnoDB`行锁</a>

`InnoDB`⽀持三种行锁定方式：
- 记录锁（`Record Lock`）：也被称为记录锁，属于单个行记录上的锁。
- 间隙锁（`Gap Lock`）：锁定一个范围，不包括记录本身。
- 临键锁（`Next-key Lock`）：`Record Lock`+`Gap Lock`，锁定一个范围，包含记录本身，记录锁锁已经存在的，间隙锁锁新插入的。

`InnoDB`的默认隔离级别`RR`（可重读）是可以解决幻读问题发生的，主要有下面两种情况：

- 快照读（一致性⾮锁定读）：由`MVCC`机制来保证不出现幻读。
- 当前读（一致性锁定读）：使用`Next-Key Lock`进行加锁来保证不出现幻读。

### 当前读和快照读

快照读（一致性⾮锁定读）就是单纯的`SELECT`语句，不包括下面的`SELECT`语句：
```sql
SELECT ... FOR UPDATE
SELECT ... LOCK IN SHARE MODE
```
快照即记录的历史版本，每行记录可能存在多个历史版本（多版本技术）。

快照读的情况下，如果读取的记录正在执行 UPDATE/DELETE 操作，读取操作不会因此去等待记录上`X`锁的释放，而是会去读取行的一个快照。

只有在事务隔离级别`RC`（读取已提交）和`RR`（可重读）下，`InnoDB`才会使用一致性⾮锁定读：

- 在`RC`级别下，对于快照数据，一致性⾮锁定读总是读取被锁定行的最新一份快照数据。
- 在`RR`级别下，对于快照数据，一致性⾮锁定读总是读取本事务开始时的行数据版本。

快照读比较适合对于数据一致性要求不是特别⾼且追求极致性能的业务场景。

当前读 （一致性锁定读）就是给行记录加`X`锁或`S`锁。

当前读的一些常⻅`SQL`语句类型如下：
```sql
-- 对读的记录加一个X锁
SELECT...FOR UPDATE
-- 对读的记录加一个S锁
SELECT...LOCK IN SHARE MODE
-- 对修改的记录加一个X锁
INSERT...
UPDATE...
DELETE...
```

### `MVCC`机制

`MVCC`(Multi-Version Concurrency Control)叫多版本并发控制，是`InnoDB`存储引擎中用于处理事务并发的关键机制之一。
`MVCC`允许在读取数据的同时进行更新操作，从而提高了系统的并发性能。

#### 基本原理

`MVCC`主要通过记录多个版本的数据来支持并发读取和写入操作

- 当事务读取一行数据时，它可以看到符合其事务开始时刻的数据版本
- 当事务更新一行数据时，`InnoDB`会保存旧版本的数据，并创建一个新的版本

这样，不同的事务可以看到不同版本的数据，从而避免了数据冲突。

#### 实现

`MVCC`的实现依赖于：隐式字段、`Undo log`（撤销日志）、`Read View`（读视图）

**隐式字段**

在内部，`InnoDB`向数据库中存储的每一行添加三个字段：

- `DB_ROW_ID`：`6 byte`，隐藏的自增 ID。（如果数据表中没有主键，那么InnoDB会自动生成单调递增的隐藏主键（表中有主键或者非NULL的UNIQUE键时都不会包含 DB_ROW_ID列））
- `DB_TRX_ID` ：`6 byte`，插入或更新行的最后一个事务ID。（用于MVCC的ReadView判断事务id, 删除在内部被视为更新，其中行中的一个特殊位被设置为将其标记为已删除）
- `DB_ROLL_PTR`：`7 byte`，回滚指针。（用于MVCC中指向undo log记录，指向已写入回滚段(`rollback segment`)的一条`undo log`记录, 记录着行(`row`)更新前的副本）

**`undo log`（撤销日志）**

`undo log`是各个事务修改同一条记录的时候生成的历史记录，，这些记录保存在`undo log`里，这些日志通过回滚指针串联在一起，方便回滚，同时会生成一条版本链。

数据分为两类
- `Insert undo log`：`insert`生成的日志，仅在事务回滚中需要，并且可以在事务提交后立即丢弃。
- `Update undo log`：`update`、`delete`生成的日志，除了用于事务回滚，还用于一致性读取，只有不存在`innodb`为其分配快照的事务之后才能丢弃它们，在一致读取中可能需要`update undo log`中的信息来构建数据库行的早期版本。

删除操作实际上不会直接删除，而只是标记为删除，最终的删除操作是`purge`线程完成的

`InnoDB`中，事务中的`Delete`操作实际上并不是真正的删除掉数据行，而是一种`Delete Mark`操作，在记录上标识删除，真正的删除工作需要后台`purge`线程去完成。

`purge`线程作用
- 清理`undo log`
- 清除`page`里面带有`Delete_Bit`标识的数据行

使用`InnoDB`存储引擎的表，它的聚簇记录中包含
- `TRX_ID`：每次事务对聚簇记录进行修改的时候，就会将该事务的`id`复制给`TRX_ID`隐藏列
- `ROLL_PTR`：每次对每条聚簇索引进行改动的时候，都会将旧的版本信息写入`undo log`中，通过回滚指针就能找到记录修改前的信息。

`undo log`存储在`InnoDB`的内部数据结构中
- `undo`表空间
  - `undo log`存储在一个特殊的表空间中，称为`undo`表空间。
  - 通过配置参数`innodb_undo_directory`指定`undo`表空间所在的目录。
  - 通过配置参数`innodb_undo_logs`指定`undo`表空间中`undo`段的数量，默认为`128`。
- `undo`段
  - `undo log`是按照`undo`段来组织的。
  - 每个`undo`段包含多个页，每个页上存储着`undo log`记录。
  - 一个`undo`段可以容纳多个`undo log`记录，每个记录对应一个事务的操作。
- `undo`记录
  - 每个`undo log`记录都包含有关事务操作的信息，包括操作前的数据值、事务`ID`、回滚指针等。
  - 回滚指针指向同一个`undo`段中的前一个`undo log`记录，形成一个链表。

**`Read View`（读视图）**

`Read View`它代表了事务开始时可见的数据版本集合，用于确定哪些版本的数据对当前事务可见。

主要的字段：
- `m_low_limit_id`：尚未分配的最小事务`ID`，等于它的, 都不可见
- `m_up_limit_id`：最小活跃未提交事务`ID`，小于它的, 都可见
- `m_creator_trx_id`：创建`Read View`的事务`ID`，等于它的, 都可见
- `m_ids`：创建`Read View`时，正活跃未提交的事务`ids`，在`m_ids`里面不可见，否则可见

事务在读取数据时会检查数据的事务`ID`是否在`Read View`中，只有符合条件的数据版本才会被读取。

`m_low_limit_id`不是`m_ids`的最大值，而是系统能够分配的事务`ID`最大值，事务`ID`是递增分配的，并且只有事务在进行增删改操作的时候才会分配事务`ID`。

如：有`1、2、3`三个事务，`3`的事务提交后，一个新事务在生成`Read View`的时候，`m_ids`里是`1、2`，`m_up_limit_id`是`1`，`m_low_limit_id`就是`4`

**`Read View`的判断流程**：当查询一条数据的时候
- 首先获取查询操作的事务的版本号
- 获取当前系统的`Read View`
- 将查询到的数据与`Read View`中的事务版本号进行比较
- 如果不符合`Read View`的规则，则通过回滚指针形成的`undo log`版本链从`undo log`中获取符合规则的历史快照
- 返回符合规则的数据

**`MVCC`的行为受到事务隔离级别的影响，不同隔离级别使用`Read View`**
- 读未提交：能够读取未提交的事务修改的数据，所以直接读取最新的记录就可以，不必使用`MVCC`。
- 读已提交：不能读取未提交的事务修改的数据，并且不能进行重复读取，事务中，每次快照读都会新生成一个快照和`Read View`，这就是在`RC`级别下的事务中可以看到别的事务提交的更新的原因。
- 可重复读：不能读取未提交的事务修改的数据，并且能进行重复读取，所以只在第一次查询的时候获取一次`Read View`，之后查询都只查看已经生成的`Read View`副本。
- 可串行化：`MVCC`被禁用，`InnoDB`规定使用加锁的方式来访问记录，通过加锁的方式让所有`sql`都串行化执行了，也是读最新的，不存在快照读`Read View`。

**例：**
假设有一个简单的表`orders`，包含`id`和`status`两列，现在有两个事务`T1`和`T2`同时运行：
- 事务`T1`更新订单状态
  - `T1`开始事务。
  - `T1`更新订单状态：`UPDATE orders SET status = 'SHIPPED' WHERE id = 1;`
  - `T1`提交事务。
- 事务`T2`读取订单状态
- `T2`开始事务。
- `T2`读取订单状态：`SELECT * FROM orders WHERE id = 1;`
- `T2`读取到的数据取决于事务隔离级别。

示例代码：
```java
public class MVCCExample {

    public static void main(String[] args) throws SQLException {
        Connection conn1 = DriverManager.getConnection("jdbc:mysql://localhost:3306/testdb", "root", "password");
        Connection conn2 = DriverManager.getConnection("jdbc:mysql://localhost:3306/testdb", "root", "password");

        // 设置事务隔离级别为 REPEATABLE READ
        conn1.setTransactionIsolation(Connection.TRANSACTION_REPEATABLE_READ);
        conn2.setTransactionIsolation(Connection.TRANSACTION_REPEATABLE_READ);

        // 事务 T1
        try (PreparedStatement ps1 = conn1.prepareStatement("UPDATE orders SET status = ? WHERE id = ?")) {
            ps1.setString(1, "SHIPPED");
            ps1.setInt(2, 1);
            ps1.executeUpdate();
            conn1.commit();
        }

        // 事务 T2
        try (PreparedStatement ps2 = conn2.prepareStatement("SELECT * FROM orders WHERE id = ?")) {
            ps2.setInt(1, 1);
            ResultSet rs = ps2.executeQuery();

            while (rs.next()) {
                System.out.println("Order ID: " + rs.getInt("id"));
                System.out.println("Status: " + rs.getString("status"));
            }
        }

        conn1.close();
        conn2.close();
    }
}
```
示例中，我们创建了两个事务`T1`和`T2`。`T1`更新了一条订单记录的状态，而`T2`试图读取这条记录的状态。
事务隔离级别被设置为`REPEATABLE READ`，这意味着`T2`在其事务开始后不会看到`T1`的更改。

总结来说，`MVCC`是`InnoDB`存储引擎中用于处理并发读取和写入操作的关键机制。通过维护多个数据版本，它可以有效地支持高并发环境下的事务处理。

## <a id="sy">索引</a>

索引是一种数据结构，可以帮助我们快速的进行数据的查找。

索引的数据结构和具体存储引擎的实现有关，在`MySQL`中使用较多的索引有`Hash`索引，`B+`树索引等

`InnoDB`存储引擎的默认索引实现为：`B+`树索引

### 索引分类

- 单值索引：即一个索引只包含单个列，一个表可以有多个单列索引
  - 建表时，加上`key`(列名) 指定
  - 单独创建，`create index 索引名 on 表名(列名)`
  - 单独创建，`alter table 表名 add index 索引名(列名)`

- 唯一索引：索引列的值必须唯一，但允许有`null`且`null`可以出现多次
  - 建表时，加上`unique(列名)`指定
  - 单独创建，`create unique index idx_表名_列名 on 表名(列名)`
  - 单独创建，`alter table 表名 add unique 索引名(列名)`

- 主键索引：设定为主键后数据库会自动建立索引，`Innodb`为聚簇索引，值必须唯一且不能为`null`
  - 建表时，加上`primary key(列名)`指定

- 复合索引：即一个索引包含多个列
  - 建表时，加上`key(列名列表)`指定
  - 单独创建，`create index 索引名 on 表名(列名列表)`
  - 单独创建，`alter table 表名 add index 索引名(列名列表)`

###  唯一索引和普通索引

唯一索引不一定比普通索引快，还可能慢。

- 查询时， 在未使用`limit 1`的情况下，在匹配到一条数据后，唯一索引即返回，普通索引会继续匹配下一条数据，发现不匹配后返回。
  - 唯一索引少了一次匹配，但实际上这个消耗微乎其微。

- 更新时，比较复杂
  - 普通索引将记录放到`change buffer`中语句就执行完了。
  - 唯一索引，必须要校验唯一性，必须将数据页读入内存确定没有冲突，然后才能继续操作。

写多读少的情况，普通索引利用`change buffer`有效减少了对磁盘的访问次数，普通索引性能要高于唯一索引。

### `B-Tree`和`B+Tree`

#### 区别

- 存放结构
  - `B-Tree`的关键字和记录是放在一起的，叶子节点可以看作外部节点，不包含任何信息
  - `B+Tree`的非叶子节点中只有关键字和指向下一个节点的索引，记录只放在叶子节点中

- 查找时间
  - `B-Tree`中，越靠近根节点的记录查找时间越快，只要找到关键字即可确定记录的存在
  - `B+Tree`中每个记录的查找时间基本是一样的，都需要从根节点走到叶子节点，而且在叶子节点中还要再比较关键字

从查找时间看`B-Tree`的要比`B+Tree`好，在实际应用中是`B+Tree`的要好些。

因为`B+Tree`的非叶子节点不存放实际的数据，这样每个节点可容纳的元素个数比`B-Tree`多，树高比`B-Tree`小，这样带来的好处是减少磁盘访问次数。

`B+Tree`找到一个记录所需的比较次数要比`B-Tree`多，但是一次磁盘访问的时间相当于成百上千次内存比较的时间，
因此实际中`B+Tree`的性能还会好些，
而且`B+Tree`的叶子节点使用指针连接在一起，方便顺序遍历（例如查看一个目录下的所有文件，一个表中的所有记录等），
这也是很多数据库和文件系统使用`B+Tree`的缘故。

#### 总结

- `B+Tree`的磁盘读写代价更低
  - `B+Tree`的内部结点并没有指向关键字具体信息的指针。因此其内部结点相对`B-Tree`更小。
  - 如果把所有同一内部结点的关键字存放在同一盘块中，那么盘块所能容纳的关键字数量也越多。
  - 一次性读入内存中的需要查找的关键字也就越多。相对来说`IO`读写次数也就降低了。

- `B+Tree`的查询效率更加稳定
  - 由于非终结点并不是最终指向文件内容的结点，而只是叶子结点中关键字的索引。
  - 所以任何关键字的查找必须走一条从根结点到叶子结点的路。
  - 所有关键字查询的路径长度相同，导致每一个数据的查询效率相当。

### `Hash`索引和`B+`树

#### `Hash`索引

`Hash`索引底层就是`Hash`表，进行查找时，调用一次`Hash`函数就可以获取到相应的键值，之后进行回表查询获得实际数据。

#### `B+`树

`B+`树底层实现是多路平衡查找树，对于每一次的查询都是从根节点出发，查找到叶子节点方可以获得所查键值，然后根据查询判断是否需要回表查询数据。

#### 不同

`Hash`索引进行等值查询更快(一般情况下)，但是却无法进行范围查询。

因为在`Hash`索引中经过`Hash`函数建立索引之后，索引的顺序与原顺序无法保持一致，不能支持范围查询。

而`B+`树的的所有节点皆遵循(左节点小于父节点，右节点大于父节点，多叉树也类似)，天然支持范围。

`Hash`索引不支持使用索引进行排序，原理同上。

`Hash`索引不支持模糊查询以及多列索引的最左前缀匹配。原理也是因为`Hash`函数的不可预测。`AAAA`和`AAAAB`的索引没有相关性。

`Hash`索引任何时候都避免不了回表查询数据，而`B+`树在符合某些条件(聚簇索引，覆盖索引等)的时候可以只通过索引完成查询。

`Hash`索引虽然在等值查询上较快，但是不稳定，性能不可预测，当某个键值存在大量重复的时候，发生`Hash`碰撞，此时效率可能极差。

而`B+`树的查询效率比较稳定，对于所有的查询都是从根节点到叶子节点，且树的高度较低。

因此，在大多数情况下，直接选择`B+`树索引可以获得稳定且较好的查询速度。而不需要使用`Hash`索引。

### 聚簇索引

在`B+`树的索引中，叶子节点可能存储了当前的`key`值，也可能存储了当前的`key`值以及整行的数据，这就是非聚簇索引和聚簇索引。

`InnoDB`中，只有主键索引是聚簇索引，如果没有主键，则挑选一个唯一键建立聚簇索引。

如果没有唯一键，则隐式的生成一个键来建立聚簇索引。

当查询使用聚簇索引时，在对应的叶子节点，可以获取到整行数据，因此不用再次进行回表查询。

### 非聚簇索引回表查询

这涉及到查询语句所要求的字段是否全部命中了索引，如果全部命中了索引，那么就不必再进行回表查询。

举个例子：假设在员工表的年龄上建立了索引

当进行`select age from employee where age < 20`时，在索引的叶子节点上，已经包含了`age`信息，不会再次进行回表查询。

### 多个索引

`MySQL`中，对于一个`SQL`查询，一个表实际上只能使用一个索引来作为主要的查询路径。
这意味着虽然一个查询可以涉及多个表，并且每个表都可以有自己的索引，但是针对单个表而言，`MySQL`查询优化器会选择一个最有效的索引来执行查询。

`MySQL`支持一些特性，可以让一个查询利用多个索引的信息

- 索引合并 (Index Merge)：`MySQL`可以在某些情况下合并多个索引的信息来完成查询。这种情况下，`MySQL`会使用多个索引来分别获取结果集，然后再将这些结果集合并起来。这种方式通常发生在使用`UNION`或者`OR`的查询中。
- 多列索引 (Multi-Column Indexes)：通过创建包含多个列的索引，可以同时利用这些列上的索引信息。这种索引被称为复合索引。
- 索引下推 (Index Condition Pushdown, ICP)：`MySQL 5.6`引入了一项特性，可以在索引扫描过程中直接过滤掉不符合条件的行，从而减少访问表中数据的次数。这可以视为间接利用了多个索引的效果，因为它减少了全表扫描的次数。
- 覆盖索引 (Covering Index)：如果一个索引包含了查询所需的所有列，`MySQL`可以直接从索引中获取数据，而不需要访问表中的实际数据行。这可以显著提高查询速度。
- 分区索引 (Partitioned Indexes)：在使用分区表时，可以创建分区索引，这有助于优化查询性能。

**单个表的最大索引数量**

关于单个表可以创建的最大索引数量，`MySQL 5.0`以后的版本在`64`位系统上支持每个表最多`16`个索引。
每个索引的最大长度为`256`字节。请注意，这个限制可能会随着`MySQL`版本的不同而有所变化。

例：假设有一个表`employees`，并且有多个索引，例如`idx_name`、`idx_department`和`idx_salary`。
```sql
SELECT * FROM employees WHERE name = 'HaoHaoDaYouXi' AND department = 'IT';
```
`MySQL`可能会选择使用一个复合索引`idx_name_department`（如果存在）来进行查询，或者如果不存在复合索引，则可能使用索引合并策略来结合`idx_name`和`idx_department`的结果。

**总结**

虽然一个表在单个查询中只能使用一个索引来作为主要的查询路径，但是通过上述技术，MySQL 可以有效地利用多个索引来优化查询性能。
在设计索引时，考虑查询模式以及如何创建复合索引是非常重要的。

## <a id="explain">`Explain`性能分析</a>

使用`EXPLAIN`关键字可以模拟优化器执行`SQL`查询语句，可以知道`MySQL`是如何处理`SQL`语句的。分析查询语句或是表结构的性能瓶颈。

### 字段解释

- `id`：`select`查询的序列号，包含一组数字，表示查询中执行`select`子句或操作表的顺序。
  - `id`相同，执行顺序由上至下
  - `id`不同，如果是子查询，`id`的序号会递增，`id`值越大优先级越高，越先被执行
  - `id`有相同也有不同：`id`如果相同，可以认为是一组，从上往下顺序执行；在所有组中，`id`值越大，优先级越高，越先执行

`id`号每个号码，表示一趟独立的查询。一个`sql`的查询趟数越少越好。

- `select_type`：代表查询的类型，主要是用于区别普通查询、联合查询、子查询等的复杂查询
  - `simple`：表示不需要`union`操作或者不包含子查询的简单查询。
  - `primary`：表示最外层查询。
  - `union`：`union`操作中第二个及之后的查询。
  - `dependent union`：`union`操作中第二个及之后的查询，并且该查询依赖于外部查询。
  - `subquery`：子查询中的第一个查询。
  - `dependent subquery`：子查询中的第一个查询，并且该查询依赖于外部查询。
  - `derived`：派生表查询，既from字句中的子查询。
  - `materialized`：物化查询。
  - `uncacheable subquery`：无法被缓存的子查询，对外部查询的每一行都需要重新进行查询。
  - `uncacheable union`：`union`操作中第二个及之后的查询，并且该查询属于`uncacheable subquery`。

- `table`：这个数据是基于哪张表的。

- `type`：是查询的访问类型。是较为重要的一个指标，
  - 结果值从最好到最坏依次是：`system`>`const`>`eq_ref`>`ref`>`fulltext`>`ref_or_null`>`index_merge`>`unique_subquery`>`index_subquery`>`range`>`index`>`ALL`，
  - 一般来说，得保证查询至少达到`range`级别，最好能达到`ref`。
  - 常见的：`system`>`const`>`eq_ref`>`ref`>`range`>`index`>`ALL`，其他的不常见。
    - `system`：表只有一行记录（等于系统表），这是`const`类型的特列，平时不会出现，这个也可以忽略不计。
    - `const`：表示通过索引一次就找到了，`const`用于比较`primary key`或者`unique`索引。因为只匹配一行数据，所以很快。如将主键置于`where`列表中，`MySQL`就能将该查询转换为一个常量。
    - `eq_ref`：唯一性索引扫描，对于每个索引键，表中只有一条记录与之匹配。常见于主键或唯一索引扫描。
    - `ref`：非唯一性索引扫描，返回匹配某个单独值的所有行。本质上也是一种索引访问，它返回所有匹配某个单独值的行，然而，它可能会找到多个符合条件的行，所以他应该属于查找和扫描的混合体。
    - `range`：只检索给定范围的行，使用一个索引来选择行。`key`列显示使用了哪个索引一般就是在`where`语句中出现了`between`、`<`、`>`、`in`等的查询这种范围扫描索引扫描比全表扫描要好，
      因为它只需要开始于索引的某一点，而结束语另一点，不用扫描全部索引。
    - `index`：出现`index`是`sql`使用了索引但是没用索引进行过滤，一般是使用了覆盖索引或者是利用索引进行了排序分组。
    - `all`：将遍历全表以找到匹配的行。
    - 其他`type`：
      - `index_merge`：在查询过程中需要多个索引组合使用，通常出现在有`or`关键字的`sql`中。
      - `ref_or_null`：对于某个字段既需要过滤条件，也需要`null`值的情况下。查询优化器会选择用`ref_or_null`连接查询。
      - `index_subquery`：利用索引来关联子查询，不再全表扫描。
      - `unique_subquery`：该联接类型类似于`index_subquery`。子查询中的唯一索引。

- `possible_keys`：显示可能应用在这张表中的索引，一个或多个。查询涉及到的字段上若存在索引，则该索引将被列出，但不一定被查询实际使用。

- `key`：实际使用的索引。如果为`NULL`，则没有使用索引。

- `key_len`：表示索引中使用的字节数，可通过该列计算查询中使用的索引的长度。`key_len`显示的值为索引字段的最大可能长度，并非实际使用长度。
  计算`key_len`，先看索引上字段的类型 + 长度，比如：`int=4;` `varchar(20)=20;` `char(20)=20`
  如果是`varchar`或者`char`这种字符串字段，视字符集要乘不同的值，比如`utf-8`要乘`3`，`GBK`要乘`2`，`varchar`这种动态字符串要加`2`个字节，允许为空的字段要加`1`个字节

- `ref`：显示索引的哪一列被使用了，如果可能的话，是一个常数。哪些列或常量被用于查找索引列上的值。

- `rows`：显示`MySQL`认为它执行查询时必须检查的行数。越少越好！

- `Extra`：其他的额外重要的信息。
  - `Using filesort`：说明`MySQL`会对数据使用一个外部的索引排序，而不是按照表内的索引顺序进行读取。
  - `MySQL`中无法利用索引完成的排序操作称为`文件排序`。排序字段若通过索引去访问将大大提高排序速度。
  - `Using temporary`：使用临时表保存中间结果，`MySQL`在对查询结果排序时使用临时表。常见于排序`order by`和分组查询`group by`。
  - `Using index`：表示相应的`select`操作中使用了覆盖索引(`Covering Index`)，避免访问了表的数据行，效率不错！
    如果同时出现`using where`，表明索引被用来执行索引键值的查找；
    如果没有同时出现`using where`，表明索引只是用来读取数据而非利用索引执行查找。
  - `Using where`：表明使用了`where`过滤。
  - `Using join buffer`：使用了连接缓存。
  - `impossible where`：`where`子句的值总是`false`，不能用来获取任何数据。
  - `select tables optimized away`：在没有`group by`子句的情况下，基于索引优化`MIN`、`MAX`操作或者对于`MyISAM`存储引擎优化`COUNT(*)`操作，
    不必等到执行阶段再进行计算，查询执行计划生成的阶段即完成优化。
  - `distinct`：优化`distinct`操作，在找到第一匹配的元祖后即停止找同样值的动作。

## <a id="sqlyh">`SQL`优化</a>

简单概括就是：能少查就少查，能少排序就少排序，能少计算就少计算。

例如：
使用`in`时，尽量控制好数量，查询一条数据时加上`limit 1`，能使用索引就使用索引，少使用`or`，模糊查询少使用`%`开头，等等。

减少数据库的操作，就可以提高查询效率。

## <a id="ccgc">存储过程</a>

存储过程是数据库程序，可以理解为数据库函数，但区别在于存储过程可以执行多条`sql`语句，而函数只能执行一条。

通过系统表`information_schema.ROUTINES`查看存储过程的详细信，
`information_schema.ROUTINES`是数据库中一个系统表，存储了所有存储过程、函数、触发器的详细信息，包括名称、返回值类型、参数、创建时间、修改时间等。
```sql
select * from information_schema.routines where routine_name = 'test';
```
`information_schema.ROUTINES`表中的列：

- `SPECIFIC_NAME`：存储过程的具体名称，包括该存储过程的名字，参数列表。
- `ROUTINE_SCHEMA`：存储过程所在的数据库名称。
- `ROUTINE_NAME`：存储过程的名称。
- `ROUTINE_TYPE`：`PROCEDURE`表示是一个存储过程，`FUNCTION`表示是一个函数。
- `ROUTINE_DEFINITION`：存储过程的定义语句。
- `CREATED`：存储过程的创建时间。
- `LAST_ALTERED`：存储过程的最后修改时间。
- `DATA_TYPE`：存储过程的返回值类型、参数类型等。

### 使用

创建存储过程：
```sql
create procedure test()
begin
	select id,name from user;
end;
```

调用
```sql
call test();
```

查看创建存储过程的语句：
```sql
show create procedure test;
```
- Procedure：存储过程名称
- Create Procedure：创建存储过程语句
- Definer：存储过程创建者
- sql_mode：SQL模式
- character_set_client：客户端字符集
- collation_connection：连接字符集
- Database Collation：数据库字符集

删除
```sql
drop procedure if exists test;
```

存储过程还可以声明参数、传参、循环、条件判断等等。

假设我们有一个名为`employees`的表，包含员工信息，我们想要找出所有部门中薪资低于平均薪资的员工，并更新他们的薪资为平均薪资的1.1倍。
```sql
-- 定义分隔符：`DELIMITER`。这改变了MySQL命令的默认分隔符，
-- 使得存储过程中可以包含多个SQL语句。
-- $$通常用于定义存储过程或函数的开始和结束
DELIMITER $$

-- 选择数据库my：USE `my`。
-- 这指示MySQL使用名为my的数据库。
USE `my`$$

-- 删除如果已存在的存储过程：DROP PROCEDURE IF EXISTS `UpdateSalariesBelowAverage`。
-- 如果存储过程`StatisticsForDay1`已经存在，则删除它。
DROP PROCEDURE IF EXISTS `UpdateSalariesBelowAverage`$$

CREATE PROCEDURE UpdateSalariesBelowAverage(IN department_id INT)
BEGIN
    -- 声明局部变量
    DECLARE done INT DEFAULT FALSE;
    DECLARE emp_id, emp_salary, avg_salary INT;
    DECLARE cur CURSOR FOR SELECT id, salary FROM employees WHERE department_id = department_id;
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = TRUE;

    -- 计算指定部门的平均薪资
    SELECT AVG(salary) INTO avg_salary FROM employees WHERE department_id = department_id;

    -- 打开游标
    OPEN cur;

    read_loop: LOOP
        -- 从游标中获取数据
        FETCH cur INTO emp_id, emp_salary;

        -- 检查是否到达游标末尾
        IF done THEN
            LEAVE read_loop;
        END IF;

        -- 判断员工薪资是否低于平均薪资
        IF emp_salary < avg_salary THEN
            -- 更新员工薪资为平均薪资的1.1倍
            UPDATE employees SET salary = avg_salary * 1.1 WHERE id = emp_id AND department_id = department_id;
        END IF;
    END LOOP;

    -- 关闭游标
    CLOSE cur;
END$$

DELIMITER ;
```
这个存储过程中：
- 首先定义了一个输入参数`department_id`，用于指定要处理的部门。
- 声明了几个局部变量，包括`done`标记游标是否结束，`emp_id`和`emp_salary`用于存储从游标中读取的员工ID和薪资，`avg_salary`用于存储部门的平均薪资。
- 使用`CURSOR`创建了一个游标，用于遍历指定部门的所有员工。
- 计算了指定部门的平均薪资。
- 使用`LOOP`循环遍历游标中的每一项，如果员工的薪资低于平均薪资，则更新该员工的薪资为平均薪资的1.1倍。
- 最后，关闭游标，完成存储过程。

这个存储过程可以被调用，传入具体的部门`ID`，然后自动找出并更新薪资低于平均值的员工信息。

### 触发器

触发器是数据库对象，它是与表相关联的特殊程序。
它可以在特定的数据操作：例如插入（`INSERT`）、更新（`UPDATE`）或删除（`DELETE`）时触发时自动执行。
`MySQL`触发器使数据库开发人员能够在数据的不同状态之间维护一致性和完整性，并且可以为特定的数据库表自动执行操作。

触发器的作用主要有以下几个方面：

- 强制实施业务规则：触发器可以帮助确保数据表中的业务规则得到强制执行，例如检查插入或更新的数据是否符合某些规则。
- 数据审计：触发器可以声明在执行数据修改时自动记日志或审计数据变化的操作，使数据对数据库管理员和`SQL`审计人员更易于追踪和审计。
- 执行特定业务操作：触发器可以自动执行特定的业务操作，例如计算数据行的总数、计算平均值或总和等。

触发器有两种类型：`BEFORE`和`AFTER`

- `BEFORE`触发器在执行`INSERT`、`UPDATE`、`DELETE`语句之前执行
- `AFTER`触发器在执行`INSERT`、`UPDATE`、`DELETE`语句之后执行。

#### 触发器的语法

```sql
CREATE TRIGGER trigger_name
BEFORE/AFTER INSERT/UPDATE/DELETE ON table_name FOR EACH ROW
BEGIN
-- 触发器执行的 SQL 语句
END;
```
- `trigger_name`：触发器的名称
- `BEFORE/AFTER`：触发器的类型，可以是 BEFORE 或者 AFTER
- `INSERT/UPDATE/DELETE`：触发器所监控的 DML 调用类型
- `table_name`：触发器所绑定的表名
- `FOR EACH ROW`：表示触发器在每行受到 DML 的影响之后都会执行
- 触发器执行的`SQL`语句：该语句会在触发器被触发时执行

以下是一个简单的示例，该触发器会在向`employees`表中插入新记录时自动更新`departments`表中的员工计数：
```sql
DELIMITER $$

CREATE TRIGGER update_department_employee_count
AFTER INSERT ON employees
FOR EACH ROW
BEGIN
   UPDATE departments SET employee_count = employee_count + 1 WHERE department_id = NEW.department_id;
END $$

DELIMITER ;
```
- `DELIMITER $$ 和 DELIMITER ;`
  - 这些命令用于更改`SQL`语句的结束标记。默认情况下，`SQL`语句以分号(`;`)结束。因为触发器定义中可能包含分号，所以这里先将结束标记更改为`$$`，然后在定义结束后再改回分号。
- `CREATE TRIGGER`
  - 用于创建新的触发器。
- `update_department_employee_count`
  - 是触发器的名称。
- `AFTER INSERT ON employees`
  - 表明触发器将在`employees`表上执行`INSERT`操作后触发。
- `FOR EACH ROW`
  - 表示每次插入一行时触发器都会执行一次。
- `BEGIN ... END`
  - 定义触发器执行的SQL语句块。
- `UPDATE departments SET employee_count = employee_count + 1 WHERE department_id = NEW.department_id;`
  - 这条语句将更新`departments`表中与新插入记录的`department_id`相匹配的行，增加其`employee_count`字段的值。

#### 触发器的NEW和OLD关键字

`NEW`和`OLD`关键字在触发器中用于表示触发器所监控的行的新旧状态。
- `NEW`：在触发`INSERT`或`UPDATE`操作期间，`NEW`用于引用将要插入或更新到表中的新行的值。
- `OLD`：在触发`UPDATE`或`DELETE`操作期间，`OLD`用于引用更新或删除之前在表中的旧行的值。

`NEW`和`OLD`使用方法是相似的。在触发器中，可以像引用表的其他列一样引用`NEW`和`OLD`。

可以使用`OLD.column_name`从旧行中引用列值，也可以使用`NEW.column_name`从新行中引用列值。

## <a id="fkfb">分库分表</a>

分库分表是为了解决大量数据的存储和查询问题。当数据量逐渐增大，单个数据库可能无法满足存储和查询的需求，就需要对数据库进行分库分表。

- 分库指的是将数据按照一定的规则拆分到多个数据库中，每个数据库中存放一部分数据，例如：按照地区分库。通过分库可以提升存储的能力，每个数据库可以存储更多的数据量。

- 分表指的是将一张大表按照一定的规则拆分成多个小表，每个小表只存放一部分数据，例如：按照时间分表。通过分表可以提升查询的性能，每个小表查询的数据量较少，查询速度更快。

### 分库分表的优缺性

- 分库分表的好处
  - 提高性能：减少单个数据库的负载，避免热点数据引起的性能瓶颈。分散数据读写压力，提升I/O效率和响应速度。
  - 扩展能力：通过将数据分散到多个数据库或表中，可以更容易地进行水平扩展，即增加更多的数据库服务器来承载更大的数据量和更高的并发请求。
  - 故障隔离：单一数据库的故障不会影响到整个系统的运行，因为数据分布在不同的数据库上。
  - 资源利用：可以充分利用多台服务器的硬件资源，包括CPU、内存和磁盘空间。
  - 数据安全：数据分散存储可以减少单点故障的风险，提高数据安全性。
  - 灵活性：允许对不同数据库或表进行不同的配置和优化，比如使用不同的存储引擎或索引策略。

- 分库分表的坏处
  - 复杂性增加：应用层需要处理跨库跨表的事务管理，增加了开发和维护的难度。需要额外的中间件或逻辑来管理和路由数据到正确的库或表。
  - 数据一致性：跨库操作可能导致数据一致性问题，尤其是在分布式事务中。
  - 查询性能：复杂的查询可能需要从多个库或表中获取数据，这可能会降低查询性能，尤其是涉及到JOIN操作时。
  - 扩容不便：扩容时需要重新设计分配问题，这可能是一个耗时且复杂的过程。
  - 分布键选择：选择合适的分布键非常重要，错误的选择会导致数据分布不均，进而影响性能。
  - 运维成本：分库分表增加了运维的复杂度，包括监控、备份和恢复等。
  - 业务影响：分库分表可能会影响业务逻辑，特别是对于那些依赖于全局视图的应用程序。

### 拆分策略

分库分表的拆分策略主要分为：垂直拆分和水平拆分。
- 垂直拆分
  - 分库：将数据按照业务模块进行拆分，例如：用户中心库、商城库。
  - 分表：将数据按照字段进行拆分，例如：基本信息表、详细信息表。
- 水平拆分
  - 分库：按照分库规则进行数据库拆分，例如：规则是按照年份进行分，xxx_2023、xxx_2024。
  - 分表：按照分库规则进行表里的数据拆分，例如：规则是按照年份进行分，xxx_2023、xxx_2024。

垂直拆分主要根据业务场景和使用进行拆分，水平拆分主要看拆分规则，可以根据时间、ID取模、地区等等规则进行拆分。

### 分库分表工具

- `sharding-sphere`：`jar`包形式，前身是`sharding-jdbc`；
- `TDDL`：`jar`包形式，淘宝根据自身业务需求研发的；
- `Mycat`：中间件。

目前使用多的是`Mycat`

### 分库分表常见问题

- 分布式事务一致性问题：可以使用分布式事务中间件`Seata`
- 跨节点关联查询、分页、排序函数：`Mycat`支持
- 主键避重：`UUID`、`雪花算法`等

## <a id="dxfl">读写分离</a>

读写分离是为了解决读写分离的问题。当数据库的数据量很大时，需要将读和写分离到不同的服务器上，以提升性能和减少单点故障。

### 原理

读写分离的基本思想是将数据库的读取操作和写入操作分开，由不同的数据库实例处理。
这样可以有效避免写操作（如INSERT、UPDATE、DELETE）带来的锁竞争影响读操作的性能，同时也可以通过并行处理读操作来提高数据库的吞吐量。

### 实现方案

- 手动路由
  - 应用程序直接控制读写请求的路由，例如，所有写操作发送到主数据库，而读操作则发送到一个或多个从数据库。
- 数据库中间件
  - 使用代理或中间件（如ProxySQL、Amoeba、Mycat等）来自动处理读写分离。这些工具可以智能地将读写请求转发到适当的数据库实例。
- 基于连接池的解决方案
  - 利用连接池技术，如HikariCP、C3P0等，结合应用框架或ORM（对象关系映射）工具，动态选择读写数据库。
- 应用层逻辑
  - 在应用程序中实现读写分离的逻辑，如使用不同的数据源配置，根据请求类型选择正确的数据源。

### `MySQL`主从复制

读写分离的一个关键组成部分是`MySQL`的主从复制功能。
主数据库（`Master`）负责写操作，从数据库（`Slave`）通过复制主数据库的二进制日志（`Binlog`）来保持数据同步。
从数据库可以有多个，用于分担读取负载。

- 优势
  - 提高读取性能：通过并行处理读取请求，可以显著提高读取性能。
  - 提高写入性能：减少写操作对读操作的影响，提高写入性能。
  - 增强可用性：即使主数据库出现故障，读取操作仍然可以从从数据库中继续进行。
- 劣势
  - 数据延迟：从数据库的数据可能不是实时的，存在一定的复制延迟。
  - 复杂性增加：维护多个数据库实例和复制链路会增加系统的复杂性。
  - 故障转移：需要实现故障检测和自动切换机制，确保在主数据库故障时能够快速恢复服务。

#### 配置主从复制

在主服务器上，配置my.cnf或my.ini文件，启用二进制日志
```
[mysqld]
log-bin=mysql-bin
server-id=1
```

创建复制用户并授权从服务器连接到主服务器
```sql
CREATE USER 'replica'@'%' IDENTIFIED BY 'replica_password';
GRANT REPLICATION SLAVE ON *.* TO 'replica'@'%';
FLUSH PRIVILEGES;
```

查看主服务器状态，记录二进制日志名和位置点
```sql
SHOW MASTER STATUS;
```

从服务器上，配置`server-id`（不同于主服务器）
```
[mysqld]
server-id=2
```

配置从服务器以连接到主服务器并开始复制
```sql
CHANGE MASTER TO
MASTER_HOST='主服务器IP',
MASTER_USER='replica',
MASTER_PASSWORD='replica_password',
MASTER_LOG_FILE='记录的日志名',
MASTER_LOG_POS=记录的位置点;
```

从服务器上启动复制
```sql
START SLAVE;
```

检查从服务器状态，确认复制正常运行
```sql
SHOW SLAVE STATUS\G
```

可以使用`SHOW SLAVE HOSTS`命令来查看主从关系。该命令会返回当前从库已经注册的主库信息，包括主库的`IP`地址、端口号、复制用户名等。

### Mysql主从服务器时间同步问题

可以使用`NTP`（网络时间协议）来同步服务器的时间，也可以使用`MySQL`自带的`master_timestamp_offset`参数来设置主服务器和从服务器之间的时间差值偏移量。

## <a id="bfhf">备份和恢复</a>

备份和恢复是数据管理中的关键环节，旨在保护数据免受意外丢失、损坏或灾难性事件的影响。
无论是对于个人用户还是企业级应用，制定有效的备份和恢复策略都是至关重要的。

- 备份
- 全备份 (Full Backup)
  - 复制所有选定的数据和文件。这是最全面的备份类型，但也是最耗时和占用空间的。
- 增量备份 (Incremental Backup)
  - 只备份自上次备份以来更改的数据。这种方式节省空间，但恢复时可能需要多次备份文件。
- 差异备份 (Differential Backup)
  - 备份自上次全备份以来所有更改的数据。相比增量备份，差异备份在恢复时更快，但占用更多存储空间。
- 事务日志备份 (Transaction Log Backup)
  - 特别适用于数据库系统，备份自上次备份以来的所有事务日志记录，用于恢复到某个时间点。

- 恢复
- 恢复点目标 (RPO)
  - RPO定义了在数据丢失后可接受的最大数据丢失量。例如，如果RPO为1小时，则系统应该能够在1小时内恢复到最近的状态。
- 恢复时间目标 (RTO)
  - RTO定义了系统从故障状态恢复到正常运行状态所需的时间。这包括从检测到故障到系统完全恢复所有功能的时间。
- 灾难恢复计划 (DRP)
  - DRP是一系列预先定义的步骤，用于在重大灾难或中断后恢复关键业务功能和数据。

### mysqldump

`mysqldump`是一个用于备份MySQL数据库的工具。
它允许你备份整个数据库或一个或多个表，并生成一个SQL文件，其中包含创建数据库和表的语句，以及插入数据的语句。

备份数据库：
```shell
# 备份单个数据库
mysqldump -u 用户名 -p 数据库名 > 备份文件名.sql

# 备份多个数据库
mysqldump -u 用户名 -p --databases 数据库名1 数据库名2 > 备份文件名.sql

# 备份所有数据库
mysqldump -u 用户名 -p --all-databases > 备份文件名.sql
```

恢复数据库：
```shell
# 使用mysql命令恢复数据库
mysql -u 用户名 -p 数据库名 < 备份文件名.sql
```


----
