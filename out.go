// something about log writer
package logcat

import (
	"errors"
	"io"
	"os"
	"time"
)

var todaylogfile = make(map[string]*os.File)

// wrapping the writer with it, when we donot want to see the color stuff.
type cleanLog struct {
	file io.Writer
}

// just strip the '\033[xxm'.
func (clf cleanLog) Write(b []byte) (n int, err error) {
	return clf.file.Write(append(b[5:len(b)-5], byte('\n')))
}

// return a log file of today, the filename is today's date.
// prefix is the path to log file. if u leave it empty, it will be ".".
func Today(prefix string) *os.File {
	if prefix == "" {
		prefix = "."
	}
	daytime := time.Now().Format("2006-01-02")
	filename := daytime + ".log"
	filepath := prefix + "/"
	os.MkdirAll(filepath, os.ModePerm)
	logflie, err := os.OpenFile(filepath+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	// save the reference
	todaylogfile[prefix] = logflie
	return logflie
}

// close the file creating by Today().
// when you close it, the logfile would not work.
func CloseToday(prefix string) error {
	if prefix == "" {
		prefix = "."
	}
	if todaylogfile[prefix] == nil {
		return errors.New("You donot saved that log today.")
	}
	defer func() { todaylogfile[prefix] = nil }()
	return todaylogfile[prefix].Close()
}

// using a log file to saving bad and warn log. the file could be a writer
// using Today(), Example: logcat.SetLogFile(logcat.Today("."))
func SetLogFile(file io.Writer) {
	if file == os.Stdout {
		return
	}
	wrapFile := cleanLog{file: file}
	badCat.SetOutput(io.MultiWriter(os.Stdout, wrapFile))
	warnCat.SetOutput(io.MultiWriter(os.Stdout, wrapFile))
}

// set the trace out writer
func SetTraceWay(w io.Writer) {
	if w == os.Stdout {
		traceCat.SetOutput(w)
		return
	}
	traceCat.SetOutput(cleanLog{file: w})
}
