# Last Updated
{{ $log := .Log }}
{{ range $key, $date := .LogDate }}## {{ $date }}

{{ range $index, $note := index $log $date }}* [[{{ $note.Category }}] {{ $note.Title }}]({{ $note.Href }})
{{ end }}
{{ end }}