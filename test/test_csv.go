package main

import (
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"gitlab.com/agamigo/agamigo/coupler/metafarms"
)

func main() {
	csvFile, err := os.Open("example.csv")
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}

	kss, err := metafarms.NewKillsheetsFromCSV(csvFile)
	if err != nil {
		log.Fatalf("Unable to read file into ks: %v", err)
	}

	for _, ks := range kss {
		spew.Dump(ks)
	}

	err = csvFile.Close()
	if err != nil {
		log.Fatalf("Unable to close file: %v", err)
	}
}
