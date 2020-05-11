package github

import (
	"github.com/owenrumney/toolbox/commands/support"
	"github.com/owenrumney/toolbox/internal/github"
	"github.com/spf13/cobra"
)

func NewReposCommand() *cobra.Command {
	reposCommand := &cobra.Command{
		Use:   "repos",
		Short: "Get a lits of repos",
		Run: func(cmd *cobra.Command, args []string) {
			filter := support.GetFilter(args)
			github.GetRepos(filter)
		},
	}
	reposCommand.AddCommand(indexCmd)
	return reposCommand
}

var indexCmd = &cobra.Command{
	Use:   "indexrepos",
	Short: "Index all repos the user is authenticated to see",
	Run: func(cmd *cobra.Command, args []string) {
		github.IndexRepos()
	},
}
