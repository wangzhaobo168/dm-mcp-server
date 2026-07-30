// utils/db.go
package utils

import (
	"database/sql"
	"fmt"
	"net"
	"net/url"

	_ "github.com/wangzhaobo168/dm" // 引入达梦数据库驱动
)

// ConnectDMDatabase 连接达梦数据库
func ConnectDMDatabase() (*sql.DB, error) {
	username := GetUserName()
	password := GetPassWord()
	host := GetHost()
	port := GetPort()
	dataSourceName := buildDataSourceName(username, password, host, port)
	db, err := sql.Open("dm", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("打开达梦数据库失败: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("无法连接达梦数据库: %w", err)
	}
	return db, nil
}

func buildDataSourceName(username, password, host, port string) string {
	return (&url.URL{
		Scheme: "dm",
		User:   url.UserPassword(username, password),
		Host:   net.JoinHostPort(host, port),
	}).String()
}
