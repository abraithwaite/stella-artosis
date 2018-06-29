package main

import (
	"flag"
	"strings"

	"github.com/abraithwaite/stella-artosis"
)

var (
	ignore  = flag.String("ignore", "/sys,/dev,/proc", "comma separated list of files/directories to ignore")
	include = flag.String("include", ".", "comma separated list of files/directories to include")
	hash    = flag.String("hash", "sha1", "which hash to use for analysis")
	follow  = flag.Bool("s", false, "if set, follow symlinks")
)

func main() {
	flag.Parse()
	c := artosis.Config{
		Ignored:        parseFiles(*ignore),
		Include:        parseFiles(*include),
		Hash:           *hash,
		FollowSymlinks: *follow,
	}
	artosis.Scan(c)
}

func parseFiles(files string) map[string]bool {
	x := strings.Split(files, ",")
	ret := make(map[string]bool, len(x))
	for _, i := range x {
		ret[i] = true
	}
	return ret
}
