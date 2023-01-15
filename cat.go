// logcat is a simple logger with color texts in terminal,
// which just encapsulate golang's log package.
// and it's eazier for me to use. :p
// here is basic logger define
package logcat

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

// new logger
var (
	goodCat  = log.New(os.Stdout, magenta, log.LstdFlags)
	badCat   = log.New(os.Stdout, red, log.LstdFlags)
	warnCat  = log.New(os.Stdout, yellow, log.LstdFlags)
	infoCat  = log.New(os.Stdout, cyan, log.LstdFlags)
	traceCat = log.New(ioutil.Discard, blue, log.LstdFlags)
)

// set the prefix for logger
func SetCatPrefix(prefix string) {
	goodCat.SetPrefix(magenta + prefix)
	badCat.SetPrefix(red + prefix)
	warnCat.SetPrefix(yellow + prefix)
	infoCat.SetPrefix(cyan + prefix)
	traceCat.SetPrefix(blue + prefix)
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
