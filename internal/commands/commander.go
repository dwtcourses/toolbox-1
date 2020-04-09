package commands

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Run() error {

	logLevel := logrus.WarnLevel

	for  _, arg := range os.Args {
		if arg == "--debug" {
			logLevel = logrus.DebugLevel
		}
	}

	logrus.SetLevel(logLevel)

	fmt.Print("Executing root command")
	return rootCmd.Execute()
}