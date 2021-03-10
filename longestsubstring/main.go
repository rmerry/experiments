package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("USAGE: %s <dict-file> <string>\n", os.Args[0])
		os.Exit(1)
	}

	var (
		match    bool
		dictFile string = strings.ToLower(os.Args[1])
		target   string = strings.ToLower(os.Args[2])
	)

	file, err := os.Open(dictFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var (
		closestMatch int
		closestWord  string
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		w := strings.ToLower(scanner.Text())
		if target == w {
			match = true
			break
		}

		var (
			grid [][]int = make([][]int, len(target))
		)

		for i := 0; i < len(target); i++ {
			var row = make([]int, len(w))
			for l := 0; l < len(w); l++ {
				// Calculate max neighbour value.
				var (
					topNeighbour     int
					leftNeighbour    int
					topLeftNeighbour int
					maxNeighbour     int
				)
				if i > 0 {
					topNeighbour = grid[i-1][l]
				}
				if l > 0 {
					leftNeighbour = row[l-1]
				}
				if l > 0 && i > 0 {
					topLeftNeighbour = grid[i-1][l-1]
				}
				if topNeighbour > leftNeighbour {
					maxNeighbour = topNeighbour
				} else {
					maxNeighbour = leftNeighbour
				}

				// Set cell value.
				if strings.ToLower(string(target[i])) == strings.ToLower(string(w[l])) {
					row[l] = topLeftNeighbour + 1
					if row[l] > closestMatch {
						closestMatch = row[l]
						closestWord = w
					}
				} else {
					row[l] = maxNeighbour
				}
			}
			grid[i] = row
		}
		// printGrid(grid, target, w)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if match {
		fmt.Println("match found")
	} else {
		fmt.Println("no direct match found, possible matches:\n")
		fmt.Printf(" - %s\n", closestWord)
	}
}

func printGrid(g [][]int, target string, test string) {
	fmt.Printf("\n\n")
	fmt.Printf("  ")
	for _, c := range test {
		fmt.Printf("%4s", string(c))
	}
	fmt.Printf("\n\n")
	for i := 0; i < len(g); i++ {
		fmt.Printf("%s ", string(target[i]))
		for l := 0; l < len(g[i]); l++ {
			fmt.Printf("%4d", g[i][l])
		}
		fmt.Printf("\n")
	}
}
