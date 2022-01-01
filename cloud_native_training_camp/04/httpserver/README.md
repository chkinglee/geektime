# httpserver
## 模块二作业
### 作业内容
1. 接收客户端request，并将request中带的header写入response header
2. 读取当前系统的环境变量中的VERSION配置，并写入response header
3. Server端记录访问日志包括客户端IP，HTTP返回码，输出到server端的标准输出
4. 当访问localhost/healthz时，应返回200

### 编译方式
```shell
# 默认OS=linux ARCH=amd64
make build

# 针对其他系统架构，可在命令行中自定义
make build OS=darwin ARCH=arm64
```

### 运行方式
```shell
./bin/amd64/httpserver
```

## 模块三作业
### 作业内容
1. 构建本地镜像。
2. 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）。 
3. 将镜像推送至 Docker 官方镜像仓库。 
4. 通过 Docker 命令本地启动 httpserver。 
5. 通过 nsenter 进入容器查看 IP 配置。

### 制作docker镜像
```shell
# 默认TAG为1.0.0 可通过命令行自定义
make release TAG=1.0.0
```

### 推送docker镜像到官方仓库
```shell
make push
```

### 通过docker本地启动httpserver
```shell
docker run --name httpserver -d chkinglee/httpserver:1.0.0 -p 8077:8077
```

### 访问
```shell
curl 127.0.0.1:8077/healthz
```
I am running

