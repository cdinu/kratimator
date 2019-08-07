package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sync"
)

var (
	isForRealPtr *bool
)

func main(){
	isForRealPtr = flag.Bool("isForReal", false, "set this to overwrite your files")

	flag.Parse()

	var wg sync.WaitGroup
	log.Println("Starting...")
	files := flag.Args()

	if len(files) == 0 {
		log.Fatal("ERROR: You need to provide files to process")
	}

	wg.Add(len(files))
	for _, file := range files {
		filePath := file
		go func() {
			processFile(filePath);
			wg.Done()
		}()
	}

	wg.Wait()
}

func processFile(filePath string) {
	log.Printf("Kratimating %s", filePath);
	in, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create(filePath+".kratimated.jade")
	if err != nil {
		log.Fatal(err)
	}

	defer in.Close()
	defer out.Close()

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		_, err = out.WriteString(kratimate(line) + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	if *isForRealPtr {
		os.Rename(filePath+".kratimated.jade", filePath)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
