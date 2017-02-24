package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/vedhavyas/battleship"
)

func main() {

	var inputFile string
	var outputFile string
	log.SetFlags(log.Lshortfile)
	flag.StringVar(&inputFile, "input-file", "", "Input file with game details.")
	flag.StringVar(&outputFile, "output-file", "", "Output file to write result of the Game.")
	flag.Parse()

	if inputFile == "" {
		log.Fatal("Input file cannot be empty")
	}

	if outputFile == "" {
		log.Fatal("Output file cannot be empty")
	}

	gameData, err := readRawGameData(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	gameResult, err := battleship.PlayGame(gameData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(gameResult.Result)
	err = writeGameResult(gameResult, outputFile)
	if err != nil {
		log.Fatal(err)
	}
}

//readRawData returns the raw data from the input file
//We do not make any format check for raw data and program is expected to panic if
//gameData is not in correct format as expected
func readRawGameData(inputFile string) ([]string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	var data []string
	for reader.Scan() {
		data = append(data, strings.TrimSpace(reader.Text()))
	}

	return data, nil
}

//writeGameResult will write the game result to the outputFile
//if outputFile exist, then the file is overwritten with the newer data
// if not a new file is created
func writeGameResult(gameResult battleship.GameResult, outputFile string) error {
	var buf bytes.Buffer
	fmt.Fprintln(&buf, "Player1")
	fmt.Fprintln(&buf, gameResult.Player1Board)
	fmt.Fprintln(&buf, "Player2")
	fmt.Fprintln(&buf, gameResult.Player2Board)
	fmt.Fprintln(&buf, "")
	fmt.Fprintf(&buf, "P1: %v\n", gameResult.Player1Hits)
	fmt.Fprintf(&buf, "P2: %v\n", gameResult.Player2Hits)
	fmt.Fprint(&buf, gameResult.Result)
	return ioutil.WriteFile(outputFile, buf.Bytes(), 0644)
}
