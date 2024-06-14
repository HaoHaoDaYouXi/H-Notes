# Redis 高可用
高可用，我们要从为什么要用它以及它的实现方式和单机的区别，这些方方面面来看待

Redis的几种常见使用方式包括：
    
    Redis单副本；
    Redis多副本（主从）；
    Redis Sentinel（哨兵）；
    Redis Cluster；
    Redis自研。



## Redis 主从复制
redis主从复制主要分两个角色，主机(master)主要负责读写操作，从机(slave)主要负责读操作，
主机定期同步数据到从机上，保证数据一致性。

redis同步数据主要分两种，全量同步和增量同步。
主从复制不会阻塞master，在同步数据时，master还可以继续处理请求，
因为redis会生成新的进程来解决同步问题。
主从里面的从也可以是主(树形结构)，提高效率，减少主机压力。
主机可以有多个从机，从机只能有一个主机。
主从配置一般是修改redis.conf文件内的slaveof格式：slaveof ip port，
redis-cli -p 6379 info Replication，可以查看主机有几个从机。

## Redis 同步数据
主要分全量同步和增量同步，从机第一次链接一定是全量同步，
短线重连根据runid判断是否一致来执行全量同步或增量同步，
每个redis服务器都有自己的runid，主机根据runid查询有没有保存，
没有就全量同步，有就增量同步，
主从服务器会分别维护一个offset(复制偏移量)主机每次向服务器传播N个字节的数据时，
就会把自己的offset的值加N，从机每次接受到N个字节数据时，就将自己的offset加N。
	
复制积压缓冲区是主机维护的一个固定长度的先进先出的队列，默认大小1M，
主要是当主机传播命令时，把命令放入，当断开时，
主机会将缓冲区的所有数据发给从机(断开之后的数据)。
	
同步执行过程，从机链接时判断自己是否保存了主机的runid(判断是否第一次)，
没有保存就向主机发出全量同步，有保存就把runid发送给主机，主机判断是否和自己的一致，
不一致就把当前的runid在发给从机并执行全量同步，一致就会判断offset
相差有没有超过缓冲区的大小，没有就等待主机同步数据给从机，
超过主机就生成快照文件，给从机在同步缓冲区的数据。
	
全量同步分三个流程：
    
    同步快照(主机创建并发送快照给从机，从机进入快照并解析，
    主机同时将此阶段生成的新命令写入到缓冲区)，
    同步缓冲区(主机向从机同步缓冲区的写的操作命令)，
    同步增量(主机同步写操作到从机)
	
增量同步主要在从机完成初始化正常工作时，主机发生写操作就同步到从机，
正常主机每执行一个写命令就向从机发请求，从机接受并处理。

## Redis 哨兵(sentinel)机制
	
sentinel主要监控Redis集群中master的状态，当master发生故障时，
可以实现master和slave的切换，保证系统的高可用。
主从的缺点，没法对master进行动态选举，这需要sentinel机制完成。
sentinel会不断检查master和slave状态是否正常，当发现某个节点出问题时，
sentinel可以通过API向管理员或其他应用程序发送通知。
	
当master不能正常操作时，sentinel会开始一次故障转移，会将失效的master下的一个slave升级为新的master，
并让其他slave改为新的master，当客户端试图链接失效的master，集群会向客户端展示新的master地址，切换后对应的配置文件
都会有所变化，master会对一个slaveof的配置，slave对应的master也改成新的，sentinel.conf的监控对象也会改变。

## 故障判断原理
每个sentinel进程每秒钟一次的频率向整个集群中的master、slave以及
其他的sentinel进程发送一个ping的请求，
如果一个实例距离最后一次有效ping请求超过down-after-milliseconds规定的值，
这个实例就会被sentinel标记为主观下线(SDOWN)，
如果一个master被标记为主观下线，
则正在监视这个master的sentinel进程要以每秒一次的频率确定master的确进入主观下线状态，
当超过配置文件中给定的sentinel的数量，在指点的时间范围内确定master进入了主观下线状态，
则master会被标记为客观下线(ODOWN)
一般情况每个sentinel会以每10s一次的频率向集群中所有的master、slave发送info命令，
当master被标记为客观下线，sentinel会向下面所有的slave发送info的频率改为1s一次，
若没有一定数量的sentinel同意master下线，那master的客观下线状态会被移除，
若master对ping的命令有回复，master的主观下线状态也会被移除。
