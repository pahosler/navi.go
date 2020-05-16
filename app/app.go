package app

import (
	"flag"
	"fmt"

	"github.com/noriah/navi.go"
	"github.com/noriah/navi.go/lib/logging"
)

type flags struct {
	Config       string
	CreateConfig bool
	Version      bool
}

// Go runs the bot
func Go() error {
	opts := &flags{}

	doFlags(opts)

	if opts.Version {
		fmt.Printf("navi (%s) [%s]\nVersion: %s\nBuild Date: %s\nCommit: %s\n",
			navi.Bin(), navi.Arch(),
			navi.Version(), navi.Date(), navi.Commit())
		return nil
	}

	log := logging.New(
		logging.Name("navi"),
		logging.Level(logging.Trace),
	)

	log.Debug("bootstrap")

	return nil
}

func doFlags(opts *flags) {

	flag.StringVar(&opts.Config, "c", "sfxcloud.json", "path to config")
	flag.BoolVar(&opts.CreateConfig, "C", false, "create config if not exist")
	flag.BoolVar(&opts.Version, "V", false, "show version and exit")

	flag.Parse()
}
