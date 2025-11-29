/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// var random bool

// recommendCmd represents the recommend command
var recommendCmd = &cobra.Command{
	Use:   "recommend",
	Short: "Get a movie recommendation",
	Long:  `Use this command to get a movie recommendation for you to watch`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("recommend called")
		// if random {
		// 	fmt.Println("Random was called")
		// }
		random, err := cmd.Flags().GetBool("random")
		if err != nil {
			fmt.Println("We got an error")
			return
		}
		if random {
			fmt.Println("Random was called vato")
		}
	},
}

func init() {
	rootCmd.AddCommand(recommendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recommendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	recommendCmd.Flags().BoolP("random", "r", false, "Get a random (good) movie")
}
