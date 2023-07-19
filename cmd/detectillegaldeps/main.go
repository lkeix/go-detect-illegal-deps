package main

import (
	"flag"
	detectillegaldeps "github.com/lkeix/go-detect-illegal-deps"
	"log"
	"os"
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
	errs := d.Detect()
	for _, err := range errs {
		log.Println(err)
	}
	if len(errs) != 0 {
		os.Exit(1)
	}
}
