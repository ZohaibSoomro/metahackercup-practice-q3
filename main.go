package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// CB=Cheese burger
const (
	inputFileName  = "dim_sum_delivery_input.txt"
	outputFileName = "output.txt"
	NO             = "NO"
	YES            = "YES"
)

var (
	aPos        = 1
	bPos        = 1
	isAliceTurn = true
)

func main() {
	testCases, sc, closeFile := parseInput(inputFileName)
	defer closeFile()
	outFile, err := os.OpenFile(outputFileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	for i := 1; i <= testCases; i++ {
		sc.Scan()
		values := strings.Fields(sc.Text())
		R := parseInt(values[0])
		C := parseInt(values[1])
		A := parseInt(values[2])
		B := parseInt(values[3])

		possibleWin := NO
		for {
			if isAliceTurn {
				if aPos == R {
					isAliceTurn = false
					continue
				}

				possibleMove := A
				// calculate max possible move count bw A..1
				for ; possibleMove > 1; possibleMove-- {
					if possibleMove <= R-aPos {
						break
					}
				}
				//alice moves
				aPos += possibleMove
				//if new position leads to R,C
				isAliceTurn = false
				if aPos == R && bPos == C {
					possibleWin = YES
					break
				}
			} else {
				// bob's turn
				//if bob at C
				if bPos == C {
					isAliceTurn = true
					continue
				}

				possibleMove := B
				// calculate max possible move count bw A..1
				for ; possibleMove > 1; possibleMove-- {
					if possibleMove <= C-bPos {
						break
					}
				}
				//bob moves
				bPos += possibleMove
				//if new position leads to R,C
				isAliceTurn = true
				if aPos == R && bPos == C {
					possibleWin = NO
					break
				}
			}
			if aPos >= R && bPos >= C {
				possibleWin = NO
				break
			}
		}
		outFile.WriteString(fmt.Sprintf("Case #%d: %s\n", i, possibleWin))
		aPos = 1
		bPos = 1
		isAliceTurn = true
	}
}

func parseInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func parseInput(fileName string) (int, *bufio.Scanner, func() error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(file)
	sc.Scan()
	testCases, _ := strconv.Atoi(sc.Text())
	return testCases, sc, file.Close
}
