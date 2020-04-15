
package main

import (
	"github.com/owenrumney/toolbox/commands"
	"github.com/owenrumney/toolbox/internal"
)

func main() {
	internal.LoadConfig()
	commands.Execute()
}
