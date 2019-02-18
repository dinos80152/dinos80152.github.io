package main

import (
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"

	"github.com/dinos80152/dinos80152.github.io/pkg"
)

const (
	noteMdDir      string = "_notes"
	noteHTMLDir    string = "notes"
	noteURL        string = "/notes/"
	mdIndex        string = "readme.md"
	htmlIndex      string = "index.html"
	layout         string = "tpls/layout.gtpl"
	noteListLayout string = "tpls/notelist.gtpl"
	logLayout      string = "tpls/log.gtpl"
)

type noteList pkg.NoteList

var nl = noteList{}

func main() {
	cleanFolder(noteHTMLDir)
	genNotes(noteMdDir, noteHTMLDir)
	genUpdateLog()
	readmeToIndex("./")
	genSiteMap()
	// find ./ -type d -exec chmod 755 {} \;
	// find ./ -type f -exec chmod 644 {} \;
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
		genHTML(mdDir, htmlDir, f, "")
	}
	genNoteList()
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

func genUpdateLog() {
	log := make(noteList)
	logDate := make([]string, 0)

	for _, notes := range nl {
		for _, note := range notes {
			log[note.UpdatedDate] = append(log[note.UpdatedDate], note)
		}
	}

	for _, notes := range log {
		sort.Slice(notes, func(i, j int) bool {
			if notes[i].Category == notes[j].Category {
				return notes[i].Title < notes[j].Title
			}
			return notes[i].Category < notes[j].Category
		})
	}

	for date := range log {
		logDate = append(logDate, date)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(logDate)))

	f, _ := os.Create("README.md")
	defer f.Close()

	t, _ := template.ParseFiles(logLayout)
	t.Execute(f, struct {
		Log     noteList
		LogDate []string
	}{
		log,
		logDate,
	})
}

func readmeToIndex(dir string) {
	paths, _ := ioutil.ReadDir(dir)
	for _, path := range paths {
		if path.IsDir() {
			readmeToIndex(path.Name())
		}
		if strings.EqualFold(path.Name(), mdIndex) {
			if dir == noteMdDir {
				genHTML(dir, noteHTMLDir, path, htmlIndex)
			} else {
				genHTML(dir, dir, path, htmlIndex)
			}
		}
	}
}

func genSiteMap() {
	us := pkg.NewURLSet(pkg.NoteList(nl))
	us.GenSiteMap()
}

func genHTML(sourceDir, outputDir string, fi os.FileInfo, outputBaseName string) {
	note := pkg.NewNote(sourceDir, outputDir, fi, outputBaseName)
	note.ToHTML()
	nl[note.Category] = append(nl[note.Category], note)
}
