package aws

import (
	"github.com/owenrumney/toolbox/commands/support"
	"github.com/owenrumney/toolbox/internal/aws"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewEventCommand() *cobra.Command {
	eventsCommand := &cobra.Command{
		Use:   "events",
		Short: "Interact with AWS events",
		Run: func(cmd *cobra.Command, args [] string) {
			filter := support.GetFilter(args)
			flags := getEventFlags(cmd.Flags())
			aws.NewCloudWatchEventClient().GetEvents(filter, *flags)
		},
	}
	eventsCommand.AddCommand(disableCommand)
	eventsCommand.AddCommand(enableCommand)
	return eventsCommand
}

var disableCommand = &cobra.Command{
	Use: "disable",
	Short: "Disable rules",
	Run: func (cmd *cobra.Command, args []string) {
		filter := support.GetFilter(args)
		aws.NewCloudWatchEventClient().DisableEvents(filter)
	},
}

var enableCommand = &cobra.Command{
	Use: "enable",
	Short: "Enable rules",
	Run: func (cmd *cobra.Command, args []string) {
		filter := support.GetFilter(args)
		aws.NewCloudWatchEventClient().EnableEvents(filter)
	},
}

func getEventFlags(flags *pflag.FlagSet) *aws.EventFlags {
	enabled, err := flags.GetBool("enabled")
	if err != nil {
		panic(err)
	}
	disabled, err := flags.GetBool("disabled")
	if err != nil {
		panic(err)
	}
	return &aws.EventFlags{
		ShowEnabled:  enabled,
		ShowDisabled: disabled,
	}
}
