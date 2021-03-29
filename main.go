package main

import (
	"folio/cmd/cli"
)

func main() {
	err := cli.RootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
