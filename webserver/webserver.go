package webserver

import (
	"github.com/rodrigovb96/morseParser/parser"

	"net/http"
	"html/template"
	"path/filepath"
	"strings"
	"fmt"
	"os"
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

type response struct {
	ParsedText string
	Command	string
}

func processInput(command string, message string) response{

	var result string
	// From ASCII to morse 
	if command == "t" {
		result = parser.MorseCode.FromASCII(message)
	// From morse to ASCII
	} else if command == "f" {
		if strings.ContainsRune(message,' ') {
			result = parser.MorseCode.ToASCII(message)
		}else {
			result = "Invalid Input!"
		}

		command = ""
	}

	return response{ ParsedText: result, Command: command }

}

// The page with the result value
func parsePage(w http.ResponseWriter, r * http.Request) {
	command := r.URL.Path[ len("/parse/") : len("/parse/")+1 ]
	query := r.URL.Query()
	message := query.Get("input")

	resp := processInput(command,message)

	filePath, _ := filepath.Abs("files/templates/parsed.html")
	tmpl, err := template.ParseFiles(filePath)

	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}


	if err := tmpl.Execute(w,resp); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}


}

func determineListenAddress() (string,error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}

	return ":" + port, nil
}


// Simple exposed function for initializing the server
func InitServer() {
	http.HandleFunc("/",indexPage)
	http.HandleFunc("/parse/",parsePage)

	port, err := determineListenAddress()

	if err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
