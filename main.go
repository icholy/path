package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

var (
	base  = flag.Bool("b", false, "Base returns the last element of path")
	dir   = flag.Bool("d", false, "Dir returns all but the last element of path")
	ext   = flag.Bool("x", false, "Ext returns the file name extension used by path")
	join  = flag.Bool("j", false, "Join joins any number of path elements into a single path")
	abs   = flag.Bool("a", false, "Abs returns an absolute representation of path")
	clean = flag.Bool("c", false, "Clean returns the shortest path equivalent")
)

func ApplyToArgs(fn func(string) string) {
	for _, arg := range flag.Args() {
		if s := fn(arg); s != "" {
			fmt.Println(s)
		}
	}
}

func main() {
	flag.Parse()
	if *join {
		fmt.Println(filepath.Join(flag.Args()...))
		return
	}
	switch {
	case *base:
		ApplyToArgs(filepath.Base)
	case *dir:
		ApplyToArgs(filepath.Dir)
	case *ext:
		ApplyToArgs(filepath.Ext)
	case *abs:
		ApplyToArgs(func(s string) string {
			if p, err := filepath.Abs(s); err != nil {
				return "" //TODO: handle this?
			} else {
				return p
			}
		})
	case *clean:
		ApplyToArgs(filepath.Clean)
	}
}
