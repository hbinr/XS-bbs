# XS-bbs
Xiang Shou 论坛，XS-bbs 基于gin + gorm + go-redis 实战，面向接口开发。

项目早期使用过 wire 来解决依赖注入，主要是学习完 wire 后进行实践。

后来个人觉得该项目很小，引入 wire 反而增加了初学者理解该项目的难度，就全部改为 Go 原生实现依赖注入了~

另外个人在该项目中不仅使用了gin, 还使用了 GoFrame, 一个开箱即用的框架，不仅仅是web开发能用， 主要是学习完后使用了GoFame的一些模块。

这里是个人使用GoFame validate模块时写的文章：https://goframe.org/pages/viewpage.action?pageId=3673259



## 目录结构
```sh
├── cmd                # 程序入口
│   ├── main.go
│   ├── wire_gen.go    # 已删除
│   └── wire.go        # 已删除
├── docs               # swagger接口文档
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal           # 私有模块，业务代码和业务严重依赖的库
│   ├── app            # app 项目，按功能模块划分，方便后续扩展微服务
│   └── pkg            # 业务严重依赖的公共库
├── pkg                # 公共模块，和业务无关，可以对外使用的库
│   ├── cache          # 缓存初始化封装
│   ├── conf           # 配置定义及初始化封装
│   ├── database       # 数据库初始化封装
│   ├── logger         # 日志库初始化封装
│   ├── servers        # http 路由初识化、注册相关,后续可以支持 grpc server 
│   └── utils          # 一些工具封装
├── config.yaml        # 配置文件
├── go.mod
├── go.sum
├── LICENSE
├── README.md
└── script             # 脚本文件
    └── my_app.sql
```

## 技术选型

| 技术方向       | 框架名称  |
| -------------- | --------- |
| 请求、路由处理 | gin       |
| 参数校验       | gf/gvalid |
| 数据库(mysql)  | gorm      |
| 缓存           | go-redis  |
| 配置读写       | viper     |
| 日志           | zap       |
| API文档        | swagger   |
| 依赖注入       | wire      |

面向接口开发，`repo` 层和 `service` 层都封装了业务接口。

> `wire.go`定义了项目初始化的过程，通过`wire`工具生成具体的初始化，在`wire_gen.go`中。  已去掉 `wire` 的使用

## 项目运行

```sh
cd cmd/

go build

./cmd
```
