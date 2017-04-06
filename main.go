package main

import (
	"io/ioutil"
	"os"
	"text/template"

	"strings"

	"github.com/russross/blackfriday"
)

const (
	noteMdDir   string = "_notes"
	noteHTMLDir string = "notes"
	layout      string = "tpls/layout.gtpl"
	mdIndex     string = "readme.md"
	htmlIndex   string = "index.html"
)

type article struct {
	Title    string
	Content  string
	LastEdit string
}

func main() {
	readmeToIndex("./")
	cleanFolder(noteHTMLDir)
	files, _ := ioutil.ReadDir(noteMdDir)
	for _, f := range files {
		genHTML(noteMdDir, noteHTMLDir, f, "")
	}
}

func cleanFolder(path string) {
	root, _ := os.Getwd()
	files, _ := ioutil.ReadDir(path)
	os.Chdir(path)
	defer os.Chdir(root)
	for _, f := range files {
		if f.IsDir() {
			cleanFolder(f.Name())
		} else {
			os.Remove(f.Name())
		}
	}
}

func readmeToIndex(dir string) {
	paths, _ := ioutil.ReadDir(dir)
	for _, path := range paths {
		if path.IsDir() {
			readmeToIndex(path.Name())
		}
		if strings.EqualFold(path.Name(), mdIndex) {
			genHTML(dir, dir, path, htmlIndex)
		}
	}
}

func genHTML(sourceDir, outputDir string, fi os.FileInfo, outputBaseName string) {
	baseName := fi.Name()
	input, _ := ioutil.ReadFile(sourceDir + "/" + baseName)
	output := blackfriday.MarkdownCommon(input)
	content := string(output)

	// TODO: use regexp
	title := content[strings.Index(content, ">")+1 : strings.Index(content, "/")-1]
	atcl := article{
		Title:    title,
		Content:  content,
		LastEdit: fi.ModTime().Format("2006-01-02 15:04:05 MST Z07"),
	}

	if outputBaseName == "" {
		outputBaseName = fileName(baseName) + ".html"
	}
	f, _ := os.Create(outputDir + "/" + outputBaseName)
	defer f.Close()

	t, _ := template.ParseFiles(layout)
	t.Execute(f, atcl)
}

func fileName(baseName string) string {
	names := strings.Split(baseName, ".")
	return names[0]
}
