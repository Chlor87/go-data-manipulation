package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var a []int = nil
	fmt.Println(len(a))
}

func logJson(v any) {
	j, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(j))
}
