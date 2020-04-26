package ex4_13

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestSearchByTitle(t *testing.T)  {

	title := "Holmes"
	movie, err := getMovie(title)
	if err != nil {
		log.Fatal(err)
	}

	if zero := new(Movie); movie == *zero {
		fmt.Fprintf(os.Stderr, "No results for '%s'\n", title)
		os.Exit(2)
	}

	err = movie.writePoster()
	if err != nil {
		log.Fatal(err)
	}
}
