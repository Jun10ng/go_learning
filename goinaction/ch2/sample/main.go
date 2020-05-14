package main

import (
	_ "go_learning/goinaction/ch1/sample/matchers"
	"go_learning/goinaction/ch1/sample/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
