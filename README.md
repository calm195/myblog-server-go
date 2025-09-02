# My Blog Server (Go)

## 技术栈

| 技术栈            | 版本     | 作用      |
|----------------|--------|---------|
| Go             | 1.24.5 | 后端开发    |
| gin            | 1.10.1 | Web框架   |
| viper          | 1.12.1 | 配置文件读取  |
| zap            | 1.27.0 | 日志记录    |
| gorm           | 1.30.1 | ORM框架   |
| go-redisClient | 9.12.1 | Redis操作 |
| crypto         | 0.41.0 | 加密解密    |

## 项目结构

| 文件夹          | 说明            | 描述                                                |
|--------------|---------------|---------------------------------------------------|
| `api`        | api层          | api层                                              |
| `config`     | 配置包           | config.yaml对应的配置结构体                               |
| `configs`    | 配置包           | config*.yaml所在位置                                  |
| `core`       | 核心文件          | 核心组件(zap, viper, server)的初始化                      |
| `global`     | 全局对象          | 全局对象                                              |
| `initialize` | 初始化           | router,redisClient,gorm,validator, timer的初始化      |
| `middleware` | 中间件层          | 用于存放 `gin` 中间件代码                                  |
| `model`      | 模型层           | 模型对应数据表                                           |
| `packfile`   | 静态文件打包        | 静态文件打包                                            |
| `resource`   | 静态资源文件夹       | 负责存放静态文件                                          |
| `router`     | 路由层           | 路由层                                               |
| `service`    | service层      | 存放业务逻辑问题                                          |
| `source`     | source层       | 存放初始化数据的函数                                        |
| `util`       | 工具包           | 工具函数封装                                            |

