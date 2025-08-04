package main

import (
	"flag"
	"fmt"
	"github.com/marcosparreiras/fsDuplicateReport/fileMap"
	"github.com/marcosparreiras/fsDuplicateReport/ignorePathRegexs"
	"io/fs"
	"log"
	"os"
)

func main() {
	var dirPath string
	ignorePathRegexs := ignorePathRegexs.New()
	parseFlags(&dirPath, &ignorePathRegexs)
	fileSystem := os.DirFS(dirPath)
	fileMap := fileMap.New()
	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if ignorePathRegexs.MatchPath(path) {
			return nil
		}
		if d.IsDir() {
			return nil
		}
		info, err := os.Stat(dirPath + "/" + path)
		if err != nil {
			log.Fatal(err)
		}
		fileMap.AddFileInfo(d.Name(), path, info.Size())
		return nil
	})
	fmt.Println(fileMap)
}

func parseFlags(dirPath *string, ignoreFolders *ignorePathRegexs.IgnorePathRegexs) {
	flag.Var(ignoreFolders, "ignore", "Regex to ignore path")
	flag.StringVar(dirPath, "path", "/", "Directory path to scan")
	flag.Parse()
}
