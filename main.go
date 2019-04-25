package main

import (
	"fmt"
	"os"
	// "github.com/davecgh/go-spew/spew"
	"html/template"
	"net/http"
)

var templates = template.Must(template.New("t").ParseGlob("templates/**/*.html"))

var errorTemplate = `
<html>
	<body>
		<h1>Error rendering template %s</h1>
		<p>%s</p>
	</body>
</html>
`

type data struct {
	Active  string
	Bucket  string
	Backend string
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, r, "index/home", data{Active: "home"})
	})

	mux.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		// load bucket url from env
		imagesURL := "http://" + os.Getenv("BUCKET")
		renderTemplate(w, r,
			"index/images",
			data{
				Active: "images",
				Bucket: imagesURL,
			},
		)
	})

	http.ListenAndServe(":3000", mux)
}

func renderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf(errorTemplate, name, err),
			http.StatusInternalServerError,
		)
	}

}
