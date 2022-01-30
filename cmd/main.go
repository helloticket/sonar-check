package main

import (
	"fmt"
	"os"

	"github.com/helloticket/sonar-check/cmd/cmd"
)

var (
	version = "[not provided]"
)

func main() {
	if err := cmd.Execute(version); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
