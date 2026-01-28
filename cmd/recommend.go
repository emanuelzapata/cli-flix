/*
Copyright © 2026 Emanuel Zapata
*/
package cmd

import (
	"cli-flix/internal"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var random bool

// recommendCmd represents the recommend command
var recommendCmd = &cobra.Command{
	Use:   "recommend",
	Short: "This command will recommend movies",
	Long:  `Use this command to get recommendations on what to watch!`,

	Run: func(cmd *cobra.Command, args []string) {
		viper.ReadInConfig()
		fmt.Println("recommend called")
		if random {
			fmt.Println("passed in value")
			apiKey := viper.GetString("THE_MOVIE_DB_API_KEY")
			if apiKey == "" {
				fmt.Println("Please set an API key")
			} else {
				internal.GetRandomRecommendation(apiKey)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(recommendCmd)

	recommendCmd.Flags().BoolVarP(&random, "random", "r", false, "Get a random recommendation")
}
