package main

import (
	"github.com/rodrigovb96/morseParser/parser"
	"github.com/rodrigovb96/morseParser/webserver"

	"fmt"
	"flag"
)


var (
	inputString = flag.String("inputText","","The text that will be converted from or to morse code")
	op = flag.String("operation","to","The operation (from or to) morse code")
	mode = flag.String("mode","web","The mode that the program will operate( cli or web)")
)

func main() {

	flag.Parse()

	if *mode == "web" {

		fmt.Println("Web server mode, ignoring other flags")
		webserver.InitServer()

	} else {

		operation := parser.MorseCode.ToASCII

		if *op == "to" {
			operation = parser.MorseCode.FromASCII
		}

		var result string
		if len(*inputString) != 0 {
			result = operation(*inputString)
		} else {
			fmt.Println("No input given")
		}

		fmt.Println(result)


	}


}
