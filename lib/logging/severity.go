package logging

// Severity is a level of severity for the log entry
type Severity int32

const (
	// sevNone no log level (use default)
	sevNone = Severity(iota)
	// Trace trace log level
	Trace
	// Debug debug log level
	Debug
	// Info info log level
	Info
	// Warn warn log level
	Warn
	// Error error log level
	Error
	// sevMax max log level
	sevMax
)

var (
	sevLabel = []string{
		sevNone: "UNKWN",
		Trace:   "TRACE",
		Debug:   "DEBUG",
		Info:    "INFO ",
		Warn:    "WARN ",
		Error:   "ERROR",
		sevMax:  "UNKWN",
	}

	sevColor = []color{
		sevNone: colorMagenta,
		Trace:   colorBlue,
		Debug:   colorGreen,
		Info:    colorCyan,
		Warn:    colorYellow,
		Error:   colorRed,
		sevMax:  colorMagenta,
	}
)

func (s Severity) String() string {
	switch s {
	case Trace:
		return "trace"
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	default:
		return "?????"
	}
}
