package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
)

func startHTTPServer(goInfo string) {
	goHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, goInfo)
	}

	http.HandleFunc("/go", goHandler)
	log.Println("Listening for requests at :8000/go")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func main() {
	goInfo := fmt.Sprintf("go version: %q, GOOS: %q, GOARCH: %q",
		runtime.Version(), runtime.GOOS, runtime.GOARCH)

	fmt.Println(goInfo)

	// pass the --serve command to launch the server, otherwise it just outputs the go info to stdout
	if len(os.Args) == 2 {
		arg := os.Args[1]

		if arg == "--serve" {
			startHTTPServer(goInfo)
		}
	}
}
