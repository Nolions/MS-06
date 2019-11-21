package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type file struct {
	Name string
	AttachmentName string
	Dictionaries   []string
}

var path string

func main() {
	flag.StringVar(&path, "path", "/Users", "Path")
	flag.Parse()
	fmt.Println(path)
	for _, v := range walkDir(path) {
		fmt.Println(v)
	}

}

func walkDir(p string) []file {
	var files []file
	filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if info.IsDir() {
			fmt.Printf("======%s======\n", path)
			return nil
		}

		dictionaries := strings.Split(path, "/")
		i := strings.Split(info.Name(), ".")
		a := ""
		if len(i) > 1 {
			a = i[len(i)-1]
		}
		f := file{
			Name:           i[0],
			Dictionaries:   dictionaries[:len(dictionaries)-1],
			AttachmentName: a,
		}
		files = append(files, f)
		return nil
	})

	return files
}
