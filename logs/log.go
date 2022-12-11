package logs

import (
	"fmt"
	"log"
	"os"
)

func Info(s string) {
	fmt.Print("[INFO] ")
	log.Println(s)
}

func Infof(s string, v ...any) {
	Info(fmt.Sprintf(s, v...))
}

func Error(e error, s string) {
	fmt.Print("[ERROR] ")
	log.Printf("%s. Error = %s\n", s, e.Error())
}

func Errorf(e error, s string, v ...any) {
	Error(e, fmt.Sprintf(s, v...))
}

func Warning(s string) {
	fmt.Print("[WARNING] ")
	fmt.Println(s)
}

func Warningf(s string, v ...any) {
	Warning(fmt.Sprintf(s, v...))
}

func Fatal(e error, s string) {
	fmt.Print("[FATAL] ")
	log.Printf("%s. Error = %s\n", s, e.Error())
	os.Exit(1)
}

func Fatalf(e error, s string, v ...any) {
	Fatal(e, fmt.Sprintf(s, v...))
}
