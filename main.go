package main

import (
	"log"
	"tool/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execure err: %v", err)
	}
}
