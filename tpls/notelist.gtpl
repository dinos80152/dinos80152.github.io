# Note List
|Title | Updated Date |
|------|------------|
{{ range $key, $note := . }}| [{{ $note.Title }}]({{ $note.Href }}) | {{$note.UpdatedDate}} |
{{ end }}