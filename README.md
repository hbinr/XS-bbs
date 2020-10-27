# XS-bbs
Xiang Shou 论坛

## 目录结构
```sh
├── cmd                # 程序入口
│   ├── main.go
│   ├── wire_gen.go
│   └── wire.go
├── docs               # swagger接口文档
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal           # 私有模块，业务代码和业务严重依赖的库
│   ├── app            # app 项目，按功能模块划分，方便后续扩展微服务
│   └── pkg            # 业务严重依赖的库
├── pkg                # 公共模块，和业务无关，可以对外使用的库
│   ├── cache          # 缓存，和业务无关，可以对外使用的库
│   ├── conf           # 配置定义及初始化封装
│   ├── database       # 数据库初始化封装
│   ├── log            # 日志库初始化封装
│   └── tool           # 一些工具封装
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

面向接口开发，dao层和service层都封装了业务接口。

`wire.go`定义了项目初始化的过程，通过`wire`工具生成具体的初始化，在`wire_gen.go`中

## 项目运行

```sh
cd cmd/

go build

./cmd
```