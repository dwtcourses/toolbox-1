package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(helloCmd)
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Say hello",
	Long:  "Say hello as a way of testing everything is working okay",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Print("Hello there")
		} else {
			for _, arg := range args {
				fmt.Printf("Hello there, %v", arg)
			}
		}
	},
}
