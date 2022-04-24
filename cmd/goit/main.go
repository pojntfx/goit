package main

import "github.com/pojntfx/goit/cmd/goit/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
