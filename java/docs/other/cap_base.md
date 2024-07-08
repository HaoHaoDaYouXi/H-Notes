# <a id="cap">`CAP`理论</a>
`CAP`理论的提出者布鲁尔在提出`CAP`猜想的时候，
并没有详细定义`Consistency`、`Availability`、`Partition Tolerance`三个单词的明确定义。

⼀般⼤家的理解是：

在理论计算机科学中，`CAP`定理（`CAP theorem`）指出对于⼀个分布式系统来说，当设计读写操作时，只能同时满足以下三点中的两个：
- 一致性（`Consistency`）：所有节点访问同⼀份最新的数据副本
- 可用性（`Availability`）：⾮故障的节点在合理的时间内返回合理的响应（不是错误或者超时的响应）。
- 分区容错性（`Partition tolerance`）：分布式系统出现⽹络分区的时候，仍然能够对外提供服务。





----
