package log

import (
	"fmt"
	"log"
	"os"
)

// Severity to declare log message severities
type Severity uint

const (
	// SeverityNotApplied Severity level for tracing log entries
	SeverityNotApplied Severity = 8
	// SeverityTrace Severity level for tracing log entries
	SeverityTrace Severity = 7
	// SeverityInfo Severity level for information log entries
	SeverityInfo Severity = 6
	// SeverityNotice Severity level for notice log entries
	SeverityNotice Severity = 5
	// SeverityWarning Severity level for warning log entries
	SeverityWarning Severity = 4
	// SeverityError Severity level for error log entries
	SeverityError Severity = 3
	// SeverityCritical Severity level for critical log entries
	SeverityCritical Severity = 2
	// SeverityAlert Severity level for alert log entries
	SeverityAlert Severity = 1
	// SeverityEmergency Severity level for emergency log entries
	SeverityEmergency Severity = 0
)

var (
	// Trace logger to print Trace messages to stdout
	Trace *log.Logger
	// Info logger to print Info messages to stdout
	Info *log.Logger
	// Notice logger to print Notice messages to stdout
	Notice *log.Logger
	// Warning logger to print Warning messages to stdout
	Warning *log.Logger
	// Error logger to print Error messages to stderr
	Error *log.Logger
	// Critical logger to print Critical messages to stderr
	Critical *log.Logger
	// Alert logger to print Alert messages to stderr
	Alert *log.Logger
	// Emergency logger to print Emergency messages to stderr
	Emergency *log.Logger
)

var (
	loggers       = []**log.Logger{&Emergency, &Alert, &Critical, &Error, &Warning, &Notice, &Info, &Trace}
	logPrefixes   = []string{"EMERG: ", "ALERT: ", "CRIT:  ", "ERROR: ", "WARN:  ", "NOTICE:", "INFO:  ", "TRACE: ", "N/A:   "}
	severityNames = []string{"Emergency", "Alert", "Critical", "Error", "Warnin", "Notice", "Info", "Trace"}
)

func init() {
	isSystemD := (os.Getenv("INVOCATION_ID") != "")
	for severityLevel := Severity(0); severityLevel < SeverityNotApplied; severityLevel++ {
		writer := os.Stdout
		if severityLevel <= SeverityError {
			writer = os.Stderr
		}
		prefix := logPrefixes[severityLevel]
		flag := log.LstdFlags //log.Lshortfile | log.LstdFlags
		if isSystemD {
			prefix = fmt.Sprintf("<%v>%v", severityLevel, logPrefixes[severityLevel])
			flag = 0
		}
		*loggers[severityLevel] = log.New(writer, prefix, flag)
	}
}
