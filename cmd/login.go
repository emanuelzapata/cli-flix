/*
Copyright © 2026 Emanuel Zapata
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var apiKey string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Configure API Keys",
	Long: `Use this command to set your API key for the backend service you would like to configure 
	We currently support API Keys from TMDB. To get a key you will need to visit https://developer.themoviedb.org/docs/getting-started 
	Once there can set your API Key here for use`,

	Run: func(cmd *cobra.Command, args []string) {
		if apiKey != "" {
			key, _ := cmd.LocalFlags().GetString("apikey")
			viper.Set("THE_MOVIE_DB_API_KEY", key)
			viper.WriteConfig()
			fmt.Println("API key set")
			// Validate API Key later with simple request
		} else {
			fmt.Println("Please enter an API Key")
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&apiKey, "apikey", "a", "", "TheMovieDB API Key")
}
