# Note List

{{ range $category, $notes := . }}## {{ $category }}

|Title | Updated Date |
|------|------------|
{{ range $key, $note := $notes }}| [{{ $note.Title }}]({{ $note.Href }}) | {{$note.UpdatedDate}} |
{{ end }}
{{ end }}