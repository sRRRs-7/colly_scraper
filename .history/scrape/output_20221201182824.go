package scrape

import (
	"encoding/json"
	"log"
	"os"
)

func saveFileJson(fName string, p *info) {
	f, err := os.Create(fName)
	if err != nil {
		log.Fatal("cannot create file: ", err)
		return
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(p)
	if err != nil {
		log.Fatal("cannot encode json file: ", err)
		return
	}

	_, err = json.MarshalIndent(p, "", " ")
	if err != nil {
		log.Fatal("cannot encode json file: ", err)
		return
	}
}

func outputFile(fName string, arr []string) {
	f, err := os.Create(fName)
	if err != nil {
		log.Fatal("cannot create file: ", err)
		return
	}
	defer f.Close()
	for i, s := range arr {}
}