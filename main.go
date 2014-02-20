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
	clean = flag.Bool("c", true, "Clean returns the shortest path equivalent")
)

type pathFunc func(string) string

func applyToArgs(fn pathFunc) {
	for _, arg := range flag.Args() {
		fmt.Println(fn(arg))
	}
}

func applyToStdin(fn pathFunc) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(fn(scanner.Text()))
	}
}

func apply(fn pathFunc) {
	if flag.NArg() == 0 {
		applyToStdin(fn)
	} else {
		applyToArgs(fn)
	}
}

func main() {
	flag.Parse()
	switch {
	case *join:
		fmt.Println(filepath.Join(flag.Args()...))
	case *base:
		apply(filepath.Base)
	case *dir:
		apply(filepath.Dir)
	case *ext:
		apply(filepath.Ext)
	case *abs:
		apply(func(arg string) string {
			if s, err := filepath.Abs(arg); err != nil {
				return "" //TODO: handle this?
			} else {
				return s
			}
		})
	case *clean:
		apply(filepath.Clean)
	}
}
