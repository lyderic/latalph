package cmd

import (
	"fmt"
	"os"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func Debug(format string, args ...interface{}) {
	if viper.GetBool("debug") {
		message := fmt.Sprintf(format, args...)
		fmt.Fprintf(os.Stderr, color.CyanString(message))
	}
}
