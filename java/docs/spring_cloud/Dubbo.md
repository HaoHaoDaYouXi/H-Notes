# Dubbo

`Dubbo`是阿里巴巴开源的基于`Java`的高性能`RPC`分布式服务框架，现已成为`Apache`基金会孵化项目。

地址：https://cn.dubbo.apache.org/zh-cn/

## <a id="lct">服务注册与发现的流程图</a>

下图来自`Dubbo`官网

![dubbo_architecture.png](img/dubbo_architecture.png)

## <a id="jdjs">Dubbo节点角色</a>

| 节点        | 角色说明                |
|-----------|---------------------|
| Provider  | 暴露服务的服务提供方          |
| Consumer  | 调用远程服务的服务消费方        |
| Registry  | 服务注册与发现的注册中心        |
| Monitor   | 统计服务的调用次数和调用时间的监控中心 |
| Container | 服务运行容器              |
