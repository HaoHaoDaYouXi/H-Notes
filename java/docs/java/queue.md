# <div id="queue">`Queue`</div>
`Queue`是单端队列，只能从一端插入元素，另一端删除元素，遵循先进先出(`FIFO`)规则。

`Queue`扩展了`Collection`的接口，**因为容量问题而导致操作失败后处理方式的不同**可以分为两类方法：一种在操作失败后会抛出异常，另一种则会返回特殊值。

| `Queue`接口 | 抛出异常        | 返回特殊值        |
|-----------|-------------|--------------|
| 插入队尾      | `add(E e)`  | `offer(E e)` |
| 删除队首      | `remove()`  | `poll()`     |
| 查询队首元素    | `element()` | `peek()`     |

# <div id="deque">`Deque`</div>
`Deque`是双端队列，在队列的两端均可以插入或删除元素。

`Deque`扩展了`Queue`的接口, 增加了在队首和队尾进行插入和删除的方法，同样根据失败后处理方式的不同分为两类：

| `Deque`接口 | 抛出异常            | 返回特殊值             |
|-----------|-----------------|-------------------|
| 插入队首      | `addFirst(E e)` | `offerFirst(E e)` |
| 插入队尾      | `addLast(E e)`  | `offerLast(E e)`  |
| 删除队首      | `removeFirst()` | `pollFirst()`     |
| 删除队尾      | `removeLast()`  | `pollLast()`      |
| 查询队首元素    | `getFirst()`    | `peekFirst()`     |
| 查询队尾元素    | `getLast()`     | `peekLast()`      |

`Deque`还提供有`push()`和`pop()`等其他方法，可用于模拟栈。

## `ArrayDeque`与`LinkedList`

`ArrayDeque`和`LinkedList`都实现了`Deque`接口，两者都具有队列的功能。

- `ArrayDeque`是基于可变长的数组和双指针来实现，而`LinkedList`则通过链表来实现。
- `ArrayDeque`不支持存储`NULL`数据，但`LinkedList`支持。
- `ArrayDeque`是在`JDK1.6`被引入的，`LinkedList`在`JDK1.2`时就已经存在。
- `ArrayDeque`插入时可能存在扩容过程, 不过均摊后的插入操作依然为`O(1)`。`LinkedList`不需要扩容，但是每次插入数据时均需要申请新的堆空间，均摊性能相比更慢。

**从性能的角度上，选用`ArrayDeque`来实现队列要比`LinkedList`更好。`ArrayDeque`也可以用于实现栈。**

# <div id="priorityqueue">`PriorityQueue`</div>
`PriorityQueue`是在`JDK1.5`中被引入的, 其与`Queue`的区别在于元素出队顺序是与优先级相关的，即总是优先级最高的元素先出队。

- `PriorityQueue`利用了二叉堆的数据结构来实现的，底层使用可变长的数组来存储数据`PriorityQueue`通过堆元素的上浮和下沉，实现了在`O(logn)`的时间复杂度内插入元素和删除堆顶元素。
- `PriorityQueue`是非线程安全的，且不支持存储`NULL`和`non-comparable`的对象。
- `PriorityQueue`默认是小顶堆，但可以接收一个`Comparator`作为构造参数，从而来自定义元素优先级的先后。




----
