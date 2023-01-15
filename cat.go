// logcat is a simple logger with color texts in terminal,
// which is just be encapsulated simplely from golang's log package.
// and it's eazier for me to use. :p
package logcat

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// csi colors
const (
	red     = "\033[31m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	end     = "\033[0m"
)

// wrapping the writer with it, when we donot want to see the color stuff.
type cleanLog struct {
	file io.Writer
}

// just strip the '\033[xx m'.
func (clf cleanLog) Write(b []byte) (n int, err error) {
	return clf.file.Write(append(b[5:len(b)-5], byte('\n')))
}

// new logger
var (
	goodCat  = log.New(os.Stdout, magenta, log.Ldate|log.Ltime)
	badCat   = log.New(os.Stdout, red, log.Ldate|log.Ltime)
	warnCat  = log.New(os.Stdout, yellow, log.Ldate|log.Ltime)
	infoCat  = log.New(os.Stdout, cyan, log.Ldate|log.Ltime)
	traceCat = log.New(ioutil.Discard, blue, log.Ldate|log.Ltime)
)

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
	logflie, _ := os.OpenFile(filepath+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	return logflie
}

// using a log file to saving bad and warn log. the file could be a writer
// with Today func, Example: logcat.SetLogFile(logcat.Today("."))
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

// format print good log
func Goodf(format string, v ...any) {
	goodCat.Printf("[Good] "+format+end, v...)
	Tracef(format, v...)
}

// print line good log
func Goodln(v ...any) {
	goodCat.Println("[Good] "+fmt.Sprint(v...), end)
	Traceln(v...)
}

// format print info log
func Infof(format string, v ...any) {
	infoCat.Printf("[Info] "+format+end, v...)
	Tracef(format, v...)
}

// print line info log
func Infoln(v ...any) {
	infoCat.Println("[Info] "+fmt.Sprint(v...), end)
	Traceln(v...)
}

// format print bad log
func Badf(format string, v ...any) {
	badCat.Printf("[Bad] "+format+end, v...)
	Tracef(format, v...)
}

// print line bad log
func Badln(v ...any) {
	badCat.Println("[Bad] "+fmt.Sprint(v...), end)
	Traceln(v...)
}

// format print warn log
func Warnf(format string, v ...any) {
	warnCat.Printf("[Warn] "+format+end, v...)
	Tracef(format, v...)
}

// print line warn log
func Warnln(v ...any) {
	warnCat.Println("[Warn] "+fmt.Sprint(v...), end)
	Traceln(v...)
}

// format print trace log
func Tracef(format string, v ...any) {
	traceCat.Printf("[Trace] "+format+end, v...)
}

// print line trace log
func Traceln(v ...any) {
	traceCat.Println("[Trace] "+fmt.Sprint(v...), end)
}
