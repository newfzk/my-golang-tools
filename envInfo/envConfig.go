// Package envInfo
package envInfo

// 环境配置信息

type EnvConfig struct {
	EnvName  string         `yaml:"envName"`
	Database DatabaseConfig `yaml:"database"`
}
