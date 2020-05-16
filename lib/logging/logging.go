package logging

// Logger is an interface describing logger objects
type Logger interface {
	// Name returns the name of this logger
	Name() string

	// Colored returns true if we color our output, false otherwise.
	Colored() bool

	// Sub returns a new Logger that is sub to this one, but
	// has its name appended with the name provided to Sub
	Sub(name string) Logger

	// SubName returns a new Logger that is sub to this one,
	// but has its name replaced with the name given to SubNamed
	SubNamed(name string) Logger

	// Is takes a Level and returns true if an entry will
	// be emitted at the level specified
	Is(level Severity) bool

	// IsTrace is equal to Is(Trace)
	IsTrace() bool

	// IsDebug is equal to Is(Debug)
	IsDebug() bool

	// IsInfo is equal to Is(Info)
	IsInfo() bool

	// IsWarn is equal to Is(Warn)
	IsWarn() bool

	// IsError is equal to Is(Error)
	IsError() bool

	// Log emits a message
	Log(level Severity, args ...interface{})

	// Log a level regardless of the verbosity `Info` level
	Print(args ...interface{})

	// Log with `Trace` level
	Trace(args ...interface{})

	// Log with `Debug` level
	Debug(args ...interface{})

	// Log with `Info` level
	Info(args ...interface{})

	// Log with `Warn` level
	Warn(args ...interface{})

	// Log with `Error` level
	Error(args ...interface{})
}
