package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"
	"time"

	"github.com/russross/blackfriday"
)

const layout string = "tpls/layout.gtpl"

type Note struct {
	Title          string
	Content        string
	Category       string
	UpdatedAt      string
	UpdatedDate    string
	Href           string
	sourceDir      string
	outputDir      string
	fi             os.FileInfo
	outputBaseName string
	baseName       string
	filePath       string
}

type NoteList map[string][]*Note

func NewNote(sourceDir, outputDir string, fi os.FileInfo, outputBaseName string) *Note {
	note := new(Note)
	note.sourceDir = sourceDir
	note.outputDir = outputDir
	note.fi = fi
	note.outputBaseName = outputBaseName
	note.baseName = fi.Name()
	note.filePath = path.Join(sourceDir, note.baseName)
	return note
}

func (nt *Note) ToHTML() {
	input, _ := ioutil.ReadFile(nt.filePath)
	output := markdown(input)
	content := string(output)

	// TODO: use regexp
	title := content[strings.Index(content, ">")+1 : strings.Index(content, "/")-1]

	nt.Title = title
	nt.Content = content
	nt.Category = nt.outputDir

	t, err := nt.modifiedTime()
	if err != nil {
		fmt.Println(nt.filePath)
		panic(err)
	}
	nt.UpdatedAt = t.Format("2006-01-02 15:04:05 MST Z07")
	nt.UpdatedDate = t.Format("2006-01-02")

	// for note
	if nt.outputBaseName == "" {
		nt.outputBaseName = nt.fileName() + ".html"
	}

	nt.Href = path.Join("/", nt.outputDir, nt.outputBaseName)
	category := strings.Title(path.Base(nt.sourceDir))
	nt.Category = category

	f, _ := os.Create(path.Join(nt.outputDir, nt.outputBaseName))
	defer f.Close()

	tpl, _ := template.ParseFiles(layout)
	tpl.Execute(f, nt)
}

func (nt *Note) modifiedTime() (time.Time, error) {
	cmd := exec.Command("git", "log", "-n 1", "--pretty=%ai", nt.filePath)
	output, err := cmd.Output()
	if err != nil {
		return time.Time{}, err
	}
	t, err := time.Parse("2006-01-02 15:04:05 -0700", strings.TrimSpace(string(output)))
	if err != nil {
		return time.Time{}, err
	}
	return t, err
}

func (nt *Note) fileName() string {
	names := strings.Split(nt.baseName, ".")
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
