package cli

import (
	"os"
	"fmt"
	"flag"
)

type CLIConfig struct {
	Command  string
	Value    string
	MaxPages int
}

func ParseCLI() (*CLIConfig, error) {
	if len(os.Args) < 3 {
		return nil, fmt.Errorf(
			"usage: gc-agent <command> <value> [--max-pages N]",
		)
	}

	cfg := &CLIConfig{
		Command: os.Args[1],
		Value:   os.Args[2],
	}

	fs := flag.NewFlagSet(cfg.Command, flag.ContinueOnError)

	fs.IntVar(
		&cfg.MaxPages,
		"max-pages",
		5,
		"maximum number of pages to parse",
	)

	if err := fs.Parse(os.Args[3:]); err != nil {
		return nil, err
	}

	return cfg, nil
}
