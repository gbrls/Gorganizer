package org

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gbrls/Gorganizer/cmd/cfg"
)

//Organizer holds data for the files organizations
type Organizer struct{}

//FindFiles does a Breath-First-Search to look for all the files
func FindFiles() []string {
	filePaths := make([]string, 0)
	folders := make([]string, 0)

	folders = append(folders, ".")
	//BFS
	for len(folders) > 0 {

		d := folders[0]

		ff, _ := ioutil.ReadDir(d)
		for _, f := range ff {
			if !f.IsDir() {
				filePaths = append(filePaths, fmt.Sprintf("%s/%s", d, f.Name()))
				//fmt.Printf("Found file %s/%s\n", d, f.Name())

				continue
			}

			//fmt.Printf("Found folder %s/%s\n", d, f.Name())
			folders = append(folders, fmt.Sprintf("%s/%s", d, f.Name()))
		}

		folders = folders[1:]
	}

	return filePaths
}

//Org is the entry point for the organization
func Org(c *cfg.Config) (int, error) {

	files := FindFiles()

	ffs, _ := ioutil.ReadDir("./")
	dirs := make(map[string]int, 0)
	for _, f := range ffs {
		if f.IsDir() {
			dirs[f.Name()] = 1
		}
	}

	for k := range c.Data {
		if dirs[k] != 1 {
			os.Mkdir(k, os.ModePerm)
		}
	}

	for _, file := range files {
		for k := range c.Data {
			//ans, err := regexp.MatchString(fmt.Sprintf("%s", k),
			//	file)

			//match all cases
			//upperCase, _ := regexp.MatchString(fmt.Sprintf("%s", k),
			//	file)
			lowerCaseFile := fmt.Sprintf("%s/%s",
				filepath.Dir(file), strings.ToLower(filepath.Base(file)))

			fmt.Println(lowerCaseFile, k)

			lowerCase, err := regexp.MatchString(fmt.Sprintf("%s", k),
				lowerCaseFile)

			if err != nil {
				log.Fatalf("Error in regex (%s)\n", err)
			}

			if !lowerCase {
				continue
			}

			os.Rename(file, fmt.Sprintf("%s/%s", k, filepath.Base(file)))
		}
	}

	return 0, nil
}
