package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

type pathFunc func(string) string

func ApplyToArgs(fn pathFunc) {
	for _, arg := range flag.Args() {
		if s := fn(arg); s != "" {
			fmt.Println(s)
		}
	}
}

func ApplyToStdin(fn pathFunc) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if s := fn(scanner.Text()); s != "" {
			fmt.Println(s)
		}
	}
}

func Apply(fn pathFunc) {
	if flag.NArg() == 0 {
		ApplyToStdin(fn)
	} else {
		ApplyToArgs(fn)
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
		Apply(filepath.Base)
	case *dir:
		Apply(filepath.Dir)
	case *ext:
		Apply(filepath.Ext)
	case *abs:
		Apply(func(s string) string {
			if p, err := filepath.Abs(s); err != nil {
				return "" //TODO: handle this?
			} else {
				return p
			}
		})
	case *clean:
		Apply(filepath.Clean)
	}
}
