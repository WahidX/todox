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

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

		err = todo.Done(intIndex)

		if err != nil {
			fmt.Println("Error updating todo! Invalid index")
			return
		}
		fmt.Println("Todo marked done successfully")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
