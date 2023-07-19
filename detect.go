package detectillegaldeps

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type WhiteList map[string][]string

func (w WhiteList) IsAllowedDeps(from, to string) bool {
	if _, ok := w[from]; !ok {
		return false
	}

	for _, v := range w[from] {
		if v == to {
			return true
		}
	}

	return false
}

type Detector struct {
	basePath       string
	internalPrefix string
	Whitelist      WhiteList
}

func NewDetector(conf *Config) *Detector {
	return &Detector{
		basePath:       conf.basePath,
		internalPrefix: conf.internalPrefix,
		Whitelist:      conf.whiteList,
	}
}

func (d *Detector) Detect() error {
	reg := regexp.MustCompile(`\s*.go`)
	filepath.WalkDir(d.basePath, func(path string, dir os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		targets, err := filepath.Glob(path + "/*.go")
		if err != nil {
			return err
		}

		for _, t := range targets {
			if !reg.MatchString(t) || strings.Contains(t, "test") {
				return nil
			}

			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, t, nil, parser.ImportsOnly)
			if err != nil {
				return err
			}

			for _, i := range f.Imports {
				p := strings.Trim(i.Path.Value, "\"")
				if !d.ApplyRule(f.Name.Name, p) {
					fmt.Printf("illegal dependency found: %s -> %s\n", f.Name.Name, p)
				}
			}
		}
		return nil
	})
	return nil
}

func (d *Detector) ApplyRule(from, to string) bool {
	if !strings.HasPrefix(to, d.internalPrefix) {
		return true
	}

	to = strings.TrimPrefix(to, d.internalPrefix)
	return d.Whitelist.IsAllowedDeps(from, to)
}
