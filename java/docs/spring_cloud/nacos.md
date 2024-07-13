# Nacos

`Nacos`是阿里开放的一款中间件，它主要提供三种功能：持久化节点注册，非持久化节点注册和配置管理。

`Nacos`提供了简单易用的特性集，帮助您快速实现动态服务发现、服务配置、服务元数据及流量管理。

`Nacos`支持几乎所有主流类型的服务的发现、配置和服务管理平台，提供`注册中心`、`配置中心`
和`动态 DNS 服务`三大功能。

能够无缝对接`Springcloud`、`Spring`、`Dubbo`等流行框架。

**下图是`Nacos`的架构图：**

![nacos架构图](img/nacos.jpeg)

- `Provider APP`：服务提供者
- `Consumer APP`：服务消费者
- `Name Server`：通过`VIP`（`Virtual IP`）或`DNS`的方式实现`Nacos`高可用集群的服务路由
- `Nacos Server`：`Nacos`服务提供者，里面包含的`Open API`是功能访问入口，`Conig Service`、`Naming Service`是`Nacos`提供的配置服务、命名服务模块。`Consitency Protocol`是一致性协议，用来实现`Nacos`集群节点的数据同步，这里使用的是`Raft`算法（`Etcd`、`Redis`哨兵选举）
- `Nacos Console`：控制台

## <a id="zczxyl">注册中心原理</a>

服务实例在启动时注册到服务注册表，并在关闭时注销，
服务消费者查询服务注册表，获得可用实例，
服务注册中心需要调用服务实例的健康检查API来验证它是否能够处理请求。

在`Spring-Cloud-Common`包中有一个类`org.springframework.cloud.client.serviceregistry.ServiceRegistry`，
它是`Spring Cloud`提供的服务注册的标准。
集成到`Spring Cloud`中实现服务注册的组件，都会实现该接口。
该接口有一个实现类是`NacoServiceRegistry`。

### <a id="jcsxgc">`SpringCloud`集成`Nacos`的实现过程</a>
在`spring-clou-commons`包的`META-INF/spring.factories`中添加自动装配的配置信息，
其中`AutoServiceRegistrationAutoConfiguration`就是服务注册相关的配置类：
```java
@Configuration(proxyBeanMethods = false)
@Import(AutoServiceRegistrationConfiguration.class)
@ConditionalOnProperty(value ="spring.cloud.service-registry.auto-registration.enabled",matchIfMissing = true)
public class AutoServiceRegistrationAutoConfiguration{
    @Autowired(required = false)
    private AutoServiceRegistration autoServiceRegistration;
    @Autowired
    private AutoServiceRegistrationProperties properties;
    @PostConstruct
    protected void init() {
        if (this.autoServiceRegistration == null && this.properties.isFailFast()) {
            throw new IllegalStateException("Auto Service Registration has been requested,but there is no AutoServiceRegistration bean");
        }
    }
}
```
在`AutoServiceRegistrationAutoConfiguration`配置类中，可以看到注入了一个`AutoServiceRegistration`实例，
`AbstractAutoServiceRegistration`抽象类实现了该接口，
并且`NacosAutoServiceRegistration`继承了`AbstractAutoServiceRegistration`。

看`EventListener`我们就知道，`Nacos`是通过`Spring`的事件机制集成到`SpringCloud`中去的。

`AbstractAutoServiceRegistration`实现了`onApplicationEvent`抽象方法，并且监听`WebServerInitializedEvent`事件(当`WebServer`初始化完成之后)，调用`this.bind`方法。
```java
@Override
public void onApplicationEvent(WebServerInitializedEvent event) {
    bind(event);
}
@Deprecated
public void bind(WebServerInitializedEvent event) {
    ApplicationContext context = event.getApplicationContext();
    if (context instanceof ConfigurableWebServerApplicationContext){
        if ("management".equals(((ConfigurableWebServerApplicationContext)context).getServerNamespace )){
            return;
        }
    }
    this.port.compareAndSet(0, event.getWebServer().getPort());
    this.start();
}
```
最终会调用`NacosServiceRegistry.register()`方法进行服务注册。
```java
public void start() {
    if (!isEnabled()) {
        if (logger.isDebugEnabled()) {
            logger.debug("Discovery Lifecycle disabled. Not starting");
        }
        return;
    }
    // only initialize if nonSecurePort is greater than 0 and it isn't already running
    // because of containerPortInitializer below
    if (!this.running.get()){
        this.context.publishEvent(new InstancePreRegisteredEvent(this,getRegistration()));
        register();
        if (shouldRegisterManagement()){
            registerManagement();
        }
        this.context.publishEvent(new InstanceRegisteredEvent<>(this, getConfiguration()));
        this.running.compareAndSet(false,true);
    }
}

protected void register(){
    this.serviceRegistry.register(getRegistration());
}
```

### <a id="nacosserviceregistry">`NacosServiceRegistry`实现</a>
在`NacosServiceRegistry.registry`方法中，调用了`Nacos Client SDK`中的`namingService.registerInstance`完成服务的注册。
```java
@Override
public void register(Registration registration){
    if (StringUtils.isEmpty(registration.getServiceId())) {
        log.warn("No service to register for nacos client...");
        return;
    }
    String serviceId = registration.getServiceId();
    Instance instance = getNacosInstanceFromRegistration(registration);
    try{
        namingService.registerInstance(serviceId,instance);
        log.info("nacos registry,{} {} : {}register finished", serviceId,instance.getIp(),instance.getPort());
    }catch (Exception e) {
        log.error("nacos registry, {} register failed... {},",serviceId,registration.toString(),e);
    }
}
```
继续看`NacosNamingService`的`registerInstance()`方法：
```java
@Override
public void registerInstance(String serviceName, Instance instance) throws NacosException {
    registerInstance(serviceName,Constants.DEFAULT_GROUP,instance);
}
@Override
public void registerInstance(String serviceName, String groupName, Instance instance) throws NacosException {
    if (instance.isEphemeralO){
        BeatInfo beatInfo = new BeatInfo();
        beatInfo. setServiceName(NamingUtils.getGroupedName(serviceName, groupName));
        beatInfo.setIp(instance.getIp());
        beatInfo.setPort(instance.getPort());
        beatInfo.setCluster(instance.getClusterName());
        beatInfo.setWeight(instance.getWeight());
        beatInfo.setMetadata(instance .getMetadata());
        beatInfo.setScheduled(false);
        long instanceInterval = instance.getInstanceHeartBeatInterval();
        beatInfo.setPeriod(instanceInterval == 0 ? DEFAULT_HEART_BEAT_INTERVAL : instanceInterval);
        beatReactor.addBeatInfo(NamingUtils.getGroupedName(serviceName, groupName), beatInfo);
    }
    serverProxy.registerService(Namingutils.getGroupedName(serviceName, groupName), groupName, instance);
}
```
通过`beatReactor.addBeatInfo()`创建心跳信息实现健康检测，`Nacos Server`必须要确保注册的服务实例是健康的，而心跳检测就是服务健康检测的手段。
最后通过`serverProxy.registerService()`实现服务注册。

### <a id="xtjc">心跳机制</a>
```java
public void addBeatInfo(String serviceName, BeatInfo beatInfo) {
    NAMING_LOGGER.info("[BEAT] adding beat: {} to beat map.", beatInfo);
    dom2Beat.put(buildKey(serviceName, beatInfo.getIp(), beatInfo.getPort()), beatInfo);
    executorService.schedule(new BeatTask(beatInfo), 0, TimeUnit.MILLISECONDS);
    MetricsMonitor.getDom2BeatSizeMonitor().set(dom2Beat.size());
}
```
从上述代码看，所谓心跳机制就是客户端通过`schedule`定时向服务端发送一个数据包，然后启动一个线程不断检测服务端的回应，
如果在设定时间内没有收到服务端的回应，则认为服务器出现了故障。

`Nacos`服务端会根据客户端的心跳包不断更新服务的状态。

#### <a id="jkcjlzms">服务的健康检查分为两种模式</a>
- 客户端上报模式：客户端通过心跳上报的方式告知`Nacos`注册中心健康状态（默认心跳间隔`5s`，`Nacos`将超过超过`15s`未收到心跳的实例设置为不健康，超过`30s`将实例删除）
- 服务端主动检测：`Nacos`主动检查客户端的健康状态（默认时间间隔`20s`，健康检查失败后会设置为不健康，不会立即删除）

`Nacos`目前的`instance`有一个`ephemeral`字段属性，该字段表示实例是否是临时实例还是持久化实例。
如果是临时实例则不会在`Nacos`中持久化，需要通过心跳上报，如果一段时间没有上报心跳，则会被`Nacos`服务端删除。
删除后如果又重新开始上报，则会重新实例注册。
而持久化实例会被`Nacos`服务端持久化，此时即使注册实例的进程不存在，这个实例也不会删除，只会将健康状态设置成不健康。

这里就涉及到了`Nacos`的`AP`和`CP`模式 ，默认是`AP`，即`Nacos`的`client`的节点注册时`ephemeral=true`，那么`Nacos`集群中这个`client`节点就是`AP`，采用的是`distro`协议，而`ephemeral=false`时就是`CP`采用的是`raft`协议实现。

```properties
spring.cloud.nacos.discovery.ephemeral=true
```
`false`为永久实例，`true`表⽰临时实例开启，注册为临时实例，默认是`true`

`Nacos`的两种心跳机制是为了：
- 对于临时实例，健康检查失败，则直接删除。这种特性适合于需要应对流量突增的场景，服务可以弹性扩容，当流量过去后，服务停掉即可自动注销。
- 对于持久化实例，健康检查失败，会设置为不健康状态。它的优点就是可以实时的监控到实例的健康状态，便于后续的告警和扩容等一系列处理。

#### <a id="zwbh">自我保护</a>
`Nacos`也有自我保护机制（当前健康实例数/当前服务总实例数），值为`0-1`之间的浮点类型。

正常情况下`Nacos`只会健康的实例。
但在高并发场景，如果只返回健康实例的话，流量洪峰到来可能直接打垮剩下的健康实例，产生`雪崩效应`。

保护阈值存在的意义在于当服务`A`的`健康实例数/总实例数 < 保护阈值`时，`Nacos`会把该服务所有的实例信息（健康的+不健康的）全部提供给消费者，
消费者可能访问到不健康的实例，请求失败，但这样远比造成雪崩要好。
牺牲了请求，保证了整个系统的可用。

简单来说不健康实例的另外一个作用：防止雪崩

如果所有的实例都是临时实例，当雪崩出现时，`Nacos`的阈值保护机制是不是就没有足够的（包含不健康实例）实例返回了，其实如果有部分实例是持久化实例，即便它们已经挂掉，状态为不健康，但当触发自我保护时，还是可以起到分流的作用。

### <a id="sxzc">实现注册</a>
`Nacos`提供了`SDK`和`Open API`两种形式来实现服务注册。

**`Open API`：**
```shell
curl -X POST 'http://127.0.0.1:8848/nacos/v1/ns/instance?serviceName=nacos.naming.serviceName&ip=192.16813.1&port=8080'
```

**`SDK`：**
```java
void registerInstance(String serviceName, String ip, int port) throws NacosException;
```

这两种形式本质都一样，底层都是基于`HTTP`协议完成请求的。

所以注册服务就是发送一个`HTTP`请求：
```java
public void registerService(String serviceName, String groupName, Instance instance) throws NacosException {
	NAMING_LOGGER.info("[REGISTER-SERVICE] {} registering service {} with instance: {}",namespaceId,serviceName,instance);
	final Map<String,String> params = new HashMap<>(9);
	params.put(CommonParams.NAMESPACE_ID,namespaceId);
	params.put(CommonParams.SERVICE_NAME,serviceName);
	params.put(CommonParams.GROUP_NAME,groupName);
	params.put(CommonParams.CLUSTERNAME,instance.getClusterName());
	params.put("ip",instance.getIp());
	params .put("port",String. valueOf(instance.getPort()));
	params.put("weight",String.valueOf(instance.getWeight()));
	params.put("enable",String.valueOf(instance.isEnabled()));
	params.put("healthy",String.valueOf(instance.isHealthy()));
	params.put("ephemeral",String.valueOf(instance.isEphemeral()));
	params.put("metadata",JSON.toJSONString(instance.getMetadata()));
	
	regAPI(UtilAndComs .NACOS_URL_INSTANCE,params,HttpMethod.POSD);
}
```

对于`Nacos`服务端，对外提供的服务接口请求地址为`nacos/v1/ns/instance`，实现代码在`nacos-naming`模块下的`InstanceController`类中：
```java
@RestController
@RequestMapping(UtilsAndCommons.NACOS_NAMING_CONTEXT+"/instance")
public class InstanceController{
    //...
    @CanDistro
    @PostMapping
    public String register(HttpServletRequest request) throws Exception {
        String serviceName = WebUtils.required(request,CommonParams.SERVICENAME);
        String namespaceId = WebUtils.optional(request,CommonParams,NAMESPACE_ID,Constants.DEFAULT_NAMESPACE_ID);
        serviceManager.registerInstance(namespaceId,serviceName,parseInstance(request));
        return"ok";
    }
    //...
}
```
从请求参数汇总获得`serviceName(服务名)`和`namespaceId(命名空间Id)`

#### 调用`registerInstance`注册实例
```java
public void registerInstance(String namespaceId, String serviceName, Instance instance) throws NacosException{
    createEmptyService(namespaceId,serviceNameinstance.isEphemeral());
    Service service=getService(namespaceId,serviceName);
    if (service== null){
        throw new NacosException(NacosException.INVALID_PARAM,"service not found,namespace:"+namespaceId +",service:"+serviceName);
    }
    addInstance(namespaceId,serviceName,instance.isEphemeral(),instance);
}
```
- 创建一个控服务（在`Nacos`控制台服务列表中展示的服务信息），实际上是初始化一个`serviceMap`，它是一个`ConcurrentHashMap`集合
- `getService`，从`serviceMap`中根据`namespaceId`和`serviceName`得到一个服务对象
- 调用`addInstance`添加服务实例

```java
public void createServiceIfAbsent(String namespaceId, String serviceName, boolean local,Cluster cluster) throws NacosException {
    Service service = getService(namespaceId,serviceName);
    if(service== null){
        service= new Service();
        service.setName(serviceName);
        service.setNamespaceId(namespaceId);
        service.setGroupName(NamingUtils.getGroupName(serviceName));
        service.setLastModifiedMillis(System.currentTimeMillis());
        service.recalculateChecksum();
        if(cluster != null){
            cluster.setService(service);
            service.getClusterMap().put(cluster.getName(),cluster);
        }
        service.validate();
        putServiceAndInit(service);
        if(!local){
            addOrReplaceService(service);
        }
    }
}
```
- 根据`namespaceId`、`serviceName`从缓存中获取`Service`实例
- 如果`Service`实例为空，则创建并保存到缓存中

```java
private void putServiceAndInit(Service service) throws NacosException{
    putService(service);
    service.init();
    consistencyService.listen(KeyBuilder.buildInstanceListKey(service.getNamespaceId(),service.getName(),true),service);
    consistencyService.listen(KeyBuilder.buildInstanceListKey(service.getNamespaceId(),service.getName(),false),service);
    Loggers.SRV_LOG.info("[NEW-SERVICE]{}",service.toJSON());
}
```
- 通过`putService()`方法将服务缓存到内存
- `service.init()`建立心跳机制
- `consistencyService.listen`实现数据一致性监听
- `service.init()`方法的如下图所示，它主要通过定时任务不断检测当前服务下所有实例最后发送心跳包的时间。
- 如果超时,则设置`healthy`为`false`表示服务不健康,并且发送服务变更事件。

注意：`Nacos`客户端注册服务的同时也建立了心跳机制。

`putService`方法，它的功能是将`Service`保存到`serviceMap`中：
```java
public void putService(Service service) {
    if(!serviceMap.containsKey(service.getNamespaceId())) {
        synchronized (putServiceLock) {
            if(!serviceMap.containsKey(service,getNamespaceId())) {
                serviceMap.put(service.getNamespaceId(),new ConcurrentHashMap<>(16));
            }
        }
    }
    serviceMap.get(service.getNamespaceId()).put(service.getName(),service);
}
```

继续调用`addInstance`方法把当前注册的服务实例保存到`Service`中：
```java
addInstance(namespaceId,serviceName,instance.isEphemeral(),instance)
```

### <a id="zj">简单来说</a>
- `Nacos`客户端通过`Open API`的形式发送服务注册请求
- `Nacos`服务端收到请求后，做以下三件事：
- 构建一个`Service`对象保存到`ConcurrentHashMap`集合中
- 使用定时任务对当前服务下的所有实例建立心跳检测机制
- 基于数据一致性协议服务数据进行同步

### <a id="dzcx">服务提供者地址查询</a>

**`Open API`：**
```shell
curl -X GET127.00.1:8848/nacos/v1/ns/instance/list?serviceName=example
```

**`SDK`：**
```java
List<Instance> selectInstances(String serviceName, boolean healthy) throws NacosException;
```

`InstanceController`中的`list`方法：
```java
@GetMapping("/list")
public JSONObject list(HttpServletRequest request) throws Exception {
    String namespaceId = WebUtils.optional(request,CommonParams,NAMESPACE_ID,Constants.DEFAULT_NAMESPACE_ID);
    String serviceName = WebUtils.required(request,CommonParams.SERVICE_NAME);
    String agent =WebUtils.getUserAgent(request);
    String clusters = WebUtils.optional(request,"clusters",StringUtils.EMPTY);
    String clientIP = WebUtils.optional(request,"clientIp", StringUtils.EMPTY);
    Integer udpPort = Integer.parseInt(WebUtils.optional(request, "udpPort","0"));
    String env= WebUtils.optional(request,"env",StringUtils.EMPTY);
    boolean isCheck = Boolean.parseBoolean(WebUtils.optional(request,"isCheck","false"));
    String app= WebUtils.optional(request,"app",StringUtils.EMPTY);
    String tenant = WebUtils.optional(request,"tid",StringUtils.EMPTY);
    boolean healthyOnly = Boolean.parseBoolean(WebUtils.optional(request,"healthyOnly","false"));
    return doSrvIPXT(namespaceld, serviceName, agent, clusters, clientIP, udpPort, env,isCheck,app,tenant,healthyOnly);
}
```
#### 解析请求参数
通过`doSrvIPXT`返回服务列表数据
```java
public JSONObject doSrvIPXT(String namespaceId, String serviceName, String agent, String clusters,String clientIP,int udpPort,String env,boolean isCheck,String app,String tid,boolean healthyOnly) throws Exception {
    //...
    ClientInfo clientInfo = new ClientInfo(agent);
    JSONObject result=new JSONObject();
    Service service= serviceManager.getService(namespaceId,serviceName);
    List<Instance> srvedIPs;
    //获取指定服务下的所有实例 IP
    srvedIPs = service.srvIPs(Arrays.asList(StringUtils.split(clusters,",")));
    Map<Boolean,List<Instance>>ipMap =new HashMap<>(2);
    ipMap.put(Boolean.TRUE,new ArrayList<>());
    ipMap.put(Boolean.FALSE,new ArrayList<>());
    for (Instance ip : srvedIPs){
        ipMap.get(ip.isHealthy()).add(ip);
    }
    //遍历,完成JSON字管中的纠装
    JSONArray hosts = new JSONArray();
    for (Map.Entry<Boolean, List<Instance>> entry : ipMap.entrySet()) {
        List<Instance> ips = entry.getValue();
        if (healthyOnly && !entry.getKey()){
            continue;
        }
        for (Instance instance :ips) {
            if (!instanceisEnabled()) {
                continue;
            }
            JSONObject ipobj=new JSONObject();
            ipobj.put("ip",instance.getIp());
            ipObj.put("port",instance.getPort());
            ipObj.put("valid",entry.getKey());
            ipObj.put("healthy",entry.getKey());
            ipObj.put("marked",instance.isMarked());
            ipObj.put("instanceId",instance.getInstanceId());
            ipObj.put("metadata",instance.getMetadata());
            ipObj.put("enabled",instance.isEnabled());
            ipObj.put("weight",instance.getweight());
            ipObj.put("clusterName",instance.getClusterName());
            if(clientInfo.type== ClientInfo.ClientType.JAVA
            && clientInfo.version.compareTo(VersionUtil.parseVersion("1.0."))>=0){
                ipObj.put("serviceName",instance.getServiceName());
            }else{
                ipObj.put("serviceName",NamingUtils.getServiceName(instance.getServiceName()));
            }
            ipObj.put("ephemeral",instance.isEphemeral());
            hosts.add(ipobj);
        }
    }
    result.put("hosts",hosts);
    result.put("name",serviceName);
    result.put("cacheMillis",cacheMillis);
    result.put("lastRefTime",System.currentTimeMillis());
    result.put("checksum",service.getChecksum());
    result.put("useSpecifiedURL",false);
    result.put("clusters",clusters);
    result.put("env",env);
    result.put("metadata",service.getMetadata());
    return result;
}
```
- 根据`namespaceId`、`serviceName`获得`Service`实例
- 从`Service`实例中基于`srvIPs`得到所有服务提供者实例
- 遍历组装`JSON`字符串并返回

### <a id="dtgz">`Nacos`服务地址动态感知原理</a>
可以通过`subscribe`方法来实现监听，其中`serviceName`表示服务名、`EventListener`表示监听到的事件：
```java
void subscribe(String serviceName, EventListener listener) throws NacosException;
```

具体调用方式如下：
```java
NamingService naming = NamingFactory.createNamingService(System.getProperty("serveAddr"));
naming.subscribe("example",event->(
    if (event instanceof NamingEvent) {
        System.out.println(((NamingEvent) event).getServceName());
        System.out.printIn(((NamingEvent) event).getInstances());
    }
});
```

或者调用`selectInstance`方法，如果将`subscribe`属性设置为`true`，会自动注册监听：
```java
public List<Instance> selectInstances(String serviceName, List<String> clusters, boolean healthy,boolean subscribe){}
```

`Nacos`客户端中有一个`HostReactor`类，它的功能是实现服务的动态更新，基本原理是：
- 客户端发起时间订阅后，在`HostReactor`中有一个`UpdateTask`线程，每`10s`发送一次`Pull`请求，获得服务端最新的地址列表
- 对于服务端，它和服务提供者的实例之间维持了心跳检测，一旦服务提供者出现异常，则会发送一个`Push`消息给`Nacos`客户端，也就是服务端消费者
- 服务消费者收到请求之后，使用HostReactor中提供的`processServiceJSON`解析消息，并更新本地服务地址列表

## <a id="pzzx">配置中心原理</a>
- 客户端启动后，每`30`秒给`Server`发送一个心跳包
- `Server`拿到心跳包之后，先对比一下数据版本
- 如果版本一样说明数据没有变化，这时Server不会立即将该心跳返回，Server会一直拿着这个心跳，此时和客户端保持长连接的状态，直到数据有变化或者持有超过`29.5`秒
- 如果客户端感知到数据版本发生变化，就会主动请求`Server`拉取数据

阿里的中间件都有个特点，不像一个纯粹的中间件，更像是业务锤炼出来的产物，在`RocketMQ`，`Nacos`上特别明显
它总是会考虑非常多的业务场景，在性能与好用性方面做一个取舍
- 它也许不是纯粹的，也许不是性能最好的，但是一定是最适合拿来做业务的。

### Nacos 客户端

`Nacos`客户端所有的这个文件配置实现主要在`NacosNamingService`的类下面，这个配置中心主要在`NacosConfigService`的类下面。

该接口下面主要有一些获取配置，发布配置，增加监听器，删除配置，删除监听器等操作。
```java
public interface ConfigService {
    //获取配置
    String getConfig();
    //删除配置
    boolean removeConfig(String dataId, String group);
    //发布
    boolean publishConfig();
    //监听
    void addListener();
    //删除监听器
    void removeListener();
}
```

**`Nacos`客户端获取服务配置**

在加载完所有的`context`上下文之后，客户端就回去拉取这个注册中心里面的这个全部配置文件
```java
@Override
public String getConfig(String dataId, String group, long timeoutMs) throws NacosException {
    return getConfigInner(namespace, dataId, group, timeoutMs);
}
```

`getConfigInner`方法里面，就是具体的拉取配置这个实现
```java
private String getConfigInner(String tenant, String dataId, String group, long timeoutMs) throws NacosException{
    // 先使用本地配置
    String content = LocalConfigInfoProcessor.getFailover(agent.getName(), dataId, group, tenant);
    // 本地配置不为空，则直接返回
    if (content != null) {
        return content;
    }
    // 本地配置为空，去服务端拉取 全部的配置文件
    // 通过这个HTTP请求进行远程调用
    try{
        // 拉取需要的配置
        String[] ct = worker.getServerConfig(dataId, group, tenant, timeoutMs);
        // 保存结果到本地
        cr.setContent(ct[0]);
    }
}
```

通过`getFailover`方法实现读取本地配置
```java
public static String getFailover(String serverName, String dataId, String group, String tenant) {
    // 获取本地文件
    File localPath = getFailoverFile(serverName, dataId, group, tenant);
    // 如果本地文件为空，则直接return返回
    if (!localPath.exists() || !localPath.isFile()) {
        return null;
    }
    // 本地文件不为空，则读取
    return readFile(localPath);
}
```

`getServerConfig`的方法
```java
public String[] getServerConfig(String dataId, String group, String tenant, long readTimeout) throws NacosException {
    HttpRestResult<String> result = null;
    // HTTP请求
    result = agent.httpGet(Constants.CONFIG_CONTROLLER_PATH, null, params, agent.getEncode(), readTimeout);
    switch (result.getCode()) {
        case HttpURLConnection.HTTP_OK: // ...
        case HttpURLConnection.HTTP_NOT_FOUND:
        case HttpURLConnection.HTTP_CONFLICT:
        case HttpURLConnection.HTTP_FORBIDDEN:
        default:
    }
}
```

`Nacos`的服务配置监听
在整个容器启动完成之后，就会去调用这个监听器。
`Nacos`主要在`NacosContextRefresher`类实现监听，`ApplicationListener`接口是`Nacos`的上下文的刷新流。

构造方法如下：
```java
public NacosContextRefresher(NacosRefreshProperties refreshProperties, NacosRefreshHistory refreshHistory, ConfigService configService) {
    //刷新配置文件
    this.refreshProperties = refreshProperties;
    //刷新历史文件
    this.refreshHistory = refreshHistory;
    this.configService = configService;
}
```

类里面会调用一个`onApplicationEvent`的事件方法，里面会进行`Nacos`的监听注册。
```java
@Override
public void onApplicationEvent(ApplicationReadyEvent event) {
    // many Spring context
    if (this.ready.compareAndSet(false, true)) {
        // 监听注册
        this.registerNacosListenersForApplications();
    }
}
```

注册`Nacos`的监听器方法如下：
- 获取`Nacos`的全部的配置文件
- 获取id之后，通过id对服务进行一个监听。
```java
private void registerNacosListenersForApplications() {
    if (refreshProperties.isEnabled()) {
        for (NacosPropertySource nacosPropertySource : NacosPropertySourceRepository.getAll()) {
            // 获取id
            String dataId = nacosPropertySource.getDataId();
            registerNacosListener(nacosPropertySource.getGroup(), dataId);
        }
    }
}
```

监听`Nacos`的主要方法`registerNacosListener`的具体实现如下：
- 当配置发生变化时，监听方法就会发起一个调用，对应的配置进行更新和替换。
- 每一次更新都会有一个历史版本
```java
private void registerNacosListener(final String group, final String dataId) {
    Listener listener = listenerMap.computeIfAbsent(dataId, i -> new Listener() {
        // 当配置发生变化时，监听方法就会发起一个调用
        @Override
        public void receiveConfigInfo(String configInfo) {
            // 记录历史版本
            refreshHistory.add(dataId, md5);
            // 发布监听事件
            applicationContext.publishEvent(new RefreshEvent(this, null, "Refresh Nacos config"));
        }
    });
}
```

最后调用一个`refresh`方法，进行环境的刷新，会将新的参数和原来的参数进行比较，通过发布环境变更事件，对做出改变的值进行更新操作。
```java
public synchronized Set<String> refresh() {
    Set<String> keys = refreshEnvironment();
    this.scope.refreshAll();
    return keys;
}
```
如果感知对应的配置有改变的操作后，会清除当前的配置实例，并将新的实例重新通过这个`bean`工厂重新`getBean`。

#### 客户端总结
- 客户端启动的时候，会优先拉取本地配置
- 如果本地配置不存在，就和服务端建立`HTTP`请求，拉取服务端的全部配置，就是配置中心的全部配置
- 拉取到全部配置之后，会获取每一个配置文件的`dataId`，通过`dataId`对服务端的每一个配置文件进行监听
- 当服务端的配置文件出现更新时，可以通过监听器进行到感知，客户端也会对对应的配置文件进行更新
- 每一次更新的配置都会存储在`Nacos`配置文件里面，作为一个历史文件保留

### `Nacos` 服务端

#### 服务端获取全部配置
是在`ConfigController`类，在服务端`nacos-config`模块。

`getConfig`方法
```java
@GetMapping
@Secured(action = ActionTypes.READ, parser = ConfigResourceParser.class)
public void getConfig(HttpServletRequest request, HttpServletResponse response,
@RequestParam("dataId") String dataId, @RequestParam("group") String group,
@RequestParam(value = "tenant", required = false, defaultValue = StringUtils.EMPTY) String tenant,
@RequestParam(value = "tag", required = false) String tag)
throws IOException, ServletException, NacosException {
    // check tenant
    ParamUtils.checkTenant(tenant);
    tenant = NamespaceUtil.processNamespaceParameter(tenant);
    // check params
    ParamUtils.checkParam(dataId, group, "datumId", "content");
    ParamUtils.checkParam(tag);
    final String clientIp = RequestUtil.getRemoteIp(request);
    // 获取配置信息
    inner.doGetConfig(request, response, dataId, group, tenant, tag, clientIp);
}
```
`doGetConfig`方法从本地文件读取配置，而不是读取数据库的配置。
文件主要存储在这个`Nacos`的`data`的文件目录下
```java
public String doGetConfig(HttpServletRequest request, HttpServletResponse response, String dataId, String group, String tenant, String tag, String clientIp) throws IOException, ServletException{
    File file = null;
    // md5 加密
    md5 = cacheItem.getMd54Beta();
    // 从磁盘获取文件
    file = DiskUtil.targetBetaFile(dataId, group, tenant);
}
```

#### 服务端将配置存储磁盘

在`DumpService`抽象类，有从内存中将全部配置文件存入到磁盘里面。
抽象类有两个实现类，分别是`EmbeddedDumpService`和`ExternalDumpService`。

实现类里面有一个初始化方法，通过`bean`的前置处理器去初始化实例。
通过`dumpOperate`方法来实现具体的配置文件的存储。
```java
@PostConstruct
@Override
protected void init() throws Throwable {
    // 存储配置文件
    dumpOperate(processor, dumpAllProcessor, dumpAllBetaProcessor, dumpAllTagProcessor);
}
```
在`dumpOperate`方法里面，来实现存储，其主要是一些全量加载和一些增量加载。
```java
protected void dumpOperate(){
    TimerContext.start(dumpFileContext);
    try{
        Runnable dumpAll = () -> dumpAllTaskMgr.addTask(DumpAllTask.TASK_ID, new DumpAllTask());
        Runnable dumpAllBeta = () -> dumpAllTaskMgr.addTask(DumpAllBetaTask.TASK_ID, new DumpAllBetaTask());
        Runnable dumpAllTag = () -> dumpAllTaskMgr.addTask(DumpAllTagTask.TASK_ID, new DumpAllTagTask());
    } catch (Throwable e) {
    }
    Runnable clearConfigHistory = () -> {
        LOGGER.warn("clearConfigHistory start");
        if (canExecute()) {
            try {
                Timestamp startTime = getBeforeStamp(TimeUtils.getCurrentTime(), 24 * getRetentionDays());
                // 用于分页，每次获取磁盘里的1000行数据
                int totalCount = persistService.findConfigHistoryCountByTime(startTime);
                if (totalCount > 0) {
                    int pageSize = 1000;
                    int removeTime = (totalCount + pageSize - 1) / pageSize;
                    while (removeTime > 0) {
                        persistService.removeConfigHistory(startTime, pageSize);
                        removeTime--;
                    }
                }
            } catch (Throwable e) {     
            }
        }
    };
    // 加载配置信息
    try {
        // 判断是增量获取还是全量获取，主要是通过时间是否大于6小时
        dumpConfigInfo(dumpAllProcessor);
    } catch (Throwable e) {
    }          
}
```

#### 服务端总结

- 每个配置文件在注册之后，都会存在`Nacos`的数据库里，最后会将数据库的数据存入到磁盘里面，
- 客户端来拉取这个配置信息的时候，就会直接去读这个本地磁盘里面的数据。

# <a id="qb">Nacos和其他注册中心的区别</a>

| 区别项           | Nacos                 | Eureka      | Consul            | CoreDNS    | Zookeeper  |
|---------------|-----------------------|-------------|-------------------|------------|------------|
| 一致性协议         | CP+AP                 | AP          | CP                | —          | CP         |
| 健康检查          | TCP/HTTP/MYSQL/Client | Client Beat | TCP/HTTP/gRPC/Cmd | -          | Keep Alive |
| 负载均衡策略        | 权重/metadata/Selector  | Ribbon      | Fabio             | RoundRobin | —          |
| 雪崩保护          | 有                     | 有           | 无                 | 无          | 无          |
| 自动注销实例        | 支持                    | 支持          | 不支持               | 不支持        | 支持         |
| 访问协议          | HTTP/DNS              | HTTP        | HTTP/DNS          | DNS        | TCP        |
| 监听支持          | 支持                    | 支持          | 支持                | 不支持        | 支持         |
| 多数据中心         | 支持                    | 支持          | 支持                | 不支持        | 不支持        |
| 跨注册中心同步       | 支持                    | 不支持         | 支持                | 不支持        | 不支持        |
| SpringCloud集成 | 支持                    | 支持          | 支持                | 不支持        | 支持         |
| Dubbo集成       | 支持                    | 不支持         | 不支持               | 不支持        | 支持         |
| K8S集成         | 支持                    | 不支持         | 支持                | 支持         | 不支持        |


----
