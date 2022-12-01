package scrape

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func saveFileJson(fName string, arr *info) {
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

	b, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		log.Fatal("cannot encode json file: ", err)
		return
	}
	fmt.Println(string(b))
}