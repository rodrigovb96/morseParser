package main

import (
	"parser"
	"net/http"
	"strings"
	"fmt"
)

func procFrom(w http.ResponseWriter, r * http.Request) {
	message := r.URL.Path
	fmt.Println(message)
	message = strings.TrimPrefix(message,"/")
	fmt.Println(message)

	message = parser.MorseCode.FromASCII(message)

	fmt.Println(message)
	w.Write([]byte(message))

}

func main() {

	http.HandleFunc("/from",procFrom)

	if err := http.ListenAndServe(":8080",nil); err != nil {
		panic(err)
	}

	fmt.Println(parser.MorseCode.ToASCII(".-. --- -.. .-. .. --. --- / ...- .- .-.. . -. - . / -... . .-. -. .- .-. -.. . ... "))

	fmt.Println(parser.MorseCode.FromASCII("amanda ferro"))


}
