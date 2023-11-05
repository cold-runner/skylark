# skylark开发规范
## 目录结构
```bash
├─api # swagger 文档生成处
├─conf  # skylark 项目配置
│      config.example.yaml  # skylark 配置文件
├─doc   # skylark 项目目录
│  │  architecture.png  # 项目架构图
│  │  development.md    # 开发文档
│  └─api
│          api_specification.md # 后端api接口详情
│          code_specification.md  # skylark 错误码设计规范
│          generic.md # skylark 错误码设计规范
│          README  # 总览
│          user_api.md  # 后端用户模块api接口
│
├─internal
│  ├─controller # controlle层
│  │  ├─mw      # 中间件
│  │  └─v1      # api
│  │          post.go
│  │          comment.go
│  │          user.go
│  │
│  ├─core       # 服务器核心代码
│  │      router.go
│  │      server.go
│  │
│  ├─model      # model 层
│  │  ├─post
│  │  │      post.go
│  │  │
│  │  ├─consts   # 常量
│  │  │      consts.go
│  │  │
│  │  └─user    # 用户 model
│  │          tourists.go
│  │          user.go
│  │
│  ├─pkg        # 项目内共享包
│  │  ├─errCode    # 错误码
│  │  │      base.go
│  │  │      errCode.go
│  │  │      register.go
│  │  │      resp.go
│  │  │      user.go
│  │  │
│  │  ├─config  # 项目内共享的各个组件配置
│  │  │ …………
│  │  │
    …………
│  ├─service    # 服务层
    …………
│  │
│  └─db      # 持久层
│      │  post.go  # 文章持久层接口
│      │  comment.go  # 评论持久层接口
│      │  db.go    # 存储层接口
│      │  user.go     # 用户层接口
│      │
│      └─mysql  # 接口实现层
│              mysql.go
│              user.go
└─pkg         # 项目内外共享包
   ………………
```
项目维护者代码必须按照功能放在相应的目录

## git commit规范
本项目遵从[angular提交规范](https://docs.google.com/document/d/1QrDFcIiPjSLDn3EL15IJygNPiHORgU1_OOAqWjiDU5Y/edit#heading=h.greljkmo14y0)
暂设scope如下
    
| scope     | 模块      |
|-----------|---------|
| core      | 服务器核心配置 |
| user      | 用户模块    |

推荐安装[Commitizen](https://github.com/commitizen/cz-cli)工具结合本项目中的`.cz-config.js`中的提示内容进行代码提交

## 错误处理规范
skylark 项目使用了[marmotedu](https://github.com/marmotedu/errors)的error错误处理包
- 定义新错误

    定义新错误时，要在`internal/pkg/code`目录下的目标文件中通过const常量注册错误，要求符合常量值[服务和模块说明](./api/code_specification.md) 中的规范。
    新注册模块时请先修改doc文档表明新错误模块标识号，然后创建错误模块文件
- 追加错误

    在符合[服务和模块说明](./api/code_specification.md)的前提下，将新错误码追加到响应错误模块文件中
- 错误码格式
```go
const ( 
    // ErrXxx - xxx: xxxxxxx
    ErrXxx = iota + XxXxXx
    ......
    // ErrXxx - xxx: xxxxxxx
    ErrXxx
)
```
`internal/pkg/code`目录下的错误内容大致如上，新注册错误时，要求必须存在如上述示例的**注释格式**！


### 修改错误码相关内容后，一定要使用go generate命令重新生成代码和文档！！！
### 修改错误码相关内容后，一定要使用go generate命令重新生成代码和文档！！！
### 修改错误码相关内容后，一定要使用go generate命令重新生成代码和文档！！！

## 测试相关
开发完新功能后，必须编写单元测试并通过测试才可以提交！
## 日志规范
日志当中只记录必要的调试信息
- 在出现错误时就打印日志
- 在功能入口打印日志
- 在分支处打印日志