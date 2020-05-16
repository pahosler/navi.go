package logging

import (
	"fmt"
)

type color byte

const (
	colorReset = color(0)
	colorBlack = color(iota + 29)
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
	colorCyan
	colorWhite
)

const (
	colorWrapFmt   = "\x1b[%dm%s\x1b[0m"
	colorStringFmt = "\x1b[%dm"
)

func (c color) Wrap(s string) string {
	return fmt.Sprintf(colorWrapFmt, byte(c), s)
}

func (c color) String() string {
	return fmt.Sprintf(colorStringFmt, byte(c))
}
