# Note List
|Title | Updated Date |
|------|------------|
{{ range $category, $notes := . }}|  **{{ $category }}** | |
{{ range $key, $note := $notes }}| [{{ $note.Title }}]({{ $note.Href }}) | {{$note.UpdatedDate}} |
{{ end }}{{ end }}