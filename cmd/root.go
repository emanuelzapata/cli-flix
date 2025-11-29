/*
Copyright © 2025 Emanuel Zapata
*/
package cmd

import (
	"os"

	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli-flix",
	Short: "🎬  Your personal IMDB in the terminal!",
	Long: `
 _____  _     _____     ______ _     _______   __
/  __ \| |   |_   _|    |  ___| |   |_   _\ \ / /
| /  \/| |     | |______| |_  | |     | |  \ V / 
| |    | |     | |______|  _| | |     | |  /   \ 
| \__/\| |_____| |_     | |   | |_____| |_/ /^\ \
 \____/\_____/\___/     \_|   \_____/\___/\/   \/
                                                 
                                                 	
───────────────────────────────────────────────────
🍿   CLI-Flix — Movie & TV Info Without Leaving Code
───────────────────────────────────────────────────

Ever been deep in a coding flow and thought:
“Man, I could use a movie break...”

CLI-Flix has you covered.

✨  Search movies and shows by name
🎥  Instantly fetch details, ratings, and release info
🧠  Decide if it's worth your precious downtime

───────────────────────────────────────────────────
Because even developers deserve a good movie break.
🍿
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	viper.SetConfigName("cli-flix-config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/cli-flix/")
	viper.AddConfigPath("$HOME/.cli-flix")
	if _, err := os.Stat("cli-flix-config.env"); os.IsNotExist(err) {
		fmt.Println("Creating file")
		file, fileError := os.Create("cli-flix-config.env")
		if fileError != nil {
			log.Fatal(fileError)
		}
		file.Close()
		viper.SetDefault("THE_MOVIE_DB_API_KEY", "")
		viper.WriteConfig()
	} else {
		fmt.Println("File exists")
		err := viper.ReadInConfig()
		fmt.Println(viper.Get("THE_MOVIE_DB_API_KEY"))
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
		if viper.Get("THE_MOVIE_DB_API_KEY") == "" {
			fmt.Println("IMDB Api key is empty")
		}
	}
}
