# Go Community Backend

一个基于Go语言开发的社区后端服务，提供用户认证、文章管理和表单处理等功能。

## 技术栈

- Go 1.x
- Gin Web Framework
- MySQL
- JWT认证
- CORS支持
- Viper配置管理

## 主要功能

- 用户认证
  - 登录
  - 注册
  - JWT token验证

- 表单处理
  - 创建表单
  - 获取表单列表
  - 获取表单详情
  - 更新表单
  - 删除表单

- 用户管理
  - 获取用户列表
  - 获取用户详情

## 快速开始

### 环境要求

- Go 1.x
- MySQL

### 配置

在 `config/config.yml` 中配置应用参数：

```yaml
app:
  name: go_community
  port: 8080

front:
  url: http://localhost:5173

database:
  dsn: root:1234@tcp(127.0.0.1:3306)/go_community?parseTime=true
  MaxIdleConns: 10
  MaxOpenConns: 100
```

### 安装依赖

```bash
go mod download
```

### 运行

```bash
go run main.go
```

服务将在配置的端口上启动（默认8080）。

## API接口

### 认证接口

- `POST /api/auth/login` - 用户登录
- `POST /api/auth/register` - 用户注册

### 需要认证的接口

以下接口需要在请求头中携带JWT token：

- 文章相关
  - `POST /api/articles` - 创建文章
  - `GET /api/articles` - 获取文章列表
  - `GET /api/articles/:id` - 获取文章详情

- 表单相关
  - `POST /api/forms` - 创建表单
  - `GET /api/forms` - 获取表单列表
  - `GET /api/forms/:id` - 获取表单详情
  - `PUT /api/forms/:id` - 更新表单
  - `DELETE /api/forms/:id` - 删除表单

- 用户相关
  - `GET /api/users` - 获取用户列表
  - `GET /api/users/:id` - 获取用户详情

## 项目结构

```
.
├── config/          # 配置文件和初始化
├── controllers/     # 控制器
├── global/         # 全局变量
├── middlewares/    # 中间件
├── models/         # 数据模型
├── router/         # 路由配置
├── utils/          # 工具函数
├── go.mod          # Go模块文件
└── main.go         # 程序入口
```

## 安全性

- 使用JWT进行用户认证
- CORS配置限制允许的源
- 密码加密存储

## 贡献

欢迎提交Issue和Pull Request！

## 许可证

[MIT License](LICENSE)