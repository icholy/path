# Path
---

> Simple tool for dealing with paths

```
Usage of path:
  -a=false: Abs returns an absolute representation of path
  -b=false: Base returns the last element of path
  -c=false: Clean returns the shortest path equivalent
  -d=false: Dir returns all but the last element of path
  -j=false: Join joins any number of path elements into a single path
  -x=false: Ext returns the file name extension used by path
```

Examples:

``` sh
$ ls | xargs path -a
/home/icholy/code/go/src/github.com/icholy/path/main.go
/home/icholy/code/go/src/github.com/icholy/path/path
/home/icholy/code/go/src/github.com/icholy/path/README.md
```

``` sh
$ path -x `ls`
```
