package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spcameron/test-site/internal/commands"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	switch os.Args[1] {
	case "build":
		written, code, err := commands.RunBuild(os.Args[2:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(code)
		}

		for _, v := range written {
			fmt.Printf("build: wrote %s\n", v)
		}

	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n\n", os.Args[1])
		usage()
		os.Exit(2)
	}
}

func usage() {
	msg := strings.Join([]string{
		"Usage:",
		"\t<app> build [--out <dir>]",
		"\t<app> serve [--dir <dir>] [--addr <host:port>]",
		"",
		"Commands:",
		"\tbuild    Generate static site output",
		// "\tserve    Serve a directory over HTTP for local preview",
	}, "\n")

	fmt.Fprint(os.Stderr, msg)
}

// func OLDmain() {
// 	log.SetFlags(0)
//
// 	if len(os.Args) < 2 {
// 		usage()
// 		os.Exit(2)
// 	}
//
// 	switch os.Args[1] {
// 	case "build":
// 		written, code, err := commands.RunBuild(os.Args[2:])
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "%v\n", err)
// 			os.Exit(code)
// 		}
//
// 		for _, v := range written {
// 			fmt.Printf("build: wrote %s\n", v)
// 		}
// 	case "serve":
// 		code, err := commands.RunServe(os.Args[2:])
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "%v\n", err)
// 			os.Exit(code)
// 		}
// 	case "-h", "--help", "help":
// 		usage()
// 	default:
// 		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", os.Args[1])
// 		usage()
// 		os.Exit(2)
// 	}
// }
