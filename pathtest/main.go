package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func handleFunc(path string, f os.FileInfo, err error) error {
	fmt.Println(path)
	if f.IsDir() {
		log.Println(f.Name(), " is a dir")
		err := os.Remove(path)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

func main() {
	path := "../pathtest"
	fi, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		return
	}
	if fi.IsDir() {
		filepath.Walk(path, handleFunc)
	}
}
