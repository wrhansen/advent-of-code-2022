/*
Advent of Code: Day 6
https://adventofcode.com/2022/day/6

Facts:

  - Start of packet marker is indicated by a sequence of the first four different
    characters

  - Identify where the four most recently received characters were all different

  - Specifically, report the number of characters from the beginning of the buffer to the end of the first such four-character marker.

*
*/
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	filename := os.Args[1]

	// Read file input
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	buffer := string(contents)
	fmt.Println(buffer)

	markerSize := 4
	markerPosition := getStartOfPacketMarker(buffer, markerSize)
	fmt.Printf("Start of Packet: %d\n", markerPosition)

	startOfMessageMarker := getStartOfPacketMarker(buffer, 14)
	fmt.Printf("Start of Packet: %d\n", startOfMessageMarker)

}

func getStartOfPacketMarker(buffer string, markerSize int) int {
	for idx := 0; idx < len(buffer)-markerSize; idx++ {
		markerFrame := buffer[idx : idx+markerSize]
		fmt.Printf("%d]: %v\n", idx, markerFrame)
		if allUniqueCharacters(markerFrame) == true {
			return idx + markerSize
		}
	}
	return -1
}

func allUniqueCharacters(frame string) bool {
	hashmap := make(map[string]bool)
	for _, char := range frame {
		_, ok := hashmap[string(char)]
		if ok {
			return false
		} else {
			hashmap[string(char)] = true
		}
	}
	return true
}
