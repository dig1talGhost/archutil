package logger

import (
	"os"

	"github.com/charmbracelet/log"
)

var Standard = log.NewWithOptions(os.Stderr, log.Options{
	ReportTimestamp: false,
	Prefix:          ":",
})

var TimeStamped = log.NewWithOptions(os.Stderr, log.Options{
	ReportTimestamp: true,
	Prefix:          ":",
})
