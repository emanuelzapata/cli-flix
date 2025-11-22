/*
Copyright © 2025 Emanuel Zapata
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
	Short: "Login command to set your API key",
	Long: `Use this login command to set your local API key for OMDB which can be gotten from https://www.omdbapi.com/ 
	// Currently this version of cli-flix only supports OMDB but other services will be implemented soon!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
		fmt.Println(apiKey)
		viper.Set("OMDB_API_KEY", apiKey)
		viper.WriteConfig()
		// viper.SetConfigName("cli-flix-conf")
		// err := viper.ReadInConfig()
		// if err != nil {
		// 	panic(fmt.Errorf("fatal error config file: %w", err))
		// }
		// viper.SetDefault("API_KEY", "")
		// viper.WriteConfig()
		// viper.
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loginCmd.Flags().StringVarP(&apiKey, "apikey", "a", "", "OMDB API Key")
}
