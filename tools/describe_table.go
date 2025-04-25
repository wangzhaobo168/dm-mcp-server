package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/wangzhaobo168/dm-mcp-server/utils"
)

const (
	DescribeTableToolName = "describe_table"
)

var DescribeTableTool = mcp.NewTool(DescribeTableToolName,
	mcp.WithDescription("显示表的结构"),
	mcp.WithString(
		"table",
		mcp.Description("表名称"),
		mcp.Required(),
	),
	mcp.WithString(
		"database",
		mcp.Description("数据库名称（如果未指定则使用DM_SCHEMA）"),
	),
)

func DescribeTableToolHandelFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	dbName, ok := request.Params.Arguments["schema"].(string)
	if !ok {
		dbName = utils.GetSchema()
	}
	tableName, ok := request.Params.Arguments["table"].(string)
	if !ok {
		return mcp.NewToolResultError("表名称不能为空"), nil
	}
	db, err := utils.ConnectDMDatabase()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("连接达梦数据库失败: %v", err)), nil
	}
	defer db.Close()

	// 使用达梦数据库系统表查询表结构
	query := fmt.Sprintf(`
		SELECT COLUMN_NAME, DATA_TYPE, DATA_LENGTH, NULLABLE
		FROM ALL_TAB_COLUMNS 
		WHERE OWNER = '%s' AND TABLE_NAME = '%s'
		ORDER BY COLUMN_ID`, dbName, tableName)

	rows, err := db.Query(query)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("获取表结构失败: %v", err)), nil
	}
	defer rows.Close()

	var schema []string
	for rows.Next() {
		var columnName, dataType, nullable string
		var dataLength int
		if err := rows.Scan(&columnName, &dataType, &dataLength, &nullable); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("扫描列信息失败: %v", err)), nil
		}
		columnDesc := fmt.Sprintf("%s %s(%d) %s",
			columnName,
			dataType,
			dataLength,
			nullable)
		schema = append(schema, columnDesc)
	}

	return mcp.NewToolResultText(strings.Join(schema, "\n")), nil
}
