package main

import (
	"html/template"
	"net/http"
	"os"
)

var index = `<!DOCTYPE html>
<html>
  <head>
    <style>
      body {
        background-color: {{ .bg_clr }};
      }
    </style>
  </head>
  <body></body>
</html>`

var indexTemplate *template.Template

func init() {
	indexTemplate = template.Must(template.New("index").Parse(index))
}

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet || r.URL.Path != "/" {
		return
	}

	color := map[string]string{
		"bg_clr": getenv("HELLO_BG_COLOR", "blue"),
	}
	_ = indexTemplate.Execute(w, color)
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
