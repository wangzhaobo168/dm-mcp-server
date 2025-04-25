package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wangzhaobo168/dm-mcp-server/tools"
	"github.com/wangzhaobo168/dm-mcp-server/utils"

	"github.com/mark3labs/mcp-go/server"
	_ "github.com/wangzhaobo168/dm"
)

var (
	Version = utils.Version
)

// 添加工具
func addTools(s *server.MCPServer) {
	// 添加 list_tables 工具
	s.AddTool(tools.ListTablesTool, tools.ListTablesToolHandelFunc)
	// 添加 execute_query 工具
	s.AddTool(tools.ExecuteQueryTool, tools.ExecuteQueryToolHandelFunc)
	// 添加 describe_table 工具
	s.AddTool(tools.DescribeTableTool, tools.DescribeTableToolHandelFunc)
}

// 创建一个新的MCP服务器
func newMCPServer() *server.MCPServer {
	return server.NewMCPServer(
		"达梦数据库MCP Server",
		Version,
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)
}

func run() error {
	s := newMCPServer()
	addTools(s)
	// 启动服务器
	if err := server.ServeStdio(s); err != nil {
		return fmt.Errorf("服务器错误: %v", err)
	}
	return nil
}
func main() {
	username := flag.String("username", "", "用户名")
	password := flag.String("password", "", "密码")
	host := flag.String("host", "", "主机")
	port := flag.String("port", "", "端口")
	schema := flag.String("schema", "", "模式名")
	showVersion := flag.Bool("version", false, "显示版本号")
	flag.Parse()
	if *username != "" {
		utils.SetUserName(*username)
	} else {
		*username = utils.GetUserName()
	}

	if *password != "" {
		utils.SetPassword(*password)
	} else {
		*password = utils.GetPassWord()
	}
	if *host != "" {
		utils.SetHost(*host)
	} else {
		*host = utils.GetHost()
	}
	if *port != "" {
		utils.SetPort(*port)
	} else {
		*port = utils.GetPort()
	}
	if *schema != "" {
		utils.SetSchema(*schema)
	} else {
		*schema = utils.GetSchema()
	}
	if *showVersion {
		fmt.Printf("Dm MCP Server\n")
		fmt.Printf("Version: %s\n", Version)
		os.Exit(0)
	}
	if err := run(); err != nil {
		panic(err)
	}
}
