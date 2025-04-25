
---

# 达梦数据库 MCP 服务

## 项目简介
本项目是一个基于达梦数据库的 MCP（Microservice Communication Protocol）服务，提供了以下功能：
- **列出数据库中的表**：通过 `list_tables` 工具列出指定数据库中的所有表。
- **执行只读 SQL 查询**：通过 `execute_query` 工具执行只读 SQL 查询（仅支持 `SELECT`、`SHOW`、`DESCRIBE` 和 `EXPLAIN` 语句）。
- **显示表结构**：通过 `describe_table` 工具显示指定表的结构。
## 使用方法

windows 下可以直接下载 dm-mcp-server 二进制文件，并添加到 PATH 中。

### 使用 go install 安装
   ```bash
   go install github.com/wangzhaobo168/dm-mcp-server
   ```
### 达梦MCP  Server 配置示例： 
```
{
    "mcpServers": {
        "db-mcp-server": {
        "command": "dm-mcp-server",
        "args": [
        ],
        "env": {
            "DM_PORT": "端口号,默认5236",
            "DM_HOST": "主机地址",
            "DM_USERNAME": "账号",
            "DM_PASSWORD": "密码",
            "DM_SCHEMA": "模式名"
        }
    }
}
```
## 检查 dm-mcp-server 版本：

```bash
dm-mcp-server --version
```

## 快速开始

### 1. 环境要求
- **Go 版本**：1.16 或更高版本。
- **达梦数据库**：已安装并配置好达梦数据库。
- **依赖库**：确保安装了 `github.com/mark3labs/mcp-go` 库。

### 2. 安装依赖
在项目根目录下运行以下命令初始化 Go Modules 并下载依赖：

```bash
go mod init dm-mcp-server
go mod tidy
```


### 3. 配置数据库连接
在运行服务之前，需要配置达梦数据库的连接信息。可以通过以下方式设置：
- **命令行参数**：启动服务时通过命令行参数指定。
- **环境变量**：通过环境变量设置。

#### 命令行参数
```bash
go run main.go \
  -username <用户名> \
  -password <密码> \
  -host <主机地址> \
  -port <端口> \
  -schema <模式名>
```

## 项目结构
```
dm-mcp-server/
├── main.go                # 主程序入口
├── tools/                 # 工具实现
│   ├── list_tables.go     # 列出表工具
│   ├── execute_query.go   # 执行查询工具
│   └── describe_table.go  # 显示表结构工具
├── utils/                 # 工具函数
│   ├── config.go          # 配置管理
│   └── db.go              # 数据库连接
├── go.mod                 # Go Modules 文件
└── README.md              # 项目文档
```


## 贡献指南
欢迎提交 Issue 和 Pull Request 来改进本项目。请确保代码风格一致，并通过测试。

## 许可证
本项目采用 [MIT 许可证](LICENSE)。

---

### 补充说明
- 如果 `github.com/mark3labs/mcp-go` 库不可用，请替换为其他兼容的 MCP 实现。
- 如果需要支持更多数据库操作，可以在 `tools` 包中添加新的工具。

---
