/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

import (
	"github.com/sunnydonaldson/thoth/common"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <question>",
	Short: "Add a question to the bank.",
	Long: "Add a question to the bank, with optional tags",
	Example: "thoth add \"How do you find the area of a circle?\" -t maths,geometry",
	Args: cobra.ExactArgs(1),
	Run:  run,
	
}

func run(cmd *cobra.Command, args []string) {
	tags, err := cmd.Flags().GetStringSlice("tags")
	if err != nil {
		panic(err)
	}
	question := common.Question{Text: args[0], Tags: tags}
	cmd.Println(question)
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringSliceP("tags", "t", []string{}, "provide a list of tags")
}
