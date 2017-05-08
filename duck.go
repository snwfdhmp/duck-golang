package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s <action> [opts]", os.Args[0])
    fmt.Println("")
    fmt.Println("Available options :")
    fmt.Println("")
    fmt.Println("name             usage                      description")
    fmt.Println("---------------  -------------------------  ---------------------------------")
    fmt.Println("init             duck init <name> [opt]     initialize a new duck project")
    fmt.Println("deploy           duck deploy [opt]          deploy duck architecture")
    fmt.Println("config           duck config <action> [opt] project configuration tools")
    fmt.Println("class            duck class <action> [opt]  tools for classes")
    fmt.Println("compile          duck compile [env]         run project compiler")
    fmt.Println("run              duck run [target-version]  run project (no arg->last version)")
    fmt.Println("project-version  duck pv [show|set|inc]     configuration for project version")
    fmt.Println("tar              duck tar                   backup 'src/' dir into a tarball")
    fmt.Println("quick-commit     duck qc [custom-msg]       alias git add *, commit, push")
    fmt.Println("doc              duck doc [command]         shows command's help message")
    fmt.Println("help             duck help                  shows this message")
    os.Exit(1)
}

func main() {
	for _, arg := range os.Args {
		switch(arg) {
		case "help":
			usage()
			os.Exit(0)
		}
	}
}