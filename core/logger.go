package core

import (
	"log"
	"fmt"
	"os"
	"strconv"
)

// Used colors for the unix console
const (
	// output colors
	ColorRed		= 1
	ColorMagenta	= 5
	ColorCyan		= 6
	ColorWhite		= 7
	ColorDefault	= 8
	// log levels
	LevelDebug		= 0
	LevelInfo		= 1
	LevelWarning	= 2
	LevelError		= 3
)

// Log wrapper, inherit from Logger
type Logger struct {
	*log.Logger
	verbose	bool
	level	int
	prefix	string
}

// Creates a new Logger
func NewLogger(filename string, verbose bool, level int, prefix string) *Logger {
	// open log file
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	l := log.New(file, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	return &Logger{
		l,
		verbose,
		level,
		prefix,
	}
}

// Main public function to use
func (l *Logger) Log(msg string, level int) {
	if level >= l.level {
		switch level {
		case LevelDebug:
			l.debug(msg)
		case LevelInfo:
			l.info(msg)
		case LevelWarning:
			l.warning(msg)
		case LevelError:
			l.error(msg)
		}
	}
}

// Print to stdout, only if in verbose mode
func (l *Logger) printStdout(msg string, color int) {
	if l.verbose {
		// add color to a string for printing it out in the console
		var colorized = "\x1b[3" + strconv.Itoa(color) + ";1m" + msg + "\x1b[0m"
		fmt.Println(colorized)
	}
}

// Log Debug
func (l *Logger) debug(msg string) {
	msg = "DEBUG (" + l.prefix + ") " + msg
	l.Println(msg)
	l.printStdout(msg, ColorCyan)
}

// Log Info
func (l *Logger) info(msg string) {
	msg = "INFO (" + l.prefix + ") " + msg
	l.Println(msg)
	l.printStdout(msg, ColorWhite)
}

// Log Warning
func (l *Logger) warning(msg string) {
	msg = "WARNING (" + l.prefix + ") " + msg
	l.Println(msg)
	l.printStdout(msg, ColorMagenta)
}

// Log Error
func (l *Logger) error(msg string) {
	msg = "ERROR (" + l.prefix + ") " + msg
	l.Println(msg)
	l.printStdout(msg, ColorRed)
}