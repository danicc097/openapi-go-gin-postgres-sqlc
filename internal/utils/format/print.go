package format

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(obj interface{}) {
	bytes, _ := json.MarshalIndent(obj, "  ", "  ")
	fmt.Println(string(bytes))
}
