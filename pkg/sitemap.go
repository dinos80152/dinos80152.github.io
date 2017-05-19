package pkg

import (
	"encoding/xml"
	"net/url"
	"os"
	"strconv"
	"strings"
	"text/template"
)

const hostURL string = "http://dinolai.com"

type urlSet struct {
	XMLName        xml.Name `xml:"urlset"`
	XMLNS          string   `xml:"xmlns,attr"`
	XSI            string   `xml:"xmlns:xsi,attr"`
	SchemaLocation string   `xml:"xsi:schemaLocation,attr"`
	URLs           []*xmlURL
}

func NewURLSet(nl NoteList) *urlSet {
	us := new(urlSet)
	us.XMLNS = "http://www.sitemaps.org/schemas/sitemap/0.9"
	us.XSI = "http://www.w3.org/2001/XMLSchema-instance"
	us.SchemaLocation = "http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd"
	us.createURLs(nl)
	return us
}

func (us *urlSet) createURLs(nl NoteList) {
	for _, notes := range nl {
		for _, note := range notes {
			u := &xmlURL{
				Location:   hostURL + note.Href,
				Lastmod:    note.UpdatedDate,
				ChangeFreq: "daily",
			}
			u.calculatePriority()
			us.URLs = append(us.URLs, u)
		}
	}
}

func (us *urlSet) GenSiteMap() {
	output, _ := xml.MarshalIndent(us, "", "    ")

	f, _ := os.Create("sitemap.xml")
	defer f.Close()

	t, _ := template.ParseFiles("tpls/sitemap.gtpl")
	t.Execute(f, string(output))

}

type xmlURL struct {
	XMLName    xml.Name `xml:"url"`
	Location   string   `xml:"loc"`
	Lastmod    string   `xml:"lastmod"`
	ChangeFreq string   `xml:"changefreq"`
	Priority   float64  `xml:"priority"`
}

func (xu *xmlURL) calculatePriority() {
	u, _ := url.Parse(xu.Location)
	c := strings.Count(u.Path, "/")
	result := 0.5 + 0.5/float64(c)
	str := strconv.FormatFloat(result, 'f', 2, 64)
	xu.Priority, _ = strconv.ParseFloat(str, 64)
}
