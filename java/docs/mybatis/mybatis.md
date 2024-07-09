# Mybatis

`Mybatis`是一个半`ORM`（对象关系映射）框架，它内部封装了`JDBC`，开发时只需要关注`sql`语句本身，
不需要花费精力去处理加载驱动、创建连接、创建`statement`等繁杂的过程。
开发人员直接编写原生态`sql`，可以严格控制`sql`执行性能，灵活度高。

`MyBatis`可以使用`xml`或注解来配置和映射原生信息，将`POJO`映射成数据库中的记录，避免了几乎所有的`JDBC`代码和手动设置参数以及获取结果集。

通过`xml`文件或注解的方式将要执行的各种`statement`配置起来，并通过`java`对象和`statement`中`sql`的动态参数进行映射生成最终执行的`sql`语句，
最后由`MyBatis`框架执行`sql`并将结果映射为`java`对象并返回。（从执行`sql`到返回`result`的过程）。

## <a id="yqd">Mybatis的优缺点</a>

**优点：**

- 基于`sql`语句编程，相当灵活，不会对应用程序或者数据库的现有设计造成任何影响，`sql`写在`xml`
  里，解除`sql`与程序代码的耦合，便于统一管理；提供`XML`标签，支持编写动态`sql`语句，并可重用。
- 代码少，与`JDBC`相比，减少了一半以上的代码量，消除了`JDBC`大量冗余的代码，不需要手动开关连接；
- 很好的与各种数据库兼容（因为`MyBatis`使用`JDBC`来连接数据库，所以只要`JDBC`支持的数据库`MyBatis`都支持）。
- 能够与`Spring`很好的集成；
- 提供映射标签，支持对象与数据库的`ORM`字段关系映射；提供对象关系映射标签，支持对象关系组件维护。

**缺点：**

- `sql`语句的编写工作量较大，尤其当字段多、关联表多时，对开发人员编写`sql`语句能力有一定要求。
- `sql`语句依赖于数据库，导致数据库移植性差，不能随意更换数据库。

## <a id="qb">Mybatis和Hibernate的区别</a>

**相同点**

都是对`jdbc`的封装，都是持久层的框架，都用于`dao`层的开发。

**不同点**

- `MyBatis`是一个半自动映射的框架，配置`Java`对象与`sql`语句执行结果的对应关系，多表关联关系配置简单
- `Hibernate`是一个全表映射的框架，配置`Java`对象与数据库表的对应关系，多表关联关系配置复杂

**`sql`优化和移植性**

- `Hibernate`对`sql`语句封装，提供了日志、缓存、级联（级联比`MyBatis`强大）等特性，</br>
  此外还提供 `hql`（`Hibernate Query Language`）操作数据库，数据库无关性支持好，但会多消耗性能。</br>
  如果项目需要支持多种数据库，代码开发量少，但`sql`语句优化困难。

- `MyBatis`需要手动编写`sql`，支持动态`sql`、处理列表、动态生成表名、支持存储过程。</br>
  开发工作量相对大些。直接使用`sql`语句操作数据库，不支持数据库无关性，但`sql`语句优化容易。

## <a id="cybq">Mybatis 常用标签</a>

最常见的无非就是`crud`（增删改查）此类标签：

- `insert`：新增
- `update`：修改
- `delete`：删除
- `select`：查询

除了以上还有很多：

- `resultMap`：结果映射
- `parameterMap`：参数映射
- `resultType`：结果类型
- `parameterType`：参数类型
- `sql`：sql片段
- `include`：引用sql片段
- `selectKey`：主键生成策略，获取自增主键id的值并进行设置
- `association`：一对一关联
- `collection`：一对多关联
- `discriminator`：多表继承
- `set`、`where`、`if`、`foreach`、`trim`、`choose`、`when`、`otherwise`、`bind`：一般写动态`sql`涉及到的标签
- 等等

## <a id="fhqb">Mybatis `$()`和`#()`的区别</a>

- `${}`是字符串替换，是`Properties`⽂件中的变量占位符，它可以⽤于标签属性值和`sql`内部，属于静态⽂本替换
  - `Mybatis`在处理`${}`时，就是把`${}`直接替换成变量的值，这种会出现`sql`注入的风险。
- `#{}`是预编译处理，是`sql`的参数占位符，
  - `Mybatis`在处理`#{}`时，会将`sql`中的`#{}`替换为`?`号，调用`PreparedStatement`的`set`方法来赋值

## <a id="mhcx">Mybatis 模糊查询</a>

`Mybatis`的模糊查询一般存在两种写法，一种是使用`${}`，另一种是使用`#{}`。

使用`${}`的存在`sql`注入的风险，一般不推荐使用

一般的写法：
```sql
select * from user where name like CONCAT('%', #{name}, '%')
```

`like`也可以根据使用情况替换成`likeLeft`、`likeRight`
- `likeLeft`：使用`%`作为通配符，只能用在字符串的开头。如：`name likeLeft '%T'`，表示查询`name`字段以`T`结尾的记录。
- `likeRight`：使用`%`作为通配符，只能用在字符串的末尾。如：`name likeRight 'T%'`，表示查询`name`字段以`T`开头的记录

## <a id="qtcx">Mybatis 嵌套查询</a>

`MyBatis`嵌套查询通常指的是在一个查询中嵌套另一个查询的结果。

这可以通过使用`<select>`标签嵌套来实现，也可以通过在映射文件中使用`<collection>`、`<association>`等复杂类型的属性来实现。

以下是一个使用`<collection>`进行嵌套查询的例子：

假设我们有两个表：学生表`student`、班级表`clazz`

现在，我们查询一个班级和它的所有学生的信息。

在班级类中定义一个集合类型的`allStudents`属性来放所有学生：
```java
public class Clazz {
    private Long id;
    private String name;
    // 其他字段...
    private List<Student> allStudents;
    // getters、setters...
}
```
然后，在映射文件中定义查询并使用`<collection>`进行嵌套查询：
```xml
<mapper namespace="com.xxx.mapper.ClazzMapper">
  <!-- 结果映射 -->
  <resultMap id="ClazzMap" type="Clazz">
    <id property="id" column="id"/>
    <result property="title" column="title"/>
    <!-- 其他字段映射... -->
    <collection property="allStudents"
                ofType="Student"
                select="selectStudentByClazzId"
                column="id"/>
  </resultMap>
 
  <!-- 查询班级 -->
  <select id="selectById" resultType="Clazz">
      SELECT * FROM clazz WHERE id = #{id}
  </select>

  <!-- 查询学生，并嵌套在帖子查询中 -->
  <select id="selectStudentByClazzId" resultType="Student">
      SELECT * FROM student WHERE clazz_id = #{id}
  </select>
 
</mapper>
```
在`<collection>`标签中，`property`属性指定了嵌套查询的属性名，`ofType`属性指定了集合中元素的类型，`select`属性指定了用于查询集合的嵌套查询的`ID`，`column`属性指定了嵌套查询使用的外键列。

最后调用`selectById`方法来查询班级信息，同时会自动执行嵌套查询`selectStudentByClazzId`，并将结果映射到`Clazz`对象的`allStudents`属性中。

## <a id="hc">Mybatis 缓存</a>

`Mybatis`中有一级缓存和二级缓存，默认情况下一级缓存是开启的，而且是不能关闭的。

**一级缓存**是指`SqlSession`级别的缓存，当在同一个`SqlSession`中进行相同的`sql`语句查询时，
第二次以后的查询不会从数据库查询，而是直接从缓存中获取，一级缓存最多缓存`1024`条`sql`。

**二级缓存**是指可以跨`SqlSession`的缓存。是`mapper`级别的缓存，对于`mapper`级别的缓存不同的`SqlSession`是可以共享的。

### 一级缓存原理（SqlSession级别）
第一次发出一个查询`sql`，`sql`查询结果写入`SqlSession`的一级缓存中，缓存使用的数据结构是一个`map`。
- `key`：`MapperID`+`offset`+`limit`+`sql`+所有的入参
- `value`：用户信息

同一个`SqlSession`再次发出相同的`sql`，就从缓存中取出数据。</br>
如果两次中间出现`commit`操作（修改、添加、删除），本`SqlSession`中的一级缓存区域全部清空，</br>
下次再去缓存中查询不到，就要从数据库查询，从数据库查询到再写入缓存。

### 二级缓存原理（mapper基本）
二级缓存的范围是`mapper`级别（`mapper`同一个命名空间），`mapper`以命名空间为单位创建缓存数据结构，结构是`map`。
- `key`：`MapperID`+`offset`+`limit`+`sql`+所有的入参

`mybatis`的二级缓存是通过`CacheExecutor`实现的。

`CacheExecutor`其实是`Executor`的代理对象。

所有的查询操作，在`CacheExecutor`中都会先匹配缓存中是否存在，不存在则查询数据库。

具体使用需要配置：
- `Mybatis`全局配置中启用二级缓存配置
- 在对应的`Mapper.xml`中配置`cache`节点
- 在对应的`select`查询节点中添加`useCache=true`
- 属性类需要实现`Serializable`序列化接口

## <a id="gzyl">Mybatis 工作原理</a>

- 读取`MyBatis`配置文件：`mybatis-config.xml`为`MyBatis`的全局配置文件，配置了`MyBatis`的运行环境等信息，例如数据库连接信息。
- 加载`sql`映射文件。文件中包含了操作数据库的`sql`语句，需要在`MyBatis`配置文件`mybatis-config.xml`中加载。
  - `mybatis-config.xml`文件可以加载多个映射文件，每个文件对应数据库中的一张表。
- 构造会话工厂：通过`MyBatis`的环境等配置信息构建会话工厂`SqlSessionFactory`。
- 创建会话对象：由会话工厂创建`SqlSession`对象，该对象中包含了执行`sql`语句的所有方法。
- `Executor`执行器：`MyBatis`底层定义了一个`Executor`接口来操作数据库，
  根据`SqlSession`传递的参数动态地生成需要执行的`sql`语句，同时负责查询缓存的维护。
- `MappedStatement`对象：在`Executor`接口的执行方法中有一个`MappedStatement`类型的参数，
  是对映射信息的封装，用于存储要映射的`sql`语句的`id`、`参数`等信息。
- 输入参数映射：输入参数类型可以是`Map`、`List`等集合类型，也可以是基本数据类型和`POJO`类型。
  输入参数映射过程类似于`JDBC`对`preparedStatement`对象设置参数的过程。
- 输出结果映射：输出结果类型可以是`Map`、`List`等集合类型，也可以是基本数据类型和`POJO`类型。
  输出结果映射过程类似于`JDBC`对结果集的解析过程。





----
