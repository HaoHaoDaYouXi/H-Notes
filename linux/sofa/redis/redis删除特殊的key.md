# redis删除特殊的key
### 

```
对于意外产生的redis的key中包含的特殊字符的key,某些可视化工具会显示为类似\xE9\xA2\x98这种的字符
一般可视化软件能看的到详细信息的key是可以删除的,但是由于redis本身的del命令不支持批量删除,所以需要借助
redis-cli来完成,然后由于key中包含特殊符号的特殊性,直接使用keys获取到也无法删除掉,
需要借助下面的lua脚本去完成,命令如下,普通的cli命令加上执行的lua脚本
```

```
./redis-cli -p 6379 -a 123456 -n 0 --eval ./de.lua '*test*'
./redis-cli -p 端口名 -a 密码 -n 库 --eval ./de.lua 后面跟匹配key的规则
```

需要的lua脚本
```
local key=KEYS[1]
local list=redis.call("keys", key);
for i,v in ipairs(list) do
    redis.call("del", v);
end

```
