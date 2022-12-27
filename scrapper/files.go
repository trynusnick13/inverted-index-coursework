package scrapper

import (
	"fmt"
	"io/ioutil"
	"log"
)

func GetAllFilesToRead(dirName string) []string {
	paths := make([]string, 0, 10)
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		paths = append(paths, fmt.Sprintf("%s/%s", dirName, file.Name()))
	}

	return paths
}
