package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/wangzhaobo168/dm-mcp-server/utils"
)

const (
	ExecuteQueryToolName = "execute_query"
)

var ExecuteQueryTool = mcp.NewTool(ExecuteQueryToolName,
	mcp.WithDescription("执行只读SQL查询"),
	mcp.WithString(
		"query",
		mcp.Description("SQL查询（仅允许SELECT、SHOW、DESCRIBE和EXPLAIN语句）"),
		mcp.Required(),
	),
	mcp.WithString(
		"database",
		mcp.Description("数据库名称（如果未指定则使用配置的DM_SCHEMA）"),
	),
)

func ExecuteQueryToolHandelFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	query := request.Params.Arguments["query"].(string)
	dbName, ok := request.Params.Arguments["schema"].(string)
	if !ok {
		dbName = utils.GetSchema()
	}
	// 验证查询语句是否为只读操作
	queryUpper := strings.ToUpper(strings.TrimSpace(query))
	if !strings.HasPrefix(queryUpper, "SELECT") &&
		!strings.HasPrefix(queryUpper, "SHOW") &&
		!strings.HasPrefix(queryUpper, "DESCRIBE") &&
		!strings.HasPrefix(queryUpper, "EXPLAIN") {
		return mcp.NewToolResultError("仅允许SELECT、SHOW、DESCRIBE和EXPLAIN语句"), nil
	}

	db, err := utils.ConnectDMDatabase()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("连接达梦数据库失败: %v", err)), nil
	}
	defer db.Close()

	// 设置当前schema
	if _, err := db.Exec(fmt.Sprintf("SET SCHEMA %s", dbName)); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("设置schema失败: %v", err)), nil
	}

	rows, err := db.Query(query)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("执行查询失败: %v", err)), nil
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("获取列失败: %v", err)), nil
	}

	// 构建表头
	var results []string
	results = append(results, strings.Join(columns, "\t"))

	// 获取数据行
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range columns {
			valuePtrs[i] = &values[i]
		}
		if err := rows.Scan(valuePtrs...); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("扫描行失败: %v", err)), nil
		}
		var rowValues []string
		for _, v := range values {
			rowValues = append(rowValues, fmt.Sprintf("%v", v))
		}
		results = append(results, strings.Join(rowValues, "\t"))
	}

	return mcp.NewToolResultText(strings.Join(results, "\n")), nil
}
