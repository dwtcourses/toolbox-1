package github

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/liamg/clinch/prompt"
	"github.com/owenrumney/toolbox/internal/action"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type repoListing struct {
	Repos []*github.Repository `json:"repos"`
}

var repoFile string

func init() {
	homeDir, _ := os.UserHomeDir()
	repoFile = path.Join(homeDir, ".toolbox", "repos.json")
}

func GetRepos(filter string) {
	repoList := getRepositories()
	var filteredRepos []*github.Repository
	var filteredRepoChoices []string

	for _, repo := range repoList.Repos {
		if strings.Contains(string(*repo.Name), filter) {
			filteredRepos = append(filteredRepos, repo)
			filteredRepoChoices = append(filteredRepoChoices, string(*repo.Name))
		}
	}
	if len(filteredRepos) == 0 {
		println("No choices returned.")
		return
	}
	index, _, _ := prompt.ChooseFromList("Select repo", filteredRepoChoices)
	if index == -1 {
		fmt.Println("Cancelled")
		return
	}

	result := filteredRepos[index]
	chooseAction(result)
}

func chooseAction(repo *github.Repository) {
	actions := []action.Action{
		{Description: "Open in browser", Action: func() { action.OpenInBrowser(repo.GetHTMLURL()) }},
		{Description: "Try to find locally", Action: func() { println("Open locally") }},
		{Description: "Show outstanding pull requests", Action: func() { println("Show outstanding pull requests") }},
		{Description: "Print repo json", Action: func() { action.PrintJson(repo) }},
	}

	var actionChoices []string
	for _, a := range actions {
		actionChoices = append(actionChoices, a.Description)
	}

	index, _, _ := prompt.ChooseFromList(fmt.Sprintf("What do you want to do with %q", string(*repo.Name)), actionChoices)
	if index == -1 {
		println("Cancelled")
		return
	}
	actions[index].Action()
}

func getRepositories() repoListing {
	file, err := ioutil.ReadFile(repoFile)
	if err != nil {
		panic(err)
	}
	repoList := &repoListing{}
	err = json.Unmarshal(file, repoList)
	if err != nil {
		panic(err)
	}
	return *repoList
}

func IndexRepos() {
	ctx := context.Background()
	client := GetClient(ctx)

	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 50},
	}

	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.List(ctx, "", opt)
		if err != nil {
			panic(err)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	writeRepositoriesToFile(allRepos)
}

func writeRepositoriesToFile(allRepos []*github.Repository) {
	repoList := &repoListing{
		Repos: allRepos,
	}
	file, err := json.MarshalIndent(repoList, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Writing %d repositories to file %q\n", len(allRepos), repoFile)
	err = ioutil.WriteFile(repoFile, file, 0644)
	if err != nil {
		panic(err)
	}
}
