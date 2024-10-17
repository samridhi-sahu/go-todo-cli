/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"

	controllers "github.com/samridhi-sahu/go-cli/controller"
)

// var (
// 	deleteTaskName string
// 	updateTaskName string
// 	getTaskName    string
// 	addTaskName    string
// )

// deleteTaskCmd represents the getTask command
var deleteTaskCmd = &cobra.Command{
	Use:   "delete",
	Short: "This flag will delete the desired task",
	Long:  `This delete flag will delete and return the deleted task by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		controllers.DeleteTask(args)
	},
}

// updateTaskCmd represents the getTask command
var updateTaskCmd = &cobra.Command{
	Use:   "update",
	Short: "This flag will update the desired task",
	Long:  `This update flag will update and return the desired task by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		controllers.UpdateTasks(args)
	},
}

// getTaskCmd represents the getTask command
var getTaskCmd = &cobra.Command{
	Use:   "get",
	Short: "This flag will get the desired task",
	Long:  `This get flag will return the desired task by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		controllers.GetTasks(args)
	},
}

var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "This flag will add the desired task",
	Long:  `This add flag will add the desired task in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		controllers.AddTask(args)
	},
}

func init() {
	rootCmd.AddCommand(getTaskCmd)
	rootCmd.AddCommand(addTaskCmd)
	rootCmd.AddCommand(updateTaskCmd)
	rootCmd.AddCommand(deleteTaskCmd)

	// deleteTaskCmd.Flags().StringVarP(&deleteTaskName, "name", "d", "", "Name of the task to delete")
	// updateTaskCmd.Flags().StringVarP(&updateTaskName, "name", "u", "", "Name of the task to update")
	// getTaskCmd.Flags().StringVarP(&getTaskName, "name", "g", "", "Name of the task to get")
	// addTaskCmd.Flags().StringVarP(&addTaskName, "name", "n", "", "Name of the task to add")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
