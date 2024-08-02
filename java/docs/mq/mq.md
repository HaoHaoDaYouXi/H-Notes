# 消息队列 MQ

消息队列（`Message Queue`，简称`MQ`）是一种用于在分布式系统中进行消息传递的软件中间件。
它允许应用程序之间通过消息进行通信，而不需要直接调用对方。
消息队列提供了异步处理、解耦、缓冲、负载均衡等功能，是构建可扩展、高可用和高性能系统的关键组件之一。

**主要特性：**

- 异步通信：消息队列允许发送者与接收者异步操作，即发送者不必等待接收者处理消息即可继续执行后续任务。
- 解耦：发送者和接收者之间没有直接依赖，即使接收者暂时不可用，发送者仍然可以发送消息。
- 缓冲：消息队列可以作为消息的临时存储，当接收者处理能力不足时，可以暂时存储消息直到被处理。
- 可靠性：消息队列通常提供持久化存储，确保消息不会因为系统故障而丢失。
- 负载均衡：消息队列可以将消息均匀地分发给多个接收者，以平衡工作负载。
- 广播：一个消息可以被多个消费者同时消费，这在发布/订阅模式中非常有用。

**常见的消息队列中间件：**

- `ActiveMQ`：一个开源的消息中间件，支持多种消息协议。
- `RabbitMQ`：基于`AMQP`协议的开源消息队列服务。
- `Kafka`：主要用于构建实时数据管道和流应用。
- `RocketMQ`：由阿里巴巴集团开发，适用于大规模分布式系统中的消息传输和处理。

**使用场景：**

- 微服务间通信：在微服务架构中，不同的服务可以通过消息队列进行通信，提高系统的松耦合性和可扩展性。
- 日志收集：收集和聚合不同系统产生的日志，便于统一分析和监控。
- 事件驱动架构：基于事件触发的业务流程，如用户行为触发后台处理逻辑。
- 大数据处理：作为数据流的传输通道，将数据从源系统传输到数据仓库或数据分析系统。

## <a id="qb">ActiveMQ、RabbitMQ、RocketMQ、Kafka区别</a>

### ActiveMQ

特点：

- 支持多种消息协议，如`AMQP`、`STOMP`、`MQTT`等。
- 提供丰富的管理界面和工具。
- 具有较好的文档和支持社区。
- 支持`JMS`标准，适合`Java`应用程序。

适用场景：

- 对于需要多种消息协议支持的场景。
- 需要遵循`JMS`标准的应用程序。

### RabbitMQ

特点：

- 基于`Erlang`开发，具有很高的并发能力和低延迟。
- 支持`AMQP`协议，也支持多种其他协议。
- 提供了丰富的插件和客户端库。

适用场景：

- 需要高性能和低延迟的场景。
- 微服务架构中作为消息中间件。

### RocketMQ

特点：

- 高性能，单机吞吐量可达十万级别。
- 支持高可用性的分布式架构。
- 由阿里巴巴开发，特别适合大规模分布式系统。

适用场景：

- 需要极高吞吐量和高可用性的场景。
- 大型电商、金融等业务场景。

### Kafka

特点：

- 高吞吐量，适用于大数据处理和流式处理。
- 分布式架构，支持水平扩展。
- 数据持久化，提供数据冗余和容错机制。

适用场景：

- 实时数据管道和流应用。
- 日志收集和分析系统。
- 大数据处理平台。

### 总结

- 吞吐量：`Kafka`和`RocketMQ`的吞吐量较高，适合大数据和高并发场景；`ActiveMQ`和`RabbitMQ`的吞吐量相对较低，但更注重功能的全面性和灵活性。
- 延时：`RabbitMQ`和`RocketMQ`的延时更低，适合对实时性要求较高的场景。
- 可用性：所有四个系统都支持高可用性，但`RocketMQ`和`Kafka`的分布式架构可能提供更高的可用性和更好的容错性。
- 功能特性：`ActiveMQ`和`RabbitMQ`提供了更丰富的功能和协议支持，而`Kafka`和`RocketMQ`更专注于高吞吐量和分布式处理。

目前主流还是基于`Kafka`和`RocketMQ`，`Kafka`做日志系统和大数据处理相关的，`RocketMQ`做消息队列业务系统使用

## <a id="mdx">保证消息的幂等性</a>

消息队列消息的幂等性是指在使用消息队列进行消息传递时，对于同一条消息的处理不会因为重复消费或处理而导致系统状态的错误或不一致。

如果一个消息队列消息是幂等的，那么在同一条消息被处理多次时，系统状态不会受到任何负面影响。

消息队列出现消息幂等性问题的主要原因是消息重复发送。这种情况可能发生在以下情况下：

- 生产者重复发送消息：由于网络不稳定或其他异常情况，生产者可能会发送同样的消息多次。
- 消息队列本身的问题：由于消息队列本身的问题，消息可能会被重复发送。
- 消费者的问题：由于消费者的问题，消息可能会被重复消费。

无论是哪种情况，都会导致消息的重复处理，从而破坏了消息处理的幂等性。因此，在设计消息队列时需要考虑如何保证消息的幂等性，以避免这种问题的发生。

下面是一些实现消息幂等性的策略：
- 消息去重
  - `ID`校验：为每条消息分配一个全局唯一的`ID`，在接收端维护一个已处理消息的`ID`列表，如果接收到的消息`ID`已经存在于列表中，则忽略该消息。
  - 数据库校验：在数据库中为每个业务操作设置唯一键，如果尝试插入一条记录时发现唯一键冲突，则说明该消息已经被处理过。
- 状态机
  - 将业务逻辑设计为状态机，每次消息处理都根据当前
  - 状态和消息内容来更新状态。状态机的设计可以确保即使消息被重复处理，系统状态也不会改变。
- 使用幂等命令
  - 设计业务逻辑时，尽量使用幂等的操作，例如，增加或减少某个计数器，或者更新某个字段，只要条件相同，多次执行结果不变。
- 消息确认机制
  - 在消息队列中，只有当消息成功处理后才确认消息，这样可以避免消息因未被正确处理而被丢弃。如果消息处理失败，可以重新投递消息。
- 事务处理
  - 对于关键操作，可以使用事务来保证操作的原子性和一致性，确保即使在消息重复处理的情况下，也能保持数据的正确性。
- 消息队列特性利用
  - 死信队列：如果消息处理失败，可以将消息发送到死信队列，然后人工检查或重试。
  - 延迟队列：对于需要定时重复执行的任务，可以使用延迟队列来控制消息的重复发送时间。

具体的策略还是得根据具体的业务场景和需求来确定。

`RocketMQ`实现消息的幂等性：
```java
public class IdempotentProducer {
    private final DefaultMQProducer producer;
    private ConcurrentHashMap<String, Boolean> messageTrack = new ConcurrentHashMap<>();
 
    public IdempotentProducer(String producerGroup) throws MQClientException {
        producer = new DefaultMQProducer(producerGroup);
        producer.start();
    }
 
    public void sendMessage(Message msg) throws MQClientException, RemotingException, MQBrokerException, InterruptedException {
        String messageKey = msg.getKeys(); // 假设消息的唯一标识存储在msg.getKeys()中
        if (messageTrack.containsKey(messageKey)) {
            // 如果已经发送过该消息，则不再重复发送
            System.out.println("Message already sent, skipping: " + messageKey);
            return;
        }
        SendResult sendResult = producer.send(msg);
        if (sendResult.getSendStatus() == SendStatus.SEND_OK) {
            // 消息发送成功，记录消息已被发送
            messageTrack.put(messageKey, Boolean.TRUE);
        }
    }
 
    public void shutdown() {
        producer.shutdown();
    }
}
 
public class IdempotentConsumer {
    private final DefaultMQPushConsumer consumer;
    private ConcurrentHashMap<String, Boolean> messageTrack = new ConcurrentHashMap<>();
 
    public IdempotentConsumer(String consumerGroup, String namesrvAddr, String topic, String tag) throws MQClientException {
        consumer = new DefaultMQPushConsumer(consumerGroup);
        consumer.setNamesrvAddr(namesrvAddr);
        consumer.subscribe(topic, tag);
        consumer.registerMessageListener((msg, context) -> {
            String messageKey = msg.getKeys();
            if (messageTrack.containsKey(messageKey)) {
                // 如果已经处理过该消息，则不再重复处理
                System.out.println("Message already processed, skipping: " + messageKey);
                return ConsumeConcurrentlyStatus.CONSUME_SUCCESS;
            }
            // 处理消息的业务逻辑
            // ...
 
            messageTrack.put(messageKey, Boolean.TRUE);
            return ConsumeConcurrentlyStatus.CONSUME_SUCCESS;
        });
        consumer.start();
    }
}
```
示例中，`IdempotentProducer`类负责发送消息，使用`ConcurrentHashMap`来跟踪已发送的消息。如果尝试发送已经跟踪的消息，它将不会实际发送消息。

`IdempotentConsumer`类负责消息的消费，它在处理消息之前检查`ConcurrentHashMap`来确定是否已经处理了该消息。如果消息已经被处理，它将不会再次处理该消息。

这个方案确保了消息不会被重复发送或处理，但请注意，这种方案不是`RocketMQ`本身提供的，可能会有内存使用问题，尤其是在处理大量消息时。
另外，这种方案不保证消息的绝对顺序性，因为它可能在存储状态时丢失。在实际应用中，需要将跟踪状态持久化到一个可靠的存储系统中。


## <a id="bds">保证消息的不丢失</a>

常见的策略和技术：
- 持久化存储：将消息存储到磁盘或持久化存储系统中，以防止在系统崩溃或断电时数据丢失。
- 确认机制：发送方在发送消息后等待接收方的确认，如果在一定时间内没有收到确认，则重新发送消息。
- 重试策略：当消息发送失败时，实施重试策略，如指数退避，以避免网络瞬态故障导致的消息丢失。
- 冗余备份：在多个节点上复制消息，即使某个节点失败，其他节点仍然可以继续处理消息。
- 事务支持：使用事务来确保消息处理的原子性，即要么全部成功，要么全部失败，从而避免部分处理状态下的数据不一致。
- 消息队列的使用：消息队列如`RabbitMQ`、`Kafka`等提供了内置的持久化和可靠性保障机制，确保消息在传输过程中不会丢失。
- 死信队列：对于无法处理的消息，将其转移到死信队列中，以便后续分析和处理，防止消息被无限次地重复处理。
- 监控与报警：建立监控系统，对消息处理的各个环节进行监控，一旦发现异常立即报警，及时处理问题，减少消息丢失的风险。
- 幂等性设计：确保消息处理操作具有幂等性，即多次执行同一操作的结果与执行一次相同，避免因重复处理导致的数据错误。

按照`RocketMQ`可以考虑：
- 持久化
   - 消息存储：`RocketMQ`将消息存储在`Broker`的磁盘上，确保即使`Broker`重启，消息也不会丢失。
   - 刷盘策略：`RocketMQ`支持同步刷盘（`SYNC_FLUSH`）和异步刷盘（`ASYNC_FLUSH`）。同步刷盘保证消息写入磁盘后再返回成功，异步刷盘则在写入缓存后即返回成功，但会定期将缓存中的数据刷盘。
- 主从复制
  - 主从架构：`RocketMQ`采用主从架构，主`Broker`负责接收消息并同步到从`Broker`，从`Broker`可以作为备用读取点，提高系统的可用性和容灾能力。
  - 同步复制：在同步复制模式下，主`Broker`必须等到从`Broker`确认接收到消息后才认为消息发送成功。
  - 异步复制：异步复制模式下，主`Broker`在发送消息后不需要等待从`Broker`的确认，这提高了性能，但增加了消息丢失的风险。
- 消息重试
  - 重试机制：`RocketMQ`支持消息重试，如果`Consumer`处理消息失败，消息会被重新投递到队列中，直到达到最大重试次数。
- 事务消息
  - 事务消息：`RocketMQ`支持事务消息，确保消息的发送与本地事务的提交或回滚保持一致。事务消息在发送时先发送为半消息（`half message`），只有当事务状态确认后，消息才会变为可消费状态。
- 监控与报警
  - 监控系统：`RocketMQ`提供了监控指标，可以监控消息的发送、接收、处理情况，以及`Broker`的健康状态。
  - 报警机制：配置报警规则，当系统出现异常时及时通知运维人员，以便快速响应和解决问题。
- 消费者确认
  - 应答机制：`Consumer`在成功处理完消息后需要向`Broker`发送确认，`Broker`根据确认情况决定是否重新投递消息。
- 配置合理的超时和重试策略
  - 合理设置：根据业务需求合理配置消息的超时时间和重试策略，避免不必要的资源浪费和消息积压。

`RocketMQ`发送同步刷盘的消息代码示例：
```java
public class Producer {
    public static void main(String[] args) throws Exception {
        // 创建生产者
        DefaultMQProducer producer = new DefaultMQProducer("producer_group");
        // 指定Namesrv地址
        producer.setNamesrvAddr("localhost:9876");
        // 设置刷盘策略为同步刷盘
        producer.setFlushDiskType(DefaultMQProducer.FlushDiskType.SYNC_FLUSH);

        // 启动生产者
        producer.start();
 
        // 创建消息
        Message msg = new Message("topic_test", "tag_test", "message body".getBytes());
        // 发送消息
        SendResult sendResult = producer.send(msg);
        
        // 打印发送结果
        System.out.println(sendResult);
 
        // 关闭生产者
        producer.shutdown();
    }
}
```

## <a id="sxx">保证消息的顺序性</a>

保证消息的顺序性在消息队列中是一个重要的需求，尤其是在那些需要按照特定顺序处理消息的场景下，比如交易流水、用户操作记录等。

在`RocketMQ`中，保证消息的顺序性可以通过以下几种方式实现：
- 全局顺序消息
  - 全局顺序消息是指在整个消息队列中，所有消息都按照发送的顺序被消费。为了实现这一点，RocketMQ要求生产者只向一个队列发送消息，并且所有的消费者都必须从同一个队列中消费消息。
  - 这种方式适用于消息量较小的场景，因为所有消息都通过单一队列处理，可能会成为瓶颈。
- 分区顺序消息
  - 分区顺序消息允许在多个队列之间进行顺序控制，但保证的是每个分区内的消息顺序。这意味着，如果消息被标记为属于同一个分区，它们将按照发送顺序被消费。
  - 生产者在发送消息时，可以通过设置`MessageQueueSelector`选择特定的队列，通常使用消息键（`messageKey`）来决定消息应该发送到哪个队列。
  - 消费者端也需要配置为顺序消费模式，确保消息按顺序处理。
- 消息键（`Message Key`）
  - 使用消息键可以将相关联的消息路由到相同的队列中，这样就可以保证这些消息在该队列中是有序的。例如，如果消息与特定的用户ID关联，可以使用用户ID作为消息键，确保同一用户的所有消息都在同一队列中处理。
- 单线程消费
  - 为了保证消息的顺序性，消费者可以配置为单线程消费，这意味着在处理消息时只有一个线程在工作，从而避免了多线程并发处理导致的顺序混乱。
- 顺序消费组
  - `RocketMQ`允许创建专门的顺序消费组，这些组中的消费者实例会按照顺序消费消息，而不会进行并行处理。
- 消息队列的选择
  - 生产者在发送消息时，可以通过自定义选择器来指定消息应该进入哪一个队列，从而控制消息的顺序。

在实际应用中，选择哪种方式取决于具体的需求和场景。
全局顺序消息适合消息量小的情况，而分区顺序消息更适合大规模消息处理，同时保持一定程度的顺序性。

`RocketMQ`发送和消费顺序消息，代码示例：
```java
public class OrderedProducer {
    public static void main(String[] args) throws MQClientException, InterruptedException, RemotingException, MQBrokerException {
        // 创建消息生产者，并指定组名
        DefaultMQProducer producer = new DefaultMQProducer("groupName");
        producer.setNamesrvAddr("localhost:9876"); // 设置NameServer地址
        producer.start();
 
        // 发送消息到同一个Topic和Queue中
        String topic = "OrderedTopic";
        for (int i = 0; i < 10; i++) {
            Message msg = new Message(topic, "TagA", "OrderedMessage" + i, ("Hello RocketMQ " + i).getBytes(RemotingHelper.DEFAULT_CHARSET));
            SendResult sendResult = producer.send(msg);
            System.out.printf("%s%n", sendResult);
        }
 
        producer.shutdown();
    }
}
 
public class OrderedConsumer {
    public static void main(String[] args) throws MQClientException {
        // 创建消息消费者，并指定组名
        DefaultMQPushConsumer consumer = new DefaultMQPushConsumer("groupName");
        consumer.setNamesrvAddr("localhost:9876"); // 设置NameServer地址
        consumer.subscribe("OrderedTopic", "TagA"); // 订阅Topic和Tag
        consumer.setMessageModel(MessageModel.CLUSTERING); // 设置集群消费模式
 
        // 注册消息监听器
        consumer.registerMessageListener((MessageListenerConcurrently) (msgs, context) -> {
            try {
                for (Message msg : msgs) {
                    // 处理消息
                    System.out.printf("Consume Thread:%s, QueueID:%d, Message:%s%n", Thread.currentThread().getName(), msg.getQueueId(), new String(msg.getBody()));
                }
                return ConsumeConcurrentlyStatus.CONSUME_SUCCESS;
            } catch (Exception e) {
                e.printStackTrace();
                return ConsumeConcurrentlyStatus.RECONSUME_LATER;
            }
        });
 
        consumer.start();
        System.out.printf("Consumer Started.%n");
    }
}
```

## <a id="gqsx">消息的过期失效问题</a>

RocketMQ中的消息过期失效问题通常是指消息在一定时间内未被消费则会过期。

RocketMQ提供了两种方式来设置消息的过期时间：

- 消息存活时间（Message Age）：通过设置消息属性putUserProperty("__STARTDELIVERTIME", String.valueOf(System.currentTimeMillis() + delayLevel))，可以指定消息的存活时间。

- 消息队列的消息过期时间（Queue Max Size & Message In Memory Size）：在broker配置文件中设置队列的最大消息数量和内存中消息的最大大小，超过这些值的消息会被自动清除。

想要处理消息过期失效的问题，可以考虑以下策略：

- 增加消费者的消费能力，确保它们能够及时处理消息。

- 使用延时消息，通过设置不同的延时级别来让消费者有足够的时间处理消息。

- 对于长时间处理的消息，可以考虑使用定时任务检查消息的状态，如果消费者长时间未处理则重新推送或者标记为过期。

设置消息的存活时间：
```java
// 创建消息
Message msg = new Message("topic", "tag", "message body".getBytes(RemotingHelper.DEFAULT_CHARSET));
 
// 设置消息存活时间，例如30分钟
long currentTime = System.currentTimeMillis();
long delayTime = 30 * 60 * 1000;
msg.putUserProperty("__STARTDELIVERTIME", String.valueOf(currentTime + delayTime));
 
// 发送消息
producer.send(msg);
```


----
