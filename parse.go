package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MyKey struct {
	streetAddress string
	town          string
	saleDate      string
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Usage: ./Parse <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	rawFile, e := readFileIntoArray(filename)
	if e != nil {
		fmt.Println(e)
		return
	}
	println(len(rawFile))
	testOne(rawFile)
	testTwo(rawFile)
	testThree(rawFile)
}

/*In the case of duplicates, use the last encountered record.*/
func testOne(rawFile [][]string) {

	m := make(map[MyKey][]string)

	for i := range rawFile {
		key := extractAddressAndDate(rawFile[i])
		m[key] = rawFile[i]
	}
	//It says to print the list, but this looks the same.
	for _, value := range m {
		fmt.Println(value)
	}
}

/*Modify the code in case of duplicates to use the first encountered record.*/
func testTwo(rawFile [][]string) {

	m := make(map[MyKey][]string)
	//Test for the existence of a key in the map, and if it doesn't exist, add
	for i := range rawFile {
		key := extractAddressAndDate(rawFile[i])

		_, ok := m[key]
		if !ok {
			m[key] = rawFile[i]
		}
	}

	fmt.Println("Test two")
}

/*Instead of inserting the last record, make sure that no duplicates are entered at all. */
func testThree(rawFile [][]string) {

	//look here http://www.dotnetperls.com/duplicates-go
	//fmt.Println(rawFile)
	//fmt.Println("test three")
}

func filterByCheapProperties() {}
func filterOutCertainStreets() {}
func filterOneInTen()          {}

// Test #4
// Modify the codebase to run the following filters:
// Filter out cheap properties (anything under 400k)
// Filter out properties that are avenues, crescents, or places (AVE, CRES, PL) cos those guys are just pretentious...
// Filter out every 10th property (to keep our users on their toes!)

func readFileIntoArray(filename string) ([][]string, error) {
	file, e := os.Open(filename)

	if e != nil {
		fmt.Println(e)
		return nil, e
	}

	defer file.Close()
	var lines [][]string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if scanner.Text() != "" { //Gotta check for empty lines.
			line := strings.Split(scanner.Text(), "\t")
			lines = append(lines, line)
		}
	}
	return lines, e
}

func extractAddressAndDate(propertySale []string) MyKey {
	return MyKey{streetAddress: propertySale[1], town: propertySale[2], saleDate: propertySale[3]}
}
