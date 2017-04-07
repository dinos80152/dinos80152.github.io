package main

import (
	"io/ioutil"
	"os"
	"text/template"

	"strings"

	"github.com/russross/blackfriday"
)

const (
	noteMdDir      string = "_notes"
	noteHTMLDir    string = "notes"
	noteURL        string = "/notes/"
	noteListLayout string = "tpls/notelist.gtpl"
	layout         string = "tpls/layout.gtpl"
	mdIndex        string = "readme.md"
	htmlIndex      string = "index.html"
)

type article struct {
	Title       string
	Content     string
	UpdatedAt   string
	UpdatedDate string
	Href        string
}

var noteList []article

func main() {
	cleanFolder(noteHTMLDir)
	genNotes(noteMdDir, noteHTMLDir)
	readmeToIndex("./")
}

func readmeToIndex(dir string) {
	paths, _ := ioutil.ReadDir(dir)
	for _, path := range paths {
		if path.IsDir() {
			readmeToIndex(path.Name())
		}
		if strings.EqualFold(path.Name(), mdIndex) {
			if dir == noteMdDir {
				md2HTML(dir, noteHTMLDir, path, htmlIndex)
			} else {
				md2HTML(dir, dir, path, htmlIndex)
			}
		}
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

func genNotes(mdDir, htmlDir string) {
	files, _ := ioutil.ReadDir(mdDir)
	for _, f := range files {
		if strings.EqualFold(f.Name(), "readme.md") {
			continue
		}
		md2HTML(mdDir, htmlDir, f, "")
	}
	genNoteList()
}

func md2HTML(sourceDir, outputDir string, fi os.FileInfo, outputBaseName string) {
	baseName := fi.Name()
	input, _ := ioutil.ReadFile(sourceDir + "/" + baseName)
	output := blackfriday.MarkdownCommon(input)
	content := string(output)

	// TODO: use regexp
	title := content[strings.Index(content, ">")+1 : strings.Index(content, "/")-1]
	atcl := article{
		Title:       title,
		Content:     content,
		UpdatedAt:   fi.ModTime().Format("2006-01-02 15:04:05 MST Z07"),
		UpdatedDate: fi.ModTime().Format("2006-01-02"),
	}

	if outputBaseName == "" {
		outputBaseName = fileName(baseName) + ".html"
	}

	if sourceDir == noteMdDir {
		atcl.Href = noteURL + outputBaseName
		noteList = append(noteList, atcl)
	}

	f, _ := os.Create(outputDir + "/" + outputBaseName)
	defer f.Close()

	t, _ := template.ParseFiles(layout)
	t.Execute(f, atcl)
}

func genNoteList() {
	f, _ := os.Create(noteMdDir + "/readme.md")
	defer f.Close()

	t, _ := template.ParseFiles(noteListLayout)
	t.Execute(f, noteList)
}

func fileName(baseName string) string {
	names := strings.Split(baseName, ".")
	return names[0]
}
