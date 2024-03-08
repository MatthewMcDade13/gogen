package gogen

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"slices"
	"strings"

	"github.com/MatthewMcDade13/gogen/src/config"
	"github.com/MatthewMcDade13/gogen/src/util"
)

const (
	ROOT_INIT = iota
	ROOT_NEW
)

const GOMAIN_TEMPL string = (`package main

import "fmt"

func main() {
	fmt.Println("Ayyyyeeeee lmao")
}

`)

func package_string(modname string) string {
	return fmt.Sprintln("package", modname) + "\n"
}

func gomod_templ(modname string) string {
	templ := bytes.Buffer{}

	templ.WriteString(package_string(modname))

	funcdecl := fmt.Sprintln("func", util.ToTitleCase(modname), "() {\n\n}")
	templ.WriteString(funcdecl)

	return templ.String()
}

func gomod_templ_test(modname string) string {

	templ := bytes.Buffer{}

	templ.WriteString(package_string(modname))

	templ.WriteString("import \"testing\"\n\n")

	funcdecl := fmt.Sprintln("func", "Test"+util.ToTitleCase(modname), "(t *testing.T) {\n    t.Fatal(\"test not yet implemented!\")\n}")
	templ.WriteString(funcdecl)

	return templ.String()
}

func makefile_templ() string {
	templ := bytes.Buffer{}

	templ.WriteString(".PHONY : run\n")
	templ.WriteString("run:\n")
	templ.WriteString("\tgo run ./src/\n\n")

	templ.WriteString(".PHONY : build\n")
	templ.WriteString("build:\n")
	templ.WriteString("\tmkdir ./bin && go build -o ./bin/gogen ./src/*.go\n\n")

	return templ.String()
}

// NO MUTATE PLS!
var valid_genargs = []string{"new", "n", "mod", "m", "init"}

func IsValidTypeArg(ty string) bool {
	return slices.Contains(valid_genargs, ty)
}

func IsNewArg(ty string) bool {
	return ty == "new" || ty == "n"
}

func IsModArg(ty string) bool {
	return ty == "mod" || ty == "m"
}

func ValidArgsString() string {
	return strings.Join(valid_genargs, "|")
}

func Write(ty string, name string) error {
	if IsNewArg(ty) {
		return gen_project(name, ROOT_NEW)
	} else if IsModArg(ty) {
		return gen_module(name)
	} else if ty == "init" {
		gen_project(name, ROOT_INIT)
	}

	return fmt.Errorf("Invaid argument: %v", ty)
}

func gen_project(name string, root_type int) error {
	_, err := os.Stat(strings.ToLower(name))
	if errors.Is(err, fs.ErrExist) {
		return fmt.Errorf("folder %v already exists in current working directory: %v", name, cwd())
	}

	if fs_exists(name) {
		return fmt.Errorf("Folder: %v, already exists in current working directory: %v", name, cwd())
	}

	var root_path string
	var proj_path string
	switch root_type {
	case ROOT_NEW:
		root_path = path.Join(cwd(), name)
		proj_path = path.Join(root_path, "src")
		break
	case ROOT_INIT:
		root_path = cwd()
		proj_path = path.Join(root_path, "src")
	}

	if err := os.MkdirAll(proj_path, util.DEFAULT_FS_PERM); err != nil {

		return err
	}

	main_path := path.Join(proj_path, "main.go")

	f, err := os.Create(main_path)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(GOMAIN_TEMPL)

	prefix := config.GoModPrefix()
	gomod_name := path.Join(prefix, name)

	// run go mod init {prefix}/{name}
	cmd := exec.Command("go", "mod", "init", gomod_name)
	cmd.Dir = path.Join(cwd(), name)
	if err := cmd.Run(); err != nil {
		return err
	}

	makefile_path := path.Join(root_path, "Makefile")
	makefile_bytes := []byte(makefile_templ())

	if err := os.WriteFile(makefile_path, makefile_bytes, util.DEFAULT_FILE_PERM); err != nil {
		return err
	}

	return nil
}

func gen_module(name string) error {
	var gen_path string
	if fs_exists("go.mod") {
		gen_path = path.Join(cwd(), "src")
	} else {
		gen_path = cwd()
	}

	fullpath := path.Join(gen_path, name)

	if err := os.Mkdir(fullpath, util.DEFAULT_FS_PERM); err != nil {

		return err
	}

	gofile_path := path.Join(fullpath, name+".go")
	gofile_path_test := path.Join(fullpath, name+"_test.go")

	gofile, err := os.Create(gofile_path)
	if err != nil {
		return err
	}
	defer gofile.Close()
	gofile.WriteString(gomod_templ(name))

	gofile_test, err := os.Create(gofile_path_test)
	if err != nil {
		return err
	}
	defer gofile_test.Close()
	gofile_test.WriteString(gomod_templ_test(name))

	return nil

}

func cwd() string {
	if cwd, err := os.Getwd(); err == nil {
		return cwd
	} else {
		return fmt.Sprint("CWD ERROR => ", err)
	}
}

func fs_exists(name string) bool {

	fullpath := path.Join(cwd(), name)

	if _, err := os.Stat(fullpath); err == nil {
		return true
	}
	return false
}
