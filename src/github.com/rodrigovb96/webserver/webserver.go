package webserver

import (
	"net/http"
	"html/template"
	"path/filepath"
	"strings"
	"parser"
)

// The index page
func indexPage(w http.ResponseWriter, r * http.Request) {
	filePath, _ := filepath.Abs("files/templates/index.html")
	tmpl, err := template.ParseFiles(filePath)

	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}


	if err := tmpl.Execute(w,nil); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

}

// The page with the result value
func parsePage(w http.ResponseWriter, r * http.Request) {
	command := r.URL.Path[ len("/parse/") : len("/parse/")+1 ]
	query := r.URL.Query()
	message := query.Get("input")

	var result string
	// From ASCII to morse 
	if command == "t" {
		result = parser.MorseCode.FromASCII(message)

	// From morse to ASCII
	} else if command == "f" {
		message := strings.Replace(message,"+"," ",-1)
		result = parser.MorseCode.ToASCII(message)
	}

	w.Write([]byte(result))
}

// Simple exposed function for initializing the server
func InitServer() {
	http.HandleFunc("/",indexPage)
	http.HandleFunc("/parse/",parsePage)

	if err := http.ListenAndServe(":8080",nil); err != nil {
		panic(err)
	}
}
