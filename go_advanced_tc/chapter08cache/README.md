# chapter8-分布式缓存&分布式事务

## 作业内容
1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。 
2. 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

## 思考分析

### 代码文件与执行方式

**环境部署**

在本机部署了redis-server，使用默认的配置文件，127.0.0.1:6379，无密码
```shell
$ redis-server -v
Redis server v=6.2.6 sha=00000000:0 malloc=libc bits=64 build=c6f3693d1aced7d9
```
执行`sh homework.sh`即可完成作业中要求的压测与数据统计

### 值得注意的点

**1 - 压测过程**

根据同一个value长度，压测`set`和`get`命令，本组完成后`flushall`清理数据

为了保证key数量尽可能接近指定的`key_num`，在redis-benchmark的`-n`参数里多加了一个0，即`10`倍于`key_num`的请求量

**2 - 数据统计**

因为redis-server本身会占用一部分内存，所以在压测前需要先收集这部分内存大小，在压测后统计时对`used_memory`减去这部分内存

