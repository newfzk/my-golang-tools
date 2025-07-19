// Package envInfo
package envInfo

import "rs10.com/rs10-commands/constant"

// 数据库配置信息

type DatabaseConfig struct {
	Type     constant.DatabaseType `yaml:"type"`
	Driver   string                `yaml:"driver"`
	Host     string                `yaml:"host"`
	Port     string                `yaml:"port"`
	Username string                `yaml:"username"`
	Password string                `yaml:"password"`
	DbName   string                `yaml:"dbname"`
}
