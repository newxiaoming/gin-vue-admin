## 目录


```
.
├── docker-compose.yaml.example ## docker-compose.yaml.example 示例文件
├── Dockerfile
├── Makefile
├── README.MD
└── src ## 源码文件夹
    ├── app ## 应用和路由
    │   ├── corrention
    │   │   └── corrention.go
    │   └── router ## 路由
    │       ├── api.go
    │       └── initrouter.go ## 基本路由
    ├── bin ## 编译后的文件
    ├── build ## 持续集成、持续部署、其他脚本
    │   ├── gocd-ci
    │   └── shell
    ├── cmd ## 入口
    │   ├── api
    │   │   └── api.go
    │   └── root.go
    ├── config ## 配置文件
    │   ├── config.yaml
    │   └── config.yaml.example
    ├── doc ## 文档
    │   └── v1
    │       └── README.MD
    ├── go.mod
    ├── go.sum
    ├── internal ## 私有库
    │   ├── middleware ## 中间件
    │   │   ├── correntionparams.go
    │   │   └── xfyunparams.go
    │   └── pkg
    ├── LICENSE
    ├── main.go ## main入口文件
    ├── pkg ## 公共库
    │   ├── response
    │   │   ├── response.go
    │   │   └── textcorrention.go
    │   └── xfyun
    │       ├── authorization
    │       │   └── authorization.go
    │       └── textcorrention
    │           └── textcorrention.go
    ├── README.MD
    ├── test
    └── tool ## 工具库
        ├── config.go
        ├── gpool ## 线程池
        │   └── gpool.go
        ├── http.go
        ├── interface.go
        └── string.go
```