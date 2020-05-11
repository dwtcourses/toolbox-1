package commands

import (
	"fmt"
	"github.com/owenrumney/toolbox/commands/terraform"
	"github.com/spf13/cobra"
)

var terraformCmd = &cobra.Command{
	Use:   "terraform",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("terraform called")
	},
}

func init() {
	rootCmd.AddCommand(terraformCmd)
	terraformCmd.AddCommand(terraform.NewStateCommand())
}
