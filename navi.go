package navi

import (
	"fmt"
	"os"
	"path/filepath"
	rt "runtime"
	"time"
)

var (
	start   time.Time = time.Now()
	runtime string    = rt.Version()
	arch    string    = fmt.Sprintf("%s %s", rt.GOOS, rt.GOARCH)

	version string = "dev"
	date    string = "dev"
	commit  string = "dev"
)

// App is exposed so we dont have to expose the whole instance
type App interface {
	Go() error
	Name() string
}

// Details are data that can be sent with json
type Details struct {
	Version string `json:"version"`
	Date    string `json:"date"`
	Commit  string `json:"commit"`
	Runtime string `json:"-"`
	Arch    string `json:"-"`
}

// RunDuration returns how long since program start
func RunDuration() time.Duration {
	return time.Since(start)
}

// StartupTime returns the time that we booted
func StartupTime() time.Time {
	return start
}

// Runtime returns the go runtime
func Runtime() string {
	return runtime
}

// Arch returns the arch built on
func Arch() string {
	return arch
}

// Version get the version of this
func Version() string {
	return version
}

// Date get the build date of this
func Date() string {
	return date
}

// Commit get the commit hash of this build
func Commit() string {
	return commit
}

// Bin returns the name of the executing binary.
// Only for information.
// DO NOT USE FOR LOGIC!
func Bin() string {
	return filepath.Base(os.Args[0])
}

// Data returns a struct of the current build info
func Data() Details {
	return Details{
		Version: version,
		Date:    date,
		Commit:  commit,
		Runtime: runtime,
		Arch:    arch,
	}
}
