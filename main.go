package main

import (
	"io/ioutil"
	"os"
	"path"
	"text/template"

	"strings"

	"sort"

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

type note struct {
	Title       string
	Content     string
	UpdatedAt   string
	UpdatedDate string
	Href        string
}

type noteList map[string][]note

var nl = noteList{}

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
	files, _ := ioutil.ReadDir(path)
	root, _ := os.Getwd()
	os.Chdir(path)
	defer os.Chdir(root)
	for _, f := range files {
		if f.IsDir() {
			os.RemoveAll(f.Name())
		} else {
			os.Remove(f.Name())
		}
	}
}

func genNotes(mdDir, htmlDir string) {
	files, _ := ioutil.ReadDir(mdDir)
	for _, f := range files {
		if f.IsDir() {
			nextMdDir := path.Join(mdDir, f.Name())
			nextHTMLDir := path.Join(htmlDir, f.Name())
			os.Mkdir(nextHTMLDir, 755)
			genNotes(nextMdDir, nextHTMLDir)
			continue
		}
		if strings.EqualFold(f.Name(), mdIndex) {
			continue
		}
		md2HTML(mdDir, htmlDir, f, "")
	}
	genNoteList()
}

func md2HTML(sourceDir, outputDir string, fi os.FileInfo, outputBaseName string) {
	baseName := fi.Name()
	input, _ := ioutil.ReadFile(path.Join(sourceDir, baseName))
	output := markdown(input)
	content := string(output)

	// TODO: use regexp
	title := content[strings.Index(content, ">")+1 : strings.Index(content, "/")-1]
	nt := note{
		Title:       title,
		Content:     content,
		UpdatedAt:   fi.ModTime().Format("2006-01-02 15:04:05 MST Z07"),
		UpdatedDate: fi.ModTime().Format("2006-01-02"),
	}

	if outputBaseName == "" {
		outputBaseName = fileName(baseName) + ".html"
	}

	if path.Dir(sourceDir) == noteMdDir {
		nt.Href = path.Join("/", outputDir, outputBaseName)
		category := strings.Title(path.Base(sourceDir))
		nl[category] = append(nl[category], nt)
	}

	f, _ := os.Create(path.Join(outputDir, outputBaseName))
	defer f.Close()

	t, _ := template.ParseFiles(layout)
	t.Execute(f, nt)
}

func genNoteList() {

	for _, notes := range nl {
		sort.Slice(notes, func(i, j int) bool {
			return notes[i].UpdatedAt > notes[j].UpdatedAt
		})
	}

	f, _ := os.Create(path.Join(noteMdDir, "/readme.md"))
	defer f.Close()

	t, _ := template.ParseFiles(noteListLayout)
	t.Execute(f, nl)
}

func fileName(baseName string) string {
	names := strings.Split(baseName, ".")
	return names[0]
}

func markdown(input []byte) []byte {
	htmlFlags := 0 |
		blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_FRACTIONS |
		blackfriday.HTML_SMARTYPANTS_DASHES |
		blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

	extensions := 0 |
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS |
		blackfriday.EXTENSION_AUTO_HEADER_IDS |
		blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
		blackfriday.EXTENSION_DEFINITION_LISTS

	renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")
	return blackfriday.MarkdownOptions(input, renderer, blackfriday.Options{
		Extensions: extensions})
}
