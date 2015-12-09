package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sync"
)

var (
	filelist  = flag.String("fl", "", "Filepath list of files to downloaded.")
	storepath = flag.String("sp", "", "Directory path to save all files.")
	files     []string
	wg        sync.WaitGroup
)

func init() {
	flag.Parse()
}

func catchError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ignoreError(err error) {
	if err != nil {
		log.Printf("Catch an error but ignore it: %v", err)
	}
}

func main() {

	// We need to know where the list of files location
	if len(*filelist) < 1 {
		log.Fatal("Need to set filelist parameter.")
	}

	// We need to know where the location to save downloaded files
	if len(*storepath) < 1 {
		log.Fatal("Need to set storepath parameter.")
	}

	file, errFile := os.Open(*filelist)

	//catch an error on reading file
	catchError(errFile)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		files = append(files, scanner.Text())
	}

	//catch an error on reading content file
	catchError(scanner.Err())

	// Start workers to grab the file only if the container not empty
	if len(files) >= 1 {

		// number of workers depends on number of files
		for _, f := range files {

			wg.Add(1)

			/*
				Put downloader process into another thread
				for each file.
			*/
			go func(f string) {

				defer wg.Done()

				download(f)

			}(f)

		}

	}

	// wait for all channels until they finish their jobs
	wg.Wait()
}
