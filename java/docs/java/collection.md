# 集合

集合类存放于 Java.util 包中，主要有 3 种：set(集)、list(列表包含 Queue)和 map(映射)。

1. Collection：Collection 是集合 List、Set、Queue 的最基本的接口。
2. Iterator：迭代器，可以通过迭代器遍历集合中的数据
3. Map：是映射表的基础接口

![collection.png](img/collection.png)

## <div id="jh_szdqb">`List`，`Set`，`Queue`，`Map`四者的区别？</div>

- `List`(对付顺序的好帮⼿): 存储的元素是有序的、可重复的。
  - `ArrayList`
    - `Object[]`数组
    - 排列有序，可以重复
    - 新增删除慢，`getter setter`快
    - 线程不安全
    - 扩容：容量*1.5+1
  - `Vector`
    - `Object[]`数组
    - 排列有序，可以重复
    - 新增删除慢
    - 线程安全，效率低
    - 扩容：容量*2
  - `LinkedList`
    - 双向链表(`JDK1.6`之前为循环链表，`JDK1.7`取消了循环)
    - 排列有序，可以重复
    - 新增删除快，查询慢
    - 线程不安全
- `Set`(注重独⼀⽆⼆的性质): 存储的元素是⽆序的、不可重复的。
  - `HashSet`(⽆序，唯⼀)
    - 基于`HashMap`实现的，底层采⽤`HashMap`来保存元素，`hash`表实现
    - 排列无序，不可以重复
    - 存取速度快
  - `LinkedHashSet`
    - `LinkedHashSet`是`HashSet`的⼦类，并且其内部是通过`LinkedHashMap`来实现的。
  - `TreeSet`(有序，唯⼀)
    - 内部是`TreeMap`的`SortedSet`实现，底层采⽤红⿊树(⾃平衡的排序⼆叉树)
    - 排列无序，不可以重复，排列存储
- `Queue`(实现排队功能的叫号机)：按特定的排队规则来确定先后顺序，存储的元素是有序的、可重复的。
  - `PriorityQueue`
    - `Object[]`数组来实现⼆叉堆
  - `ArrayQueue`
    - `Object[]`数组 + 双指针
- `Map`(⽤`key`来搜索的专家)
  - `HashMap`
    - `JDK1.8`之前`HashMap`由数组+链表组成的，数组是`HashMap`的主体，链表则
      是主要为了解决哈希冲突⽽存在的(“拉链法”解决冲突)。`JDK1.8`以后在解决哈希冲突时有了
      ᫾⼤的变化，当链表⻓度⼤于阈值(默认为 8)(将链表转换成红⿊树前会判断，如果当前数组
      的⻓度⼩于`64`，那么会选择先进⾏数组扩容，⽽不是转换为红⿊树)时，将链表转化为红⿊
      树，以减少搜索时间
    - `key`不可以重复，`value`可以重复，都可以为`null`
    - 线程不安全
    - `LinkedHashMap`
      - `LinkedHashMap`继承⾃`HashMap`，所以它的底层仍然是基于拉链式散列结构即由数组和链表或红⿊树组成。
        另外，`LinkedHashMap`在上⾯结构的基础上，增加了⼀条双向链表，使得上⾯的结构可以保持键值对的插⼊顺序，同时通过对链表进⾏相应的操作，实现了访问顺序相关逻辑。
  - `Hashtable`
    - 数组+链表组成的，数组是`Hashtable`的主体，链表则是主要为了解决哈希冲突而存在的
    - `key`不可以重复，`value`可以重复，都不可以为`null`
    - 线程安全
  - `TreeMap`
    - 红⿊树(⾃平衡的排序⼆叉树)
    - `key`不可以重复，`value`可以重复

----
