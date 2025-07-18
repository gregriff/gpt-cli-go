/*
Copyright © 2025 Greg Griffin <greg.griffin2@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gpt-cli-go",
	Short: "A minimal LLM chat interface",
	Long: `gpt-cli-go is a terminal-based chat interface to the LLM-provider API's (Anthropic, OpenAI).
It aims to provide a minimal feature-set with a polished UX, and supports Markdown rendering of responses.

Keybinds:
- Quit : ctrl+d
- Clear History/Quit : ctrl+c
- Toggle Focus : esc
- Text Input Controls : ctrl+a,u,k,e,n,p,b,f,h,m,t,w,d
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// TODO: use funcs from 'os' to make it work on Windows
		viper.SetConfigName("gpt-cli-go")
		viper.SetConfigType("toml")
		viper.AddConfigPath(home + "/.config/gpt-cli-go/")
		viper.AddConfigPath(".")
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				fmt.Println("No config file found")
			} else {
				fmt.Println("Error reading config file: ", err)
				os.Exit(1)
			}
		}
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $XDG_CONFIG_HOME/gpt-cli-go/gpt-cli-go.toml)")
}
