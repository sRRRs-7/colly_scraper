package scrape

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
)

func saveFileJson(fName string, p *info) {
	f, err := os.Create(fName)
	if err != nil {
		log.Fatal("cannot create file: ", err)
		return
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")
	err = enc.Encode(p)
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

	for _, s := range arr {
		_, err := f.WriteString(s)
		if err != nil {
			log.Fatal("failed to write file txt")
		}
	}

	file, err := os.Open(fName)
	if err != nil {
		log.Fatal("cannot open file: ", err)
		return
	}
	defer file.Close()

	var lines [][]string
	s := bufio.NewScanner(file)
	newline := true

	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if len(line) == 0 {
			newline = true
			continue
		}
		if newline {
			newline = false
			lines = append(lines, make([]string, 0))
		}
		last := len(lines) - 1
		lines[last] = append(lines[last], line)
	}
	if err := s.Err(); err != nil {
		log.Fatal("trim file error", err)
		return
	}

	var str []string
	for _, l := range lines {
		str = append(str, strings.Join(l, " | "))
	}

	f, err = os.Create(fName)
	if err != nil {
		log.Fatal("cannot create file: ", err)
		return
	}
	defer f.Close()

	for _, s := range str {
		_, err := f.WriteString(s)
		f.WriteString("\n")
		if err != nil {
			log.Fatal("failed to write file txt")
		}
	}
}