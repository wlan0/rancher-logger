package rancher_formatter

import (
	"runtime"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type RancherFormatter struct {
	*log.TextFormatter

	// Include detailed stacktrace, useful for debugging
	IncludeStackTrace bool
}

func (rf *RancherFormatter) Format(entry *log.Entry) ([]byte, error) {
	fields := entry.Data
	fields["caller"] = context()

	if rf.IncludeStackTrace {
		fields["trace"] = trace()
	}
	return rf.TextFormatter.Format(entry)
}

func context() string {
	if _, file, line, ok := runtime.Caller(7); ok {
		return strings.Join([]string{file, strconv.Itoa(line)}, ":")
	}
	return "unavailable"
}

func trace() string {
	stack := make([]byte, 2048)
	size := runtime.Stack(stack, false)
	return string(stack[:size])
}
