package main

import (
	"log"

	"github.com/7c/mygobase/domain"
	"github.com/fatih/color"
	"github.com/sanity-io/litter"
)

func main() {
	got, err := domain.ParseDomain("foo.blogspot.co.uk")
	if err != nil {
		log.Fatal(err)
	}
	color.Set(color.FgYellow)
	litter.Dump(got)
}
