package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/epy0n0ff/ssh-config-gen"
)

var toml = flag.String("toml", "", "hosts.toml file path")
var tpl = flag.String("tpl", "", "ssh configration template file path")

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
}
func main() {
	flag.Parse()

	tml, err := sshconfig.LoadToml(*toml)
	if err != nil {
		fmt.Errorf("error:%s", err)
		os.Exit(1)
	}
	config, err := sshconfig.GetSshConfig(*tpl, tml)
	if err != nil {
		fmt.Errorf("error:%s", err)
		os.Exit(1)
	}

	fmt.Print(config)
}
