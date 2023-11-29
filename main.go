package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func handle(w http.ResponseWriter, r *http.Request) {
	os := runtime.GOOS
	fmt.Fprintf(w, "Hello World from [%s] container!\n", os)
}

func main() {

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
	Name 	 string
}

func (ide IDE) getAllFunctionality() []string {
	s := []string{ide.Editor, ide.Debugger, ide.Build, ide.Name}

	return s
}
