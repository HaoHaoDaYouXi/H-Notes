# Redis概述
- Redis是一个开源的使用ANSI C语言编写、支持网络、可基于内存亦可持久化的日志型、Key-Value（键值型）数据库（非关系型数据库），并提供多种语言的API。

- Redis是一个高性能的Key-Value数据库。它的出现很大程度补偿来Memcached这类Key-Value型存储的不足，在部分场合下可以对关系型数据库起到很好的补充作用。它提供来Java、C/C++、PHP、JavaScript、Perl、Object-C、Python、Ruby、Erlang等客户端，使用方便。

- Redis支持主从同步，Redis能够借助于Sentinel（哨兵，Redis自带的）工具来监控主从节点，当主节点发生故障时，会自己提升另外一个从节点成为新的主节点。

## 1. 支持的数据类型

- 和Memcached类似，但它支持存储的Value类型相对更多，包括String（字符串）、List（列表）、Sets（集合）、Sorted Sets（有序集合）和Hash（哈希类型、关联数组）、Bitmaps（位图）和HyperLoglog。

## 2. 性能

- 100万较小的键存储字符串，大概消耗100M内存；

- 由于Redis是单线程，如果服务器主机上有多个CPU，只有一个能够使用，但并不意味着CPU会成为瓶颈，因为Redis是一个比较简单的K-V数据存储，CPU通常不会成为瓶颈的；

- 在常见的linux服务器上，500K（50万）的并发，只需要一秒钟处理，如果主机硬件较好的情况下，每秒钟可以达到上百万的并发.

## 3. Redis与Memcache对比
- Memcache只能使用内存来缓存对象。而Redis除了可以使用内存来缓存对像，还可以周期性的将数据保存到磁盘上，对数据进行永久存储。
当服务器突然断电或死机后，redis基于磁盘中的数据进行恢复；
- Redis是单线程服务器，只有一个线程来响应所有的请求。Memcache是多线程的；
- Redis支持更多的数据类型。