package tools

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(i interface{}) {
	fmt.Println(PrettyFormat(i))
}

func PrettyFormat(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
