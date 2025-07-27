// Package cmd
/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"

	myTea "rs10.com/rs10-commands/bubbleTeaTest"
)

// bubbleTeaTestCmd represents the bubbleTeaTest command
var bubbleTeaTestCmd = &cobra.Command{
	Use:   "bubbleTeaTest",
	Short: "珍珠奶茶测试命令",
	Long:  `珍珠奶茶-TUI命令测试`,
	Run: func(cmd *cobra.Command, args []string) {
		myTea.TeaTest()
	},
}

func init() {
	rootCmd.AddCommand(bubbleTeaTestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bubbleTeaTestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bubbleTeaTestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
