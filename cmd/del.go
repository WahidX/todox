/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"todox/todo"

	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a todo item by index",
	Long: `Delete a todo item by index
Example: todox del 1`,
	Run: func(cmd *cobra.Command, args []string) {
		index := args[0]
		if index == "" {
			fmt.Println("Please provide an index")
			return
		}

		intIndex, err := strconv.Atoi(index)
		if err != nil {
			fmt.Println("Invalid index")
			return
		}

		err = todo.Del(intIndex)

		if err != nil {
			fmt.Println("Error deleting todo! Invalid index")
			return
		}
		fmt.Println("Todo deleted successfully")
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
