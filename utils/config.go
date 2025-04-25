package utils

import (
	"os"
)

var (
	Version     = "1.0.0"
	DefaultPort = "5236"
	userName    string
	password    string
	host        string
	port        string
	schema      string
)

func GetUserName() string {
	if userName != "" {
		return userName
	}
	return os.Getenv("DM_USERNAME")
}

func GetPassWord() string {
	if password != "" {
		return password
	}
	return os.Getenv("DM_PASSWORD")
}

func GetHost() string {
	if host != "" {
		return host
	}
	return os.Getenv("DM_HOST")
}

func GetSchema() string {
	if schema != "" {
		return schema
	}
	return os.Getenv("DM_SCHEMA")
}

func GetPort() string {
	if port != "" {
		return port
	}
	if port := os.Getenv("DM_PORT"); port != "" {
		return port
	}
	return DefaultPort
}

func SetHost(h string) {
	host = h
}

func SetPort(p string) {
	port = p
}

func SetUserName(name string) {
	userName = name
}

func SetPassword(pwd string) {
	password = pwd
}

func SetSchema(s string) {
	schema = s
}
