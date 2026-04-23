package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/spcameron/press"
	"github.com/spcameron/test-site/internal/site"
)

func RunBuild(args []string) ([]string, int, error) {
	fs := flag.NewFlagSet("build", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	out := fs.String("out", "build/public", "output directory")
	if err := fs.Parse(args); err != nil {
		return nil, 2, err
	}

	opts := press.BuildOptions{
		OutDir: *out,
	}

	r := site.Renderers()

	written, err := press.Build(opts, r)
	if err != nil {
		return nil, 1, fmt.Errorf("build: %w", err)
	}

	return written, 0, nil
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
