package commands

import (
	"github.com/owenrumney/toolbox/commands/aws"
	"github.com/spf13/cobra"
)

var awsCommand = &cobra.Command{
	Use:   "aws",
	Short: "AWS Commands",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(awsCommand)
	eventCommand := createEventCommand()
	awsCommand.AddCommand(eventCommand)
}

func createEventCommand() *cobra.Command {
	eventCommand := aws.NewEventCommand()
	eventCommand.Flags().Bool("enabled", false, "Show enabled")
	eventCommand.Flags().Bool("disabled", false, "Show enabled")
	return eventCommand
}
