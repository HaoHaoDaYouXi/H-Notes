# 搜索引擎

现在的系统，基本都有通用查询功能，就例如：商城网站的搜索，要可以根据商品名称、商品描述等信息进行搜索。这时就需要搜索引擎了。

搜索引擎是一种根据用户需求与特定算法，运用多种技术检索出指定信息并反馈给用户的信息检索工具。

目前Java开发中涉及到的主流有：Elasticsearch、Solr

## <a id="elasticsearch">Elasticsearch</a>

`Elasticsearch`是一个分布式、`RESTful`风格的搜索和数据分析引擎，能够解决不断涌现出的各种用例。
作为`Elastic Stack`的核心，它集中存储您的数据，帮助您发现意料之中以及意料之外的情况。

`ElasticSearch`是基于`Lucene`的搜索服务器。它提供了一个分布式多用户能力的全文搜索引擎，基于`RESTful Web`接口。

`Elasticsearch`是用`Java`开发的，并作为`Apache`许可条款下的开放源码发布，是当前流行的企业级搜索引擎。

核心特点如下：
- 分布式的实时文件存储，每个字段都被索引且可用于搜索。
- 分布式的实时分析搜索引擎，海量数据下近实时秒级响应。
- 简单的`RESTful API`，天生的兼容多语言开发。
- 易扩展，处理`PB`级结构化或非结构化数据。

基本概念
- 文档（`Document`）：`Elasticsearch`中的基本单位，通常是一个`JSON`格式的对象。
- 索引（`Index`）：类似于数据库的概念，用于存储文档。
- 类型（`Type`）：在`Elasticsearch 7.x`版本之前，一个索引可以包含多个类型；`7.x`版本之后，一个索引只能包含一种类型。
- 映射（`Mapping`）：定义了文档的字段及其数据类型，类似于数据库表的结构定义。
- 分片（`Shard`）：`Elasticsearch`中的最小可寻址单元，每个索引可以被划分为多个分片，以实现水平扩展。
- 副本（`Replica`）：分片的副本，用于提高数据的可用性和容错性。

### `ELK Stack`

`ELK`分别代表`Elasticsearch`、`Logstash`和`Kibana`，这三个工具共同构成了一个完整的日志管理和分析解决方案。

#### `Elasticsearch`

- 简介：`Elasticsearch`是一个基于`Lucene`的分布式搜索和分析引擎，用于存储、搜索和分析大量的数据。
- 功能：
  - 索引：将日志数据索引化，以便快速检索。
  - 搜索：提供高效的全文搜索能力。
  - 聚合：支持对数据进行聚合分析，如统计、分组等。
- 特点：
  - 高性能：能够处理大量的数据和高并发的查询请求。
  - 可扩展性：支持水平扩展，可以通过添加更多的节点来提高性能和存储容量。
  - 易用性：提供`RESTful API`，支持多种编程语言的客户端。

#### `Logstash`

- 简介：`Logstash`是一个用于收集、解析和丰富日志数据的工具。
- 功能：
  - 输入：支持多种数据源，如文件、网络、数据库等。
  - 过滤：提供丰富的插件来解析和转换日志数据。
  - 输出：支持多种输出方式，如`Elasticsearch`、数据库等。
- 特点：
- 灵活性：支持多种插件，可以根据需求定制数据处理流程。
- 实时性：能够实时处理数据流。
- 可扩展性：支持集群部署，可以处理大量数据。

#### `Kibana`

- 简介：`Kibana`是一个用于可视化`Elasticsearch`数据的工具，提供了一个直观的`Web`界面。
- 功能：
  - 仪表板：创建自定义的仪表板来展示数据。
  - 图表：支持多种类型的图表，如折线图、柱状图、饼图等。
  - 搜索：提供高级搜索功能，可以进行复杂查询。
- 特点：
  - 易用性：提供图形界面，无需编程知识即可使用。
  - 交互性：支持实时刷新数据，可以动态调整视图。
  - 可定制性：支持自定义仪表板和视图。

#### `ELK Stack`的工作流程

- 数据收集：`Logstash`从各种数据源收集日志数据。
- 数据处理：`Logstash`对收集到的数据进行解析、转换和丰富。
- 数据存储：处理后的数据被发送到`Elasticsearch`中进行存储和索引。
- 数据可视化：`Kibana`从`Elasticsearch`中读取数据，并以图表和仪表板的形式展示。

#### 应用场景

- 日志分析：收集和分析系统日志，监控系统运行状态。
- 安全监控：监控网络流量和安全事件，发现潜在的安全威胁。
- 业务分析：分析用户行为数据，优化产品和服务。
- 性能监控：监控应用性能，定位性能瓶颈。

### <a id="cxlx">`Elasticsearch`查询类型</a>

`Elasticsearch`支持多种查询类型，这些查询类型可以大致分为两大类：精确匹配和全文检索匹配。

每种查询类型都有其特定的应用场景和特点
- 精确匹配用于：是否完全一致
- 全文检索用于：是否相关

下面是`Elasticsearch`中一些常用的查询类型及其简要说明：

精确匹配查询主要用于查找与查询条件完全一致的文档。这类查询通常用于数值比较、精确字符串匹配等场景
- `term`（精确）：用于查找包含特定值的文档。例如，查找所有`status`字段值为`0`的文档。
- `terms`（精确）：用于查找包含多个特定值的文档。例如，查找所有`status`字段值为`0`或`1`的文档。
- `term set`（精确）：用于查找包含特定值的文档，但值可以是一个集合。例如，查找所有`status`字段值为`0`或`1`或`2`的文档。
- `range`（范围）：用于查找字段值落在指定范围内的文档。例如，查找价格在`100`到`1000`之间的商品。
- `exists`（包含）：用于查找包含特定字段的文档。例如，查找所有包含`test`字段的文档。
- `prefix`（前缀）：用于查找字段值以特定前缀开头的文档。例如，查找所有`name`字段以`test`开头的。
- `ids`（精确）：用于查找包含特定ID的文档。例如，查找所有ID为`123`的文档。
- `fuzzy`（模糊）：根据字段中的模糊匹配进行查询，可以通过设置`fuzziness`参数来控制模糊程度，自动模式格式为：`AUTO:[low],[high]`，默认为`AUTO`相当于`AUTO:3,6`。
- `wildcard`（通配符）：用于通配符符号（`*`和`?`）进行模糊匹配。
- `regexp`（正则）：用于查找字段值符合正则表达式的文档。

全文检索，如：
- `match`：用于全文搜索，会对查询字段进行分词，匹配文档中包含指定词项的文档。例如，查找所有标题中包含`elasticsearch`的文档。
- `match_phrase`：用于短语搜索，会对查询字段进行分词，匹配包含指定短语的文档。例如，查找所有描述中包含`elasticsearch cluster`的文档。
- `match_phrase_prefix`：`match_phrase`的前缀匹配，对查询字段进行分词，匹配包含指定短语的文档，并允许短语中的词项之间存在短语前缀。例如，查找所有描述中包含`elasticsearch cluster`的文档，但允许`cluster`的前缀。
- `multi_match`：用于多字段全文搜索，对多个查询字段进行分词，匹配包含指定词项的文档。例如，查找所有标题和描述中包含`elasticsearch`的文档。
- `query_string`：用于全文搜索，对查询字段进行分词，匹配包含指定词项的文档，并支持模糊匹配、短语匹配等。例如，查找所有标题中包含`elasticsearch`的文档，并允许短语匹配。
- `more_like_this`：用于查找与指定文档相似的其他文档。例如，查找与一篇博客文章主题相似的其他文章。
- `bool`：用于组合多个查询条件和过滤器，支持必须（`must`）、应该（`should`）、不允许（`must_not`）等逻辑组合。

其他查询，如：
- `script_score`：用于根据自定义脚本来评分文档。
- `function_score`：用于根据自定义函数来评分文档。
- `has_child`：用于查找拥有特定子文档的文档。
- `has_parent`：用于查找属于特定父文档的文档。
- `nested`：用于在嵌套字段中进行查询。
- ...

match 查询示例，用于查找所有标题中包含 "elasticsearch" 的文档：
```
GET /my-index/_search
{
  "query": {
    "match": {
      "title": "elasticsearch"
    }
  }
}
``` 

### <a id="cxlx">`Elasticsearch`聚合</a>

聚合有助于从搜索中使用的查询中收集数据，聚合为各种统计指标，便于统计信息或做其他分析。
聚合可求得：平均值、最大值、最小值、总和，等等指标。

- 分桶`Bucket`聚合
  - 根据字段值，范围或其他条件将文档分组为桶（也称为箱）。
- 指标`Metric`聚合
  - 从字段值计算指标（例如总和或平均值）的指标聚合。
- 管道`Pipeline`聚合
  - 子聚合，从其他聚合（而不是文档或字段）获取输入。

#### 分桶聚合示例

**日期区间（Date Histogram）**：按照日期区间分组。
```
GET /my-index/_search
{
  "size": 0,
  "aggs": {
    "sales_by_month": {
      "date_histogram": {
        "field": "order_date",
        "calendar_interval": "month"
      },
      "aggs": {
        "total_sales": {
          "sum": {
            "field": "price"
          }
        }
      }
    }
  }
}
```

**区间（Range）**：按照数值区间分组。
```
GET /my-index/_search
{
  "size": 0,
  "aggs": {
    "price_ranges": {
      "range": {
        "field": "price",
        "ranges": [
          { "to": 100 },
          { "from": 100, "to": 500 },
          { "from": 500 }
        ]
      }
    }
  }
}
```

**直方图（Histogram）**：按照数值间隔分组。
```
GET /my-index/_search
{
  "size": 0,
  "aggs": {
    "age_groups": {
      "histogram": {
        "field": "age",
        "interval": 10
      }
    }
  }
}
```

#### 指标聚合示例

**最大值（Max）**：找出某个字段的最大值。
```
GET /my-index/_search
{
  "size": 0,
  "aggs": {
    "max_price": {
      "max": {
        "field": "price"
      }
    }
  }
}
```

**平均值（Average）**：计算某个字段的平均值。
```
GET /my-index/_search
{
  "size": 0,
  "aggs": {
    "avg_price": {
      "avg": {
        "field": "price"
      }
    }
  }
}
```

**基数（Cardinality）**：计算某个字段的唯一值数量。
```
GET /my-index/_search
{
  "size": 0,
  "aggs": {
    "unique_customers": {
      "cardinality": {
        "field": "customer_id"
      }
    }
  }
}
```

#### 管道聚合示例

**百分比（Derivative Aggregation）**：计算每个区间的销售占比。
```
GET /my-index/_search
{
  "size": 0,
  "aggs": {
    "price_ranges": {
      "range": {
        "field": "price",
        "ranges": [
          { "to": 100 },
          { "from": 100, "to": 500 },
          { "from": 500 }
        ]
      },
      "aggs": {
        "total_sales": {
          "sum": {
            "field": "price"
          }
        },
        "percentage_of_total": {
          "bucket_script": {
            "buckets_path": {
              "total_sales": "total_sales",
              "total_sales_all": "_total.total_sales.value"
            },
            "script": "params.total_sales / params.total_sales_all * 100"
          }
        }
      }
    }
  }
}
```

### <a id="jq">集群</a>

`Elasticsearch`集群是由一个或多个`Elasticsearch`节点组成的集合，它们共同协作来存储和处理数据。
集群的设计目的是为了提供高可用性、容错能力和可扩展性。

#### 集群组成

- 节点（`Node`）：`Elasticsearch`集群中的一个实例，可以运行在一个或多个物理机器上。
- 主节点（`Master Node`）：负责管理集群的元数据，如索引设置和分片分配，不负责文档级别的管理。
- 数据节点（`Data Node`）：负责存储实际的数据，可以关闭`http`功能。
- 客户端节点（`Client Node`）：不存储数据，只转发请求到数据节点。

#### 集群配置

- 单机部署
  - 对于开发和测试环境，可以使用单机部署`Elasticsearch`。只需要安装并启动`Elasticsearch`实例即可。
- 集群部署
  - 对于生产环境，通常需要部署多个节点以实现高可用性和负载均衡。

#### 配置步骤

- 解压`Elasticsearch`：从官方下载`Elasticsearch`压缩包，解压出多个实例。
- 修改配置：编辑每个实例的`config/elasticsearch.yml`文件，配置必要的参数。
  - `cluster.name`：设置集群名称。
  - `node.name`：设置节点名称。
  - `network.host`：设置节点绑定的`IP`地址。
  - `discovery.seed_hosts`：设置集群中其他节点的地址，用于节点发现。
  - `cluster.initial_master_nodes`：设置初始主节点的名称，用于选举主节点。
- 启动节点：启动每个`Elasticsearch`实例。

#### 集群管理

- 查看集群状态：可以使用`_cat API：GET _cat/nodes?v=true&h=ip,node.role,master`或者`Kibana`来查看集群的状态。
- 查看集群健康状况：`GET _cluster/health`
- 查看集群配置：`GET _cluster/settings`

#### 分片与副本

- 分片（`Shard`）：`Elasticsearch`将索引分成多个分片，每个分片都是一个独立的`Lucene`实例。
- 副本（`Replica`）：分片的副本，用于提高数据的可用性和容错性。

可以通过配置索引的`number_of_shards`和`number_of_replicas`参数来控制分片和副本的数量。

#### 集群操作

- 添加新节点
  - 配置新节点：确保新节点的配置文件正确配置了集群名称和其他必要的参数。
  - 启动新节点：启动新节点后，它将自动加入集群。
- 移除节点
  - 关闭节点：停止节点的服务。
  - 等待重新平衡：集群会自动重新分配分片。
- 重新平衡分片
  - 如果需要手动触发分片重新平衡，可以更改集群的`cluster.routing.allocation.enable`设置。

#### 故障转移
集群具有内置的故障转移机制。当主节点或数据节点失败时，集群会自动选举新的主节点，并恢复数据的可用性。

##### `Master`选举

前置条件：
- 只有是候选主节点（`master：true`）的节点才能成为主节点。
- 最小主节点数（`min_master_nodes`）的目的是防止脑裂。

`Elasticsearch`选主是`ZenDiscovery`模块负责的

主要包含`Ping`（节点之间通过这个RPC来发现彼此）和`Unicast`（单播模块包含一个主机列表以控制哪些节点需要 ping 通）

获取主节点的核心入口为`findMaster`，选择主节点成功返回对应`Master`，否则返回`null`。

选举流程大致描述如下：
- 确认候选主节点数达标，`elasticsearch.yml`设置的值`discovery.zen.minimum_master_nodes`
- 对所有候选主节点根据`nodeId`字典排序，每次选举每个节点都把自己所知道节点排一次序，然后选出第一个（第`0`位）节点，暂且认为它是`master`节点。
- 如果对某个节点的投票数达到一定的值（候选主节点数`n/2+1`）并且该节点自己也选举自己，那这个节点就是`master`。否则重新选举一直到满足上述条件。

###### `Master`节点和候选`Master`节点

主节点负责集群相关的操作，例如创建或删除索引，跟踪哪些节点是集群的一部分，以及决定将哪些分片分配给哪些节点。
- 稳定的主节点是衡量集群健康的重要标志。
- 候选主节点是被选具备候选资格，可以被选为主节点的那些节点。

当集群候选`Master`数量不小于`3`个，可以通过设置最少投票通过数量（`discovery.zen.minimum_master_nodes`）超过所有候选节点一半以上来解决脑裂问题（两个候选节点各占一半）。
当候选数量为两个时，只能修改为唯一的一个`Master`候选，其他作为`data`节点，避免脑裂问题。

### <a id="cyml">常用的`cat`命令</a>

| 含义    | 命令                       |
|-------|--------------------------|
| 别名    | GET _cat/aliases?v       |
| 分配相关  | GET _cat/allocation      |
| 计数    | GET _cat/count?v         |
| 字段数据  | GET _cat/fielddata?v     |
| 运行状况  | GET_cat/health?          |
| 索引相关  | GET _cat/indices?v       |
| 主节点相关 | GET _cat/master?v        |
| 节点属性  | GET _cat/nodeattrs?v     |
| 节点    | GET _cat/nodes?v         |
| 待处理任务 | GET _cat/pending_tasks?v |
| 插件    | GET _cat/plugins?v       |
| 恢复    | GET _cat / recovery?v    |
| 存储库   | GET _cat /repositories?v |
| 段     | GET _cat /segments?v     |
| 分片    | GET _cat/shards?v        |
| 快照    | GET _cat/snapshots?v     |
| 任务    | GET _cat/tasks?v         |
| 模板    | GET _cat/templates?v     |
| 线程池   | GET _cat/thread_pool?v   |

### <a id="fc">分词</a>

分词是文本分析的一个重要组成部分，它涉及到将文本分割成一系列的单词（`tokens`）或术语（`terms`）。
这个过程对于建立索引和执行全文搜索至关重要。

#### 分词的概念

分词是指将文本转换成一系列单词（`term or token`）的过程，也可以叫做文本分析，在`Elasticsearch`中称为`analysis`。
分词会在创建或更新文档时进行，对相应的文档进行分词处理，以便后续的索引和搜索。

#### 分词器（Analyzer）

分词器是用于定义如何将文本转换为一系列词项的规则集。

- 字符过滤器（`Char Filter`）：在分词之前对文本进行预处理，如去除`HTML`标签、转换大小写等。
- 分词器（`Tokenizer`）：将文本分割成一系列词项。
- 标记过滤器（`Token Filter`）：对分词器产生的词项进行进一步处理，如去除停用词、转换大小写、词干提取等。

#### 分词器

`Elasticsearch`提供了一些分词器，这些分词器可以满足大多数常见的文本分析需求。

- `Standard Analyzer`：标准分词器，这个是默认的分词器，使用Unicode文本分割算法，它会去除标点符号将文本按单词切分并且转换为小写。
- `Simple Analyzer`：简单分词器，它会将文本分割成词项，并将它们转换为小写。
- `Stop Analyzer`：停用词分词器，类似于`Simple Analyzer`，增加了停用词过滤（如 a、an、and、are、as、at、be、but 等）。
- `Whitespace Analyzer`：空白分词器，仅按空白符分割文本，并不进行小写转换。
- `Keyword Analyzer`：关键字分词器，不会对文本进行任何分词处理，而是将整个文本作为一个词项。
- `Pattern Analyzer`：模式分词器，使用正则表达式来定义分词规则，默认使用 \W+ (非字符分隔)，支持小写转换和停用词删除。
- `Language Analyzer`：针对特定语言的分词器，如英语、德语等。
- `Customer Analyzer`：自定义分词器，除了内置分词器外，还可以自定义分词器来满足特定的需求。这可以通过组合不同的字符过滤器、分词器和标记过滤器来实现。

#### 分词器配置

分词器可以在索引级别的`settings`中进行配置。
示例：创建一个索引，并定义一个自定义分词器`my_analyzer`，该分词器使用`standard`分词器，并添加一个`lowercase`标记过滤器。
```
PUT my_index
{
  "settings": {
    "analysis": {
      "analyzer": {
        "my_analyzer": {
          "tokenizer": "standard",
          "filter": ["lowercase"]
        }
      }
    }
  }
}
```

#### 分析文档

可以使用`_analyze API`来测试分词器的效果。
示例：使用上面定义的`my_analyzer`来分析一段文本。
```
GET my_index/_analyze
{
 "analyzer": "my_analyzer",
 "text": "Hello World! This is a test."
}
```

#### 分词器的工作流程

分词器的工作流程一般如下：
- 字符过滤器：对原始文本进行预处理。
- 分词器：将处理后的文本分割成一系列词项。
- 标记过滤器：对词项进行进一步处理，如去除停用词、转换大小写、词干提取等。

#### 停用词

停用词是一些常见的词汇，如冠词、介词等，在搜索时通常会被忽略。可以使用停用词过滤器来去除这些词项。
示例：定义一个停用词过滤器，并将其应用于分词器。
```
PUT my_index
{
 "settings": {
   "analysis": {
     "analyzer": {
       "my_analyzer": {
         "tokenizer": "standard",
         "filter": ["lowercase", "stop"]
       }
     },
     "filter": {
       "stop": {
         "type": "stop",
         "stopwords": ["the", "is", "a"]
       }
     }
   }
 }
}
```

#### 总结

分词是文本分析的核心部分，通过合理的配置分词器，可以有效地提高搜索的准确性和效率。
内置分词器可以满足大多数需求，而对于更复杂的情况，可以通过自定义分词器来实现。

### <a id="sy">索引</a>



## <a id="solr">Solr</a>








---- 
