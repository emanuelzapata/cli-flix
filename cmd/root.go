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

var title string
var year int
var apikey string

// var cfgFile string

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
		// fmt.Println(viper.Get("THE_MOVIE_DB_API_KEY"))
		// if apikey != "" {
		// 	fmt.Println(apikey)
		// } else if title != "" && year != 0 {
		// 	fmt.Println("Movie was given with a year")
		// 	fmt.Println(title)
		// 	fmt.Println(year)
		// } else if title != "" {
		// 	fmt.Println("Movie title was given only")
		// } else if title == "" && year != 0 {
		// 	fmt.Println("Only year was given, please give us a title")
		// }
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-flix.yaml)")
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
		// Handle errors
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
		if viper.Get("THE_MOVIE_DB_API_KEY") == "" {
			fmt.Println("IMDB Api key is empty")
		}
	}

	// viper.SetConfigName("cli-flix-config")
	// // configFile := viper.New()
	// // configFile.SetDefault("user.name", "emanuelzapata")
	// // configFile.SetConfigFile("test-config-data.yaml")
	// // configFile.SafeWriteConfig()
	// // viper.SetConfigName("test-config")
	// // viper.SetConfigType("yaml")
	// viper.AddConfigPath(".")
	// viper.AddConfigPath("/etc/cli-flix/")
	// viper.AddConfigPath("$HOME/.cli-flix")
	// viper.SetDefault("THE_MOVIE_DB_API_KEY", "")
	// viper.WriteConfig()

	// err := viper.ReadInConfig()
	// fmt.Println(viper.Get("THE_MOVIE_DB_API_KEY"))
	// // Handle errors
	// if err != nil {
	// 	panic(fmt.Errorf("fatal error config file: %w", err))
	// }
	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	// TODO: Delete
	// rootCmd.Flags().StringVarP(&title, "title", "t", "", "Movie or TV Show Title")
	// rootCmd.Flags().IntVarP(&year, "year", "y", 0, "Year of the TV Show or Movie")
	// rootCmd.Flags().StringVarP(&apikey, "login", "", "", "Login Using Users API Key")
	// rootCmd.Flags().BoolP("help", "h", false, "Get help")
}
