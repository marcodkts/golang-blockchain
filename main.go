package main

import (
	"os"

	"github.com/marcodkts/golang-blockchain/cli"
)

func main() {
	defer os.Exit(0)
	cli := cli.CommandLine{}
	cli.Run()
}
