package logging

import (
	"io"
	"os"
)

type options struct {
	name       string
	level      Severity
	out        io.Writer
	color      bool
	forceColor bool
}

func defaults() *options {
	return &options{
		level: Info,
		out:   os.Stderr,
		color: true,
	}
}

// Option modifes the logger
type Option func(*options)

// Name sets the name of the logger. Default is ""
func Name(name string) Option {
	return func(opts *options) {
		opts.name = name
	}
}

// WithColor sets color enabled or not. Default is to enable color if
// the output is a terminal
func WithColor(color bool) Option {
	return func(opts *options) {
		opts.color = color
	}
}

// ForceColor forces the logger to use or not use color.
func ForceColor(color bool) Option {
	return func(opts *options) {
		opts.forceColor = color
	}
}

// Output sets the output. Default is stderr
func Output(out io.Writer) Option {
	return func(opts *options) {
		opts.out = out
	}
}

// Level sets the level of the logger. Default is Info
func Level(level Severity) Option {
	if level >= sevMax || level <= sevNone {
		panic("invalid level supplied")
	}

	return func(opts *options) {
		opts.level = level
	}
}
