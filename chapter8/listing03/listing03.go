// This sample program demonstrates how to use the base log package.
package main

import (
	"log"
	"os"
)
var debuger log.Logger
func init() {
	fileName := "Info_First.log"
	logFile, err := os.Create(fileName)
	//defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error")
	}
	debuger = *log.New(logFile, "[Info]", log.Llongfile)
	debuger.SetPrefix("TRACE: ")

	//log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {

	debuger.Println("message")
	// Println writes to the standard logger.
	//debuger.Println("message")

	// Fatalln is Println() followed by a call to os.Exit(1).
	//debuger.Fatalln("fatal message")

	// Panicln is Println() followed by a call to panic().
	//debuger.Panicln("panic message")
}
