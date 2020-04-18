package commands

import (
	"github.com/owenrumney/toolbox/internal/github"
	"github.com/spf13/cobra"
)

var githubCmd = &cobra.Command{
	Use: "github",
	Short: "Get all Github Repos",
	Run: func(cmd *cobra.Command, args []string) {
		filter := getFilter(args)
		github.GetRepos(filter)
	},
}

var repoCmd = &cobra.Command{
	Use: "repos",
	Short: "Get a lits of repos",
	Run: func(cmd *cobra.Command, args []string) {
		filter := getFilter(args)
		github.GetRepos(filter)
	},
}

var indexRepos = &cobra.Command{
	Use: "indexrepos",
	Short: "Index all repos the user is authenticated to see",
	Run: func(cmd *cobra.Command, args []string) {
		github.IndexRepos()
	},
}

func getFilter(args []string) string {
	filter := ""
	if len(args) > 0 {
		filter = args[0]
	}
	return filter
}

func init() {
	rootCmd.AddCommand(githubCmd)
	githubCmd.AddCommand(repoCmd)
	githubCmd.AddCommand(indexRepos)

}