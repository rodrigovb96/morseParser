package main

import (
	"strings"
	"os"
	"fmt"
	"bufio"
)


// A node for the tree
type node struct {
	value rune
	left  *node
	right *node
}

type morseParser struct {
	morseTree *node
	lookup map[rune]string
}

func (parser *morseParser) fromDotOrDashToChar(morseLetter string) rune {

	aux := parser.morseTree

	for _, dotOrDash := range morseLetter {

		if dotOrDash == '.' {
			aux = aux.left
		} else if dotOrDash == '-'{
			aux = aux.right
		// For spaces
		} else if dotOrDash == '/' {
			return ' '
		}

	}

	return aux.value
}

func (parser *morseParser) insert(value rune, directions string) bool {

	aux := parser.morseTree

	for _ , d := range directions {
		if d == '.' {
			if aux.left == nil {
				aux.left = new(node)
			}
			aux = aux.left
		} else if d == '-' {
			if aux.right == nil {
				aux.right = new(node)
			}
			aux = aux.right
		} else {
			return false
		}
	}


	aux.value = value

	return true
}



func (parser *morseParser)ToASCII(morseCode string)(asciiStr string) {

	morseLetters := strings.Split(morseCode," ")

	for _, letter := range morseLetters {
		asciiStr += string(parser.fromDotOrDashToChar(letter))
	}


	return
}


func (parser *morseParser) initTree() {
	morseFile, err := os.Open("init_morse.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(morseFile)

	scanner.Split(bufio.ScanLines)

	for sucess := scanner.Scan(); sucess != false; sucess = scanner.Scan() {
		valueKey := strings.Split(scanner.Text()," ")

		parser.lookup[rune(valueKey[0][0])] = valueKey[1]
		parser.insert(rune(valueKey[0][0]),valueKey[1])
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

func (parser *morseParser) FromASCII(asciiStr string)( morseCode string) {
	for _, letter := range strings.ToUpper(asciiStr){
		if letter == ' ' {
			morseCode += " "
		} else {
			morseCode += parser.lookup[rune(letter)]
		}

	}

	return
}


var Parser *morseParser

func init() {
	Parser = new(morseParser)
	Parser.morseTree = new(node)
	Parser.lookup = make(map[rune]string)
	Parser.initTree()

}


func main() {

	fmt.Println(Parser.ToASCII(".-. --- -.. .-. .. --. --- / ...- .- .-.. . -. - . / -... . .-. -. .- .-. -.. . ... "))

	fmt.Println(Parser.FromASCII("amanda ferro"))


}
