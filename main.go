package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)

func main() {
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/upload", uploadHandle)

	http.ListenAndServe(":8080", nil)
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "upload.html")
}

func uploadHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		file, handle, err := r.FormFile("file")
		defer file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		f, err := os.OpenFile("./uploads/" + handle.Filename, os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		fmt.Fprintf(w, "Upload Complete")
	} else {
		http.ServeFile(w, r, "upload.html")
	}
}