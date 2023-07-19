package main

import (
	"flag"
	detectillegaldeps "github.com/lkeix/go-detect-illegal-deps"
	"log"
)

func main() {
	yaml := flag.String("config-yaml", "go-illegal-deps.yaml", "")
	basePath := flag.String("base-path", ".", "")
	flag.Parse()

	c, err := detectillegaldeps.NewConfig(*yaml, *basePath)
	if err != nil {
		log.Fatal(err)
	}
	d := detectillegaldeps.NewDetector(c)
	d.Detect()
}
