package terraform

import (
	"fmt"
	"github.com/liamg/clinch/prompt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type tfeState struct {
	Hostname  string
	Org       string
	Workspace string
}

func GetState() {
	bytes, err := ioutil.ReadFile("state.tf")
	if err != nil {
		fmt.Println("No state.tf file found")
		return
	}
	fmt.Printf("%v", string(bytes))
}

func WriteState() {
	if !checkForTerraformFiles() {
		response := prompt.EnterInputWithDefault("There are no terraform files here, are you sure?", "yes")
		if response != "yes" {
			println("The state.tf file will not be created")
			return
		}
	}
	hostname := prompt.EnterInputWithDefault("TFE Hostname", viper.GetString("tfe_hostname"))
	org := prompt.EnterInput("TFE Organisation: ")
	workspace := prompt.EnterInput("TFE Workspace: ")

	stateContent := &tfeState{
		Hostname:  hostname,
		Org:       org,
		Workspace: workspace,
	}

	tmpl, err := generateTemplate()
	if err != nil {
		panic(err)
	}
	writeStateFile(tmpl, stateContent)
}

func generateTemplate() (*template.Template, error) {
	tmpl, err := template.New("statefile").Parse(`"terraform" {
	  backend "remote" {
	    hostname = "{{.Hostname}}"
	    organization = "{{.Org}}"
	    workspaces {
	      name = "{{.Workspace}}"
	    }
	  }
	}`)
	return tmpl, err
}

func writeStateFile(tmpl *template.Template, stateContent *tfeState) {
	file, err := os.Create("state.tf")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = tmpl.Execute(file, stateContent)
	if err != nil {
		panic(err)
	}
}

func checkForTerraformFiles() bool {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".tf") {
			return true
		}
	}
	return false
}
