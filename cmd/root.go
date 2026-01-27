/*
Copyright © 2026 Emanuel Zapata
*/
package cmd

import (
	"errors"
	"os"

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
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-flix.yaml)")

	viper.SetConfigName("cli-flix-config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/cli-flix/")
	viper.AddConfigPath("$HOME/.cli-flix")
	_, err := os.Stat("cli-flix-config.env")
	if errors.Is(err, os.ErrNotExist) {
		os.Create("cli-flix-config.env")
		viper.SetDefault("THE_MOVIE_DB_API_KEY", "")
		viper.WriteConfig()
	}
}
