package utils

import (
	"player"
	"os"
	"bufio"
	"path/filepath"
	"time"
)


func ApplyFuncinFile(fileName string, procFunc func(string)string) string{
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


func PlayMorseCode(morseCode string) {

	for _, letter := range morseCode{
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
