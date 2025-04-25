package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/wangzhaobo168/dm-mcp-server/utils"
)

const (
	ListTablesToolName = "list_tables"
)

var ListTablesTool = mcp.NewTool(ListTablesToolName,
	mcp.WithDescription("列出指定数据库中的所有表"),
	mcp.WithString(
		"database",
		mcp.Description("数据库名称（可选）"),
	),
)

func ListTablesToolHandelFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	dbName, ok := request.Params.Arguments["schema"].(string)
	if !ok {
		dbName = utils.GetSchema()
	}
	db, err := utils.ConnectDMDatabase()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("连接达梦数据库失败: %v", err)), nil
	}
	defer db.Close()

	// 使用达梦数据库系统表查询指定schema下的所有表
	query := fmt.Sprintf("SELECT TABLE_NAME FROM ALL_TABLES WHERE OWNER = '%s'", dbName)
	rows, err := db.Query(query)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("获取表失败: %v", err)), nil
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("扫描表名称失败: %v", err)), nil
		}
		tables = append(tables, tableName)
	}

	return mcp.NewToolResultText(strings.Join(tables, ", ")), nil
}
