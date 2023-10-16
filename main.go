package main

import (
	"bytes"
	"embed"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

//go:embed index.html
var TEMPLATE string

//go:embed static/bg.jpg
//go:embed static/favicon.ico
var static embed.FS

func renderTemplate(cluster string, show404 bool) (string, error) {
	t := template.Must(template.New("index-html").Parse(TEMPLATE))
	var tpl bytes.Buffer
	err := t.Execute(&tpl, map[string]interface{}{
		"Cluster": cluster,
		"Show404": show404,
	})
	if err != nil {
		return "", err
	}
	return tpl.String(), nil
}

func Server(port int, domain string, cluster string) error {
	HTML_200, _ := renderTemplate(cluster, false)
	HTML_404, _ := renderTemplate(cluster, true)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if r.Host == domain {
			w.WriteHeader(200)
			fmt.Fprint(w, HTML_200)
		} else {
			w.WriteHeader(404)
			fmt.Fprint(w, HTML_404)
		}
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, "ok\n")
	})
	fs := http.FileServer(http.FS(static))
	http.Handle("/static/", fs)

	fmt.Printf("Server started on 0.0.0.0:%d, see http://127.0.0.1:%d\n", port, port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func main() {
	Server(8000, os.Getenv("DOMAIN"), os.Getenv("CLUSTER"))
}
