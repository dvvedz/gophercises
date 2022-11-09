package story

import (
	"encoding/json"
	"net/http"
	"text/template"
)

var defaultTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Text adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
     
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}} 

    <ul>
        {{range .Options}}
        <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
    </ul>
</body>
</html>
`

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

type handler struct {
	s Story
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("").Parse(defaultTemplate))
}

func HandleJson(jd []byte) (Story, error) {
	var story Story
	if err := json.Unmarshal(jd, &story); err != nil {
		return nil, err
	}
	return story, nil
}
