# chapter4-工程化实践
## 作业内容
1. 按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。
## 思考分析

### 代码文件与执行方式

本目录下的project.go

```shell
go mod tidy
go mod vendor
go run project.go
```

目录文件结构
```shell
.
├── README.md
├── cmd
├── configs
│   └── test.toml
├── docs
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   └── project-api
│   │       ├── api
│   │       │   ├── handler
│   │       │   │   └── student_handler
│   │       │   │       └── handler.go
│   │       │   ├── repo
│   │       │   │   └── student_repo
│   │       │   │       ├── model.go
│   │       │   │       └── repo.go
│   │       │   └── service
│   │       │       └── student_service
│   │       │           └── service.go
│   │       └── router
│   │           ├── handler.go
│   │           ├── router.go
│   │           └── router_api.go
│   └── configs
│       ├── configs.go
│       └── constants.go
├── logs
│   └── chapter04project-access.log
├── pkg
│   ├── logger
│   │   └── logger.go
│   ├── repo
│   │   └── mysql
│   │       └── mysql.go
│   └── shutdown
└── project.go
```

github仓库中没有空目录

1. configs 存放工程配置文件
2. internal 存放项目业务代码
   1. app 按服务划分
   2. configs 加载配置文件和全局变量 
3. logs 日志目录
4. pkg 存放工程公共组件
5. project.go 工程入口

### 代码分析

业务框架采用的是传统的MVC思想（DDD从网上看了几篇教程，思想基本了解，但实践没实践，不明觉厉）

因为没有使用任何开源框架，基本也没设计VO DTO

正常的话，
1. Handler中应有VO struct：request 和 response
2. Service中应有DTO struct，供Handler根据request创建DTO，service接收并处理

没有使用wire，思想倒是看明白了，但感觉除了使代码看上去更简洁……逻辑细节的正确性还是要人为保证..大概就是有用wire的功夫，我代码就已经写完了..