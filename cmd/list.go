/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"todox/todo"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		todos := todo.List()
		done := []*todo.Todo{}
		undone := []*todo.Todo{}

		for _, t := range todos {
			if t.Done {
				done = append(done, t)
			} else {
				undone = append(undone, t)
			}
		}

		all, _ := cmd.Flags().GetBool("all")

		fmt.Println("Available todos:")
		for _, t := range undone {
			fmt.Printf("\t%d. %s\n", t.ID, t.Content)
		}

		if all {
			fmt.Println("Completed:")
			for _, t := range done {
				fmt.Printf("\t%d. \033[9m %s\033[0m \n", t.ID, t.Content)
			}
		}

		fmt.Println("")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().BoolP("all", "a", false, "All todos will be listed")
}
