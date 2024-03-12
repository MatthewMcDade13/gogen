package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/MatthewMcDade13/gogen/src/config"
	"github.com/MatthewMcDade13/gogen/src/gen"
	"github.com/charmbracelet/log"
)

var gomod_prefix string

// TODO :: Implement this prefix flag. Right now this app just ignores it.
func init() {
	flag.StringVar(&gomod_prefix, "p", "", "gogen new -p github.com/githubuser project_name")
}

func main() {

	version := flag.Bool("v", false, "Prints version of gogen")

	flag.Parse()

	if *version {
		msg := fmt.Sprintf("gogen %v\n", config.PROJECT_VERSION)
		fmt.Println(msg)
		os.Exit(0)
	}

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
