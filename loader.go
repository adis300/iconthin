package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

var relativePath = ""

func getRelativePath() string {
	if relativePath == "" {
		filename := os.Args[0] // get command line first parameter

		filedirectory := filepath.Dir(filename)

		relativePath, _ = filepath.Abs(filedirectory)
		return relativePath
	}
	return relativePath
}

func loadFile(filepath string) []byte {
	data, _ := ioutil.ReadFile(filepath)
	return data
}

func loadView(viewName string) []byte {
	filepath := relativePath + "/public/" + viewName + ".html"
	return loadFile(filepath)
}
