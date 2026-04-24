package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/spcameron/press"
	"github.com/spcameron/test-site/internal/site"
)

func RunBuild(args []string) (int, error) {
	fs := flag.NewFlagSet("build", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	out := fs.String("out", "build/public", "output directory")
	static := fs.String("static", "static", "static assets directory")
	assets := fs.String("assets", "assets", "output assets base path")

	cleanFlag := fs.Bool("clean", false, "clean output directory before building")
	noCleanFlag := fs.Bool("no-clean", false, "do not clean output directory before building")

	if err := fs.Parse(args); err != nil {
		return 2, err
	}

	if fs.NArg() != 0 {
		return 2, fmt.Errorf("unexpected build argument: %s", fs.Arg(0))
	}

	clean := true
	switch {
	case *cleanFlag && *noCleanFlag:
		return 2, fmt.Errorf("cannot specify both --clean and --no-clean")
	case *cleanFlag:
		clean = true
	case *noCleanFlag:
		clean = false
	}

	opts := press.BuildOptions{
		OutDir:         *out,
		Clean:          clean,
		StaticDir:      *static,
		AssetsBasePath: *assets,
		OnWrite: func(p string) {
			fmt.Printf("build: wrote %s\n", p)
		},
	}

	r := site.Renderers()

	err := press.Build(opts, r)
	if err != nil {
		return 1, fmt.Errorf("build: %w", err)
	}

	return 0, nil
}

func RunServe(args []string) (int, error) {
	fs := flag.NewFlagSet("serve", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	dir := fs.String("dir", "build/public", "directory to serve")
	addr := fs.String("addr", "127.0.0.1:8080", "listen address")

	if err := fs.Parse(args); err != nil {
		return 2, err
	}

	opts := press.ServeOptions{
		Dir:  *dir,
		Addr: *addr,
	}

	if err := press.Serve(opts); err != nil {
		return 1, fmt.Errorf("serve: %w", err)
	}

	return 0, nil
}
