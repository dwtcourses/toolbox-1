package terraform

import (
	"github.com/owenrumney/toolbox/internal/terraform"
	"github.com/spf13/cobra"
)

func NewStateCommand() *cobra.Command {
	stateCommand := &cobra.Command{
		Use:   "state",
		Short: "Interact with state for terraform",
		Run: func(cmd *cobra.Command, args []string) {
			terraform.GetState()
		},
	}
	stateCommand.AddCommand(createCmd)
	return stateCommand
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a state file",
	Run: func(cmd *cobra.Command, args []string) {
		terraform.WriteState()
	},
}
