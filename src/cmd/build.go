package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fns := template.FuncMap{
		"getenv": os.Getenv,
	}

	if len(os.Args) != 3 {
		log.Fatal("expected 2 arguments (source and destination directories)")
	}
	srcDir := os.Args[1]
	dstDir := os.Args[2]

	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filename := file.Name()
		extOnly := filepath.Ext(filename)
		if extOnly != ".html" {
			continue
		}

		fmt.Println("generating template ", filename, "...")
		out, err := os.Create(filepath.Join(dstDir, filename))
		if err != nil {
			log.Fatal(err)
		}
		tpl, err := template.New(filename).Funcs(fns).ParseFiles(filepath.Join(srcDir, filename))
		if err != nil {
			log.Fatal(err)
		}
		if err := tpl.Execute(out, nil); err != nil {
			log.Fatal(err)
		}
	}
}
