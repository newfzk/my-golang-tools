/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"rs10.com/rs10-commands/constant"
	"rs10.com/rs10-commands/envInfo"
	"rs10.com/rs10-commands/sql"
)

// runSqlFilesCmd represents the runSqlFiles command
var runSqlFilesCmd = &cobra.Command{
	Use:   "runSqlFiles",
	Short: "rs10-sql脚本批量执行",
	Long: `rs10-sql脚本批量执行
			rs10-commands sql runSqlFiles <sql文件根目录>
			`,
	Run: func(cmd *cobra.Command, args []string) {
		sqlDir := args[0]
		envName := viper.GetString("envName")
		databaseConfig := envInfo.DatabaseConfig{}
		databaseConfig.Host = viper.GetString("database.host")
		databaseConfig.Port = viper.GetString("database.port")
		databaseConfig.Type = constant.DatabaseType(viper.GetString("database.type"))
		databaseConfig.Username = viper.GetString("database.username")
		databaseConfig.Password = viper.GetString("database.password")
		databaseConfig.DbName = viper.GetString("database.dbname")

		log.Printf("对%s环境批量执行%s路径下的所有sql文件", envName, sqlDir)
		log.Printf("数据库地址:%s 数据库类型:%s 数据库名:%s 数据库用户名:%s",
			databaseConfig.Host, databaseConfig.Type, databaseConfig.DbName, databaseConfig.Username)

		sql.Run(&databaseConfig, sqlDir)
	},
}

func init() {
	sqlCmd.AddCommand(runSqlFilesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runSqlFilesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runSqlFilesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
