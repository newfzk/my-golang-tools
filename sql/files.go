// Package sql
package sql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"rs10.com/rs10-commands/constant"
	"rs10.com/rs10-commands/envInfo"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// 批量执行sql文件

// Run 批量执行sql文件
// 入参1 数据库配置信息 databaseConfig
// 入参2 要执行的sql文件的文件名 //TODO 后续改成目录名 进行批量执行
func Run(databaseConfig *envInfo.DatabaseConfig, sqlDir string) {
	if databaseConfig.Type == constant.MySQL {
		runSqlFiles4Mysql(databaseConfig, sqlDir)
	}
	return
}

// 批量执行sql文件到 Mysql 数据库
func runSqlFiles4Mysql(databaseConfig *envInfo.DatabaseConfig, sqlDir string) {
	host := databaseConfig.Host
	port := databaseConfig.Port
	dbName := databaseConfig.DbName
	username := databaseConfig.Username
	password := databaseConfig.Password

	// 构建连接字符串
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

	// 连接数据库
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	defer db.Close()

	// 测试连接
	err = db.Ping()
	if err != nil {
		log.Fatalf("数据库连接测试失败: %v", err)
	}
	log.Println("成功连接数据库")

	// 读取SQL文件
	sqlFile := sqlDir
	content, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Fatalf("读取SQL文件失败: %v", err)
	}

	// 分割SQL语句 TODO 这样分割不行 sql语句中可能有 分号
	queries := strings.Split(string(content), ";")

	// 执行每条SQL语句
	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}

		_, err := db.Exec(query)
		if err != nil {
			log.Printf("执行SQL失败: %v\n语句: %s", err, query)
			continue
		}
		fmt.Printf("执行成功: %s\n", getQueryType(query))
	}
	fmt.Println("SQL文件执行完成")
}

// 辅助函数：获取SQL类型
func getQueryType(query string) string {
	query = strings.ToUpper(strings.TrimSpace(query))
	switch {
	case strings.HasPrefix(query, "DROP"):
		return "DROP操作"
	case strings.HasPrefix(query, "CREATE"):
		return "CREATE操作"
	case strings.HasPrefix(query, "INSERT"):
		return "INSERT操作"
	case strings.HasPrefix(query, "DELETE"):
		return "DELETE操作"
	default:
		return "其他操作"
	}
}
