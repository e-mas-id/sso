// Let's borrow logger package
// https://github.com/rl404/go-malscraper/blob/master/pkg/mallogger/mallogger.go
package utils

import (
	"fmt"
	"strings"
	"time"
)

const (
	// TimeFormat is timestamp format for utils.
	TimeFormat = "2006/01/02 15:04:05.000"
	// SuccessIcon is success utils indicator.
	SuccessIcon = "[S]"
	// SuccessType is key to write success utils.
	SuccessType = "success"
	// TraceIcon is trace utils indicator.
	TraceIcon = "[T]"
	// TraceType is key to write trace utils.
	TraceType = "trace"
	// DebugIcon is debug utils indicator.
	DebugIcon = "[D]"
	// DebugType is key to write debug utils.
	DebugType = "debug"
	// InfoIcon is info utils indicator.
	InfoIcon = "[I]"
	// InfoType is key to write info utils.
	InfoType = "info"
	// WarnIcon is warning utils indicator.
	WarnIcon = "[W]"
	// WarnType is key to write warn utils.
	WarnType = "warn"
	// ErrorIcon is error utils indicator.
	ErrorIcon = "[E]"
	// ErrorType is key to write error utils.
	ErrorType = "error"
	// FatalIcon is fatal error utils indicator.
	FatalIcon = "[F]"
	// FatalType is key to write fatal utils.
	FatalType = "fatal"
)

// Foreground text colors. The output colors may vary on different OS.
// Taken from https://en.wikipedia.org/wiki/ANSI_escape_code.
const (
	Reset         = "\033[0m"
	Red           = "\033[31m"
	Green         = "\033[32m"
	Yellow        = "\033[33m"
	Blue          = "\033[34m"
	Magenta       = "\033[35m"
	Cyan          = "\033[36m"
	White         = "\033[37m"
	BrightBlack   = "\033[90m"
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"
)

// Print to print colored text in console/terminal.
func Print(textColor ...string) {
	fmt.Print(strings.Join(textColor, "") + Reset)
}

// Println to print colored text with new line in console/terminal.
func Println(textColor ...string) {
	fmt.Println(strings.Join(textColor, "") + Reset)
}

// Trace to write trace utils in console/terminal.
func Trace(text string) {
	timestampStr := time.Now().Format(TimeFormat)
	Println(timestampStr, " ", Blue, TraceIcon, " ", text)
}

// Info to write info utils in console/terminal.
func Info(text string) {
	timestampStr := time.Now().Format(TimeFormat)
	Println(timestampStr, " ", BrightBlue, InfoIcon, " ", text)
}

// Error to write error utils in console/terminal.
func Error(text string) {
	timestampStr := time.Now().Format(TimeFormat)
	Println(timestampStr, " ", BrightRed, ErrorIcon, " ", text)
}

// Debug to write debug utils in console/terminal.
func Debug(text string) {
	timestampStr := time.Now().Format(TimeFormat)
	Println(timestampStr, " ", Yellow, DebugIcon, " ", text)
}

// Success to write success utils in console/terminal.
func Success(text string) {
	timestampStr := time.Now().Format(TimeFormat)
	Println(timestampStr, " ", Green, SuccessIcon, " ", text)
}

// Warn to write warning utils in console/terminal.
func Warn(text string) {
	timestampStr := time.Now().Format(TimeFormat)
	Println(timestampStr, " ", BrightYellow, WarnIcon, " ", text)
}

// Fatal to write fatal error utils in console/terminal.
func Fatal(text string) {
	timestampStr := time.Now().Format(TimeFormat)
	Println(timestampStr, " ", Red, FatalIcon, " ", text)
}

// Log is helper function for easier call for all utils depends on the boolean (isNeeded).
// Not writing utils by default.
func Log(logType string, text string, isNeeded ...bool) {
	if len(isNeeded) > 0 && !isNeeded[0] {
		return
	}

	switch strings.ToLower(logType) {
	case TraceType:
		Trace(text)
	case DebugType:
		Debug(text)
	case SuccessType:
		Success(text)
	case InfoType:
		Info(text)
	case WarnType:
		Warn(text)
	case ErrorType:
		Error(text)
	case FatalType:
		Fatal(text)
	default:
		Println(text)
	}
}

// LogFmt to create string with format for utils.
func LogFmt(text ...string) string {
	return strings.Join(text, " | ")
}
