package main

import (
	"net/http"
	"fmt"
	"os"
)

func servespecial(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/":
		filerespond(w, r, "index.html")
	default:
		http.NotFound(w, r)
	}
}

func basicserve(w http.ResponseWriter, r *http.Request) {
	filerespond(w, r, r.URL.Path[1:])
}

func filerespond(w http.ResponseWriter, r *http.Request, path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	stat, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.ServeContent(w, r, path, stat.ModTime(), file)
}

func main() {
	// We expect exactly one argument - port number
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <port_number>\n", os.Args[0])
		os.Exit(1)
	}

	// Static page handlers
	http.HandleFunc("/", servespecial)
	http.HandleFunc("/style.css", basicserve)
	http.HandleFunc("/Strasua.ttf", basicserve)
	http.HandleFunc("/robots.txt", basicserve)
	http.HandleFunc("/public.key", basicserve)
	http.HandleFunc("/bg.png", basicserve)
	
	// Dynamic handlers for content shared in the /public folder
	http.HandleFunc("/public/", basicserve)

	// Go away and serve pages forever
	err := http.ListenAndServe(":"+os.Args[1], nil)
	if err != nil { fmt.Println(err) }
}
