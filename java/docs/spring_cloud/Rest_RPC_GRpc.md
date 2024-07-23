# Rest、RPC、GRpc

在微服务架构中，服务与服务之间的通信是关键部分，主要通信方式可以分为同步和异步两大类。

一般业务功能中直接调用的都是同步通信方式，有REST、RPC、GRpc三种，其他的像是消息、通知等等就是异步通信。

## `REST`

`REST`（Representational State Transfer）是一种软件架构风格，用于设计网络应用程序。
它基于`HTTP`协议，使用`URI`（Uniform Resource Identifier）来定位资源，
使用`HTTP`方法（`GET`、`POST`、`PUT`、`DELETE`等）来操作资源。

`RESTful API`的设计风格使得接口简洁、易懂，方便使用。
