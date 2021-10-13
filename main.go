package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
)

func getGoInfo() string {
	goInfo := fmt.Sprintf("go version: %q, GOOS: %q, GOARCH: %q",
		runtime.Version(), runtime.GOOS, runtime.GOARCH)

	return goInfo
}

func startHTTPServer(servePort int) {
	port := strconv.Itoa(servePort)

	goHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, getGoInfo())
	}

	http.HandleFunc("/go", goHandler)
	log.Println("Listening for requests on port: " + port + " path: /go")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	fmt.Println(getGoInfo())

	// pass the --serve PORT command to launch the server, otherwise it just outputs the go info to stdout
	if len(os.Args) == 3 {
		arg1 := os.Args[1]

		if arg1 == "--serve" {
			servePort, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal("Expected numeric port number for --serve command (e.g. --serve 8000)")
			}
			startHTTPServer(servePort)
		}
	}
}
