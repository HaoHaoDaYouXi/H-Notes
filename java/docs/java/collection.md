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

## <div id="jh_hashmap">`HashMap`(数组+链表+红黑树)</div>

`HashMap`根据键的`hashCode`值存储数据，大多数情况下可以直接定位到它的值，因而具有很快的访问速度，但遍历顺序却是不确定的。

`HashMap`最多只允许一条记录的键为`null`，允许多条记录的值为`null`。

`HashMap`非线程安全，即任一时刻可以有多个线程同时写`HashMap`，可能会导致数据的不一致。

如果需要满足线程安全，可以用`Collections`的`synchronizedMap`方法使`HashMap`具有线程安全的能力，或者使用`ConcurrentHashMap`。

### `HashMap`1.8之前

![hashMap_java7.png](img/hashMap_java7.png)

大方向上，`HashMap`里面是一个数组，然后数组中每个元素是一个单向链表。
上图中，每个绿色的实体是嵌套类`Entry`的实例，`Entry`包含四个属性：`key`，`value`，`hash`值和用于单向链表的`next`。

1. `capacity`：当前数组容量，始终保持`2^n`，可以扩容，扩容后数组大小为当前的`2`倍。
2. `loadFactor`：负载因子，默认为`0.75`。
3. `threshold`：扩容的阈值，等于`capacity`*`loadFactor`

### `HashMap`1.8之后
`Java8`对`HashMap`进行了一些修改，最大的不同就是利用了红黑树，所以其由`数组+链表+红黑树`组成。

根据`HashMap`1.8之前的介绍，我们知道，查找的时候，根据`hash`值我们能够快速定位到数组的具体下标，但是之后的话，需要顺着链表一个个比较下去才能找到我们需要的，时间复杂度取决于链表的长度，为`O(n)`。

为了降低这部分的开销，在`Java8`中，当链表中的元素超过了`8`个以后，会将链表转换为红黑树，在这些位置进行查找的时候可以降低时间复杂度为`O(logN)`。

![hashMap_java8.png](img/hashMap_java8.png)

## <div id="jh_concurrenthashmap">`ConcurrentHashMap`</div>

----
