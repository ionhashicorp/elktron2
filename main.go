package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"net/http"
)

var (
	//go:embed react/build
	react embed.FS
)

func frontendHomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "<h1>Hello frontend</h1>")
}

// func frontendReactHandler(rw http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(rw, "<h1>Hello frontend</h1>")
// }

func backendHomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "<h1>Hello backend</h1>")
}

func main() {
	isFrontend := flag.Bool("frontend", false, "Usage: true to start frontend, false do not start frontend. Default is false")
	isBackend := flag.Bool("backend", false, "Usage: true to start backend, false to not start backend. Default is false")
	flag.Parse()

	dirBuild, _ := fs.Sub(react, "react/build")
	if *isFrontend {
		fmt.Println("starting frontend on port :3000")
		http.HandleFunc("/", frontendHomeHandler)
		http.Handle("/react", http.FileServer(http.FS(dirBuild))) // access with http://localhost:3000/react/index.html
		http.ListenAndServe(":3000", nil)
	}

	if *isBackend {
		fmt.Println("starting backend on port :8080")
		http.HandleFunc("/", backendHomeHandler)
		http.ListenAndServe(":8080", nil)
	}

	fmt.Println("You need to select either backend or frontend")
}
