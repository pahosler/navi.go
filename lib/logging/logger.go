package logging

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/term"
)

const (
	timeFormat = "2006-01-02T15:04:05.000Z0700"
)

type logger struct {
	name  string
	level *int32
	out   *writer
	mu    *sync.Mutex
	color bool
}

// New returns a new logger
func New(opts ...Option) Logger {
	return newLogger(opts...)
}

func newLogger(opts ...Option) Logger {
	o := defaults()
	for _, opt := range opts {
		opt(o)
	}

	var color bool

	if o.forceColor {
		color = true
	} else if o.color {
		if f, ok := o.out.(*os.File); ok {
			if term.IsTerminal(int(f.Fd())) {
				color = true
			}
		}
	}

	l := &logger{
		name:  o.name,
		level: new(int32),
		mu:    new(sync.Mutex),
		out:   &writer{out: o.out},
		color: color,
	}

	atomic.StoreInt32(l.level, int32(o.level))

	return l
}

func (l *logger) log(name string, level Severity, force bool, args ...interface{}) {
	if len(args) < 1 {
		return
	}

	if level >= sevMax {
		level = Error
	} else if level <= sevNone {
		level = Info
	}

	if !force && level < Severity(atomic.LoadInt32(l.level)) {
		return
	}

	now := time.Now()

	l.mu.Lock()
	defer l.mu.Unlock()

	l.out.WriteString(now.Format(timeFormat))

	label := sevLabel[level]
	reset := false

	if l.color {
		color := sevColor[level]
		if level < Warn {
			label = color.Wrap(label)
		} else {
			l.out.WriteString(color.String())
			reset = true
		}
	}

	l.out.WriteString(" [")
	l.out.WriteString(label)
	l.out.WriteString("] {")

	if name != "" {
		l.out.WriteString(name)
		l.out.WriteString("} >>> ")
	}

	prefixLen := l.out.Len()

	if msg, ok := args[0].(string); ok && len(args) > 1 {
		fmt.Fprintf(l.out, msg, args[1:]...)
	} else {
		fmt.Fprint(l.out, args...)
	}

	lastLF := false

	if b, s := l.out.Bytes(), l.out.Len(); (s-prefixLen) <= 0 || b[s-1] != '\n' {
		lastLF = true
	}

	if reset {
		l.out.WriteString(colorReset.String())
	}

	if lastLF {
		l.out.WriteRune('\n')
	}

	l.out.Flush()
}

func (l *logger) Name() string {
	return l.name
}

func (l *logger) Colored() bool {
	return l.color
}

func (l *logger) Sub(name string) Logger {
	nl := *l
	if nl.name != "" {
		nl.name = nl.name + "." + name
	} else {
		nl.name = name
	}

	return &nl
}

//
func (l *logger) SubNamed(name string) Logger {
	nl := *l
	nl.name = name
	return &nl
}

func (l *logger) Is(level Severity) bool {
	return Severity(atomic.LoadInt32(l.level)) <= level
}

func (l *logger) IsTrace() bool {
	return Severity(atomic.LoadInt32(l.level)) == Trace
}

func (l *logger) IsDebug() bool {
	return Severity(atomic.LoadInt32(l.level)) <= Debug
}

func (l *logger) IsInfo() bool {
	return Severity(atomic.LoadInt32(l.level)) <= Info
}

func (l *logger) IsWarn() bool {
	return Severity(atomic.LoadInt32(l.level)) <= Warn
}

func (l *logger) IsError() bool {
	return Severity(atomic.LoadInt32(l.level)) <= Error
}

func (l *logger) Log(level Severity, args ...interface{}) {
	l.log(l.Name(), level, false, args...)
}

func (l *logger) Print(args ...interface{}) {
	l.log(l.Name(), Info, true, args...)
}

func (l *logger) Trace(args ...interface{}) {
	l.log(l.Name(), Trace, false, args...)
}

func (l *logger) Debug(args ...interface{}) {
	l.log(l.Name(), Debug, false, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.log(l.Name(), Info, false, args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.log(l.Name(), Warn, false, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.log(l.Name(), Error, false, args...)
}
