package main

import (
	"fmt"
	"log"

	"github.com/AlekSi/vscode-delve-example/lib"
)

func main() {
	v1, err := lib.NewVersion("1.1.0")
	if err != nil {
		log.Fatal(err)
	}
	v2, err := lib.NewVersion("1.2.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(v1.Less(v2))
}
