/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"todox/todo"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo item",
	Long:  "Usage: todox add <content...>",
	Run: func(cmd *cobra.Command, args []string) {
		content := ""
		for _, arg := range args {
			content += arg + " "
		}
		todo.Add(content)

		fmt.Println("Todo added successfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
