package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
)

func main() {
	goInfo := fmt.Sprintf("go version: %q, GOOS: %q, GOARCH: %q",
		runtime.Version(), runtime.GOOS, runtime.GOARCH)

	fmt.Println(goInfo)

	goHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, goInfo)
	}

	http.HandleFunc("/go", goHandler)
	log.Println("Listening for requests at http://localhost:8000/go")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
