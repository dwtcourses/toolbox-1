package commands

import (
	"fmt"
	"github.com/owenrumney/toolbox/internal/terraform"
	"github.com/spf13/cobra"
)

var terraformCmd = &cobra.Command{
	Use:   "terraform",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("terraform called")
	},
}

var stateCmd = &cobra.Command{
	Use: "state",
	Short: "Interact with state for terraform",
	Run: func(cmd *cobra.Command, args []string) {
		terraform.GetState()
	},
}

var createCmd = &cobra.Command{
	Use: "create",
	Short: "Create a state file",
	Run: func(cmd *cobra.Command, args []string) {
		terraform.WriteState()
	},
}

func init() {
	rootCmd.AddCommand(terraformCmd)
	terraformCmd.AddCommand(stateCmd)
	stateCmd.AddCommand(createCmd)
}
