package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/MatthewMcDade13/gogen/src/gogen"
	"github.com/charmbracelet/log"
)

var gomod_prefix string

func init() {
	flag.StringVar(&gomod_prefix, "p", "", "gogen new -p github.com/githubuser project_name")
}

func main() {
	flag.Parse()

	if s, err := parse_state(); err == nil {
		gogen.Write(s.ty, s.name)
	} else {
		log.Fatal(err)
	}

}

type state struct {
	ty   string
	name string
}

func parse_state() (*state, error) {
	args := os.Args[1:]

	if len(args) == 0 {
		return nil, fmt.Errorf("No arguments passed to gogen. Example usage: `gogen mod client`")
	} else if len(args) < 2 {
		return nil, fmt.Errorf("Not enough arguments passed to gogen. Ex: `gogen new my_go_project`")
	}

	ty := strings.ToLower(args[0])

	if !gogen.IsValidTypeArg(ty) {
		return nil, fmt.Errorf("invalid gen-type argument: %v. Expected: %v", ty, gogen.ValidArgsString())
	}

	name := strings.ToLower(args[1])

	return &state{ty, name}, nil

}
