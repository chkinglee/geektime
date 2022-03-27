# chapter9-Go语言实践-网络编程

## 作业内容
1. 总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。
2. 实现一个从 socket connection 中解码出 goim 协议的解码器。
## 思考分析

### 粘包场景

相对简单的场景是，在client连续发包时，server没有及时接收包（或其他原因）导致server接收到的包内容为连续几个包的内容

在终端1中启动server端接收包
```go
go run stick/server.go
```

在终端2中启动client端发送包
```go
go run stick/client.go
```

当client发送完毕后，查看server输出
![img.png](img.png)

从上图可以看出，server接收包时出现粘包

### 解决方案

常见的三类解决方案

- client对原始包封装包头，包头包括一个`标记字符串`和`数据长度`。server解析时先检查是否遇到`标记字符串`，然后获取当前包的`数据长度`，并检查后面的数据段的长度是否满足`数据长度`，当`标记字符串`和`数据长度`均匹配时，则为1个完整包
- client和server约定一个固定的数据长度，client发送该长度的数据，server截取该长度的数据。但这种方式下，数据的准确性得不到保证。
- client和server约定一个分隔符，client发送的数据包尾部添加该分隔符，server按照该分隔符分割数据。但这种方式下，server可能误分。

**LengthField**

约定包头字符串为`ConstHeader = "geektime"`，并用`ConstSaveDataLength = 4`字节存储数据长度

在终端1中启动server端接收包
```go
go run length_field/server.go
```

在终端2中启动client端发送包
```go
go run length_field/client.go
```

![img_1.png](img_1.png)

**FixLength**

约定数据长度`ConstFixLength = 100`，当client发送的数据长度小于该值时，使用`ConstFixLengthFullChar = "@"`对其尾部进行填充。

server在解析时，截取约定长度的数据，并删除尾部的填充字节

在终端1中启动server端接收包
```go
go run fix_length/server.go
```

在终端2中启动client端发送包
```go
go run fix_length/client.go
```

预期输出与上图一致

**Delimiter**

约定分隔符为`ConstDelimiter = '@'`。client发送数据包时在其尾部追加该分隔符。server在解析数据包时读取该分隔符。

在终端1中启动server端接收包
```go
go run delimiter/server.go
```

在终端2中启动client端发送包
```go
go run delimiter/client.go
```

预期输出与上图一致


### 作业2
没研究goim……