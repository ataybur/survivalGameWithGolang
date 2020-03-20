package httputils

import (
	"bytes"
	"fmt"
	"game/funcs"
	"game/structs"
	"io"
	"net/http"
	"strings"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	var buf bytes.Buffer
	io.Copy(&buf, file)
	// do something with the contents...
	// I normally have a struct defined and unmarshal into a struct, but this will
	// work as an example
	contents := buf.String()
	lines := strings.Split(contents, "\n")
	contextP := &structs.Context{}
	contextP.Init()
	contextP.Fill(lines)
	funcs.Play(contextP)
	buf.Reset()
	fmt.Fprintf(w, contents+"\n"+contextP.GetLog())
	fmt.Println(contents)
	// I reset the buffer in case I want to use it again
	// reduces memory allocations in more intense projects

}

func SetupRoutes() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/upload", UploadFile)

	if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Could not listen: %v\n", err)
	}
	fmt.Println("Server stopped")
}
