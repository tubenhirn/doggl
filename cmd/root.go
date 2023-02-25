package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	apitoken    string
	userLicense string
	AppVersion  string
	rootCmd     = &cobra.Command{
		Use:   "doggl",
		Short: "Doogl - A simple toggl cli.",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.doggl)")

	rootCmd.PersistentFlags().StringVarP(&apitoken, "api_token", "t", "", "A toggl api token.")
	viper.BindPFlag("api_token", rootCmd.PersistentFlags().Lookup("api_token"))

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".doggl".
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".doggl")
	}

	// look for DOG_ prefixed env vars
	viper.SetEnvPrefix("DOG")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}
