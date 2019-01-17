package pkg

import (
	"encoding/xml"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

const hostURL string = "https://dinolai.com"

type urlSet struct {
	XMLName        xml.Name `xml:"urlset"`
	XMLNS          string   `xml:"xmlns,attr"`
	XSI            string   `xml:"xmlns:xsi,attr"`
	SchemaLocation string   `xml:"xsi:schemaLocation,attr"`
	URLs           []*xmlURL
}

func NewURLSet(nl NoteList) *urlSet {
	us := new(urlSet)
	us.XMLNS = "https://www.sitemaps.org/schemas/sitemap/0.9"
	us.XSI = "https://www.w3.org/2001/XMLSchema-instance"
	us.SchemaLocation = "https://www.sitemaps.org/schemas/sitemap/0.9 https://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd"
	us.createURLs(nl)
	return us
}

func (us *urlSet) createURLs(nl NoteList) {
	sortCategories := make([]string, len(nl))
	i := 0
	for category := range nl {
		sortCategories[i] = category
		i++
	}
	sort.Sort(sort.StringSlice(sortCategories))
	for _, category := range sortCategories {
		for _, note := range nl[category] {
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
