# gogen
CLI tool to Generate New Go projects and submodules


## Installing

```bash
git clone https://github.com/MatthewMcDade13/gogen.git
cd gogen

# Linux
./install

# Install script for Windows not yet implemented

```

## Usage

```bash
# New go project
gogen new my_project

# Gen new file module
gogen mod server
```

## Project Structure

project_name/
    src/
        main.go
    go.mod
    Makefile

## FileModule Structure

filemodule_name/
    filemodule.go
    filemodule_test.go
