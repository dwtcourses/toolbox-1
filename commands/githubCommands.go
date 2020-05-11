package commands

import (
	"github.com/owenrumney/toolbox/commands/github"
	"github.com/spf13/cobra"
)

var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "Get all Github Repos",
}

func init() {
	rootCmd.AddCommand(githubCmd)
	githubCmd.AddCommand(github.NewReposCommand())
}
