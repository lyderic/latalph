package cmd

import (
	"os"

	. "github.com/lyderic/tools"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Version:               VERSION,
	Use:                   "latalph",
	Short:                 "latalph v." + VERSION,
	DisableFlagsInUseLine: true,
	SilenceErrors:         true,
	SilenceUsage:          true,
	CompletionOptions:     cobra.CompletionOptions{HiddenDefaultCmd: true},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return process()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		Redln(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	config := os.Getenv("HOME") + "/.latalph.yaml"
	rootCmd.PersistentFlags().StringP("config", "", config, "config file")
	rootCmd.PersistentFlags().BoolP("debug", "", false, "show debugging information")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}

func initConfig() {
	if viper.IsSet("config") {
		viper.SetConfigFile(viper.GetString("config"))
	} else {
		viper.AddConfigPath(os.Getenv("HOME"))
		viper.SetConfigName(".latalph")
	}
	//viper.SetDefault("foo", "bar")
	viper.SetEnvPrefix("latalph")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		Debug("Using config file: %s\n", viper.ConfigFileUsed())
	}
	Debug("%v\n", viper.AllSettings())
}
