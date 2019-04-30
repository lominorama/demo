package main

import (
	"database/sql"
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"os"
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
	Active    string
	FirstImg  string
	SecondImg string
	ThirdImg  string
}

type image struct {
	ID   int
	Name string
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, r, "index/home", data{Active: "home"})
	})

	mux.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		var bucket, dbName, dbUser, dbPass, dbHost, dbPort string

		if os.Getenv("BUCKET") != "" {
			bucket = os.Getenv("BUCKET")
		} else {
			panic("Missing BUCKET env variable")
		}

		if os.Getenv("DB_NAME") != "" {
			dbName = os.Getenv("DB_NAME")
		} else {
			dbName = "demo"
		}

		if os.Getenv("DB_USER") != "" {
			dbUser = os.Getenv("DB_USER")
		} else {
			dbUser = "demo"
		}

		if os.Getenv("DB_PASSWORD") != "" {
			dbPass = os.Getenv("DB_PASSWORD")
		} else {
			dbPass = "demo"
		}

		if os.Getenv("DB_HOST") != "" {
			dbHost = os.Getenv("DB_HOST")
		} else {
			dbHost = "localhost"
		}

		if os.Getenv("DB_PORT") != "" {
			dbPort = os.Getenv("DB_PORT")
		} else {
			dbPort = "3306"
		}

		imagesURL := "http://" + bucket
		d := data{Active: "images"}

		connectionString := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
		db, err := sql.Open("mysql", connectionString)
		if err != nil {
			panic(err.Error())
		}

		defer db.Close()

		_, err = db.Query("SELECT id, name FROM images")
		if err != nil {
			log.Println(err.Error())
			//If an error occurs fallback to v2 images
			d.FirstImg = imagesURL + "/image-1.jpg"
			d.SecondImg = imagesURL + "/image-2.jpg"
			d.ThirdImg = imagesURL + "/image-3.jpg"

		} else {
			var img image

			result := db.QueryRow("SELECT id, name FROM images where id=1")
			err = result.Scan(&img.ID, &img.Name)
			d.FirstImg = imagesURL + "/" + img.Name

			result = db.QueryRow("SELECT id, name FROM images where id=2")
			err = result.Scan(&img.ID, &img.Name)
			d.SecondImg = imagesURL + "/" + img.Name

			result = db.QueryRow("SELECT id, name FROM images where id=3")
			err = result.Scan(&img.ID, &img.Name)
			d.ThirdImg = imagesURL + "/" + img.Name

		}

		renderTemplate(w, r, "index/images", d)
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
