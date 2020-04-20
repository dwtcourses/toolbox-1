package action

import (
	"encoding/json"
	"fmt"
)

func PrintJson(jsonObj interface{}) {
	content, err := json.MarshalIndent(jsonObj, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}