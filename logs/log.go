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

func Error(s string, e error) {
	fmt.Print("[ERROR] ")
	log.Printf("%s. Error = %s\n", s, e.Error())
}

func Warning(s string) {
	fmt.Print("[WARNING] ")
	fmt.Println(s)
}

func Fatal(s string, e error) {
	fmt.Print("[FATAL] ")
	log.Printf("%s. Error = %s\n", s, e.Error())
	os.Exit(1)
}
