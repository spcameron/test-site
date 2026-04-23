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
