package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	logg "github.com/sirupsen/logrus"
)

func handle(w http.ResponseWriter, r *http.Request) {
	os := runtime.GOOS
	_, err := fmt.Fprintf(w, "Hello World from [%s] container!\n", os)
	if err != nil {
		return
	}
}

func main() {

	first := 1
	second := 1
	third := 1
	four := 1

	_, _, _, _ = first, second, third, four

	fmt.Println("\033[0;31mred\033[0m")
	fmt.Println("\033[0;32mgreen\033[0m")
	fmt.Println("\033[0;33myellow\033[0m")
	fmt.Println("\033[0;34mblue\033[0m")
	fmt.Println("\033[0;34mblue\033[0m")

	logg.WithFields(logg.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	log.Print("Trying to start server")

	http.HandleFunc("/", handle)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}

type IDE struct {
	Editor   string
	Debugger string
	Build    string
	Name     string
}

func (ide IDE) getAllFunctionality() []string {
	s := []string{ide.Editor, ide.Debugger, ide.Build, ide.Name}

	return s
}
