package utils

import "testing"

// TestPrint to test log print function.
func TestPrint(t *testing.T) {
	Print(Red, "Nice")
}

// TestPrintln to test log print with new line function.
func TestPrintln(t *testing.T) {
	Println(Blue, "Nice new line")
}

// TestTrace to test Trace function.
func TestTrace(t *testing.T) {
	Info("Trace sample")
}

// TestInfo to test Info function.
func TestInfo(t *testing.T) {
	Info("Starting app")
}

// TestDebug to test Debug function.
func TestDebug(t *testing.T) {
	Debug("Good response")
}

// TestWarn to test Warn function.
func TestWarn(t *testing.T) {
	Warn("Deprecated function")
}

// TestSuccess to test Success function.
func TestSuccess(t *testing.T) {
	Success("Created user")
}

// TestError to test Error function.
func TestError(t *testing.T) {
	Error("Failed encode")
}

// TestFatal to test Fatal function.
func TestFatal(t *testing.T) {
	Fatal("Server dead")
}

// LogTest is simple model for Log test.
type LogTest struct {
	Type     string
	Text     string
	IsNeeded bool
}

// TestLog to test Log function.
func TestLog(t *testing.T) {
	tests := []LogTest{
		{"success", "nice", false},
		{"trace", "nice", true},
		{"success", "nice", true},
		{"info", "nice", true},
		{"debug", "nice", true},
		{"warn", "nice", true},
		{"error", "nice", true},
		{"fatal", "nice", true},
		{"magic", "nice", true},
	}

	for _, ts := range tests {
		Log(ts.Type, ts.Text, ts.IsNeeded)
	}
}

// LogFmtTest is simple model for LogFmt test.
type LogFmtTest struct {
	Texts  []string
	Result string
}

// TestLogFmt to test LogFmt function.
func TestLogFmt(t *testing.T) {
	tests := []LogFmtTest{
		{
			Texts:  []string{},
			Result: "",
		},
		{
			Texts:  []string{"a"},
			Result: "a",
		},
		{
			Texts:  []string{"a", "b", "c"},
			Result: "a | b | c",
		},
	}

	for _, test := range tests {
		res := LogFmt(test.Texts...)
		if res != test.Result {
			t.Errorf("Expected %v got %v", test.Result, res)
		}
	}
}
