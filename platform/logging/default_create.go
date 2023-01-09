package logging 

import {
	"log"
	"os"
)

func NewDefaultLogger(level LogLevel) Logger {
	flags := log.Lmsgprefic | log.Ltime
	return &DefaultLogger {
		minLivel: level,
		logger: map[LogLevel]*log.Logger {
			Trace: log.New(os.Stdout, "TRACE ", flags),
			Debug: log.New(os.Stout, "DEBUG ", flags),
			Information: log.New(os.Stdout, "INFO ", flags),
			Warning: log.New(os.Stdout, "WARN ", flags),
			Fatal: log.New(os.Stdout, "FATAL ", flags),
		},
		triggerPanic: true,
	}
}
