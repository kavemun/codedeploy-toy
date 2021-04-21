package main

import (
	"flag"

	"github.com/kavemun/toy-robot/cmd"
)

// Map
// (0,4)             (4,4)
//
//
//
// (0,0)             (4,0)

func main() {
	filename := flag.String("file", "", "filename")
	flag.Parse()
	cmd.Play(*filename)
}
