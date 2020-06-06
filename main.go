package main

import (
	"github.com/rodrigovb96/parser"
	"github.com/rodrigovb96/webserver"
	"github.com/rodrigovb96/utils"

	"fmt"
	"flag"
)


var (
	fileName = flag.String("file","","Path to file from which the text will be converted from or to morse code")
	inputString = flag.String("inputText","","The text that will be converted from or to morse code")
	needToPlayAudio = flag.Bool("playAudio",false,"If setted and the operation=to, the program will play the morse code in audio")
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
		} else if len(*fileName) != 0 {
			result = utils.ApplyFuncinFile(*fileName,operation)
		} else {
			fmt.Println("No input given")
		}

		fmt.Println(result)

		if *needToPlayAudio && (*op== "to") {
			utils.PlayMorseCode(result)
		}

	}


}
