package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

//download used to access file via url to get the response status
//and bytes
func download(file string) {
	resp, errResp := http.Get(file)
	catchError(errResp)

	if resp.StatusCode == 200 {

		contentType := resp.Header.Get("Content-Type")
		filename := getFileName(file)
		filetype := getFileType(contentType)

		// put the file creation into another thread
		go buildFile(filename, filetype, resp.Body)

	} else {
		log.Println("Response not good, ignore it...")
	}
}

//buildFile used to create file in local disk
func buildFile(filename, filetype string, content io.ReadCloser) {

	// close the source here, because this process happened
	// in other thread
	defer content.Close()

	file := filename + "." + filetype
	filepath := *storepath + file
	f, err := os.Create(filepath)

	ignoreError(err)
	defer f.Close()

	//copying file content into disk
	size, errIO := io.Copy(f, content)
	ignoreError(errIO)

	log.Printf("%v downloaded with size %v", file, size)
}

// getFileType used to get file extension
func getFileType(response string) string {
	extract := strings.Split(response, "/")
	return extract[len(extract)-1]
}

// getFileName used to extract file name from url
func getFileName(file string) string {
	extract := strings.Split(file, "/")
	filename := extract[len(extract)-1]
	splitFiletype := strings.Split(filename, ".")

	return splitFiletype[0]
}
