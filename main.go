package main

import (
	"parser"
	"net/http"
	"strings"
	"fmt"
	"flag"
	"os"
	"bufio"
	"path/filepath"
	"player"
	"time"
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

func helloWorld(w http.ResponseWriter, r * http.Request) {
	message := r.URL.Path
	fmt.Println(message)
	message = strings.TrimPrefix(message,"/")
	fmt.Println(message)

	message = parser.MorseCode.FromASCII(message)

	fmt.Println(message)
	w.Write([]byte(message))

}

func applyFuncinFile(fileName string, procFunc func(string)string) string{
	absPath, _ := filepath.Abs(fileName)
	inputFile, err := os.Open(absPath)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	var finalString string
	for sucess := scanner.Scan(); sucess != false; sucess = scanner.Scan() {
		finalString += procFunc(scanner.Text()) + "\n"

	}


	return finalString

}


func main() {

	//http.HandleFunc("/",helloWorld)
	//http.HandleFunc("/from",procFrom)

	/*if err := http.ListenAndServe(":8080",nil); err != nil {
		panic(err)
	}*/

	fileName := flag.String("file","","Path to file from which the text will be converted from or to morse code")
	inputString := flag.String("inputText","","The text that will be converted from or to morse code")
	op := flag.String("operation","from","The operation (from or to) morse code")
	needToPlayAudio := flag.Bool("playAudio",false,"If setted and the operation=to, the program will play the morse code in audio")

	flag.Parse()

	operation := parser.MorseCode.ToASCII

	if *op == "to" {
		operation = parser.MorseCode.FromASCII
	}

	var result string
	if len(*inputString) != 0 {
		result = operation(*inputString)
	} else if len(*fileName) != 0 {
		result = applyFuncinFile(*fileName,operation)
	} else {
		fmt.Println("No input given")
	}

	fmt.Println(result)

	if *needToPlayAudio && (*op== "to") {

		for _, letter := range result {
			delay := time.Duration(0)

			if letter == '.' {
				player.PlayAudioFromFile("files/dot.wav")
			} else if letter == '-' {
				player.PlayAudioFromFile("files/dash.wav")
			} else if letter == '/' {
				delay = 100
			}


			time.Sleep(delay * time.Millisecond)
		}
	}



}
