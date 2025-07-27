// Package cmd
/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"rs10.com/rs10-commands/k8s"
)

// k8sCmd represents the k8s command
var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "RS10-K8s相关运维命令",
	Long:  `k8s相关可视化运维命令.`,
	Run: func(cmd *cobra.Command, args []string) {
		k8s.TeaRun()
	},
}

func init() {
	rootCmd.AddCommand(k8sCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// k8sCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// k8sCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
