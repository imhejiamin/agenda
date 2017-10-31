package loghelper

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	// Info : Discard
	Info *log.Logger
	// Warning : Stdout
	Warning *log.Logger
	// Error : Stderr
	Error *log.Logger
)

var errlog *os.File

func set(
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func init() {
	errlog = getErrLogFile()
	set(ioutil.Discard, os.Stdout, errlog)

	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}

// Free : close log file
func Free() {
	errlog.Close()
}

func getErrLogFile() *os.File {
	logPath := filepath.Join(os.Getenv("GOPATH"), "/src/Agenda/data/error.log")
	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file open error : %v", err)
	}
	return file
	// defer file.Close()
}
