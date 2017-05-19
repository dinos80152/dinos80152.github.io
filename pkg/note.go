package pkg

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"

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
}

type NoteList map[string][]*Note

func NewNote(sourceDir, outputDir string, fi os.FileInfo, outputBaseName string) *Note {
	note := new(Note)
	note.sourceDir = sourceDir
	note.outputDir = outputDir
	note.fi = fi
	note.outputBaseName = outputBaseName
	return note
}

func (nt *Note) ToHTML() {
	baseName := nt.fi.Name()
	input, _ := ioutil.ReadFile(path.Join(nt.sourceDir, baseName))
	output := markdown(input)
	content := string(output)

	// TODO: use regexp
	title := content[strings.Index(content, ">")+1 : strings.Index(content, "/")-1]

	nt.Title = title
	nt.Content = content
	nt.Category = nt.outputDir
	nt.UpdatedAt = nt.fi.ModTime().Format("2006-01-02 15:04:05 MST Z07")
	nt.UpdatedDate = nt.fi.ModTime().Format("2006-01-02")

	// for note
	if nt.outputBaseName == "" {
		nt.outputBaseName = fileName(baseName) + ".html"
	}

	nt.Href = path.Join("/", nt.outputDir, nt.outputBaseName)
	category := strings.Title(path.Base(nt.sourceDir))
	nt.Category = category

	f, _ := os.Create(path.Join(nt.outputDir, nt.outputBaseName))
	defer f.Close()

	t, _ := template.ParseFiles(layout)
	t.Execute(f, nt)
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
