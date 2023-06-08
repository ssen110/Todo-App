package utilfunc

import (
	"fmt"
	"log"
	"os"
)

func PrintMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func CheckErrWhileScanning(err error) {
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile("myLOG.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    log.SetOutput(file)
}

func CheckErrWhileInterractingWithDB(err error) {
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile("myLOG.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    log.SetOutput(file)
}